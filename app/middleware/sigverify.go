package middleware

import (
	"context"
	"encoding/hex"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	sdktx "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authmiddleware "github.com/cosmos/cosmos-sdk/x/auth/middleware"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/gogo/protobuf/proto"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	// simulation signature values used to estimate gas consumption
	key                = make([]byte, secp256k1.PubKeySize)
	simSecp256k1Pubkey = &secp256k1.PubKey{Key: key}
	simSecp256k1Sig    [64]byte
)

func init() {
	// This decodes a valid hex string into a sepc256k1Pubkey for use in transaction simulation
	bz, _ := hex.DecodeString("035AD6810A47F073553FF30D2FCC7E0D3B1C0B74B61A1AAA2582344037151E143A")
	copy(key, bz)
	simSecp256k1Pubkey.Key = key
}

// SignatureVerificationGasConsumer is the type of function that is used to both
// consume gas when verifying signatures and also to accept or reject different types of pubkeys
// This is where apps can define their own PubKey
type SignatureVerificationGasConsumer = func(meter sdk.GasMeter, sig signing.SignatureV2, params types.Params) error

var _ sdktx.Handler = sigVerificationTxHandler{}

type sigVerificationTxHandler struct {
	cdc             codec.Codec
	ak              authmiddleware.AccountKeeper
	signModeHandler authsigning.SignModeHandler
	next            sdktx.Handler
}

// SigVerificationMiddleware verifies all signatures for a tx and return an error if any are invalid. Note,
// the sigVerificationTxHandler middleware will not get executed on ReCheck.
//
// CONTRACT: Pubkeys are set in context for all signers before this middleware runs
// CONTRACT: Tx must implement SigVerifiableTx interface
func SigVerificationMiddleware(cdc codec.Codec, ak authmiddleware.AccountKeeper, signModeHandler authsigning.SignModeHandler) sdktx.Middleware {
	return func(h sdktx.Handler) sdktx.Handler {
		return sigVerificationTxHandler{
			cdc:             cdc,
			ak:              ak,
			signModeHandler: signModeHandler,
			next:            h,
		}
	}
}

// OnlyLegacyAminoSigners checks SignatureData to see if all
// signers are using SIGN_MODE_LEGACY_AMINO_JSON. If this is the case
// then the corresponding SignatureV2 struct will not have account sequence
// explicitly set, and we should skip the explicit verification of sig.Sequence
// in the SigVerificationMiddleware's middleware function.
func OnlyLegacyAminoSigners(sigData signing.SignatureData) bool {
	switch v := sigData.(type) {
	case *signing.SingleSignatureData:
		return v.SignMode == signing.SignMode_SIGN_MODE_LEGACY_AMINO_JSON
	case *signing.MultiSignatureData:
		for _, s := range v.Signatures {
			if !OnlyLegacyAminoSigners(s) {
				return false
			}
		}
		return true
	default:
		return false
	}
}

func (svd sigVerificationTxHandler) sigVerify(ctx context.Context, req sdktx.Request, isReCheckTx, simulate bool) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	// no need to verify signatures on recheck tx
	if isReCheckTx {
		return nil
	}
	sigTx, ok := req.Tx.(authsigning.SigVerifiableTx)
	if !ok {
		return sdkerrors.Wrap(sdkerrors.ErrTxDecode, "invalid transaction type")
	}

	// stdSigs contains the sequence number, account number, and signatures.
	// When simulating, this would just be a 0-length slice.
	sigs, err := sigTx.GetSignaturesV2()
	if err != nil {
		return err
	}

	signerAddrs := sigTx.GetSigners()

	// check that signer length and signature length are the same
	if len(sigs) != len(signerAddrs) {
		return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "invalid number of signer;  expected: %d, got %d", len(signerAddrs), len(sigs))
	}

	for i, sig := range sigs {
		acc, err := authmiddleware.GetSignerAcc(sdkCtx, svd.ak, signerAddrs[i])
		if err != nil {
			return err
		}

		// retrieve pubkey
		pubKey := acc.GetPubKey()
		if !simulate && pubKey == nil {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidPubKey, "pubkey on account is not set")
		}

		// Check account sequence number.
		if sig.Sequence != acc.GetSequence() {
			return sdkerrors.Wrapf(
				sdkerrors.ErrWrongSequence,
				"account sequence mismatch, expected %d, got %d", acc.GetSequence(), sig.Sequence,
			)
		}

		// retrieve signer data
		genesis := sdkCtx.BlockHeight() == 0
		chainID := sdkCtx.ChainID()
		var accNum uint64
		if !genesis {
			accNum = acc.GetAccountNumber()
		}

		signerData := authsigning.SignerData{
			Address:       signerAddrs[i].String(),
			ChainID:       chainID,
			AccountNumber: accNum,
			Sequence:      acc.GetSequence(),
			PubKey:        pubKey,
		}

		if !simulate {
			err := authsigning.VerifySignature(pubKey, signerData, sig.Data, svd.signModeHandler, req.Tx)
			if err != nil {
				// try verifying signature with etherum
				if ethErr := VerifyEthereumSignature(svd.cdc, pubKey, signerData, sig.Data, svd.signModeHandler, req.Tx); ethErr == nil {
					return nil
				} else {
					fmt.Printf("ethereum signature verification failed; %s", ethErr.Error())
				}

				var errMsg string
				if OnlyLegacyAminoSigners(sig.Data) {
					// If all signers are using SIGN_MODE_LEGACY_AMINO, we rely on VerifySignature to check account sequence number,
					// and therefore communicate sequence number as a potential cause of error.
					errMsg = fmt.Sprintf("signature verification failed; please verify account number (%d), sequence (%d) and chain-id (%s)", accNum, acc.GetSequence(), chainID)
				} else {
					errMsg = fmt.Sprintf("signature verification failed; please verify account number (%d) and chain-id (%s)", accNum, chainID)
				}
				return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, errMsg)
			}
		}
	}

	return nil
}

// CheckTx implements tx.Handler.CheckTx.
func (svd sigVerificationTxHandler) CheckTx(ctx context.Context, req sdktx.Request, checkReq sdktx.RequestCheckTx) (sdktx.Response, sdktx.ResponseCheckTx, error) {
	if err := svd.sigVerify(ctx, req, checkReq.Type == abci.CheckTxType_Recheck, false); err != nil {
		return sdktx.Response{}, sdktx.ResponseCheckTx{}, err
	}

	return svd.next.CheckTx(ctx, req, checkReq)
}

// DeliverTx implements tx.Handler.DeliverTx.
func (svd sigVerificationTxHandler) DeliverTx(ctx context.Context, req sdktx.Request) (sdktx.Response, error) {
	if err := svd.sigVerify(ctx, req, false, false); err != nil {
		return sdktx.Response{}, err
	}

	return svd.next.DeliverTx(ctx, req)
}

// SimulateTx implements tx.Handler.SimulateTx.
func (svd sigVerificationTxHandler) SimulateTx(ctx context.Context, req sdktx.Request) (sdktx.Response, error) {
	if err := svd.sigVerify(ctx, req, false, true); err != nil {
		return sdktx.Response{}, err
	}

	return svd.next.SimulateTx(ctx, req)
}

func VerifyEthereumSignature(cdc codec.Codec, pubKey cryptotypes.PubKey, signerData authsigning.SignerData, sigData signing.SignatureData, handler authsigning.SignModeHandler, tx sdk.Tx) error {
	switch data := sigData.(type) {
	case *signing.SingleSignatureData:
		signBytes, err := handler.GetSignBytes(data.SignMode, signerData, tx)
		if err != nil {
			return err
		}

		fmt.Println("data.SignMode", data.SignMode)
		fmt.Println("defaultSignBytes", string(signBytes))
		if data.SignMode == signing.SignMode_SIGN_MODE_DIRECT {
			signDoc := sdktx.SignDoc{}
			err = signDoc.Unmarshal(signBytes)
			if err != nil {
				return err
			}

			signDocMetamask := SignDocForMetamask{
				Body:          &sdktx.TxBody{},
				AuthInfo:      &sdktx.AuthInfo{},
				ChainId:       signerData.ChainID,
				AccountNumber: signerData.AccountNumber,
			}

			err = proto.Unmarshal(signDoc.BodyBytes, signDocMetamask.Body)
			if err != nil {
				return err
			}

			err = proto.Unmarshal(signDoc.AuthInfoBytes, signDocMetamask.AuthInfo)
			if err != nil {
				return err
			}

			signBytes = cdc.MustMarshalJSON(&signDocMetamask)
			fmt.Println("finalSignBytes", string(signBytes))
		}

		signatureData := data.Signature
		if len(signatureData) <= crypto.RecoveryIDOffset {
			return fmt.Errorf("not a correct ethereum signature")
		}
		signatureData[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1
		recovered, err := crypto.SigToPub(accounts.TextHash(signBytes), signatureData)
		if err != nil {
			return err
		}

		recoveredPubKey := secp256k1.PubKey{Key: crypto.CompressPubkey(recovered)}
		recoveredAddress := sdk.AccAddress(recoveredPubKey.Address())
		recoveredEthereumAddr := crypto.PubkeyToAddress(*recovered)

		sigTx, ok := tx.(authsigning.SigVerifiableTx)
		if !ok {
			return sdkerrors.Wrap(sdkerrors.ErrTxDecode, "invalid transaction type")
		}

		signerAddrs := sigTx.GetSigners()
		if len(signerAddrs) != 1 {
			return fmt.Errorf("only 1 signer transaction supported: got %d", len(signerAddrs))
		}

		address := signerAddrs[0]

		if recoveredAddress.String() != address.String() {
			return fmt.Errorf("mismatching recovered address and sender.\n  recovered: %s\n  sender: %s\n  recovered ethereum address: %s", recoveredAddress.String(), address.String(), recoveredEthereumAddr.String())
		}
		return nil
	default:
		return fmt.Errorf("unexpected SignatureData %T", sigData)
	}
}

// signatureDataToBz converts a SignatureData into raw bytes signature.
// For SingleSignatureData, it returns the signature raw bytes.
// For MultiSignatureData, it returns an array of all individual signatures,
// as well as the aggregated signature.
func signatureDataToBz(data signing.SignatureData) ([][]byte, error) {
	if data == nil {
		return nil, fmt.Errorf("got empty SignatureData")
	}

	switch data := data.(type) {
	case *signing.SingleSignatureData:
		return [][]byte{data.Signature}, nil
	case *signing.MultiSignatureData:
		sigs := [][]byte{}
		var err error

		for _, d := range data.Signatures {
			nestedSigs, err := signatureDataToBz(d)
			if err != nil {
				return nil, err
			}
			sigs = append(sigs, nestedSigs...)
		}

		multisig := cryptotypes.MultiSignature{
			Signatures: sigs,
		}
		aggregatedSig, err := multisig.Marshal()
		if err != nil {
			return nil, err
		}
		sigs = append(sigs, aggregatedSig)

		return sigs, nil
	default:
		return nil, sdkerrors.ErrInvalidType.Wrapf("unexpected signature data type %T", data)
	}
}
