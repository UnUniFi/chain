<!-- This file is auto-generated. Please do not modify it yourself. -->
# Protobuf Documentation
<a name="top"></a>

## Table of Contents

- [auction/auction.proto](#auction/auction.proto)
    - [BaseAuction](#ununifi.auction.BaseAuction)
    - [CollateralAuction](#ununifi.auction.CollateralAuction)
    - [DebtAuction](#ununifi.auction.DebtAuction)
    - [Params](#ununifi.auction.Params)
    - [SurplusAuction](#ununifi.auction.SurplusAuction)
    - [WeightedAddress](#ununifi.auction.WeightedAddress)
  
- [auction/genesis.proto](#auction/genesis.proto)
    - [GenesisState](#ununifi.auction.GenesisState)
  
- [auction/query.proto](#auction/query.proto)
    - [QueryAllAuctionRequest](#ununifi.auction.QueryAllAuctionRequest)
    - [QueryAllAuctionResponse](#ununifi.auction.QueryAllAuctionResponse)
    - [QueryGetAuctionRequest](#ununifi.auction.QueryGetAuctionRequest)
    - [QueryGetAuctionResponse](#ununifi.auction.QueryGetAuctionResponse)
    - [QueryParamsRequest](#ununifi.auction.QueryParamsRequest)
    - [QueryParamsResponse](#ununifi.auction.QueryParamsResponse)
  
    - [Query](#ununifi.auction.Query)
  
- [auction/tx.proto](#auction/tx.proto)
    - [MsgPlaceBid](#ununifi.auction.MsgPlaceBid)
    - [MsgPlaceBidResponse](#ununifi.auction.MsgPlaceBidResponse)
  
    - [Msg](#ununifi.auction.Msg)
  
- [cdp/cdp.proto](#cdp/cdp.proto)
    - [AugmentedCdp](#ununifi.cdp.AugmentedCdp)
    - [Cdp](#ununifi.cdp.Cdp)
    - [CollateralParam](#ununifi.cdp.CollateralParam)
    - [DebtParam](#ununifi.cdp.DebtParam)
    - [Deposit](#ununifi.cdp.Deposit)
    - [Params](#ununifi.cdp.Params)
  
- [cdp/genesis.proto](#cdp/genesis.proto)
    - [GenesisAccumulationTime](#ununifi.cdp.GenesisAccumulationTime)
    - [GenesisState](#ununifi.cdp.GenesisState)
    - [GenesisTotalPrincipal](#ununifi.cdp.GenesisTotalPrincipal)
  
- [cdp/query.proto](#cdp/query.proto)
    - [QueryAllAccountRequest](#ununifi.cdp.QueryAllAccountRequest)
    - [QueryAllAccountResponse](#ununifi.cdp.QueryAllAccountResponse)
    - [QueryAllCdpRequest](#ununifi.cdp.QueryAllCdpRequest)
    - [QueryAllCdpResponse](#ununifi.cdp.QueryAllCdpResponse)
    - [QueryAllDepositRequest](#ununifi.cdp.QueryAllDepositRequest)
    - [QueryAllDepositResponse](#ununifi.cdp.QueryAllDepositResponse)
    - [QueryGetCdpRequest](#ununifi.cdp.QueryGetCdpRequest)
    - [QueryGetCdpResponse](#ununifi.cdp.QueryGetCdpResponse)
    - [QueryParamsRequest](#ununifi.cdp.QueryParamsRequest)
    - [QueryParamsResponse](#ununifi.cdp.QueryParamsResponse)
  
    - [Query](#ununifi.cdp.Query)
  
- [cdp/tx.proto](#cdp/tx.proto)
    - [MsgCreateCdp](#ununifi.cdp.MsgCreateCdp)
    - [MsgCreateCdpResponse](#ununifi.cdp.MsgCreateCdpResponse)
    - [MsgDeposit](#ununifi.cdp.MsgDeposit)
    - [MsgDepositResponse](#ununifi.cdp.MsgDepositResponse)
    - [MsgDrawDebt](#ununifi.cdp.MsgDrawDebt)
    - [MsgDrawDebtResponse](#ununifi.cdp.MsgDrawDebtResponse)
    - [MsgLiquidate](#ununifi.cdp.MsgLiquidate)
    - [MsgLiquidateResponse](#ununifi.cdp.MsgLiquidateResponse)
    - [MsgRepayDebt](#ununifi.cdp.MsgRepayDebt)
    - [MsgRepayDebtResponse](#ununifi.cdp.MsgRepayDebtResponse)
    - [MsgWithdraw](#ununifi.cdp.MsgWithdraw)
    - [MsgWithdrawResponse](#ununifi.cdp.MsgWithdrawResponse)
  
    - [Msg](#ununifi.cdp.Msg)
  
- [derivatives/derivatives.proto](#derivatives/derivatives.proto)
    - [EventPriceIsNotFeeded](#ununifi.derivatives.EventPriceIsNotFeeded)
    - [Market](#ununifi.derivatives.Market)
    - [PoolMarketCap](#ununifi.derivatives.PoolMarketCap)
    - [PoolMarketCap.Breakdown](#ununifi.derivatives.PoolMarketCap.Breakdown)
    - [PoolParams](#ununifi.derivatives.PoolParams)
    - [PoolParams.Asset](#ununifi.derivatives.PoolParams.Asset)
    - [Position](#ununifi.derivatives.Position)
    - [QueriedPosition](#ununifi.derivatives.QueriedPosition)
  
    - [PositionType](#ununifi.derivatives.PositionType)
  
- [derivatives/perpetual_futures.proto](#derivatives/perpetual_futures.proto)
    - [EventPerpetualFuturesPositionClosed](#ununifi.derivatives.EventPerpetualFuturesPositionClosed)
    - [EventPerpetualFuturesPositionLevied](#ununifi.derivatives.EventPerpetualFuturesPositionLevied)
    - [EventPerpetualFuturesPositionLiquidated](#ununifi.derivatives.EventPerpetualFuturesPositionLiquidated)
    - [EventPerpetualFuturesPositionOpened](#ununifi.derivatives.EventPerpetualFuturesPositionOpened)
    - [PerpetualFuturesNetPositionOfMarket](#ununifi.derivatives.PerpetualFuturesNetPositionOfMarket)
    - [PerpetualFuturesParams](#ununifi.derivatives.PerpetualFuturesParams)
    - [PerpetualFuturesPosition](#ununifi.derivatives.PerpetualFuturesPosition)
    - [PerpetualFuturesPositionInstance](#ununifi.derivatives.PerpetualFuturesPositionInstance)
  
- [derivatives/perpetual_options.proto](#derivatives/perpetual_options.proto)
    - [PerpetualOptionsParams](#ununifi.derivatives.PerpetualOptionsParams)
    - [PerpetualOptionsPositionInstance](#ununifi.derivatives.PerpetualOptionsPositionInstance)
  
    - [OptionType](#ununifi.derivatives.OptionType)
  
- [derivatives/params.proto](#derivatives/params.proto)
    - [Params](#ununifi.derivatives.Params)
  
- [derivatives/genesis.proto](#derivatives/genesis.proto)
    - [GenesisState](#ununifi.derivatives.GenesisState)
  
- [derivatives/query.proto](#derivatives/query.proto)
    - [QueryAddressPositionsRequest](#ununifi.derivatives.QueryAddressPositionsRequest)
    - [QueryAddressPositionsResponse](#ununifi.derivatives.QueryAddressPositionsResponse)
    - [QueryAllPositionsRequest](#ununifi.derivatives.QueryAllPositionsRequest)
    - [QueryAllPositionsResponse](#ununifi.derivatives.QueryAllPositionsResponse)
    - [QueryDLPTokenRateRequest](#ununifi.derivatives.QueryDLPTokenRateRequest)
    - [QueryDLPTokenRateResponse](#ununifi.derivatives.QueryDLPTokenRateResponse)
    - [QueryEstimateDLPTokenAmountRequest](#ununifi.derivatives.QueryEstimateDLPTokenAmountRequest)
    - [QueryEstimateDLPTokenAmountResponse](#ununifi.derivatives.QueryEstimateDLPTokenAmountResponse)
    - [QueryEstimateRedeemAmountRequest](#ununifi.derivatives.QueryEstimateRedeemAmountRequest)
    - [QueryEstimateRedeemAmountResponse](#ununifi.derivatives.QueryEstimateRedeemAmountResponse)
    - [QueryLiquidityProviderTokenNominalAPYRequest](#ununifi.derivatives.QueryLiquidityProviderTokenNominalAPYRequest)
    - [QueryLiquidityProviderTokenNominalAPYResponse](#ununifi.derivatives.QueryLiquidityProviderTokenNominalAPYResponse)
    - [QueryLiquidityProviderTokenRealAPYRequest](#ununifi.derivatives.QueryLiquidityProviderTokenRealAPYRequest)
    - [QueryLiquidityProviderTokenRealAPYResponse](#ununifi.derivatives.QueryLiquidityProviderTokenRealAPYResponse)
    - [QueryParamsRequest](#ununifi.derivatives.QueryParamsRequest)
    - [QueryParamsResponse](#ununifi.derivatives.QueryParamsResponse)
    - [QueryPerpetualFuturesMarketRequest](#ununifi.derivatives.QueryPerpetualFuturesMarketRequest)
    - [QueryPerpetualFuturesMarketResponse](#ununifi.derivatives.QueryPerpetualFuturesMarketResponse)
    - [QueryPerpetualFuturesPositionSizeRequest](#ununifi.derivatives.QueryPerpetualFuturesPositionSizeRequest)
    - [QueryPerpetualFuturesPositionSizeResponse](#ununifi.derivatives.QueryPerpetualFuturesPositionSizeResponse)
    - [QueryPerpetualFuturesRequest](#ununifi.derivatives.QueryPerpetualFuturesRequest)
    - [QueryPerpetualFuturesResponse](#ununifi.derivatives.QueryPerpetualFuturesResponse)
    - [QueryPerpetualOptionsMarketRequest](#ununifi.derivatives.QueryPerpetualOptionsMarketRequest)
    - [QueryPerpetualOptionsMarketResponse](#ununifi.derivatives.QueryPerpetualOptionsMarketResponse)
    - [QueryPerpetualOptionsRequest](#ununifi.derivatives.QueryPerpetualOptionsRequest)
    - [QueryPerpetualOptionsResponse](#ununifi.derivatives.QueryPerpetualOptionsResponse)
    - [QueryPoolRequest](#ununifi.derivatives.QueryPoolRequest)
    - [QueryPoolResponse](#ununifi.derivatives.QueryPoolResponse)
    - [QueryPositionRequest](#ununifi.derivatives.QueryPositionRequest)
    - [QueryPositionResponse](#ununifi.derivatives.QueryPositionResponse)
  
    - [Query](#ununifi.derivatives.Query)
  
- [derivatives/tx.proto](#derivatives/tx.proto)
    - [MsgBurnLiquidityProviderToken](#ununifi.derivatives.MsgBurnLiquidityProviderToken)
    - [MsgBurnLiquidityProviderTokenResponse](#ununifi.derivatives.MsgBurnLiquidityProviderTokenResponse)
    - [MsgClosePosition](#ununifi.derivatives.MsgClosePosition)
    - [MsgClosePositionResponse](#ununifi.derivatives.MsgClosePositionResponse)
    - [MsgMintLiquidityProviderToken](#ununifi.derivatives.MsgMintLiquidityProviderToken)
    - [MsgMintLiquidityProviderTokenResponse](#ununifi.derivatives.MsgMintLiquidityProviderTokenResponse)
    - [MsgOpenPosition](#ununifi.derivatives.MsgOpenPosition)
    - [MsgOpenPositionResponse](#ununifi.derivatives.MsgOpenPositionResponse)
    - [MsgReportLevyPeriod](#ununifi.derivatives.MsgReportLevyPeriod)
    - [MsgReportLevyPeriodResponse](#ununifi.derivatives.MsgReportLevyPeriodResponse)
    - [MsgReportLiquidation](#ununifi.derivatives.MsgReportLiquidation)
    - [MsgReportLiquidationResponse](#ununifi.derivatives.MsgReportLiquidationResponse)
  
    - [Msg](#ununifi.derivatives.Msg)
  
- [ecosystem-incentive/ecosystem_incentive.proto](#ecosystem-incentive/ecosystem_incentive.proto)
    - [IncentiveUnit](#ununifi.ecosystemincentive.IncentiveUnit)
    - [IncentiveUnitIdsByAddr](#ununifi.ecosystemincentive.IncentiveUnitIdsByAddr)
    - [RewardStore](#ununifi.ecosystemincentive.RewardStore)
    - [SubjectInfo](#ununifi.ecosystemincentive.SubjectInfo)
  
- [ecosystem-incentive/event.proto](#ecosystem-incentive/event.proto)
    - [EventDeletedNftIdRecordedForFrontendReward](#ununifi.ecosystemincentive.EventDeletedNftIdRecordedForFrontendReward)
    - [EventFailedParsingMemoInputs](#ununifi.ecosystemincentive.EventFailedParsingMemoInputs)
    - [EventNotRecordedNftId](#ununifi.ecosystemincentive.EventNotRecordedNftId)
    - [EventNotRegisteredIncentiveUnitId](#ununifi.ecosystemincentive.EventNotRegisteredIncentiveUnitId)
    - [EventRecordedIncentiveUnitId](#ununifi.ecosystemincentive.EventRecordedIncentiveUnitId)
    - [EventRegister](#ununifi.ecosystemincentive.EventRegister)
    - [EventUpdatedReward](#ununifi.ecosystemincentive.EventUpdatedReward)
    - [EventVersionUnmatched](#ununifi.ecosystemincentive.EventVersionUnmatched)
    - [EventWithdrawAllRewards](#ununifi.ecosystemincentive.EventWithdrawAllRewards)
    - [EventWithdrawReward](#ununifi.ecosystemincentive.EventWithdrawReward)
  
- [ecosystem-incentive/params.proto](#ecosystem-incentive/params.proto)
    - [Params](#ununifi.ecosystemincentive.Params)
    - [RewardParams](#ununifi.ecosystemincentive.RewardParams)
    - [RewardRate](#ununifi.ecosystemincentive.RewardRate)
  
    - [RewardType](#ununifi.ecosystemincentive.RewardType)
  
- [ecosystem-incentive/genesis.proto](#ecosystem-incentive/genesis.proto)
    - [GenesisState](#ununifi.ecosystemincentive.GenesisState)
  
- [ecosystem-incentive/memo.proto](#ecosystem-incentive/memo.proto)
    - [MemoInputs](#ununifi.ecosystemincentive.MemoInputs)
  
- [ecosystem-incentive/query.proto](#ecosystem-incentive/query.proto)
    - [QueryAllRewardsRequest](#ununifi.ecosystemincentive.QueryAllRewardsRequest)
    - [QueryAllRewardsResponse](#ununifi.ecosystemincentive.QueryAllRewardsResponse)
    - [QueryIncentiveUnitIdsByAddrRequest](#ununifi.ecosystemincentive.QueryIncentiveUnitIdsByAddrRequest)
    - [QueryIncentiveUnitIdsByAddrResponse](#ununifi.ecosystemincentive.QueryIncentiveUnitIdsByAddrResponse)
    - [QueryIncentiveUnitRequest](#ununifi.ecosystemincentive.QueryIncentiveUnitRequest)
    - [QueryIncentiveUnitResponse](#ununifi.ecosystemincentive.QueryIncentiveUnitResponse)
    - [QueryParamsRequest](#ununifi.ecosystemincentive.QueryParamsRequest)
    - [QueryParamsResponse](#ununifi.ecosystemincentive.QueryParamsResponse)
    - [QueryRecordedIncentiveUnitIdRequest](#ununifi.ecosystemincentive.QueryRecordedIncentiveUnitIdRequest)
    - [QueryRecordedIncentiveUnitIdResponse](#ununifi.ecosystemincentive.QueryRecordedIncentiveUnitIdResponse)
    - [QueryRewardRequest](#ununifi.ecosystemincentive.QueryRewardRequest)
    - [QueryRewardResponse](#ununifi.ecosystemincentive.QueryRewardResponse)
  
    - [Query](#ununifi.ecosystemincentive.Query)
  
- [ecosystem-incentive/tx.proto](#ecosystem-incentive/tx.proto)
    - [MsgRegister](#ununifi.ecosystemincentive.MsgRegister)
    - [MsgRegisterResponse](#ununifi.ecosystemincentive.MsgRegisterResponse)
    - [MsgWithdrawAllRewards](#ununifi.ecosystemincentive.MsgWithdrawAllRewards)
    - [MsgWithdrawAllRewardsResponse](#ununifi.ecosystemincentive.MsgWithdrawAllRewardsResponse)
    - [MsgWithdrawReward](#ununifi.ecosystemincentive.MsgWithdrawReward)
    - [MsgWithdrawRewardResponse](#ununifi.ecosystemincentive.MsgWithdrawRewardResponse)
  
    - [Msg](#ununifi.ecosystemincentive.Msg)
  
- [ethereum/signdoc.proto](#ethereum/signdoc.proto)
    - [SignDocForMetamask](#ununifi.ethereum.SignDocForMetamask)
  
- [incentive/incentive.proto](#incentive/incentive.proto)
    - [BaseClaim](#ununifi.incentive.BaseClaim)
    - [BaseMultiClaim](#ununifi.incentive.BaseMultiClaim)
    - [CdpMintingClaim](#ununifi.incentive.CdpMintingClaim)
    - [Multiplier](#ununifi.incentive.Multiplier)
    - [Params](#ununifi.incentive.Params)
    - [RewardIndex](#ununifi.incentive.RewardIndex)
    - [RewardPeriod](#ununifi.incentive.RewardPeriod)
  
- [incentive/genesis.proto](#incentive/genesis.proto)
    - [GenesisAccumulationTime](#ununifi.incentive.GenesisAccumulationTime)
    - [GenesisDenoms](#ununifi.incentive.GenesisDenoms)
    - [GenesisState](#ununifi.incentive.GenesisState)
  
- [incentive/query.proto](#incentive/query.proto)
    - [QueryParamsRequest](#ununifi.incentive.QueryParamsRequest)
    - [QueryParamsResponse](#ununifi.incentive.QueryParamsResponse)
  
    - [Query](#ununifi.incentive.Query)
  
- [incentive/tx.proto](#incentive/tx.proto)
    - [MsgClaimCdpMintingReward](#ununifi.incentive.MsgClaimCdpMintingReward)
    - [MsgClaimCdpMintingRewardResponse](#ununifi.incentive.MsgClaimCdpMintingRewardResponse)
  
    - [Msg](#ununifi.incentive.Msg)
  
- [nftmarket/nftmarket.proto](#nftmarket/nftmarket.proto)
    - [BidId](#ununifi.nftmarket.BidId)
    - [EventBorrow](#ununifi.nftmarket.EventBorrow)
    - [EventCancelBid](#ununifi.nftmarket.EventCancelBid)
    - [EventCancelListNfting](#ununifi.nftmarket.EventCancelListNfting)
    - [EventEndListNfting](#ununifi.nftmarket.EventEndListNfting)
    - [EventExpandListingPeriod](#ununifi.nftmarket.EventExpandListingPeriod)
    - [EventLiquidate](#ununifi.nftmarket.EventLiquidate)
    - [EventListNft](#ununifi.nftmarket.EventListNft)
    - [EventPayFullBid](#ununifi.nftmarket.EventPayFullBid)
    - [EventPlaceBid](#ununifi.nftmarket.EventPlaceBid)
    - [EventRepay](#ununifi.nftmarket.EventRepay)
    - [EventSellingDecision](#ununifi.nftmarket.EventSellingDecision)
    - [Liquidation](#ununifi.nftmarket.Liquidation)
    - [Liquidations](#ununifi.nftmarket.Liquidations)
    - [ListedClass](#ununifi.nftmarket.ListedClass)
    - [ListedNft](#ununifi.nftmarket.ListedNft)
    - [Loan](#ununifi.nftmarket.Loan)
    - [NftBid](#ununifi.nftmarket.NftBid)
    - [NftIdentifier](#ununifi.nftmarket.NftIdentifier)
    - [NftListing](#ununifi.nftmarket.NftListing)
    - [Params](#ununifi.nftmarket.Params)
    - [PaymentStatus](#ununifi.nftmarket.PaymentStatus)
    - [borrowing](#ununifi.nftmarket.borrowing)
  
    - [ListingState](#ununifi.nftmarket.ListingState)
    - [ListingType](#ununifi.nftmarket.ListingType)
  
- [nftmarket/genesis.proto](#nftmarket/genesis.proto)
    - [GenesisState](#ununifi.nftmarket.GenesisState)
  
- [nftmarket/query.proto](#nftmarket/query.proto)
    - [QueryBidderBidsRequest](#ununifi.nftmarket.QueryBidderBidsRequest)
    - [QueryBidderBidsResponse](#ununifi.nftmarket.QueryBidderBidsResponse)
    - [QueryCDPsListRequest](#ununifi.nftmarket.QueryCDPsListRequest)
    - [QueryCDPsListResponse](#ununifi.nftmarket.QueryCDPsListResponse)
    - [QueryLiquidationRequest](#ununifi.nftmarket.QueryLiquidationRequest)
    - [QueryLiquidationResponse](#ununifi.nftmarket.QueryLiquidationResponse)
    - [QueryListedClassRequest](#ununifi.nftmarket.QueryListedClassRequest)
    - [QueryListedClassResponse](#ununifi.nftmarket.QueryListedClassResponse)
    - [QueryListedClassesRequest](#ununifi.nftmarket.QueryListedClassesRequest)
    - [QueryListedClassesResponse](#ununifi.nftmarket.QueryListedClassesResponse)
    - [QueryListedNftsRequest](#ununifi.nftmarket.QueryListedNftsRequest)
    - [QueryListedNftsResponse](#ununifi.nftmarket.QueryListedNftsResponse)
    - [QueryLoanRequest](#ununifi.nftmarket.QueryLoanRequest)
    - [QueryLoanResponse](#ununifi.nftmarket.QueryLoanResponse)
    - [QueryLoansRequest](#ununifi.nftmarket.QueryLoansRequest)
    - [QueryLoansResponse](#ununifi.nftmarket.QueryLoansResponse)
    - [QueryNftBidsRequest](#ununifi.nftmarket.QueryNftBidsRequest)
    - [QueryNftBidsResponse](#ununifi.nftmarket.QueryNftBidsResponse)
    - [QueryNftListingRequest](#ununifi.nftmarket.QueryNftListingRequest)
    - [QueryNftListingResponse](#ununifi.nftmarket.QueryNftListingResponse)
    - [QueryParamsRequest](#ununifi.nftmarket.QueryParamsRequest)
    - [QueryParamsResponse](#ununifi.nftmarket.QueryParamsResponse)
    - [QueryPaymentStatusRequest](#ununifi.nftmarket.QueryPaymentStatusRequest)
    - [QueryPaymentStatusResponse](#ununifi.nftmarket.QueryPaymentStatusResponse)
    - [QueryRewardsRequest](#ununifi.nftmarket.QueryRewardsRequest)
    - [QueryRewardsResponse](#ununifi.nftmarket.QueryRewardsResponse)
  
    - [Query](#ununifi.nftmarket.Query)
  
- [nftmarket/tx.proto](#nftmarket/tx.proto)
    - [MsgBorrow](#ununifi.nftmarket.MsgBorrow)
    - [MsgBorrowResponse](#ununifi.nftmarket.MsgBorrowResponse)
    - [MsgCancelBid](#ununifi.nftmarket.MsgCancelBid)
    - [MsgCancelBidResponse](#ununifi.nftmarket.MsgCancelBidResponse)
    - [MsgCancelNftListing](#ununifi.nftmarket.MsgCancelNftListing)
    - [MsgCancelNftListingResponse](#ununifi.nftmarket.MsgCancelNftListingResponse)
    - [MsgEndNftListing](#ununifi.nftmarket.MsgEndNftListing)
    - [MsgEndNftListingResponse](#ununifi.nftmarket.MsgEndNftListingResponse)
    - [MsgListNft](#ununifi.nftmarket.MsgListNft)
    - [MsgListNftResponse](#ununifi.nftmarket.MsgListNftResponse)
    - [MsgMintNft](#ununifi.nftmarket.MsgMintNft)
    - [MsgMintNftResponse](#ununifi.nftmarket.MsgMintNftResponse)
    - [MsgPayFullBid](#ununifi.nftmarket.MsgPayFullBid)
    - [MsgPayFullBidResponse](#ununifi.nftmarket.MsgPayFullBidResponse)
    - [MsgPlaceBid](#ununifi.nftmarket.MsgPlaceBid)
    - [MsgPlaceBidResponse](#ununifi.nftmarket.MsgPlaceBidResponse)
    - [MsgRepay](#ununifi.nftmarket.MsgRepay)
    - [MsgRepayResponse](#ununifi.nftmarket.MsgRepayResponse)
    - [MsgSellingDecision](#ununifi.nftmarket.MsgSellingDecision)
    - [MsgSellingDecisionResponse](#ununifi.nftmarket.MsgSellingDecisionResponse)
  
    - [Msg](#ununifi.nftmarket.Msg)
  
- [nftmarketv1/nftmarket.proto](#nftmarketv1/nftmarket.proto)
    - [EventBorrow](#ununifi.nftmarketv1.EventBorrow)
    - [EventCancelBid](#ununifi.nftmarketv1.EventCancelBid)
    - [EventCancelListNfting](#ununifi.nftmarketv1.EventCancelListNfting)
    - [EventEndListNfting](#ununifi.nftmarketv1.EventEndListNfting)
    - [EventExpandListingPeriod](#ununifi.nftmarketv1.EventExpandListingPeriod)
    - [EventLiquidate](#ununifi.nftmarketv1.EventLiquidate)
    - [EventListNft](#ununifi.nftmarketv1.EventListNft)
    - [EventPayFullBid](#ununifi.nftmarketv1.EventPayFullBid)
    - [EventPlaceBid](#ununifi.nftmarketv1.EventPlaceBid)
    - [EventRepay](#ununifi.nftmarketv1.EventRepay)
    - [EventSellingDecision](#ununifi.nftmarketv1.EventSellingDecision)
    - [ListedClass](#ununifi.nftmarketv1.ListedClass)
    - [ListedNft](#ununifi.nftmarketv1.ListedNft)
    - [Loan](#ununifi.nftmarketv1.Loan)
    - [NftBid](#ununifi.nftmarketv1.NftBid)
    - [NftIdentifier](#ununifi.nftmarketv1.NftIdentifier)
    - [NftListing](#ununifi.nftmarketv1.NftListing)
    - [Params](#ununifi.nftmarketv1.Params)
    - [PaymentStatus](#ununifi.nftmarketv1.PaymentStatus)
  
    - [ListingState](#ununifi.nftmarketv1.ListingState)
    - [ListingType](#ununifi.nftmarketv1.ListingType)
  
- [nftmarketv1/genesis.proto](#nftmarketv1/genesis.proto)
    - [GenesisState](#ununifi.nftmarketv1.GenesisState)
  
- [nftmarketv1/query.proto](#nftmarketv1/query.proto)
    - [QueryBidderBidsRequest](#ununifi.nftmarketv1.QueryBidderBidsRequest)
    - [QueryBidderBidsResponse](#ununifi.nftmarketv1.QueryBidderBidsResponse)
    - [QueryCDPsListRequest](#ununifi.nftmarketv1.QueryCDPsListRequest)
    - [QueryCDPsListResponse](#ununifi.nftmarketv1.QueryCDPsListResponse)
    - [QueryListedClassRequest](#ununifi.nftmarketv1.QueryListedClassRequest)
    - [QueryListedClassResponse](#ununifi.nftmarketv1.QueryListedClassResponse)
    - [QueryListedClassesRequest](#ununifi.nftmarketv1.QueryListedClassesRequest)
    - [QueryListedClassesResponse](#ununifi.nftmarketv1.QueryListedClassesResponse)
    - [QueryListedNftsRequest](#ununifi.nftmarketv1.QueryListedNftsRequest)
    - [QueryListedNftsResponse](#ununifi.nftmarketv1.QueryListedNftsResponse)
    - [QueryLoanRequest](#ununifi.nftmarketv1.QueryLoanRequest)
    - [QueryLoanResponse](#ununifi.nftmarketv1.QueryLoanResponse)
    - [QueryLoansRequest](#ununifi.nftmarketv1.QueryLoansRequest)
    - [QueryLoansResponse](#ununifi.nftmarketv1.QueryLoansResponse)
    - [QueryNftBidsRequest](#ununifi.nftmarketv1.QueryNftBidsRequest)
    - [QueryNftBidsResponse](#ununifi.nftmarketv1.QueryNftBidsResponse)
    - [QueryNftListingRequest](#ununifi.nftmarketv1.QueryNftListingRequest)
    - [QueryNftListingResponse](#ununifi.nftmarketv1.QueryNftListingResponse)
    - [QueryParamsRequest](#ununifi.nftmarketv1.QueryParamsRequest)
    - [QueryParamsResponse](#ununifi.nftmarketv1.QueryParamsResponse)
    - [QueryPaymentStatusRequest](#ununifi.nftmarketv1.QueryPaymentStatusRequest)
    - [QueryPaymentStatusResponse](#ununifi.nftmarketv1.QueryPaymentStatusResponse)
    - [QueryRewardsRequest](#ununifi.nftmarketv1.QueryRewardsRequest)
    - [QueryRewardsResponse](#ununifi.nftmarketv1.QueryRewardsResponse)
  
    - [Query](#ununifi.nftmarketv1.Query)
  
- [nftmarketv1/tx.proto](#nftmarketv1/tx.proto)
    - [MsgBorrow](#ununifi.nftmarketv1.MsgBorrow)
    - [MsgBorrowResponse](#ununifi.nftmarketv1.MsgBorrowResponse)
    - [MsgBurnStableCoin](#ununifi.nftmarketv1.MsgBurnStableCoin)
    - [MsgBurnStableCoinResponse](#ununifi.nftmarketv1.MsgBurnStableCoinResponse)
    - [MsgCancelBid](#ununifi.nftmarketv1.MsgCancelBid)
    - [MsgCancelBidResponse](#ununifi.nftmarketv1.MsgCancelBidResponse)
    - [MsgCancelNftListing](#ununifi.nftmarketv1.MsgCancelNftListing)
    - [MsgCancelNftListingResponse](#ununifi.nftmarketv1.MsgCancelNftListingResponse)
    - [MsgEndNftListing](#ununifi.nftmarketv1.MsgEndNftListing)
    - [MsgEndNftListingResponse](#ununifi.nftmarketv1.MsgEndNftListingResponse)
    - [MsgExpandListingPeriod](#ununifi.nftmarketv1.MsgExpandListingPeriod)
    - [MsgExpandListingPeriodResponse](#ununifi.nftmarketv1.MsgExpandListingPeriodResponse)
    - [MsgLiquidate](#ununifi.nftmarketv1.MsgLiquidate)
    - [MsgLiquidateResponse](#ununifi.nftmarketv1.MsgLiquidateResponse)
    - [MsgListNft](#ununifi.nftmarketv1.MsgListNft)
    - [MsgListNftResponse](#ununifi.nftmarketv1.MsgListNftResponse)
    - [MsgMintNft](#ununifi.nftmarketv1.MsgMintNft)
    - [MsgMintNftResponse](#ununifi.nftmarketv1.MsgMintNftResponse)
    - [MsgMintStableCoin](#ununifi.nftmarketv1.MsgMintStableCoin)
    - [MsgMintStableCoinResponse](#ununifi.nftmarketv1.MsgMintStableCoinResponse)
    - [MsgPayFullBid](#ununifi.nftmarketv1.MsgPayFullBid)
    - [MsgPayFullBidResponse](#ununifi.nftmarketv1.MsgPayFullBidResponse)
    - [MsgPlaceBid](#ununifi.nftmarketv1.MsgPlaceBid)
    - [MsgPlaceBidResponse](#ununifi.nftmarketv1.MsgPlaceBidResponse)
    - [MsgRepay](#ununifi.nftmarketv1.MsgRepay)
    - [MsgRepayResponse](#ununifi.nftmarketv1.MsgRepayResponse)
    - [MsgSellingDecision](#ununifi.nftmarketv1.MsgSellingDecision)
    - [MsgSellingDecisionResponse](#ununifi.nftmarketv1.MsgSellingDecisionResponse)
  
    - [Msg](#ununifi.nftmarketv1.Msg)
  
- [nftmint/nftmint.proto](#nftmint/nftmint.proto)
    - [ClassAttributes](#ununifi.nftmint.ClassAttributes)
    - [ClassNameIdList](#ununifi.nftmint.ClassNameIdList)
    - [OwningClassIdList](#ununifi.nftmint.OwningClassIdList)
    - [Params](#ununifi.nftmint.Params)
  
    - [MintingPermission](#ununifi.nftmint.MintingPermission)
  
- [nftmint/event.proto](#nftmint/event.proto)
    - [EventBurnNFT](#ununifi.nftmint.EventBurnNFT)
    - [EventCreateClass](#ununifi.nftmint.EventCreateClass)
    - [EventMintNFT](#ununifi.nftmint.EventMintNFT)
    - [EventSendClassOwnership](#ununifi.nftmint.EventSendClassOwnership)
    - [EventUpdateBaseTokenUri](#ununifi.nftmint.EventUpdateBaseTokenUri)
    - [EventUpdateTokenSupplyCap](#ununifi.nftmint.EventUpdateTokenSupplyCap)
  
- [nftmint/genesis.proto](#nftmint/genesis.proto)
    - [GenesisState](#ununifi.nftmint.GenesisState)
  
- [nftmint/query.proto](#nftmint/query.proto)
    - [QueryClassAttributesRequest](#ununifi.nftmint.QueryClassAttributesRequest)
    - [QueryClassAttributesResponse](#ununifi.nftmint.QueryClassAttributesResponse)
    - [QueryClassIdsByNameRequest](#ununifi.nftmint.QueryClassIdsByNameRequest)
    - [QueryClassIdsByNameResponse](#ununifi.nftmint.QueryClassIdsByNameResponse)
    - [QueryClassIdsByOwnerRequest](#ununifi.nftmint.QueryClassIdsByOwnerRequest)
    - [QueryClassIdsByOwnerResponse](#ununifi.nftmint.QueryClassIdsByOwnerResponse)
    - [QueryNFTMinterRequest](#ununifi.nftmint.QueryNFTMinterRequest)
    - [QueryNFTMinterResponse](#ununifi.nftmint.QueryNFTMinterResponse)
    - [QueryParamsRequest](#ununifi.nftmint.QueryParamsRequest)
    - [QueryParamsResponse](#ununifi.nftmint.QueryParamsResponse)
  
    - [Query](#ununifi.nftmint.Query)
  
- [nftmint/tx.proto](#nftmint/tx.proto)
    - [MsgBurnNFT](#ununifi.nftmint.MsgBurnNFT)
    - [MsgBurnNFTResponse](#ununifi.nftmint.MsgBurnNFTResponse)
    - [MsgCreateClass](#ununifi.nftmint.MsgCreateClass)
    - [MsgCreateClassResponse](#ununifi.nftmint.MsgCreateClassResponse)
    - [MsgMintNFT](#ununifi.nftmint.MsgMintNFT)
    - [MsgMintNFTResponse](#ununifi.nftmint.MsgMintNFTResponse)
    - [MsgSendClassOwnership](#ununifi.nftmint.MsgSendClassOwnership)
    - [MsgSendClassOwnershipResponse](#ununifi.nftmint.MsgSendClassOwnershipResponse)
    - [MsgUpdateBaseTokenUri](#ununifi.nftmint.MsgUpdateBaseTokenUri)
    - [MsgUpdateBaseTokenUriResponse](#ununifi.nftmint.MsgUpdateBaseTokenUriResponse)
    - [MsgUpdateTokenSupplyCap](#ununifi.nftmint.MsgUpdateTokenSupplyCap)
    - [MsgUpdateTokenSupplyCapResponse](#ununifi.nftmint.MsgUpdateTokenSupplyCapResponse)
  
    - [Msg](#ununifi.nftmint.Msg)
  
- [pricefeed/pricefeed.proto](#pricefeed/pricefeed.proto)
    - [CurrentPrice](#ununifi.pricefeed.CurrentPrice)
    - [Market](#ununifi.pricefeed.Market)
    - [Params](#ununifi.pricefeed.Params)
    - [PostedPrice](#ununifi.pricefeed.PostedPrice)
  
- [pricefeed/genesis.proto](#pricefeed/genesis.proto)
    - [GenesisState](#ununifi.pricefeed.GenesisState)
  
- [pricefeed/query.proto](#pricefeed/query.proto)
    - [QueryAllMarketRequest](#ununifi.pricefeed.QueryAllMarketRequest)
    - [QueryAllMarketResponse](#ununifi.pricefeed.QueryAllMarketResponse)
    - [QueryAllOracleRequest](#ununifi.pricefeed.QueryAllOracleRequest)
    - [QueryAllOracleResponse](#ununifi.pricefeed.QueryAllOracleResponse)
    - [QueryAllPriceRequest](#ununifi.pricefeed.QueryAllPriceRequest)
    - [QueryAllPriceResponse](#ununifi.pricefeed.QueryAllPriceResponse)
    - [QueryAllRawPriceRequest](#ununifi.pricefeed.QueryAllRawPriceRequest)
    - [QueryAllRawPriceResponse](#ununifi.pricefeed.QueryAllRawPriceResponse)
    - [QueryGetPriceRequest](#ununifi.pricefeed.QueryGetPriceRequest)
    - [QueryGetPriceResponse](#ununifi.pricefeed.QueryGetPriceResponse)
    - [QueryParamsRequest](#ununifi.pricefeed.QueryParamsRequest)
    - [QueryParamsResponse](#ununifi.pricefeed.QueryParamsResponse)
  
    - [Query](#ununifi.pricefeed.Query)
  
- [pricefeed/tx.proto](#pricefeed/tx.proto)
    - [MsgPostPrice](#ununifi.pricefeed.MsgPostPrice)
    - [MsgPostPriceResponse](#ununifi.pricefeed.MsgPostPriceResponse)
  
    - [Msg](#ununifi.pricefeed.Msg)
  
- [ununifidist/ununifidist.proto](#ununifidist/ununifidist.proto)
    - [Params](#ununifi.ununifidist.Params)
    - [Period](#ununifi.ununifidist.Period)
  
- [ununifidist/genesis.proto](#ununifidist/genesis.proto)
    - [GenesisState](#ununifi.ununifidist.GenesisState)
  
- [ununifidist/query.proto](#ununifidist/query.proto)
    - [QueryGetBalancesRequest](#ununifi.ununifidist.QueryGetBalancesRequest)
    - [QueryGetBalancesResponse](#ununifi.ununifidist.QueryGetBalancesResponse)
    - [QueryParamsRequest](#ununifi.ununifidist.QueryParamsRequest)
    - [QueryParamsResponse](#ununifi.ununifidist.QueryParamsResponse)
  
    - [Query](#ununifi.ununifidist.Query)
  
- [Scalar Value Types](#scalar-value-types)



<a name="auction/auction.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## auction/auction.proto



<a name="ununifi.auction.BaseAuction"></a>

### BaseAuction



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [uint64](#uint64) |  |  |
| `initiator` | [string](#string) |  |  |
| `lot` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `bidder` | [string](#string) |  |  |
| `bid` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `has_received_bids` | [bool](#bool) |  |  |
| `end_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `max_end_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="ununifi.auction.CollateralAuction"></a>

### CollateralAuction



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_auction` | [BaseAuction](#ununifi.auction.BaseAuction) |  |  |
| `corresponding_debt` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `max_bid` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `lot_returns` | [WeightedAddress](#ununifi.auction.WeightedAddress) | repeated |  |






<a name="ununifi.auction.DebtAuction"></a>

### DebtAuction



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_auction` | [BaseAuction](#ununifi.auction.BaseAuction) |  |  |
| `corresponding_debt` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="ununifi.auction.Params"></a>

### Params



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `max_auction_duration` | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |
| `bid_duration` | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |
| `increment_surplus` | [string](#string) |  |  |
| `increment_debt` | [string](#string) |  |  |
| `increment_collateral` | [string](#string) |  |  |






<a name="ununifi.auction.SurplusAuction"></a>

### SurplusAuction



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_auction` | [BaseAuction](#ununifi.auction.BaseAuction) |  |  |






<a name="ununifi.auction.WeightedAddress"></a>

### WeightedAddress



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |
| `weight` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="auction/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## auction/genesis.proto



<a name="ununifi.auction.GenesisState"></a>

### GenesisState
GenesisState defines the auction module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `next_auction_id` | [uint64](#uint64) |  |  |
| `params` | [Params](#ununifi.auction.Params) |  |  |
| `auctions` | [google.protobuf.Any](#google.protobuf.Any) | repeated | this line is used by starport scaffolding # genesis/proto/state |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="auction/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## auction/query.proto



<a name="ununifi.auction.QueryAllAuctionRequest"></a>

### QueryAllAuctionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="ununifi.auction.QueryAllAuctionResponse"></a>

### QueryAllAuctionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `auctions` | [google.protobuf.Any](#google.protobuf.Any) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="ununifi.auction.QueryGetAuctionRequest"></a>

### QueryGetAuctionRequest
this line is used by starport scaffolding # 3


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [uint64](#uint64) |  |  |






<a name="ununifi.auction.QueryGetAuctionResponse"></a>

### QueryGetAuctionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `auction` | [google.protobuf.Any](#google.protobuf.Any) |  |  |






<a name="ununifi.auction.QueryParamsRequest"></a>

### QueryParamsRequest







<a name="ununifi.auction.QueryParamsResponse"></a>

### QueryParamsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#ununifi.auction.Params) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="ununifi.auction.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#ununifi.auction.QueryParamsRequest) | [QueryParamsResponse](#ununifi.auction.QueryParamsResponse) |  | GET|/ununifi/auction/params|
| `Auction` | [QueryGetAuctionRequest](#ununifi.auction.QueryGetAuctionRequest) | [QueryGetAuctionResponse](#ununifi.auction.QueryGetAuctionResponse) | this line is used by starport scaffolding # 2 | GET|/ununifi/auction/auctions/{id}|
| `AuctionAll` | [QueryAllAuctionRequest](#ununifi.auction.QueryAllAuctionRequest) | [QueryAllAuctionResponse](#ununifi.auction.QueryAllAuctionResponse) |  | GET|/ununifi/auction/auctions|

 <!-- end services -->



<a name="auction/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## auction/tx.proto



<a name="ununifi.auction.MsgPlaceBid"></a>

### MsgPlaceBid



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `auction_id` | [uint64](#uint64) |  |  |
| `bidder` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="ununifi.auction.MsgPlaceBidResponse"></a>

### MsgPlaceBidResponse






 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="ununifi.auction.Msg"></a>

### Msg


| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `PlaceBid` | [MsgPlaceBid](#ununifi.auction.MsgPlaceBid) | [MsgPlaceBidResponse](#ununifi.auction.MsgPlaceBidResponse) |  | |

 <!-- end services -->



<a name="cdp/cdp.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cdp/cdp.proto



<a name="ununifi.cdp.AugmentedCdp"></a>

### AugmentedCdp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `cdp` | [Cdp](#ununifi.cdp.Cdp) |  |  |
| `collateral_value` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `collateralization_ratio` | [string](#string) |  |  |






<a name="ununifi.cdp.Cdp"></a>

### Cdp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [uint64](#uint64) |  |  |
| `owner` | [string](#string) |  |  |
| `type` | [string](#string) |  |  |
| `collateral` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `principal` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `accumulated_fees` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `fees_updated` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `interest_factor` | [string](#string) |  |  |






<a name="ununifi.cdp.CollateralParam"></a>

### CollateralParam



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `type` | [string](#string) |  |  |
| `liquidation_ratio` | [string](#string) |  |  |
| `debt_limit` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `stability_fee` | [string](#string) |  |  |
| `auction_size` | [string](#string) |  |  |
| `liquidation_penalty` | [string](#string) |  |  |
| `prefix` | [uint32](#uint32) |  |  |
| `spot_market_id` | [string](#string) |  |  |
| `liquidation_market_id` | [string](#string) |  |  |
| `keeper_reward_percentage` | [string](#string) |  |  |
| `check_collateralization_index_count` | [string](#string) |  |  |
| `conversion_factor` | [string](#string) |  |  |






<a name="ununifi.cdp.DebtParam"></a>

### DebtParam



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `reference_asset` | [string](#string) |  |  |
| `conversion_factor` | [string](#string) |  |  |
| `debt_floor` | [string](#string) |  |  |
| `global_debt_limit` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `debt_denom` | [string](#string) |  |  |
| `surplus_auction_threshold` | [string](#string) |  |  |
| `surplus_auction_lot` | [string](#string) |  |  |
| `debt_auction_threshold` | [string](#string) |  |  |
| `debt_auction_lot` | [string](#string) |  |  |
| `circuit_breaker` | [bool](#bool) |  |  |






<a name="ununifi.cdp.Deposit"></a>

### Deposit



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `cdp_id` | [uint64](#uint64) |  |  |
| `depositor` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="ununifi.cdp.Params"></a>

### Params



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_params` | [CollateralParam](#ununifi.cdp.CollateralParam) | repeated |  |
| `debt_params` | [DebtParam](#ununifi.cdp.DebtParam) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cdp/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cdp/genesis.proto



<a name="ununifi.cdp.GenesisAccumulationTime"></a>

### GenesisAccumulationTime



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_type` | [string](#string) |  |  |
| `previous_accumulation_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `interest_factor` | [string](#string) |  |  |






<a name="ununifi.cdp.GenesisState"></a>

### GenesisState
GenesisState defines the cdp module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#ununifi.cdp.Params) |  |  |
| `cdps` | [Cdp](#ununifi.cdp.Cdp) | repeated |  |
| `deposits` | [Deposit](#ununifi.cdp.Deposit) | repeated |  |
| `starting_cdp_id` | [uint64](#uint64) |  |  |
| `gov_denom` | [string](#string) |  |  |
| `previous_accumulation_times` | [GenesisAccumulationTime](#ununifi.cdp.GenesisAccumulationTime) | repeated |  |
| `total_principals` | [GenesisTotalPrincipal](#ununifi.cdp.GenesisTotalPrincipal) | repeated | this line is used by starport scaffolding # genesis/proto/state |






<a name="ununifi.cdp.GenesisTotalPrincipal"></a>

### GenesisTotalPrincipal



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_type` | [string](#string) |  |  |
| `total_principal` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cdp/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cdp/query.proto



<a name="ununifi.cdp.QueryAllAccountRequest"></a>

### QueryAllAccountRequest







<a name="ununifi.cdp.QueryAllAccountResponse"></a>

### QueryAllAccountResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `accounts` | [google.protobuf.Any](#google.protobuf.Any) | repeated |  |






<a name="ununifi.cdp.QueryAllCdpRequest"></a>

### QueryAllCdpRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="ununifi.cdp.QueryAllCdpResponse"></a>

### QueryAllCdpResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `cdp` | [AugmentedCdp](#ununifi.cdp.AugmentedCdp) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="ununifi.cdp.QueryAllDepositRequest"></a>

### QueryAllDepositRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `owner` | [string](#string) |  |  |
| `collateral_type` | [string](#string) |  |  |






<a name="ununifi.cdp.QueryAllDepositResponse"></a>

### QueryAllDepositResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `deposits` | [Deposit](#ununifi.cdp.Deposit) | repeated |  |






<a name="ununifi.cdp.QueryGetCdpRequest"></a>

### QueryGetCdpRequest
this line is used by starport scaffolding # 3


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `owner` | [string](#string) |  |  |
| `collateral_type` | [string](#string) |  |  |






<a name="ununifi.cdp.QueryGetCdpResponse"></a>

### QueryGetCdpResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `cdp` | [AugmentedCdp](#ununifi.cdp.AugmentedCdp) |  |  |






<a name="ununifi.cdp.QueryParamsRequest"></a>

### QueryParamsRequest







<a name="ununifi.cdp.QueryParamsResponse"></a>

### QueryParamsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#ununifi.cdp.Params) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="ununifi.cdp.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#ununifi.cdp.QueryParamsRequest) | [QueryParamsResponse](#ununifi.cdp.QueryParamsResponse) |  | GET|/ununifi/cdp/params|
| `Cdp` | [QueryGetCdpRequest](#ununifi.cdp.QueryGetCdpRequest) | [QueryGetCdpResponse](#ununifi.cdp.QueryGetCdpResponse) | this line is used by starport scaffolding # 2 | GET|/ununifi/cdp/cdps/owners/{owner}/collateral-types/{collateral_type}/cdp|
| `CdpAll` | [QueryAllCdpRequest](#ununifi.cdp.QueryAllCdpRequest) | [QueryAllCdpResponse](#ununifi.cdp.QueryAllCdpResponse) |  | GET|/ununifi/cdp/cdps|
| `AccountAll` | [QueryAllAccountRequest](#ununifi.cdp.QueryAllAccountRequest) | [QueryAllAccountResponse](#ununifi.cdp.QueryAllAccountResponse) |  | GET|/ununifi/cdp/accounts|
| `DepositAll` | [QueryAllDepositRequest](#ununifi.cdp.QueryAllDepositRequest) | [QueryAllDepositResponse](#ununifi.cdp.QueryAllDepositResponse) |  | GET|/ununifi/cdp/deposits/owners/{owner}/collateral-types/{collateral_type}|

 <!-- end services -->



<a name="cdp/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cdp/tx.proto



<a name="ununifi.cdp.MsgCreateCdp"></a>

### MsgCreateCdp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `collateral` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `principal` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `collateral_type` | [string](#string) |  |  |






<a name="ununifi.cdp.MsgCreateCdpResponse"></a>

### MsgCreateCdpResponse







<a name="ununifi.cdp.MsgDeposit"></a>

### MsgDeposit



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `depositor` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |
| `collateral` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `collateral_type` | [string](#string) |  |  |






<a name="ununifi.cdp.MsgDepositResponse"></a>

### MsgDepositResponse







<a name="ununifi.cdp.MsgDrawDebt"></a>

### MsgDrawDebt



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `collateral_type` | [string](#string) |  |  |
| `principal` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="ununifi.cdp.MsgDrawDebtResponse"></a>

### MsgDrawDebtResponse







<a name="ununifi.cdp.MsgLiquidate"></a>

### MsgLiquidate



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `keeper` | [string](#string) |  |  |
| `borrower` | [string](#string) |  |  |
| `collateral_type` | [string](#string) |  |  |






<a name="ununifi.cdp.MsgLiquidateResponse"></a>

### MsgLiquidateResponse







<a name="ununifi.cdp.MsgRepayDebt"></a>

### MsgRepayDebt



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `collateral_type` | [string](#string) |  |  |
| `payment` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="ununifi.cdp.MsgRepayDebtResponse"></a>

### MsgRepayDebtResponse







<a name="ununifi.cdp.MsgWithdraw"></a>

### MsgWithdraw



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `depositor` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |
| `collateral` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `collateral_type` | [string](#string) |  |  |






<a name="ununifi.cdp.MsgWithdrawResponse"></a>

### MsgWithdrawResponse






 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="ununifi.cdp.Msg"></a>

### Msg


| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `CreateCdp` | [MsgCreateCdp](#ununifi.cdp.MsgCreateCdp) | [MsgCreateCdpResponse](#ununifi.cdp.MsgCreateCdpResponse) |  | |
| `Deposit` | [MsgDeposit](#ununifi.cdp.MsgDeposit) | [MsgDepositResponse](#ununifi.cdp.MsgDepositResponse) |  | |
| `Withdraw` | [MsgWithdraw](#ununifi.cdp.MsgWithdraw) | [MsgWithdrawResponse](#ununifi.cdp.MsgWithdrawResponse) |  | |
| `DrawDebt` | [MsgDrawDebt](#ununifi.cdp.MsgDrawDebt) | [MsgDrawDebtResponse](#ununifi.cdp.MsgDrawDebtResponse) |  | |
| `RepayDebt` | [MsgRepayDebt](#ununifi.cdp.MsgRepayDebt) | [MsgRepayDebtResponse](#ununifi.cdp.MsgRepayDebtResponse) |  | |
| `Liquidate` | [MsgLiquidate](#ununifi.cdp.MsgLiquidate) | [MsgLiquidateResponse](#ununifi.cdp.MsgLiquidateResponse) |  | |

 <!-- end services -->



<a name="derivatives/derivatives.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## derivatives/derivatives.proto



<a name="ununifi.derivatives.EventPriceIsNotFeeded"></a>

### EventPriceIsNotFeeded



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `asset` | [string](#string) |  |  |






<a name="ununifi.derivatives.Market"></a>

### Market



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_denom` | [string](#string) |  |  |
| `quote_denom` | [string](#string) |  |  |






<a name="ununifi.derivatives.PoolMarketCap"></a>

### PoolMarketCap



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `quote_ticker` | [string](#string) |  |  |
| `total` | [string](#string) |  |  |
| `breakdown` | [PoolMarketCap.Breakdown](#ununifi.derivatives.PoolMarketCap.Breakdown) | repeated |  |






<a name="ununifi.derivatives.PoolMarketCap.Breakdown"></a>

### PoolMarketCap.Breakdown



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `amount` | [string](#string) |  |  |
| `price` | [string](#string) |  |  |






<a name="ununifi.derivatives.PoolParams"></a>

### PoolParams



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `quote_ticker` | [string](#string) |  |  |
| `base_lpt_mint_fee` | [string](#string) |  |  |
| `base_lpt_redeem_fee` | [string](#string) |  |  |
| `borrowing_fee_rate_per_hour` | [string](#string) |  |  |
| `report_liquidation_reward_rate` | [string](#string) |  |  |
| `report_levy_period_reward_rate` | [string](#string) |  |  |
| `accepted_assets` | [PoolParams.Asset](#ununifi.derivatives.PoolParams.Asset) | repeated |  |






<a name="ununifi.derivatives.PoolParams.Asset"></a>

### PoolParams.Asset



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `target_weight` | [string](#string) |  |  |






<a name="ununifi.derivatives.Position"></a>

### Position



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [string](#string) |  |  |
| `market` | [Market](#ununifi.derivatives.Market) |  |  |
| `address` | [string](#string) |  |  |
| `opened_at` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `opened_height` | [uint64](#uint64) |  |  |
| `opened_base_rate` | [string](#string) |  |  |
| `opened_quote_rate` | [string](#string) |  |  |
| `remaining_margin` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `last_levied_at` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `position_instance` | [google.protobuf.Any](#google.protobuf.Any) |  |  |






<a name="ununifi.derivatives.QueriedPosition"></a>

### QueriedPosition



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `position` | [Position](#ununifi.derivatives.Position) |  |  |
| `valuation_profit` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `remaining_margin_value` | [string](#string) |  |  |
| `effective_margin` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `margin_maintenance_rate` | [string](#string) |  |  |





 <!-- end messages -->


<a name="ununifi.derivatives.PositionType"></a>

### PositionType


| Name | Number | Description |
| ---- | ------ | ----------- |
| POSITION_UNKNOWN | 0 |  |
| LONG | 1 |  |
| SHORT | 2 |  |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="derivatives/perpetual_futures.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## derivatives/perpetual_futures.proto



<a name="ununifi.derivatives.EventPerpetualFuturesPositionClosed"></a>

### EventPerpetualFuturesPositionClosed



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `position_id` | [string](#string) |  |  |
| `fee_amount` | [string](#string) |  |  |
| `trade_amount` | [string](#string) |  |  |






<a name="ununifi.derivatives.EventPerpetualFuturesPositionLevied"></a>

### EventPerpetualFuturesPositionLevied



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `reward_recipient` | [string](#string) |  |  |
| `position_id` | [string](#string) |  |  |
| `remaining_margin` | [string](#string) |  |  |
| `reward_amount` | [string](#string) |  |  |






<a name="ununifi.derivatives.EventPerpetualFuturesPositionLiquidated"></a>

### EventPerpetualFuturesPositionLiquidated



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `reward_recipient` | [string](#string) |  |  |
| `position_id` | [string](#string) |  |  |
| `remaining_margin` | [string](#string) |  |  |
| `reward_amount` | [string](#string) |  |  |






<a name="ununifi.derivatives.EventPerpetualFuturesPositionOpened"></a>

### EventPerpetualFuturesPositionOpened



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `position_id` | [string](#string) |  |  |






<a name="ununifi.derivatives.PerpetualFuturesNetPositionOfMarket"></a>

### PerpetualFuturesNetPositionOfMarket



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `market` | [Market](#ununifi.derivatives.Market) |  |  |
| `position_size` | [string](#string) |  |  |






<a name="ununifi.derivatives.PerpetualFuturesParams"></a>

### PerpetualFuturesParams



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `commission_rate` | [string](#string) |  |  |
| `margin_maintenance_rate` | [string](#string) |  |  |
| `imaginary_funding_rate_proportional_coefficient` | [string](#string) |  |  |
| `markets` | [Market](#ununifi.derivatives.Market) | repeated |  |






<a name="ununifi.derivatives.PerpetualFuturesPosition"></a>

### PerpetualFuturesPosition



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [string](#string) |  |  |
| `market` | [Market](#ununifi.derivatives.Market) |  |  |
| `address` | [string](#string) |  |  |
| `opened_at` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `opened_height` | [uint64](#uint64) |  |  |
| `opened_base_rate` | [string](#string) |  |  |
| `opened_quote_rate` | [string](#string) |  |  |
| `remaining_margin` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `last_levied_at` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `position_instance` | [PerpetualFuturesPositionInstance](#ununifi.derivatives.PerpetualFuturesPositionInstance) |  |  |






<a name="ununifi.derivatives.PerpetualFuturesPositionInstance"></a>

### PerpetualFuturesPositionInstance



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `position_type` | [PositionType](#ununifi.derivatives.PositionType) |  |  |
| `size` | [string](#string) |  |  |
| `leverage` | [uint32](#uint32) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="derivatives/perpetual_options.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## derivatives/perpetual_options.proto



<a name="ununifi.derivatives.PerpetualOptionsParams"></a>

### PerpetualOptionsParams



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `premium_commission_rate` | [string](#string) |  |  |
| `strike_commission_rate` | [string](#string) |  |  |
| `margin_maintenance_rate` | [string](#string) |  |  |
| `imaginary_funding_rate_proportional_coefficient` | [string](#string) |  |  |
| `markets` | [Market](#ununifi.derivatives.Market) | repeated |  |






<a name="ununifi.derivatives.PerpetualOptionsPositionInstance"></a>

### PerpetualOptionsPositionInstance



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `option_type` | [OptionType](#ununifi.derivatives.OptionType) |  |  |
| `position_type` | [PositionType](#ununifi.derivatives.PositionType) |  |  |
| `strike_price` | [string](#string) |  |  |
| `premium` | [string](#string) |  |  |





 <!-- end messages -->


<a name="ununifi.derivatives.OptionType"></a>

### OptionType


| Name | Number | Description |
| ---- | ------ | ----------- |
| OPTION_UNKNOWN | 0 |  |
| CALL | 1 |  |
| PUT | 2 |  |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="derivatives/params.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## derivatives/params.proto



<a name="ununifi.derivatives.Params"></a>

### Params



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pool_params` | [PoolParams](#ununifi.derivatives.PoolParams) |  |  |
| `perpetual_futures` | [PerpetualFuturesParams](#ununifi.derivatives.PerpetualFuturesParams) |  |  |
| `perpetual_options` | [PerpetualOptionsParams](#ununifi.derivatives.PerpetualOptionsParams) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="derivatives/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## derivatives/genesis.proto



<a name="ununifi.derivatives.GenesisState"></a>

### GenesisState
GenesisState defines the derivatives module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#ununifi.derivatives.Params) |  |  |
| `positions` | [Position](#ununifi.derivatives.Position) | repeated |  |
| `pool_market_cap` | [PoolMarketCap](#ununifi.derivatives.PoolMarketCap) |  |  |
| `perpetual_futures_net_position_of_market` | [PerpetualFuturesNetPositionOfMarket](#ununifi.derivatives.PerpetualFuturesNetPositionOfMarket) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="derivatives/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## derivatives/query.proto



<a name="ununifi.derivatives.QueryAddressPositionsRequest"></a>

### QueryAddressPositionsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="ununifi.derivatives.QueryAddressPositionsResponse"></a>

### QueryAddressPositionsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `positions` | [QueriedPosition](#ununifi.derivatives.QueriedPosition) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="ununifi.derivatives.QueryAllPositionsRequest"></a>

### QueryAllPositionsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="ununifi.derivatives.QueryAllPositionsResponse"></a>

### QueryAllPositionsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `positions` | [QueriedPosition](#ununifi.derivatives.QueriedPosition) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="ununifi.derivatives.QueryDLPTokenRateRequest"></a>

### QueryDLPTokenRateRequest







<a name="ununifi.derivatives.QueryDLPTokenRateResponse"></a>

### QueryDLPTokenRateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `symbol` | [string](#string) |  |  |
| `rates` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="ununifi.derivatives.QueryEstimateDLPTokenAmountRequest"></a>

### QueryEstimateDLPTokenAmountRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `mint_denom` | [string](#string) |  |  |
| `amount` | [string](#string) |  |  |






<a name="ununifi.derivatives.QueryEstimateDLPTokenAmountResponse"></a>

### QueryEstimateDLPTokenAmountResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="ununifi.derivatives.QueryEstimateRedeemAmountRequest"></a>

### QueryEstimateRedeemAmountRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `redeem_denom` | [string](#string) |  |  |
| `lpt_amount` | [string](#string) |  |  |






<a name="ununifi.derivatives.QueryEstimateRedeemAmountResponse"></a>

### QueryEstimateRedeemAmountResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `fee` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="ununifi.derivatives.QueryLiquidityProviderTokenNominalAPYRequest"></a>

### QueryLiquidityProviderTokenNominalAPYRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `before_height` | [int64](#int64) |  |  |
| `after_height` | [int64](#int64) |  |  |






<a name="ununifi.derivatives.QueryLiquidityProviderTokenNominalAPYResponse"></a>

### QueryLiquidityProviderTokenNominalAPYResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `apy` | [string](#string) |  |  |






<a name="ununifi.derivatives.QueryLiquidityProviderTokenRealAPYRequest"></a>

### QueryLiquidityProviderTokenRealAPYRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `before_height` | [int64](#int64) |  |  |
| `after_height` | [int64](#int64) |  |  |






<a name="ununifi.derivatives.QueryLiquidityProviderTokenRealAPYResponse"></a>

### QueryLiquidityProviderTokenRealAPYResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `apy` | [string](#string) |  |  |






<a name="ununifi.derivatives.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is request type for the Query/Params RPC method.






<a name="ununifi.derivatives.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is response type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#ununifi.derivatives.Params) |  | params holds all the parameters of this module. |






<a name="ununifi.derivatives.QueryPerpetualFuturesMarketRequest"></a>

### QueryPerpetualFuturesMarketRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_denom` | [string](#string) |  |  |
| `quote_denom` | [string](#string) |  |  |






<a name="ununifi.derivatives.QueryPerpetualFuturesMarketResponse"></a>

### QueryPerpetualFuturesMarketResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `price` | [string](#string) |  |  |
| `metrics_quote_ticker` | [string](#string) |  |  |
| `volume_24hours` | [string](#string) |  |  |
| `fees_24hours` | [string](#string) |  |  |
| `long_positions` | [string](#string) |  |  |
| `short_positions` | [string](#string) |  |  |






<a name="ununifi.derivatives.QueryPerpetualFuturesPositionSizeRequest"></a>

### QueryPerpetualFuturesPositionSizeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `position_type` | [PositionType](#ununifi.derivatives.PositionType) |  |  |
| `address` | [string](#string) |  |  |






<a name="ununifi.derivatives.QueryPerpetualFuturesPositionSizeResponse"></a>

### QueryPerpetualFuturesPositionSizeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `total_position_size_usd` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="ununifi.derivatives.QueryPerpetualFuturesRequest"></a>

### QueryPerpetualFuturesRequest







<a name="ununifi.derivatives.QueryPerpetualFuturesResponse"></a>

### QueryPerpetualFuturesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `metrics_quote_ticker` | [string](#string) |  |  |
| `volume_24hours` | [string](#string) |  |  |
| `fees_24hours` | [string](#string) |  |  |
| `long_positions` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `short_positions` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="ununifi.derivatives.QueryPerpetualOptionsMarketRequest"></a>

### QueryPerpetualOptionsMarketRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_denom` | [string](#string) |  |  |
| `quote_denom` | [string](#string) |  |  |






<a name="ununifi.derivatives.QueryPerpetualOptionsMarketResponse"></a>

### QueryPerpetualOptionsMarketResponse







<a name="ununifi.derivatives.QueryPerpetualOptionsRequest"></a>

### QueryPerpetualOptionsRequest







<a name="ununifi.derivatives.QueryPerpetualOptionsResponse"></a>

### QueryPerpetualOptionsResponse







<a name="ununifi.derivatives.QueryPoolRequest"></a>

### QueryPoolRequest
this line is used by starport scaffolding # 3






<a name="ununifi.derivatives.QueryPoolResponse"></a>

### QueryPoolResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `metrics_quote_ticker` | [string](#string) |  |  |
| `pool_market_cap` | [PoolMarketCap](#ununifi.derivatives.PoolMarketCap) |  |  |
| `volume_24hours` | [string](#string) |  |  |
| `fees_24hours` | [string](#string) |  |  |






<a name="ununifi.derivatives.QueryPositionRequest"></a>

### QueryPositionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `position_id` | [string](#string) |  |  |






<a name="ununifi.derivatives.QueryPositionResponse"></a>

### QueryPositionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `position` | [Position](#ununifi.derivatives.Position) |  |  |
| `valuation_profit` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `margin_maintenance_rate` | [string](#string) |  |  |
| `effective_margin` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="ununifi.derivatives.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#ununifi.derivatives.QueryParamsRequest) | [QueryParamsResponse](#ununifi.derivatives.QueryParamsResponse) | Parameters queries the parameters of the module. | GET|/ununifi/derivatives/params|
| `Pool` | [QueryPoolRequest](#ununifi.derivatives.QueryPoolRequest) | [QueryPoolResponse](#ununifi.derivatives.QueryPoolResponse) | this line is used by starport scaffolding # 2 | GET|/ununifi/derivatives/pool|
| `LiquidityProviderTokenRealAPY` | [QueryLiquidityProviderTokenRealAPYRequest](#ununifi.derivatives.QueryLiquidityProviderTokenRealAPYRequest) | [QueryLiquidityProviderTokenRealAPYResponse](#ununifi.derivatives.QueryLiquidityProviderTokenRealAPYResponse) |  | GET|/ununifi/derivatives/liquidity-providers/real-apy|
| `LiquidityProviderTokenNominalAPY` | [QueryLiquidityProviderTokenNominalAPYRequest](#ununifi.derivatives.QueryLiquidityProviderTokenNominalAPYRequest) | [QueryLiquidityProviderTokenNominalAPYResponse](#ununifi.derivatives.QueryLiquidityProviderTokenNominalAPYResponse) |  | GET|/ununifi/derivatives/liquidity-providers/nominal-apy|
| `PerpetualFutures` | [QueryPerpetualFuturesRequest](#ununifi.derivatives.QueryPerpetualFuturesRequest) | [QueryPerpetualFuturesResponse](#ununifi.derivatives.QueryPerpetualFuturesResponse) |  | GET|/ununifi/derivatives/perpetual-futures|
| `PerpetualFuturesMarket` | [QueryPerpetualFuturesMarketRequest](#ununifi.derivatives.QueryPerpetualFuturesMarketRequest) | [QueryPerpetualFuturesMarketResponse](#ununifi.derivatives.QueryPerpetualFuturesMarketResponse) |  | GET|/ununifi/derivatives/perpetual-futures/{base_denom}/{quote_denom}|
| `PerpetualOptions` | [QueryPerpetualOptionsRequest](#ununifi.derivatives.QueryPerpetualOptionsRequest) | [QueryPerpetualOptionsResponse](#ununifi.derivatives.QueryPerpetualOptionsResponse) |  | GET|/ununifi/derivatives/perpetual-options|
| `PerpetualOptionsMarket` | [QueryPerpetualOptionsMarketRequest](#ununifi.derivatives.QueryPerpetualOptionsMarketRequest) | [QueryPerpetualOptionsMarketResponse](#ununifi.derivatives.QueryPerpetualOptionsMarketResponse) |  | GET|/ununifi/derivatives/perpetual-options/{base_denom}/{quote_denom}|
| `AllPositions` | [QueryAllPositionsRequest](#ununifi.derivatives.QueryAllPositionsRequest) | [QueryAllPositionsResponse](#ununifi.derivatives.QueryAllPositionsResponse) |  | GET|/ununifi/derivatives/positions|
| `Position` | [QueryPositionRequest](#ununifi.derivatives.QueryPositionRequest) | [QueryPositionResponse](#ununifi.derivatives.QueryPositionResponse) |  | GET|/ununifi/derivatives/positions/{position_id}|
| `PerpetualFuturesPositionSize` | [QueryPerpetualFuturesPositionSizeRequest](#ununifi.derivatives.QueryPerpetualFuturesPositionSizeRequest) | [QueryPerpetualFuturesPositionSizeResponse](#ununifi.derivatives.QueryPerpetualFuturesPositionSizeResponse) |  | GET|/ununifi/derivatives/positions/perpetual-futures/{position_type}/{address}|
| `AddressPositions` | [QueryAddressPositionsRequest](#ununifi.derivatives.QueryAddressPositionsRequest) | [QueryAddressPositionsResponse](#ununifi.derivatives.QueryAddressPositionsResponse) |  | GET|/ununifi/derivatives/{address}/positions|
| `DLPTokenRates` | [QueryDLPTokenRateRequest](#ununifi.derivatives.QueryDLPTokenRateRequest) | [QueryDLPTokenRateResponse](#ununifi.derivatives.QueryDLPTokenRateResponse) |  | GET|/ununifi/derivatives/pools/dlp|
| `EstimateDLPTokenAmount` | [QueryEstimateDLPTokenAmountRequest](#ununifi.derivatives.QueryEstimateDLPTokenAmountRequest) | [QueryEstimateDLPTokenAmountResponse](#ununifi.derivatives.QueryEstimateDLPTokenAmountResponse) |  | GET|/ununifi/derivatives/estimate-dlp-token-amount/{mint_denom}/{amount}|
| `EstimateRedeemAmount` | [QueryEstimateRedeemAmountRequest](#ununifi.derivatives.QueryEstimateRedeemAmountRequest) | [QueryEstimateRedeemAmountResponse](#ununifi.derivatives.QueryEstimateRedeemAmountResponse) |  | GET|/ununifi/derivatives/estimate-redeem-amount/{redeem_denom}/{lpt_amount}|

 <!-- end services -->



<a name="derivatives/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## derivatives/tx.proto



<a name="ununifi.derivatives.MsgBurnLiquidityProviderToken"></a>

### MsgBurnLiquidityProviderToken



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `amount` | [string](#string) |  |  |
| `redeem_denom` | [string](#string) |  |  |






<a name="ununifi.derivatives.MsgBurnLiquidityProviderTokenResponse"></a>

### MsgBurnLiquidityProviderTokenResponse







<a name="ununifi.derivatives.MsgClosePosition"></a>

### MsgClosePosition



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `position_id` | [string](#string) |  |  |






<a name="ununifi.derivatives.MsgClosePositionResponse"></a>

### MsgClosePositionResponse







<a name="ununifi.derivatives.MsgMintLiquidityProviderToken"></a>

### MsgMintLiquidityProviderToken



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="ununifi.derivatives.MsgMintLiquidityProviderTokenResponse"></a>

### MsgMintLiquidityProviderTokenResponse







<a name="ununifi.derivatives.MsgOpenPosition"></a>

### MsgOpenPosition



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `margin` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `market` | [Market](#ununifi.derivatives.Market) |  |  |
| `position_instance` | [google.protobuf.Any](#google.protobuf.Any) |  |  |






<a name="ununifi.derivatives.MsgOpenPositionResponse"></a>

### MsgOpenPositionResponse







<a name="ununifi.derivatives.MsgReportLevyPeriod"></a>

### MsgReportLevyPeriod



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `position_id` | [string](#string) |  |  |
| `reward_recipient` | [string](#string) |  |  |






<a name="ununifi.derivatives.MsgReportLevyPeriodResponse"></a>

### MsgReportLevyPeriodResponse







<a name="ununifi.derivatives.MsgReportLiquidation"></a>

### MsgReportLiquidation



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `position_id` | [string](#string) |  |  |
| `reward_recipient` | [string](#string) |  |  |






<a name="ununifi.derivatives.MsgReportLiquidationResponse"></a>

### MsgReportLiquidationResponse






 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="ununifi.derivatives.Msg"></a>

### Msg


| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `MintLiquidityProviderToken` | [MsgMintLiquidityProviderToken](#ununifi.derivatives.MsgMintLiquidityProviderToken) | [MsgMintLiquidityProviderTokenResponse](#ununifi.derivatives.MsgMintLiquidityProviderTokenResponse) |  | |
| `BurnLiquidityProviderToken` | [MsgBurnLiquidityProviderToken](#ununifi.derivatives.MsgBurnLiquidityProviderToken) | [MsgBurnLiquidityProviderTokenResponse](#ununifi.derivatives.MsgBurnLiquidityProviderTokenResponse) |  | |
| `OpenPosition` | [MsgOpenPosition](#ununifi.derivatives.MsgOpenPosition) | [MsgOpenPositionResponse](#ununifi.derivatives.MsgOpenPositionResponse) |  | |
| `ClosePosition` | [MsgClosePosition](#ununifi.derivatives.MsgClosePosition) | [MsgClosePositionResponse](#ununifi.derivatives.MsgClosePositionResponse) |  | |
| `ReportLiquidation` | [MsgReportLiquidation](#ununifi.derivatives.MsgReportLiquidation) | [MsgReportLiquidationResponse](#ununifi.derivatives.MsgReportLiquidationResponse) |  | |
| `ReportLevyPeriod` | [MsgReportLevyPeriod](#ununifi.derivatives.MsgReportLevyPeriod) | [MsgReportLevyPeriodResponse](#ununifi.derivatives.MsgReportLevyPeriodResponse) |  | |

 <!-- end services -->



<a name="ecosystem-incentive/ecosystem_incentive.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ecosystem-incentive/ecosystem_incentive.proto



<a name="ununifi.ecosystemincentive.IncentiveUnit"></a>

### IncentiveUnit



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [string](#string) |  |  |
| `subject_info_lists` | [SubjectInfo](#ununifi.ecosystemincentive.SubjectInfo) | repeated |  |






<a name="ununifi.ecosystemincentive.IncentiveUnitIdsByAddr"></a>

### IncentiveUnitIdsByAddr



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |
| `incentive_unit_ids` | [string](#string) | repeated |  |






<a name="ununifi.ecosystemincentive.RewardStore"></a>

### RewardStore



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `subject_addr` | [string](#string) |  |  |
| `rewards` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="ununifi.ecosystemincentive.SubjectInfo"></a>

### SubjectInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `subject_addr` | [string](#string) |  |  |
| `weight` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="ecosystem-incentive/event.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ecosystem-incentive/event.proto



<a name="ununifi.ecosystemincentive.EventDeletedNftIdRecordedForFrontendReward"></a>

### EventDeletedNftIdRecordedForFrontendReward



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `incentive_unit_id` | [string](#string) |  |  |
| `class_id` | [string](#string) |  |  |
| `nft_id` | [string](#string) |  |  |






<a name="ununifi.ecosystemincentive.EventFailedParsingMemoInputs"></a>

### EventFailedParsingMemoInputs



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `class_id` | [string](#string) |  |  |
| `nft_id` | [string](#string) |  |  |
| `memo` | [string](#string) |  |  |






<a name="ununifi.ecosystemincentive.EventNotRecordedNftId"></a>

### EventNotRecordedNftId



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `class_id` | [string](#string) |  |  |
| `nft_id` | [string](#string) |  |  |






<a name="ununifi.ecosystemincentive.EventNotRegisteredIncentiveUnitId"></a>

### EventNotRegisteredIncentiveUnitId



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `incentive_unit_id` | [string](#string) |  |  |
| `class_id` | [string](#string) |  |  |
| `nft_id` | [string](#string) |  |  |






<a name="ununifi.ecosystemincentive.EventRecordedIncentiveUnitId"></a>

### EventRecordedIncentiveUnitId



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `incentive_unit_id` | [string](#string) |  |  |
| `class_id` | [string](#string) |  |  |
| `nft_id` | [string](#string) |  |  |






<a name="ununifi.ecosystemincentive.EventRegister"></a>

### EventRegister



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `incentive_unit_id` | [string](#string) |  |  |
| `subject_info_lists` | [SubjectInfo](#ununifi.ecosystemincentive.SubjectInfo) | repeated |  |






<a name="ununifi.ecosystemincentive.EventUpdatedReward"></a>

### EventUpdatedReward



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `incentive_unit_id` | [string](#string) |  |  |
| `earned_reward` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="ununifi.ecosystemincentive.EventVersionUnmatched"></a>

### EventVersionUnmatched



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `unmatched_version` | [string](#string) |  |  |
| `class_id` | [string](#string) |  |  |
| `nft_id` | [string](#string) |  |  |






<a name="ununifi.ecosystemincentive.EventWithdrawAllRewards"></a>

### EventWithdrawAllRewards



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `all_withdrawn_rewards` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="ununifi.ecosystemincentive.EventWithdrawReward"></a>

### EventWithdrawReward



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `withdrawn_reward` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="ecosystem-incentive/params.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ecosystem-incentive/params.proto



<a name="ununifi.ecosystemincentive.Params"></a>

### Params
Params defines the parameters for the module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `reward_params` | [RewardParams](#ununifi.ecosystemincentive.RewardParams) | repeated |  |
| `max_incentive_unit_id_len` | [uint64](#uint64) |  |  |
| `max_subject_info_num_in_unit` | [uint64](#uint64) |  |  |






<a name="ununifi.ecosystemincentive.RewardParams"></a>

### RewardParams
RewardParams defines which module have which rewards type and rate 
to maintain the correctness of the fee rate in a module
e.g. if nftmarket module have "Frontend" and "Collection" incentive, 
the combined those rates for the incentive cannot be exceed 1


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `module_name` | [string](#string) |  |  |
| `reward_rate` | [RewardRate](#ununifi.ecosystemincentive.RewardRate) | repeated |  |






<a name="ununifi.ecosystemincentive.RewardRate"></a>

### RewardRate
RewardRate defines the ratio to take reward for a specific reward_type.
The total sum of reward_rate in a module cannot be exceed 1


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `reward_type` | [RewardType](#ununifi.ecosystemincentive.RewardType) |  |  |
| `rate` | [string](#string) |  |  |





 <!-- end messages -->


<a name="ununifi.ecosystemincentive.RewardType"></a>

### RewardType
At first, we go with this one type.
NFTMARKET_FRONTEND type reward will be disributed for the creators of frontend of UnUniFi's services.

| Name | Number | Description |
| ---- | ------ | ----------- |
| NFTMARKET_FRONTEND | 0 |  |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="ecosystem-incentive/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ecosystem-incentive/genesis.proto



<a name="ununifi.ecosystemincentive.GenesisState"></a>

### GenesisState
GenesisState defines the ecosystemincentive module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#ununifi.ecosystemincentive.Params) |  |  |
| `incentive_units` | [IncentiveUnit](#ununifi.ecosystemincentive.IncentiveUnit) | repeated |  |
| `reward_stores` | [RewardStore](#ununifi.ecosystemincentive.RewardStore) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="ecosystem-incentive/memo.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ecosystem-incentive/memo.proto



<a name="ununifi.ecosystemincentive.MemoInputs"></a>

### MemoInputs



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `version` | [string](#string) |  |  |
| `incentive_unit_id` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="ecosystem-incentive/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ecosystem-incentive/query.proto



<a name="ununifi.ecosystemincentive.QueryAllRewardsRequest"></a>

### QueryAllRewardsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `subject_addr` | [string](#string) |  |  |






<a name="ununifi.ecosystemincentive.QueryAllRewardsResponse"></a>

### QueryAllRewardsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `rewards` | [RewardStore](#ununifi.ecosystemincentive.RewardStore) |  |  |






<a name="ununifi.ecosystemincentive.QueryIncentiveUnitIdsByAddrRequest"></a>

### QueryIncentiveUnitIdsByAddrRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |






<a name="ununifi.ecosystemincentive.QueryIncentiveUnitIdsByAddrResponse"></a>

### QueryIncentiveUnitIdsByAddrResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `incentive_unit_ids_by_addr` | [IncentiveUnitIdsByAddr](#ununifi.ecosystemincentive.IncentiveUnitIdsByAddr) |  |  |






<a name="ununifi.ecosystemincentive.QueryIncentiveUnitRequest"></a>

### QueryIncentiveUnitRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `incentive_unit_id` | [string](#string) |  |  |






<a name="ununifi.ecosystemincentive.QueryIncentiveUnitResponse"></a>

### QueryIncentiveUnitResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `incentive_unit` | [IncentiveUnit](#ununifi.ecosystemincentive.IncentiveUnit) |  |  |






<a name="ununifi.ecosystemincentive.QueryParamsRequest"></a>

### QueryParamsRequest







<a name="ununifi.ecosystemincentive.QueryParamsResponse"></a>

### QueryParamsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#ununifi.ecosystemincentive.Params) |  | params holds all the parameters of this module. |






<a name="ununifi.ecosystemincentive.QueryRecordedIncentiveUnitIdRequest"></a>

### QueryRecordedIncentiveUnitIdRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `class_id` | [string](#string) |  |  |
| `nft_id` | [string](#string) |  |  |






<a name="ununifi.ecosystemincentive.QueryRecordedIncentiveUnitIdResponse"></a>

### QueryRecordedIncentiveUnitIdResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `incentive_unit_id` | [string](#string) |  |  |






<a name="ununifi.ecosystemincentive.QueryRewardRequest"></a>

### QueryRewardRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `subject_addr` | [string](#string) |  |  |
| `denom` | [string](#string) |  |  |






<a name="ununifi.ecosystemincentive.QueryRewardResponse"></a>

### QueryRewardResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `reward` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="ununifi.ecosystemincentive.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#ununifi.ecosystemincentive.QueryParamsRequest) | [QueryParamsResponse](#ununifi.ecosystemincentive.QueryParamsResponse) | Parameters queries the parameters of the module. | GET|/ununifi/ecosystem_incentive/params|
| `IncentiveUnit` | [QueryIncentiveUnitRequest](#ununifi.ecosystemincentive.QueryIncentiveUnitRequest) | [QueryIncentiveUnitResponse](#ununifi.ecosystemincentive.QueryIncentiveUnitResponse) |  | GET|/ununifi/ecosystem_incentive/incentive_unit/{incentive_unit_id}|
| `AllRewards` | [QueryAllRewardsRequest](#ununifi.ecosystemincentive.QueryAllRewardsRequest) | [QueryAllRewardsResponse](#ununifi.ecosystemincentive.QueryAllRewardsResponse) |  | GET|/ununifi/ecosystem_incentive/all_rewards/{subject_addr}|
| `Reward` | [QueryRewardRequest](#ununifi.ecosystemincentive.QueryRewardRequest) | [QueryRewardResponse](#ununifi.ecosystemincentive.QueryRewardResponse) |  | GET|/ununifi/ecosystem_incentive/reward/{subject_addr}/{denom}|
| `RecordedIncentiveUnitId` | [QueryRecordedIncentiveUnitIdRequest](#ununifi.ecosystemincentive.QueryRecordedIncentiveUnitIdRequest) | [QueryRecordedIncentiveUnitIdResponse](#ununifi.ecosystemincentive.QueryRecordedIncentiveUnitIdResponse) |  | GET|/ununifi/ecosystem_incentive/recorded_incentive_unit_id/{class_id}/{nft_id}|
| `IncentiveUnitIdsByAddr` | [QueryIncentiveUnitIdsByAddrRequest](#ununifi.ecosystemincentive.QueryIncentiveUnitIdsByAddrRequest) | [QueryIncentiveUnitIdsByAddrResponse](#ununifi.ecosystemincentive.QueryIncentiveUnitIdsByAddrResponse) |  | GET|/ununifi/ecosystem_incentive/incentive_unit_ids_by_addr/{address}|

 <!-- end services -->



<a name="ecosystem-incentive/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ecosystem-incentive/tx.proto



<a name="ununifi.ecosystemincentive.MsgRegister"></a>

### MsgRegister



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `incentive_unit_id` | [string](#string) |  |  |
| `subject_addrs` | [string](#string) | repeated |  |
| `weights` | [string](#string) | repeated |  |






<a name="ununifi.ecosystemincentive.MsgRegisterResponse"></a>

### MsgRegisterResponse







<a name="ununifi.ecosystemincentive.MsgWithdrawAllRewards"></a>

### MsgWithdrawAllRewards



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |






<a name="ununifi.ecosystemincentive.MsgWithdrawAllRewardsResponse"></a>

### MsgWithdrawAllRewardsResponse







<a name="ununifi.ecosystemincentive.MsgWithdrawReward"></a>

### MsgWithdrawReward



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `denom` | [string](#string) |  |  |






<a name="ununifi.ecosystemincentive.MsgWithdrawRewardResponse"></a>

### MsgWithdrawRewardResponse






 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="ununifi.ecosystemincentive.Msg"></a>

### Msg
Msg defines the Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Register` | [MsgRegister](#ununifi.ecosystemincentive.MsgRegister) | [MsgRegisterResponse](#ununifi.ecosystemincentive.MsgRegisterResponse) |  | |
| `WithdrawAllRewards` | [MsgWithdrawAllRewards](#ununifi.ecosystemincentive.MsgWithdrawAllRewards) | [MsgWithdrawAllRewardsResponse](#ununifi.ecosystemincentive.MsgWithdrawAllRewardsResponse) |  | |
| `WithdrawReward` | [MsgWithdrawReward](#ununifi.ecosystemincentive.MsgWithdrawReward) | [MsgWithdrawRewardResponse](#ununifi.ecosystemincentive.MsgWithdrawRewardResponse) |  | |

 <!-- end services -->



<a name="ethereum/signdoc.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ethereum/signdoc.proto



<a name="ununifi.ethereum.SignDocForMetamask"></a>

### SignDocForMetamask



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `body` | [cosmos.tx.v1beta1.TxBody](#cosmos.tx.v1beta1.TxBody) |  |  |
| `auth_info` | [cosmos.tx.v1beta1.AuthInfo](#cosmos.tx.v1beta1.AuthInfo) |  |  |
| `chain_id` | [string](#string) |  |  |
| `account_number` | [uint64](#uint64) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="incentive/incentive.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## incentive/incentive.proto



<a name="ununifi.incentive.BaseClaim"></a>

### BaseClaim



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `owner` | [string](#string) |  |  |
| `reward` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="ununifi.incentive.BaseMultiClaim"></a>

### BaseMultiClaim



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `owner` | [string](#string) |  |  |
| `reward` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="ununifi.incentive.CdpMintingClaim"></a>

### CdpMintingClaim



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_claim` | [BaseClaim](#ununifi.incentive.BaseClaim) |  |  |
| `reward_indexes` | [RewardIndex](#ununifi.incentive.RewardIndex) | repeated |  |






<a name="ununifi.incentive.Multiplier"></a>

### Multiplier



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `name` | [string](#string) |  |  |
| `months_lockup` | [int64](#int64) |  |  |
| `factor` | [string](#string) |  |  |






<a name="ununifi.incentive.Params"></a>

### Params



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `cdp_minting_reward_periods` | [RewardPeriod](#ununifi.incentive.RewardPeriod) | repeated |  |
| `claim_multipliers` | [Multiplier](#ununifi.incentive.Multiplier) | repeated |  |
| `claim_end` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="ununifi.incentive.RewardIndex"></a>

### RewardIndex



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_type` | [string](#string) |  |  |
| `reward_factor` | [string](#string) |  |  |






<a name="ununifi.incentive.RewardPeriod"></a>

### RewardPeriod



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `active` | [bool](#bool) |  |  |
| `collateral_type` | [string](#string) |  |  |
| `start` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `end` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `rewards_per_second` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="incentive/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## incentive/genesis.proto



<a name="ununifi.incentive.GenesisAccumulationTime"></a>

### GenesisAccumulationTime



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_type` | [string](#string) |  |  |
| `previous_accumulation_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="ununifi.incentive.GenesisDenoms"></a>

### GenesisDenoms



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `principal_denom` | [string](#string) |  |  |
| `cdp_minting_reward_denom` | [string](#string) |  |  |






<a name="ununifi.incentive.GenesisState"></a>

### GenesisState
GenesisState defines the incentive module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#ununifi.incentive.Params) |  |  |
| `cdp_accumulation_times` | [GenesisAccumulationTime](#ununifi.incentive.GenesisAccumulationTime) | repeated |  |
| `cdp_minting_claims` | [CdpMintingClaim](#ununifi.incentive.CdpMintingClaim) | repeated |  |
| `denoms` | [GenesisDenoms](#ununifi.incentive.GenesisDenoms) |  | this line is used by starport scaffolding # genesis/proto/state |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="incentive/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## incentive/query.proto



<a name="ununifi.incentive.QueryParamsRequest"></a>

### QueryParamsRequest







<a name="ununifi.incentive.QueryParamsResponse"></a>

### QueryParamsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#ununifi.incentive.Params) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="ununifi.incentive.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#ununifi.incentive.QueryParamsRequest) | [QueryParamsResponse](#ununifi.incentive.QueryParamsResponse) | this line is used by starport scaffolding # 2 | GET|/ununifi/incentive/params|

 <!-- end services -->



<a name="incentive/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## incentive/tx.proto



<a name="ununifi.incentive.MsgClaimCdpMintingReward"></a>

### MsgClaimCdpMintingReward



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `multiplier_name` | [string](#string) |  |  |






<a name="ununifi.incentive.MsgClaimCdpMintingRewardResponse"></a>

### MsgClaimCdpMintingRewardResponse






 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="ununifi.incentive.Msg"></a>

### Msg


| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `ClaimCdpMintingReward` | [MsgClaimCdpMintingReward](#ununifi.incentive.MsgClaimCdpMintingReward) | [MsgClaimCdpMintingRewardResponse](#ununifi.incentive.MsgClaimCdpMintingRewardResponse) |  | |

 <!-- end services -->



<a name="nftmarket/nftmarket.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nftmarket/nftmarket.proto



<a name="ununifi.nftmarket.BidId"></a>

### BidId



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `nft_id` | [NftIdentifier](#ununifi.nftmarket.NftIdentifier) |  |  |
| `bidder` | [string](#string) |  |  |






<a name="ununifi.nftmarket.EventBorrow"></a>

### EventBorrow



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `borrower` | [string](#string) |  |  |
| `class_id` | [string](#string) |  |  |
| `nft_id` | [string](#string) |  |  |
| `amount` | [string](#string) |  |  |






<a name="ununifi.nftmarket.EventCancelBid"></a>

### EventCancelBid



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `bidder` | [string](#string) |  |  |
| `class_id` | [string](#string) |  |  |
| `nft_id` | [string](#string) |  |  |






<a name="ununifi.nftmarket.EventCancelListNfting"></a>

### EventCancelListNfting



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `owner` | [string](#string) |  |  |
| `class_id` | [string](#string) |  |  |
| `nft_id` | [string](#string) |  |  |






<a name="ununifi.nftmarket.EventEndListNfting"></a>

### EventEndListNfting



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `owner` | [string](#string) |  |  |
| `class_id` | [string](#string) |  |  |
| `nft_id` | [string](#string) |  |  |






<a name="ununifi.nftmarket.EventExpandListingPeriod"></a>

### EventExpandListingPeriod



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `owner` | [string](#string) |  |  |
| `class_id` | [string](#string) |  |  |
| `nft_id` | [string](#string) |  |  |






<a name="ununifi.nftmarket.EventLiquidate"></a>

### EventLiquidate



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `liquidator` | [string](#string) |  |  |
| `class_id` | [string](#string) |  |  |
| `nft_id` | [string](#string) |  |  |






<a name="ununifi.nftmarket.EventListNft"></a>

### EventListNft



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `owner` | [string](#string) |  |  |
| `class_id` | [string](#string) |  |  |
| `nft_id` | [string](#string) |  |  |






<a name="ununifi.nftmarket.EventPayFullBid"></a>

### EventPayFullBid



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `bidder` | [string](#string) |  |  |
| `class_id` | [string](#string) |  |  |
| `nft_id` | [string](#string) |  |  |






<a name="ununifi.nftmarket.EventPlaceBid"></a>

### EventPlaceBid



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `bidder` | [string](#string) |  |  |
| `class_id` | [string](#string) |  |  |
| `nft_id` | [string](#string) |  |  |
| `amount` | [string](#string) |  |  |






<a name="ununifi.nftmarket.EventRepay"></a>

### EventRepay



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `repayer` | [string](#string) |  |  |
| `class_id` | [string](#string) |  |  |
| `nft_id` | [string](#string) |  |  |
| `amount` | [string](#string) |  |  |






<a name="ununifi.nftmarket.EventSellingDecision"></a>

### EventSellingDecision



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `owner` | [string](#string) |  |  |
| `class_id` | [string](#string) |  |  |
| `nft_id` | [string](#string) |  |  |






<a name="ununifi.nftmarket.Liquidation"></a>

### Liquidation



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `liquidation_date` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="ununifi.nftmarket.Liquidations"></a>

### Liquidations



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `liquidation` | [Liquidation](#ununifi.nftmarket.Liquidation) |  |  |
| `next_liquidation` | [Liquidation](#ununifi.nftmarket.Liquidation) | repeated |  |






<a name="ununifi.nftmarket.ListedClass"></a>

### ListedClass



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `class_id` | [string](#string) |  |  |
| `nft_ids` | [string](#string) | repeated |  |






<a name="ununifi.nftmarket.ListedNft"></a>

### ListedNft



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [string](#string) |  |  |
| `uri` | [string](#string) |  |  |
| `uri_hash` | [string](#string) |  |  |






<a name="ununifi.nftmarket.Loan"></a>

### Loan



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `nft_id` | [NftIdentifier](#ununifi.nftmarket.NftIdentifier) |  |  |
| `loan` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="ununifi.nftmarket.NftBid"></a>

### NftBid



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `nft_id` | [NftIdentifier](#ununifi.nftmarket.NftIdentifier) |  |  |
| `bidder` | [string](#string) |  |  |
| `bid_amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `deposit_amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `paid_amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `bidding_period` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `deposit_lending_rate` | [string](#string) |  |  |
| `automatic_payment` | [bool](#bool) |  |  |
| `bid_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `interest_amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `borrowings` | [borrowing](#ununifi.nftmarket.borrowing) | repeated | lister borrowing bidder info |
| `id` | [BidId](#ununifi.nftmarket.BidId) |  | todo remove nft_id,bidder then use it |






<a name="ununifi.nftmarket.NftIdentifier"></a>

### NftIdentifier



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `class_id` | [string](#string) |  |  |
| `nft_id` | [string](#string) |  |  |






<a name="ununifi.nftmarket.NftListing"></a>

### NftListing



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `nft_id` | [NftIdentifier](#ununifi.nftmarket.NftIdentifier) |  |  |
| `owner` | [string](#string) |  |  |
| `listing_type` | [ListingType](#ununifi.nftmarket.ListingType) |  |  |
| `state` | [ListingState](#ununifi.nftmarket.ListingState) |  |  |
| `bid_token` | [string](#string) |  |  |
| `minimum_deposit_rate` | [string](#string) |  |  |
| `automatic_refinancing` | [bool](#bool) |  |  |
| `started_at` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `end_at` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `full_payment_end_at` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `successful_bid_end_at` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `auto_relisted_count` | [uint64](#uint64) |  |  |
| `collected_amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `minimum_bidding_period` | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |






<a name="ununifi.nftmarket.Params"></a>

### Params



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `min_staking_for_listing` | [string](#string) |  |  |
| `default_bid_active_rank` | [uint64](#uint64) |  |  |
| `bid_tokens` | [string](#string) | repeated |  |
| `auto_relisting_count_if_no_bid` | [uint64](#uint64) |  |  |
| `nft_listing_delay_seconds` | [uint64](#uint64) |  |  |
| `nft_listing_period_initial` | [uint64](#uint64) |  |  |
| `nft_listing_cancel_required_seconds` | [uint64](#uint64) |  |  |
| `nft_listing_cancel_fee_percentage` | [uint64](#uint64) |  |  |
| `nft_listing_gap_time` | [uint64](#uint64) |  |  |
| `bid_cancel_required_seconds` | [uint64](#uint64) |  |  |
| `bid_token_disburse_seconds_after_cancel` | [uint64](#uint64) |  |  |
| `nft_listing_full_payment_period` | [uint64](#uint64) |  |  |
| `nft_listing_nft_delivery_period` | [uint64](#uint64) |  |  |
| `nft_creator_share_percentage` | [uint64](#uint64) |  |  |
| `market_administrator` | [string](#string) |  |  |
| `nft_listing_commission_fee` | [uint64](#uint64) |  |  |
| `nft_listing_extend_seconds` | [uint64](#uint64) |  |  |
| `nft_listing_period_extend_fee_per_hour` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="ununifi.nftmarket.PaymentStatus"></a>

### PaymentStatus



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `nft_id` | [NftIdentifier](#ununifi.nftmarket.NftIdentifier) |  |  |
| `bidder` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `automatic_payment` | [bool](#bool) |  |  |
| `paid_amount` | [string](#string) |  |  |
| `bid_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `state` | [ListingState](#ununifi.nftmarket.ListingState) |  |  |
| `all_paid` | [bool](#bool) |  |  |






<a name="ununifi.nftmarket.borrowing"></a>

### borrowing



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `paid_interest_amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `start_at` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |





 <!-- end messages -->


<a name="ununifi.nftmarket.ListingState"></a>

### ListingState


| Name | Number | Description |
| ---- | ------ | ----------- |
| LISTING | 0 |  |
| BIDDING | 1 |  |
| SELLING_DECISION | 2 |  |
| END_LISTING | 3 |  |
| SUCCESSFUL_BID | 4 |  |



<a name="ununifi.nftmarket.ListingType"></a>

### ListingType


| Name | Number | Description |
| ---- | ------ | ----------- |
| DIRECT_ASSET_BORROW | 0 |  |
| SYNTHETIC_ASSET_CREATION | 1 |  |
| LATE_SHIPPING | 2 |  |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="nftmarket/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nftmarket/genesis.proto



<a name="ununifi.nftmarket.GenesisState"></a>

### GenesisState
GenesisState defines the nftmarket module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#ununifi.nftmarket.Params) |  |  |
| `listings` | [NftListing](#ununifi.nftmarket.NftListing) | repeated |  |
| `bids` | [NftBid](#ununifi.nftmarket.NftBid) | repeated |  |
| `cancelled_bids` | [NftBid](#ununifi.nftmarket.NftBid) | repeated |  |
| `loans` | [Loan](#ununifi.nftmarket.Loan) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="nftmarket/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nftmarket/query.proto



<a name="ununifi.nftmarket.QueryBidderBidsRequest"></a>

### QueryBidderBidsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `bidder` | [string](#string) |  |  |






<a name="ununifi.nftmarket.QueryBidderBidsResponse"></a>

### QueryBidderBidsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `bids` | [NftBid](#ununifi.nftmarket.NftBid) | repeated |  |






<a name="ununifi.nftmarket.QueryCDPsListRequest"></a>

### QueryCDPsListRequest







<a name="ununifi.nftmarket.QueryCDPsListResponse"></a>

### QueryCDPsListResponse







<a name="ununifi.nftmarket.QueryLiquidationRequest"></a>

### QueryLiquidationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `class_id` | [string](#string) |  |  |
| `nft_id` | [string](#string) |  |  |






<a name="ununifi.nftmarket.QueryLiquidationResponse"></a>

### QueryLiquidationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `liquidations` | [Liquidations](#ununifi.nftmarket.Liquidations) |  |  |






<a name="ununifi.nftmarket.QueryListedClassRequest"></a>

### QueryListedClassRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `class_id` | [string](#string) |  |  |
| `nft_limit` | [int32](#int32) |  |  |






<a name="ununifi.nftmarket.QueryListedClassResponse"></a>

### QueryListedClassResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `class_id` | [string](#string) |  |  |
| `name` | [string](#string) |  |  |
| `description` | [string](#string) |  |  |
| `symbol` | [string](#string) |  |  |
| `uri` | [string](#string) |  |  |
| `urihash` | [string](#string) |  |  |
| `nfts` | [ListedNft](#ununifi.nftmarket.ListedNft) | repeated |  |
| `nft_count` | [uint64](#uint64) |  |  |






<a name="ununifi.nftmarket.QueryListedClassesRequest"></a>

### QueryListedClassesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `nft_limit` | [int32](#int32) |  |  |






<a name="ununifi.nftmarket.QueryListedClassesResponse"></a>

### QueryListedClassesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `classes` | [QueryListedClassResponse](#ununifi.nftmarket.QueryListedClassResponse) | repeated |  |






<a name="ununifi.nftmarket.QueryListedNftsRequest"></a>

### QueryListedNftsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `owner` | [string](#string) |  |  |






<a name="ununifi.nftmarket.QueryListedNftsResponse"></a>

### QueryListedNftsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `listings` | [NftListing](#ununifi.nftmarket.NftListing) | repeated |  |






<a name="ununifi.nftmarket.QueryLoanRequest"></a>

### QueryLoanRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `class_id` | [string](#string) |  |  |
| `nft_id` | [string](#string) |  |  |






<a name="ununifi.nftmarket.QueryLoanResponse"></a>

### QueryLoanResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `loan` | [Loan](#ununifi.nftmarket.Loan) |  |  |
| `borrowing_limit` | [string](#string) |  |  |






<a name="ununifi.nftmarket.QueryLoansRequest"></a>

### QueryLoansRequest







<a name="ununifi.nftmarket.QueryLoansResponse"></a>

### QueryLoansResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `loans` | [Loan](#ununifi.nftmarket.Loan) | repeated |  |






<a name="ununifi.nftmarket.QueryNftBidsRequest"></a>

### QueryNftBidsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `class_id` | [string](#string) |  |  |
| `nft_id` | [string](#string) |  |  |






<a name="ununifi.nftmarket.QueryNftBidsResponse"></a>

### QueryNftBidsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `bids` | [NftBid](#ununifi.nftmarket.NftBid) | repeated |  |






<a name="ununifi.nftmarket.QueryNftListingRequest"></a>

### QueryNftListingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `class_id` | [string](#string) |  |  |
| `nft_id` | [string](#string) |  |  |






<a name="ununifi.nftmarket.QueryNftListingResponse"></a>

### QueryNftListingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `listing` | [NftListing](#ununifi.nftmarket.NftListing) |  |  |






<a name="ununifi.nftmarket.QueryParamsRequest"></a>

### QueryParamsRequest







<a name="ununifi.nftmarket.QueryParamsResponse"></a>

### QueryParamsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#ununifi.nftmarket.Params) |  |  |






<a name="ununifi.nftmarket.QueryPaymentStatusRequest"></a>

### QueryPaymentStatusRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `class_id` | [string](#string) |  |  |
| `nft_id` | [string](#string) |  |  |
| `bidder` | [string](#string) |  |  |






<a name="ununifi.nftmarket.QueryPaymentStatusResponse"></a>

### QueryPaymentStatusResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `paymentStatus` | [PaymentStatus](#ununifi.nftmarket.PaymentStatus) |  |  |






<a name="ununifi.nftmarket.QueryRewardsRequest"></a>

### QueryRewardsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [uint64](#uint64) |  |  |






<a name="ununifi.nftmarket.QueryRewardsResponse"></a>

### QueryRewardsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `rewards` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="ununifi.nftmarket.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#ununifi.nftmarket.QueryParamsRequest) | [QueryParamsResponse](#ununifi.nftmarket.QueryParamsResponse) |  | GET|/ununifi/nftmarket/params|
| `NftListing` | [QueryNftListingRequest](#ununifi.nftmarket.QueryNftListingRequest) | [QueryNftListingResponse](#ununifi.nftmarket.QueryNftListingResponse) |  | GET|/ununifi/nftmarket/nft_listing/{class_id}/{nft_id}|
| `ListedNfts` | [QueryListedNftsRequest](#ununifi.nftmarket.QueryListedNftsRequest) | [QueryListedNftsResponse](#ununifi.nftmarket.QueryListedNftsResponse) |  | GET|/ununifi/nftmarket/listed_nfts|
| `ListedClasses` | [QueryListedClassesRequest](#ununifi.nftmarket.QueryListedClassesRequest) | [QueryListedClassesResponse](#ununifi.nftmarket.QueryListedClassesResponse) |  | GET|/ununifi/nftmarket/listed_classes|
| `ListedClass` | [QueryListedClassRequest](#ununifi.nftmarket.QueryListedClassRequest) | [QueryListedClassResponse](#ununifi.nftmarket.QueryListedClassResponse) |  | GET|/ununifi/nftmarket/listed_class/{class_id}/{nft_limit}|
| `Loans` | [QueryLoansRequest](#ununifi.nftmarket.QueryLoansRequest) | [QueryLoansResponse](#ununifi.nftmarket.QueryLoansResponse) |  | GET|/ununifi/nftmarket/loans|
| `Loan` | [QueryLoanRequest](#ununifi.nftmarket.QueryLoanRequest) | [QueryLoanResponse](#ununifi.nftmarket.QueryLoanResponse) |  | GET|/ununifi/nftmarket/loans/{class_id}/{nft_id}|
| `CDPsList` | [QueryCDPsListRequest](#ununifi.nftmarket.QueryCDPsListRequest) | [QueryCDPsListResponse](#ununifi.nftmarket.QueryCDPsListResponse) |  | GET|/ununifi/nftmarket/cdps_list|
| `NftBids` | [QueryNftBidsRequest](#ununifi.nftmarket.QueryNftBidsRequest) | [QueryNftBidsResponse](#ununifi.nftmarket.QueryNftBidsResponse) |  | GET|/ununifi/nftmarket/nft_bids/{class_id}/{nft_id}|
| `BidderBids` | [QueryBidderBidsRequest](#ununifi.nftmarket.QueryBidderBidsRequest) | [QueryBidderBidsResponse](#ununifi.nftmarket.QueryBidderBidsResponse) |  | GET|/ununifi/nftmarket/bidder_bids/{bidder}|
| `PaymentStatus` | [QueryPaymentStatusRequest](#ununifi.nftmarket.QueryPaymentStatusRequest) | [QueryPaymentStatusResponse](#ununifi.nftmarket.QueryPaymentStatusResponse) |  | GET|/ununifi/nftmarket/payment_status/{class_id}/{nft_id}/{bidder}|
| `Rewards` | [QueryRewardsRequest](#ununifi.nftmarket.QueryRewardsRequest) | [QueryRewardsResponse](#ununifi.nftmarket.QueryRewardsResponse) |  | GET|/ununifi/nftmarket/rewards/{address}|
| `Liquidation` | [QueryLiquidationRequest](#ununifi.nftmarket.QueryLiquidationRequest) | [QueryLiquidationResponse](#ununifi.nftmarket.QueryLiquidationResponse) |  | GET|/ununifi/nftmarket/liquidation/{class_id}/{nft_id}|

 <!-- end services -->



<a name="nftmarket/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nftmarket/tx.proto



<a name="ununifi.nftmarket.MsgBorrow"></a>

### MsgBorrow



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `nft_id` | [NftIdentifier](#ununifi.nftmarket.NftIdentifier) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="ununifi.nftmarket.MsgBorrowResponse"></a>

### MsgBorrowResponse







<a name="ununifi.nftmarket.MsgCancelBid"></a>

### MsgCancelBid



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `nft_id` | [NftIdentifier](#ununifi.nftmarket.NftIdentifier) |  |  |






<a name="ununifi.nftmarket.MsgCancelBidResponse"></a>

### MsgCancelBidResponse







<a name="ununifi.nftmarket.MsgCancelNftListing"></a>

### MsgCancelNftListing



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `nft_id` | [NftIdentifier](#ununifi.nftmarket.NftIdentifier) |  |  |






<a name="ununifi.nftmarket.MsgCancelNftListingResponse"></a>

### MsgCancelNftListingResponse







<a name="ununifi.nftmarket.MsgEndNftListing"></a>

### MsgEndNftListing



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `nft_id` | [NftIdentifier](#ununifi.nftmarket.NftIdentifier) |  |  |






<a name="ununifi.nftmarket.MsgEndNftListingResponse"></a>

### MsgEndNftListingResponse







<a name="ununifi.nftmarket.MsgListNft"></a>

### MsgListNft



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `nft_id` | [NftIdentifier](#ununifi.nftmarket.NftIdentifier) |  |  |
| `listing_type` | [ListingType](#ununifi.nftmarket.ListingType) |  |  |
| `bid_token` | [string](#string) |  |  |
| `minimum_deposit_rate` | [string](#string) |  | 1% = 0.01 |
| `automatic_refinancing` | [bool](#bool) |  |  |
| `minimum_bidding_period` | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |






<a name="ununifi.nftmarket.MsgListNftResponse"></a>

### MsgListNftResponse







<a name="ununifi.nftmarket.MsgMintNft"></a>

### MsgMintNft



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `classId` | [string](#string) |  |  |
| `nftId` | [string](#string) |  |  |
| `nftUri` | [string](#string) |  |  |
| `nftUriHash` | [string](#string) |  |  |






<a name="ununifi.nftmarket.MsgMintNftResponse"></a>

### MsgMintNftResponse







<a name="ununifi.nftmarket.MsgPayFullBid"></a>

### MsgPayFullBid



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `nft_id` | [NftIdentifier](#ununifi.nftmarket.NftIdentifier) |  |  |






<a name="ununifi.nftmarket.MsgPayFullBidResponse"></a>

### MsgPayFullBidResponse







<a name="ununifi.nftmarket.MsgPlaceBid"></a>

### MsgPlaceBid



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `nft_id` | [NftIdentifier](#ununifi.nftmarket.NftIdentifier) |  |  |
| `bid_amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `bidding_period` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `deposit_lending_rate` | [string](#string) |  |  |
| `automatic_payment` | [bool](#bool) |  |  |
| `deposit_amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="ununifi.nftmarket.MsgPlaceBidResponse"></a>

### MsgPlaceBidResponse







<a name="ununifi.nftmarket.MsgRepay"></a>

### MsgRepay



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `nft_id` | [NftIdentifier](#ununifi.nftmarket.NftIdentifier) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="ununifi.nftmarket.MsgRepayResponse"></a>

### MsgRepayResponse







<a name="ununifi.nftmarket.MsgSellingDecision"></a>

### MsgSellingDecision



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `nft_id` | [NftIdentifier](#ununifi.nftmarket.NftIdentifier) |  |  |






<a name="ununifi.nftmarket.MsgSellingDecisionResponse"></a>

### MsgSellingDecisionResponse






 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="ununifi.nftmarket.Msg"></a>

### Msg


| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `MintNft` | [MsgMintNft](#ununifi.nftmarket.MsgMintNft) | [MsgMintNftResponse](#ununifi.nftmarket.MsgMintNftResponse) |  | |
| `ListNft` | [MsgListNft](#ununifi.nftmarket.MsgListNft) | [MsgListNftResponse](#ununifi.nftmarket.MsgListNftResponse) |  | |
| `CancelNftListing` | [MsgCancelNftListing](#ununifi.nftmarket.MsgCancelNftListing) | [MsgCancelNftListingResponse](#ununifi.nftmarket.MsgCancelNftListingResponse) |  | |
| `PlaceBid` | [MsgPlaceBid](#ununifi.nftmarket.MsgPlaceBid) | [MsgPlaceBidResponse](#ununifi.nftmarket.MsgPlaceBidResponse) |  | |
| `CancelBid` | [MsgCancelBid](#ununifi.nftmarket.MsgCancelBid) | [MsgCancelBidResponse](#ununifi.nftmarket.MsgCancelBidResponse) |  | |
| `SellingDecision` | [MsgSellingDecision](#ununifi.nftmarket.MsgSellingDecision) | [MsgSellingDecisionResponse](#ununifi.nftmarket.MsgSellingDecisionResponse) |  | |
| `EndNftListing` | [MsgEndNftListing](#ununifi.nftmarket.MsgEndNftListing) | [MsgEndNftListingResponse](#ununifi.nftmarket.MsgEndNftListingResponse) |  | |
| `PayFullBid` | [MsgPayFullBid](#ununifi.nftmarket.MsgPayFullBid) | [MsgPayFullBidResponse](#ununifi.nftmarket.MsgPayFullBidResponse) |  | |
| `Borrow` | [MsgBorrow](#ununifi.nftmarket.MsgBorrow) | [MsgBorrowResponse](#ununifi.nftmarket.MsgBorrowResponse) |  | |
| `Repay` | [MsgRepay](#ununifi.nftmarket.MsgRepay) | [MsgRepayResponse](#ununifi.nftmarket.MsgRepayResponse) |  | |

 <!-- end services -->



<a name="nftmarketv1/nftmarket.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nftmarketv1/nftmarket.proto



<a name="ununifi.nftmarketv1.EventBorrow"></a>

### EventBorrow



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `borrower` | [string](#string) |  |  |
| `class_id` | [string](#string) |  |  |
| `nft_id` | [string](#string) |  |  |
| `amount` | [string](#string) |  |  |






<a name="ununifi.nftmarketv1.EventCancelBid"></a>

### EventCancelBid



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `bidder` | [string](#string) |  |  |
| `class_id` | [string](#string) |  |  |
| `nft_id` | [string](#string) |  |  |






<a name="ununifi.nftmarketv1.EventCancelListNfting"></a>

### EventCancelListNfting



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `owner` | [string](#string) |  |  |
| `class_id` | [string](#string) |  |  |
| `nft_id` | [string](#string) |  |  |






<a name="ununifi.nftmarketv1.EventEndListNfting"></a>

### EventEndListNfting



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `owner` | [string](#string) |  |  |
| `class_id` | [string](#string) |  |  |
| `nft_id` | [string](#string) |  |  |






<a name="ununifi.nftmarketv1.EventExpandListingPeriod"></a>

### EventExpandListingPeriod



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `owner` | [string](#string) |  |  |
| `class_id` | [string](#string) |  |  |
| `nft_id` | [string](#string) |  |  |






<a name="ununifi.nftmarketv1.EventLiquidate"></a>

### EventLiquidate



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `liquidator` | [string](#string) |  |  |
| `class_id` | [string](#string) |  |  |
| `nft_id` | [string](#string) |  |  |






<a name="ununifi.nftmarketv1.EventListNft"></a>

### EventListNft



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `owner` | [string](#string) |  |  |
| `class_id` | [string](#string) |  |  |
| `nft_id` | [string](#string) |  |  |






<a name="ununifi.nftmarketv1.EventPayFullBid"></a>

### EventPayFullBid



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `bidder` | [string](#string) |  |  |
| `class_id` | [string](#string) |  |  |
| `nft_id` | [string](#string) |  |  |






<a name="ununifi.nftmarketv1.EventPlaceBid"></a>

### EventPlaceBid



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `bidder` | [string](#string) |  |  |
| `class_id` | [string](#string) |  |  |
| `nft_id` | [string](#string) |  |  |
| `amount` | [string](#string) |  |  |






<a name="ununifi.nftmarketv1.EventRepay"></a>

### EventRepay



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `repayer` | [string](#string) |  |  |
| `class_id` | [string](#string) |  |  |
| `nft_id` | [string](#string) |  |  |
| `amount` | [string](#string) |  |  |






<a name="ununifi.nftmarketv1.EventSellingDecision"></a>

### EventSellingDecision



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `owner` | [string](#string) |  |  |
| `class_id` | [string](#string) |  |  |
| `nft_id` | [string](#string) |  |  |






<a name="ununifi.nftmarketv1.ListedClass"></a>

### ListedClass



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `class_id` | [string](#string) |  |  |
| `nft_ids` | [string](#string) | repeated |  |






<a name="ununifi.nftmarketv1.ListedNft"></a>

### ListedNft



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [string](#string) |  |  |
| `uri` | [string](#string) |  |  |
| `uri_hash` | [string](#string) |  |  |






<a name="ununifi.nftmarketv1.Loan"></a>

### Loan



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `nft_id` | [NftIdentifier](#ununifi.nftmarketv1.NftIdentifier) |  |  |
| `loan` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="ununifi.nftmarketv1.NftBid"></a>

### NftBid



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `nft_id` | [NftIdentifier](#ununifi.nftmarketv1.NftIdentifier) |  |  |
| `bidder` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `automatic_payment` | [bool](#bool) |  |  |
| `paid_amount` | [string](#string) |  |  |
| `bid_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="ununifi.nftmarketv1.NftIdentifier"></a>

### NftIdentifier



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `class_id` | [string](#string) |  |  |
| `nft_id` | [string](#string) |  |  |






<a name="ununifi.nftmarketv1.NftListing"></a>

### NftListing



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `nft_id` | [NftIdentifier](#ununifi.nftmarketv1.NftIdentifier) |  |  |
| `owner` | [string](#string) |  |  |
| `listing_type` | [ListingType](#ununifi.nftmarketv1.ListingType) |  |  |
| `state` | [ListingState](#ununifi.nftmarketv1.ListingState) |  |  |
| `bid_token` | [string](#string) |  |  |
| `min_bid` | [string](#string) |  |  |
| `bid_active_rank` | [uint64](#uint64) |  |  |
| `started_at` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `end_at` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `full_payment_end_at` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `successful_bid_end_at` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `auto_relisted_count` | [uint64](#uint64) |  |  |






<a name="ununifi.nftmarketv1.Params"></a>

### Params



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `min_staking_for_listing` | [string](#string) |  |  |
| `default_bid_active_rank` | [uint64](#uint64) |  |  |
| `bid_tokens` | [string](#string) | repeated |  |
| `auto_relisting_count_if_no_bid` | [uint64](#uint64) |  |  |
| `nft_listing_delay_seconds` | [uint64](#uint64) |  |  |
| `nft_listing_period_initial` | [uint64](#uint64) |  |  |
| `nft_listing_cancel_required_seconds` | [uint64](#uint64) |  |  |
| `nft_listing_cancel_fee_percentage` | [uint64](#uint64) |  |  |
| `nft_listing_gap_time` | [uint64](#uint64) |  |  |
| `bid_cancel_required_seconds` | [uint64](#uint64) |  |  |
| `bid_token_disburse_seconds_after_cancel` | [uint64](#uint64) |  |  |
| `nft_listing_full_payment_period` | [uint64](#uint64) |  |  |
| `nft_listing_nft_delivery_period` | [uint64](#uint64) |  |  |
| `nft_creator_share_percentage` | [uint64](#uint64) |  |  |
| `market_administrator` | [string](#string) |  |  |
| `nft_listing_commission_fee` | [uint64](#uint64) |  |  |
| `nft_listing_extend_seconds` | [uint64](#uint64) |  |  |
| `nft_listing_period_extend_fee_per_hour` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="ununifi.nftmarketv1.PaymentStatus"></a>

### PaymentStatus



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `nft_id` | [NftIdentifier](#ununifi.nftmarketv1.NftIdentifier) |  |  |
| `bidder` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `automatic_payment` | [bool](#bool) |  |  |
| `paid_amount` | [string](#string) |  |  |
| `bid_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `state` | [ListingState](#ununifi.nftmarketv1.ListingState) |  |  |
| `all_paid` | [bool](#bool) |  |  |





 <!-- end messages -->


<a name="ununifi.nftmarketv1.ListingState"></a>

### ListingState


| Name | Number | Description |
| ---- | ------ | ----------- |
| LISTING | 0 |  |
| BIDDING | 1 |  |
| SELLING_DECISION | 2 |  |
| LIQUIDATION | 3 |  |
| END_LISTING | 4 |  |
| SUCCESSFUL_BID | 5 |  |



<a name="ununifi.nftmarketv1.ListingType"></a>

### ListingType


| Name | Number | Description |
| ---- | ------ | ----------- |
| DIRECT_ASSET_BORROW | 0 |  |
| SYNTHETIC_ASSET_CREATION | 1 |  |
| LATE_SHIPPING | 2 |  |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="nftmarketv1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nftmarketv1/genesis.proto



<a name="ununifi.nftmarketv1.GenesisState"></a>

### GenesisState
GenesisState defines the nftmarket module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#ununifi.nftmarketv1.Params) |  |  |
| `listings` | [NftListing](#ununifi.nftmarketv1.NftListing) | repeated |  |
| `bids` | [NftBid](#ununifi.nftmarketv1.NftBid) | repeated |  |
| `cancelled_bids` | [NftBid](#ununifi.nftmarketv1.NftBid) | repeated |  |
| `loans` | [Loan](#ununifi.nftmarketv1.Loan) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="nftmarketv1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nftmarketv1/query.proto



<a name="ununifi.nftmarketv1.QueryBidderBidsRequest"></a>

### QueryBidderBidsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `bidder` | [string](#string) |  |  |






<a name="ununifi.nftmarketv1.QueryBidderBidsResponse"></a>

### QueryBidderBidsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `bids` | [NftBid](#ununifi.nftmarketv1.NftBid) | repeated |  |






<a name="ununifi.nftmarketv1.QueryCDPsListRequest"></a>

### QueryCDPsListRequest







<a name="ununifi.nftmarketv1.QueryCDPsListResponse"></a>

### QueryCDPsListResponse







<a name="ununifi.nftmarketv1.QueryListedClassRequest"></a>

### QueryListedClassRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `class_id` | [string](#string) |  |  |
| `nft_limit` | [int32](#int32) |  |  |






<a name="ununifi.nftmarketv1.QueryListedClassResponse"></a>

### QueryListedClassResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `class_id` | [string](#string) |  |  |
| `name` | [string](#string) |  |  |
| `description` | [string](#string) |  |  |
| `symbol` | [string](#string) |  |  |
| `uri` | [string](#string) |  |  |
| `urihash` | [string](#string) |  |  |
| `nfts` | [ListedNft](#ununifi.nftmarketv1.ListedNft) | repeated |  |
| `nft_count` | [uint64](#uint64) |  |  |






<a name="ununifi.nftmarketv1.QueryListedClassesRequest"></a>

### QueryListedClassesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `nft_limit` | [int32](#int32) |  |  |






<a name="ununifi.nftmarketv1.QueryListedClassesResponse"></a>

### QueryListedClassesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `classes` | [QueryListedClassResponse](#ununifi.nftmarketv1.QueryListedClassResponse) | repeated |  |






<a name="ununifi.nftmarketv1.QueryListedNftsRequest"></a>

### QueryListedNftsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `owner` | [string](#string) |  |  |






<a name="ununifi.nftmarketv1.QueryListedNftsResponse"></a>

### QueryListedNftsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `listings` | [NftListing](#ununifi.nftmarketv1.NftListing) | repeated |  |






<a name="ununifi.nftmarketv1.QueryLoanRequest"></a>

### QueryLoanRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `class_id` | [string](#string) |  |  |
| `nft_id` | [string](#string) |  |  |






<a name="ununifi.nftmarketv1.QueryLoanResponse"></a>

### QueryLoanResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `loan` | [Loan](#ununifi.nftmarketv1.Loan) |  |  |
| `borrowing_limit` | [string](#string) |  |  |






<a name="ununifi.nftmarketv1.QueryLoansRequest"></a>

### QueryLoansRequest







<a name="ununifi.nftmarketv1.QueryLoansResponse"></a>

### QueryLoansResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `loans` | [Loan](#ununifi.nftmarketv1.Loan) | repeated |  |






<a name="ununifi.nftmarketv1.QueryNftBidsRequest"></a>

### QueryNftBidsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `class_id` | [string](#string) |  |  |
| `nft_id` | [string](#string) |  |  |






<a name="ununifi.nftmarketv1.QueryNftBidsResponse"></a>

### QueryNftBidsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `bids` | [NftBid](#ununifi.nftmarketv1.NftBid) | repeated |  |






<a name="ununifi.nftmarketv1.QueryNftListingRequest"></a>

### QueryNftListingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `class_id` | [string](#string) |  |  |
| `nft_id` | [string](#string) |  |  |






<a name="ununifi.nftmarketv1.QueryNftListingResponse"></a>

### QueryNftListingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `listing` | [NftListing](#ununifi.nftmarketv1.NftListing) |  |  |






<a name="ununifi.nftmarketv1.QueryParamsRequest"></a>

### QueryParamsRequest







<a name="ununifi.nftmarketv1.QueryParamsResponse"></a>

### QueryParamsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#ununifi.nftmarketv1.Params) |  |  |






<a name="ununifi.nftmarketv1.QueryPaymentStatusRequest"></a>

### QueryPaymentStatusRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `class_id` | [string](#string) |  |  |
| `nft_id` | [string](#string) |  |  |
| `bidder` | [string](#string) |  |  |






<a name="ununifi.nftmarketv1.QueryPaymentStatusResponse"></a>

### QueryPaymentStatusResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `paymentStatus` | [PaymentStatus](#ununifi.nftmarketv1.PaymentStatus) |  |  |






<a name="ununifi.nftmarketv1.QueryRewardsRequest"></a>

### QueryRewardsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [uint64](#uint64) |  |  |






<a name="ununifi.nftmarketv1.QueryRewardsResponse"></a>

### QueryRewardsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `rewards` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="ununifi.nftmarketv1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#ununifi.nftmarketv1.QueryParamsRequest) | [QueryParamsResponse](#ununifi.nftmarketv1.QueryParamsResponse) |  | GET|/ununifi/nftmarket/params|
| `NftListing` | [QueryNftListingRequest](#ununifi.nftmarketv1.QueryNftListingRequest) | [QueryNftListingResponse](#ununifi.nftmarketv1.QueryNftListingResponse) |  | GET|/ununifi/nftmarket/nft_listing/{class_id}/{nft_id}|
| `ListedNfts` | [QueryListedNftsRequest](#ununifi.nftmarketv1.QueryListedNftsRequest) | [QueryListedNftsResponse](#ununifi.nftmarketv1.QueryListedNftsResponse) |  | GET|/ununifi/nftmarket/listed_nfts|
| `ListedClasses` | [QueryListedClassesRequest](#ununifi.nftmarketv1.QueryListedClassesRequest) | [QueryListedClassesResponse](#ununifi.nftmarketv1.QueryListedClassesResponse) |  | GET|/ununifi/nftmarket/listed_classes|
| `ListedClass` | [QueryListedClassRequest](#ununifi.nftmarketv1.QueryListedClassRequest) | [QueryListedClassResponse](#ununifi.nftmarketv1.QueryListedClassResponse) |  | GET|/ununifi/nftmarket/listed_class/{class_id}/{nft_limit}|
| `Loans` | [QueryLoansRequest](#ununifi.nftmarketv1.QueryLoansRequest) | [QueryLoansResponse](#ununifi.nftmarketv1.QueryLoansResponse) |  | GET|/ununifi/nftmarket/loans|
| `Loan` | [QueryLoanRequest](#ununifi.nftmarketv1.QueryLoanRequest) | [QueryLoanResponse](#ununifi.nftmarketv1.QueryLoanResponse) |  | GET|/ununifi/nftmarket/loans/{class_id}/{nft_id}|
| `CDPsList` | [QueryCDPsListRequest](#ununifi.nftmarketv1.QueryCDPsListRequest) | [QueryCDPsListResponse](#ununifi.nftmarketv1.QueryCDPsListResponse) |  | GET|/ununifi/nftmarket/cdps_list|
| `NftBids` | [QueryNftBidsRequest](#ununifi.nftmarketv1.QueryNftBidsRequest) | [QueryNftBidsResponse](#ununifi.nftmarketv1.QueryNftBidsResponse) |  | GET|/ununifi/nftmarket/nft_bids/{class_id}/{nft_id}|
| `BidderBids` | [QueryBidderBidsRequest](#ununifi.nftmarketv1.QueryBidderBidsRequest) | [QueryBidderBidsResponse](#ununifi.nftmarketv1.QueryBidderBidsResponse) |  | GET|/ununifi/nftmarket/bidder_bids/{bidder}|
| `PaymentStatus` | [QueryPaymentStatusRequest](#ununifi.nftmarketv1.QueryPaymentStatusRequest) | [QueryPaymentStatusResponse](#ununifi.nftmarketv1.QueryPaymentStatusResponse) |  | GET|/ununifi/nftmarket/payment_status/{class_id}/{nft_id}/{bidder}|
| `Rewards` | [QueryRewardsRequest](#ununifi.nftmarketv1.QueryRewardsRequest) | [QueryRewardsResponse](#ununifi.nftmarketv1.QueryRewardsResponse) |  | GET|/ununifi/nftmarket/rewards/{address}|

 <!-- end services -->



<a name="nftmarketv1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nftmarketv1/tx.proto



<a name="ununifi.nftmarketv1.MsgBorrow"></a>

### MsgBorrow



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `nft_id` | [NftIdentifier](#ununifi.nftmarketv1.NftIdentifier) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="ununifi.nftmarketv1.MsgBorrowResponse"></a>

### MsgBorrowResponse







<a name="ununifi.nftmarketv1.MsgBurnStableCoin"></a>

### MsgBurnStableCoin



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |






<a name="ununifi.nftmarketv1.MsgBurnStableCoinResponse"></a>

### MsgBurnStableCoinResponse







<a name="ununifi.nftmarketv1.MsgCancelBid"></a>

### MsgCancelBid



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `nft_id` | [NftIdentifier](#ununifi.nftmarketv1.NftIdentifier) |  |  |






<a name="ununifi.nftmarketv1.MsgCancelBidResponse"></a>

### MsgCancelBidResponse







<a name="ununifi.nftmarketv1.MsgCancelNftListing"></a>

### MsgCancelNftListing



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `nft_id` | [NftIdentifier](#ununifi.nftmarketv1.NftIdentifier) |  |  |






<a name="ununifi.nftmarketv1.MsgCancelNftListingResponse"></a>

### MsgCancelNftListingResponse







<a name="ununifi.nftmarketv1.MsgEndNftListing"></a>

### MsgEndNftListing



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `nft_id` | [NftIdentifier](#ununifi.nftmarketv1.NftIdentifier) |  |  |






<a name="ununifi.nftmarketv1.MsgEndNftListingResponse"></a>

### MsgEndNftListingResponse







<a name="ununifi.nftmarketv1.MsgExpandListingPeriod"></a>

### MsgExpandListingPeriod



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `nft_id` | [NftIdentifier](#ununifi.nftmarketv1.NftIdentifier) |  |  |






<a name="ununifi.nftmarketv1.MsgExpandListingPeriodResponse"></a>

### MsgExpandListingPeriodResponse







<a name="ununifi.nftmarketv1.MsgLiquidate"></a>

### MsgLiquidate



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `nft_id` | [NftIdentifier](#ununifi.nftmarketv1.NftIdentifier) |  |  |






<a name="ununifi.nftmarketv1.MsgLiquidateResponse"></a>

### MsgLiquidateResponse







<a name="ununifi.nftmarketv1.MsgListNft"></a>

### MsgListNft



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `nft_id` | [NftIdentifier](#ununifi.nftmarketv1.NftIdentifier) |  |  |
| `listing_type` | [ListingType](#ununifi.nftmarketv1.ListingType) |  |  |
| `bid_token` | [string](#string) |  |  |
| `min_bid` | [string](#string) |  |  |
| `bid_active_rank` | [uint64](#uint64) |  |  |






<a name="ununifi.nftmarketv1.MsgListNftResponse"></a>

### MsgListNftResponse







<a name="ununifi.nftmarketv1.MsgMintNft"></a>

### MsgMintNft



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `classId` | [string](#string) |  |  |
| `nftId` | [string](#string) |  |  |
| `nftUri` | [string](#string) |  |  |
| `nftUriHash` | [string](#string) |  |  |






<a name="ununifi.nftmarketv1.MsgMintNftResponse"></a>

### MsgMintNftResponse







<a name="ununifi.nftmarketv1.MsgMintStableCoin"></a>

### MsgMintStableCoin



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |






<a name="ununifi.nftmarketv1.MsgMintStableCoinResponse"></a>

### MsgMintStableCoinResponse







<a name="ununifi.nftmarketv1.MsgPayFullBid"></a>

### MsgPayFullBid



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `nft_id` | [NftIdentifier](#ununifi.nftmarketv1.NftIdentifier) |  |  |






<a name="ununifi.nftmarketv1.MsgPayFullBidResponse"></a>

### MsgPayFullBidResponse







<a name="ununifi.nftmarketv1.MsgPlaceBid"></a>

### MsgPlaceBid



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `nft_id` | [NftIdentifier](#ununifi.nftmarketv1.NftIdentifier) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `automatic_payment` | [bool](#bool) |  |  |






<a name="ununifi.nftmarketv1.MsgPlaceBidResponse"></a>

### MsgPlaceBidResponse







<a name="ununifi.nftmarketv1.MsgRepay"></a>

### MsgRepay



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `nft_id` | [NftIdentifier](#ununifi.nftmarketv1.NftIdentifier) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="ununifi.nftmarketv1.MsgRepayResponse"></a>

### MsgRepayResponse







<a name="ununifi.nftmarketv1.MsgSellingDecision"></a>

### MsgSellingDecision



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `nft_id` | [NftIdentifier](#ununifi.nftmarketv1.NftIdentifier) |  |  |






<a name="ununifi.nftmarketv1.MsgSellingDecisionResponse"></a>

### MsgSellingDecisionResponse






 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="ununifi.nftmarketv1.Msg"></a>

### Msg


| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `MintNft` | [MsgMintNft](#ununifi.nftmarketv1.MsgMintNft) | [MsgMintNftResponse](#ununifi.nftmarketv1.MsgMintNftResponse) |  | |
| `ListNft` | [MsgListNft](#ununifi.nftmarketv1.MsgListNft) | [MsgListNftResponse](#ununifi.nftmarketv1.MsgListNftResponse) |  | |
| `CancelNftListing` | [MsgCancelNftListing](#ununifi.nftmarketv1.MsgCancelNftListing) | [MsgCancelNftListingResponse](#ununifi.nftmarketv1.MsgCancelNftListingResponse) |  | |
| `ExpandListingPeriod` | [MsgExpandListingPeriod](#ununifi.nftmarketv1.MsgExpandListingPeriod) | [MsgExpandListingPeriodResponse](#ununifi.nftmarketv1.MsgExpandListingPeriodResponse) |  | |
| `PlaceBid` | [MsgPlaceBid](#ununifi.nftmarketv1.MsgPlaceBid) | [MsgPlaceBidResponse](#ununifi.nftmarketv1.MsgPlaceBidResponse) |  | |
| `CancelBid` | [MsgCancelBid](#ununifi.nftmarketv1.MsgCancelBid) | [MsgCancelBidResponse](#ununifi.nftmarketv1.MsgCancelBidResponse) |  | |
| `SellingDecision` | [MsgSellingDecision](#ununifi.nftmarketv1.MsgSellingDecision) | [MsgSellingDecisionResponse](#ununifi.nftmarketv1.MsgSellingDecisionResponse) |  | |
| `EndNftListing` | [MsgEndNftListing](#ununifi.nftmarketv1.MsgEndNftListing) | [MsgEndNftListingResponse](#ununifi.nftmarketv1.MsgEndNftListingResponse) |  | |
| `PayFullBid` | [MsgPayFullBid](#ununifi.nftmarketv1.MsgPayFullBid) | [MsgPayFullBidResponse](#ununifi.nftmarketv1.MsgPayFullBidResponse) |  | |
| `Borrow` | [MsgBorrow](#ununifi.nftmarketv1.MsgBorrow) | [MsgBorrowResponse](#ununifi.nftmarketv1.MsgBorrowResponse) |  | |
| `Repay` | [MsgRepay](#ununifi.nftmarketv1.MsgRepay) | [MsgRepayResponse](#ununifi.nftmarketv1.MsgRepayResponse) |  | |
| `MintStableCoin` | [MsgMintStableCoin](#ununifi.nftmarketv1.MsgMintStableCoin) | [MsgMintStableCoinResponse](#ununifi.nftmarketv1.MsgMintStableCoinResponse) |  | |
| `BurnStableCoin` | [MsgBurnStableCoin](#ununifi.nftmarketv1.MsgBurnStableCoin) | [MsgBurnStableCoinResponse](#ununifi.nftmarketv1.MsgBurnStableCoinResponse) |  | |
| `Liquidate` | [MsgLiquidate](#ununifi.nftmarketv1.MsgLiquidate) | [MsgLiquidateResponse](#ununifi.nftmarketv1.MsgLiquidateResponse) |  | |

 <!-- end services -->



<a name="nftmint/nftmint.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nftmint/nftmint.proto



<a name="ununifi.nftmint.ClassAttributes"></a>

### ClassAttributes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `class_id` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |
| `base_token_uri` | [string](#string) |  |  |
| `minting_permission` | [MintingPermission](#ununifi.nftmint.MintingPermission) |  |  |
| `token_supply_cap` | [uint64](#uint64) |  |  |






<a name="ununifi.nftmint.ClassNameIdList"></a>

### ClassNameIdList



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `class_name` | [string](#string) |  |  |
| `class_id` | [string](#string) | repeated |  |






<a name="ununifi.nftmint.OwningClassIdList"></a>

### OwningClassIdList



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `owner` | [string](#string) |  |  |
| `class_id` | [string](#string) | repeated |  |






<a name="ununifi.nftmint.Params"></a>

### Params



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `MaxNFTSupplyCap` | [uint64](#uint64) |  |  |
| `MinClassNameLen` | [uint64](#uint64) |  |  |
| `MaxClassNameLen` | [uint64](#uint64) |  |  |
| `MinUriLen` | [uint64](#uint64) |  |  |
| `MaxUriLen` | [uint64](#uint64) |  |  |
| `MaxSymbolLen` | [uint64](#uint64) |  |  |
| `MaxDescriptionLen` | [uint64](#uint64) |  |  |





 <!-- end messages -->


<a name="ununifi.nftmint.MintingPermission"></a>

### MintingPermission


| Name | Number | Description |
| ---- | ------ | ----------- |
| OnlyOwner | 0 |  |
| Anyone | 1 | WhiteList = 2; |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="nftmint/event.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nftmint/event.proto



<a name="ununifi.nftmint.EventBurnNFT"></a>

### EventBurnNFT



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `burner` | [string](#string) |  |  |
| `class_id` | [string](#string) |  |  |
| `nft_id` | [string](#string) |  |  |






<a name="ununifi.nftmint.EventCreateClass"></a>

### EventCreateClass



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `owner` | [string](#string) |  |  |
| `class_id` | [string](#string) |  |  |
| `base_token_uri` | [string](#string) |  |  |
| `token_supply_cap` | [string](#string) |  |  |
| `minting_permission` | [MintingPermission](#ununifi.nftmint.MintingPermission) |  |  |






<a name="ununifi.nftmint.EventMintNFT"></a>

### EventMintNFT



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `class_id` | [string](#string) |  |  |
| `nft_id` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |
| `minter` | [string](#string) |  |  |






<a name="ununifi.nftmint.EventSendClassOwnership"></a>

### EventSendClassOwnership



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `receiver` | [string](#string) |  |  |
| `class_id` | [string](#string) |  |  |






<a name="ununifi.nftmint.EventUpdateBaseTokenUri"></a>

### EventUpdateBaseTokenUri



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `owner` | [string](#string) |  |  |
| `class_id` | [string](#string) |  |  |
| `base_token_uri` | [string](#string) |  |  |






<a name="ununifi.nftmint.EventUpdateTokenSupplyCap"></a>

### EventUpdateTokenSupplyCap



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `owner` | [string](#string) |  |  |
| `class_id` | [string](#string) |  |  |
| `token_supply_cap` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="nftmint/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nftmint/genesis.proto



<a name="ununifi.nftmint.GenesisState"></a>

### GenesisState
GenesisState defines the nftmint module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#ununifi.nftmint.Params) |  |  |
| `class_attributes_list` | [ClassAttributes](#ununifi.nftmint.ClassAttributes) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="nftmint/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nftmint/query.proto



<a name="ununifi.nftmint.QueryClassAttributesRequest"></a>

### QueryClassAttributesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `class_id` | [string](#string) |  |  |






<a name="ununifi.nftmint.QueryClassAttributesResponse"></a>

### QueryClassAttributesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `class_attributes` | [ClassAttributes](#ununifi.nftmint.ClassAttributes) |  |  |






<a name="ununifi.nftmint.QueryClassIdsByNameRequest"></a>

### QueryClassIdsByNameRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `class_name` | [string](#string) |  |  |






<a name="ununifi.nftmint.QueryClassIdsByNameResponse"></a>

### QueryClassIdsByNameResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `class_name_id_list` | [ClassNameIdList](#ununifi.nftmint.ClassNameIdList) |  |  |






<a name="ununifi.nftmint.QueryClassIdsByOwnerRequest"></a>

### QueryClassIdsByOwnerRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `owner` | [string](#string) |  |  |






<a name="ununifi.nftmint.QueryClassIdsByOwnerResponse"></a>

### QueryClassIdsByOwnerResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `owning_class_id_list` | [OwningClassIdList](#ununifi.nftmint.OwningClassIdList) |  |  |






<a name="ununifi.nftmint.QueryNFTMinterRequest"></a>

### QueryNFTMinterRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `class_id` | [string](#string) |  |  |
| `nft_id` | [string](#string) |  |  |






<a name="ununifi.nftmint.QueryNFTMinterResponse"></a>

### QueryNFTMinterResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `minter` | [string](#string) |  |  |






<a name="ununifi.nftmint.QueryParamsRequest"></a>

### QueryParamsRequest







<a name="ununifi.nftmint.QueryParamsResponse"></a>

### QueryParamsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#ununifi.nftmint.Params) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="ununifi.nftmint.Query"></a>

### Query


| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#ununifi.nftmint.QueryParamsRequest) | [QueryParamsResponse](#ununifi.nftmint.QueryParamsResponse) |  | GET|/ununifi/nftmint/params|
| `ClassAttributes` | [QueryClassAttributesRequest](#ununifi.nftmint.QueryClassAttributesRequest) | [QueryClassAttributesResponse](#ununifi.nftmint.QueryClassAttributesResponse) |  | GET|/ununifi/nftmint/class_owner/{class_id}|
| `NFTMinter` | [QueryNFTMinterRequest](#ununifi.nftmint.QueryNFTMinterRequest) | [QueryNFTMinterResponse](#ununifi.nftmint.QueryNFTMinterResponse) |  | GET|/ununifi/nftmint/nft_minter/{class_id}/{nft_id}|
| `ClassIdsByName` | [QueryClassIdsByNameRequest](#ununifi.nftmint.QueryClassIdsByNameRequest) | [QueryClassIdsByNameResponse](#ununifi.nftmint.QueryClassIdsByNameResponse) |  | GET|/ununifi/nftmint/class_ids_by_name/{class_name}|
| `ClassIdsByOwner` | [QueryClassIdsByOwnerRequest](#ununifi.nftmint.QueryClassIdsByOwnerRequest) | [QueryClassIdsByOwnerResponse](#ununifi.nftmint.QueryClassIdsByOwnerResponse) |  | GET|/ununifi/nftmint/class_ids_by_owner/{owner}|

 <!-- end services -->



<a name="nftmint/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nftmint/tx.proto



<a name="ununifi.nftmint.MsgBurnNFT"></a>

### MsgBurnNFT



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `class_id` | [string](#string) |  |  |
| `nft_id` | [string](#string) |  |  |






<a name="ununifi.nftmint.MsgBurnNFTResponse"></a>

### MsgBurnNFTResponse







<a name="ununifi.nftmint.MsgCreateClass"></a>

### MsgCreateClass



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `name` | [string](#string) |  |  |
| `base_token_uri` | [string](#string) |  |  |
| `token_supply_cap` | [uint64](#uint64) |  |  |
| `minting_permission` | [MintingPermission](#ununifi.nftmint.MintingPermission) |  |  |
| `symbol` | [string](#string) |  |  |
| `description` | [string](#string) |  |  |
| `class_uri` | [string](#string) |  |  |






<a name="ununifi.nftmint.MsgCreateClassResponse"></a>

### MsgCreateClassResponse







<a name="ununifi.nftmint.MsgMintNFT"></a>

### MsgMintNFT



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `class_id` | [string](#string) |  |  |
| `nft_id` | [string](#string) |  |  |
| `recipient` | [string](#string) |  |  |






<a name="ununifi.nftmint.MsgMintNFTResponse"></a>

### MsgMintNFTResponse







<a name="ununifi.nftmint.MsgSendClassOwnership"></a>

### MsgSendClassOwnership



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `class_id` | [string](#string) |  |  |
| `recipient` | [string](#string) |  |  |






<a name="ununifi.nftmint.MsgSendClassOwnershipResponse"></a>

### MsgSendClassOwnershipResponse







<a name="ununifi.nftmint.MsgUpdateBaseTokenUri"></a>

### MsgUpdateBaseTokenUri



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `class_id` | [string](#string) |  |  |
| `base_token_uri` | [string](#string) |  |  |






<a name="ununifi.nftmint.MsgUpdateBaseTokenUriResponse"></a>

### MsgUpdateBaseTokenUriResponse







<a name="ununifi.nftmint.MsgUpdateTokenSupplyCap"></a>

### MsgUpdateTokenSupplyCap



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `class_id` | [string](#string) |  |  |
| `token_supply_cap` | [uint64](#uint64) |  |  |






<a name="ununifi.nftmint.MsgUpdateTokenSupplyCapResponse"></a>

### MsgUpdateTokenSupplyCapResponse






 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="ununifi.nftmint.Msg"></a>

### Msg


| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `CreateClass` | [MsgCreateClass](#ununifi.nftmint.MsgCreateClass) | [MsgCreateClassResponse](#ununifi.nftmint.MsgCreateClassResponse) |  | |
| `SendClassOwnership` | [MsgSendClassOwnership](#ununifi.nftmint.MsgSendClassOwnership) | [MsgSendClassOwnershipResponse](#ununifi.nftmint.MsgSendClassOwnershipResponse) |  | |
| `UpdateBaseTokenUri` | [MsgUpdateBaseTokenUri](#ununifi.nftmint.MsgUpdateBaseTokenUri) | [MsgUpdateBaseTokenUriResponse](#ununifi.nftmint.MsgUpdateBaseTokenUriResponse) |  | |
| `UpdateTokenSupplyCap` | [MsgUpdateTokenSupplyCap](#ununifi.nftmint.MsgUpdateTokenSupplyCap) | [MsgUpdateTokenSupplyCapResponse](#ununifi.nftmint.MsgUpdateTokenSupplyCapResponse) |  | |
| `MintNFT` | [MsgMintNFT](#ununifi.nftmint.MsgMintNFT) | [MsgMintNFTResponse](#ununifi.nftmint.MsgMintNFTResponse) |  | |
| `BurnNFT` | [MsgBurnNFT](#ununifi.nftmint.MsgBurnNFT) | [MsgBurnNFTResponse](#ununifi.nftmint.MsgBurnNFTResponse) |  | |

 <!-- end services -->



<a name="pricefeed/pricefeed.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## pricefeed/pricefeed.proto



<a name="ununifi.pricefeed.CurrentPrice"></a>

### CurrentPrice



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `market_id` | [string](#string) |  |  |
| `price` | [string](#string) |  |  |






<a name="ununifi.pricefeed.Market"></a>

### Market



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `market_id` | [string](#string) |  |  |
| `base_asset` | [string](#string) |  |  |
| `quote_asset` | [string](#string) |  |  |
| `oracles` | [string](#string) | repeated |  |
| `active` | [bool](#bool) |  |  |






<a name="ununifi.pricefeed.Params"></a>

### Params



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `markets` | [Market](#ununifi.pricefeed.Market) | repeated |  |






<a name="ununifi.pricefeed.PostedPrice"></a>

### PostedPrice



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `market_id` | [string](#string) |  |  |
| `oracle_address` | [string](#string) |  |  |
| `price` | [string](#string) |  |  |
| `expiry` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="pricefeed/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## pricefeed/genesis.proto



<a name="ununifi.pricefeed.GenesisState"></a>

### GenesisState
GenesisState defines the pricefeed module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#ununifi.pricefeed.Params) |  |  |
| `posted_prices` | [PostedPrice](#ununifi.pricefeed.PostedPrice) | repeated | this line is used by starport scaffolding # genesis/proto/state |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="pricefeed/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## pricefeed/query.proto



<a name="ununifi.pricefeed.QueryAllMarketRequest"></a>

### QueryAllMarketRequest
this line is used by starport scaffolding # 3


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="ununifi.pricefeed.QueryAllMarketResponse"></a>

### QueryAllMarketResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `markets` | [Market](#ununifi.pricefeed.Market) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="ununifi.pricefeed.QueryAllOracleRequest"></a>

### QueryAllOracleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `market_id` | [string](#string) |  |  |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="ununifi.pricefeed.QueryAllOracleResponse"></a>

### QueryAllOracleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `oracles` | [string](#string) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="ununifi.pricefeed.QueryAllPriceRequest"></a>

### QueryAllPriceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="ununifi.pricefeed.QueryAllPriceResponse"></a>

### QueryAllPriceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `prices` | [CurrentPrice](#ununifi.pricefeed.CurrentPrice) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="ununifi.pricefeed.QueryAllRawPriceRequest"></a>

### QueryAllRawPriceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `market_id` | [string](#string) |  |  |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="ununifi.pricefeed.QueryAllRawPriceResponse"></a>

### QueryAllRawPriceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `prices` | [PostedPrice](#ununifi.pricefeed.PostedPrice) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="ununifi.pricefeed.QueryGetPriceRequest"></a>

### QueryGetPriceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `market_id` | [string](#string) |  |  |






<a name="ununifi.pricefeed.QueryGetPriceResponse"></a>

### QueryGetPriceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `price` | [CurrentPrice](#ununifi.pricefeed.CurrentPrice) |  |  |






<a name="ununifi.pricefeed.QueryParamsRequest"></a>

### QueryParamsRequest







<a name="ununifi.pricefeed.QueryParamsResponse"></a>

### QueryParamsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#ununifi.pricefeed.Params) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="ununifi.pricefeed.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#ununifi.pricefeed.QueryParamsRequest) | [QueryParamsResponse](#ununifi.pricefeed.QueryParamsResponse) |  | GET|/ununifi/pricefeed/params|
| `MarketAll` | [QueryAllMarketRequest](#ununifi.pricefeed.QueryAllMarketRequest) | [QueryAllMarketResponse](#ununifi.pricefeed.QueryAllMarketResponse) | this line is used by starport scaffolding # 2 | GET|/ununifi/pricefeed/markets|
| `OracleAll` | [QueryAllOracleRequest](#ununifi.pricefeed.QueryAllOracleRequest) | [QueryAllOracleResponse](#ununifi.pricefeed.QueryAllOracleResponse) |  | GET|/ununifi/pricefeed/markets/{market_id}/oracles|
| `Price` | [QueryGetPriceRequest](#ununifi.pricefeed.QueryGetPriceRequest) | [QueryGetPriceResponse](#ununifi.pricefeed.QueryGetPriceResponse) |  | GET|/ununifi/pricefeed/markets/{market_id}/price|
| `PriceAll` | [QueryAllPriceRequest](#ununifi.pricefeed.QueryAllPriceRequest) | [QueryAllPriceResponse](#ununifi.pricefeed.QueryAllPriceResponse) |  | GET|/ununifi/pricefeed/prices|
| `RawPriceAll` | [QueryAllRawPriceRequest](#ununifi.pricefeed.QueryAllRawPriceRequest) | [QueryAllRawPriceResponse](#ununifi.pricefeed.QueryAllRawPriceResponse) |  | GET|/ununifi/pricefeed/markets/{market_id}/raw_prices|

 <!-- end services -->



<a name="pricefeed/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## pricefeed/tx.proto



<a name="ununifi.pricefeed.MsgPostPrice"></a>

### MsgPostPrice



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `from` | [string](#string) |  |  |
| `market_id` | [string](#string) |  |  |
| `price` | [string](#string) |  |  |
| `expiry` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="ununifi.pricefeed.MsgPostPriceResponse"></a>

### MsgPostPriceResponse






 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="ununifi.pricefeed.Msg"></a>

### Msg


| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `PostPrice` | [MsgPostPrice](#ununifi.pricefeed.MsgPostPrice) | [MsgPostPriceResponse](#ununifi.pricefeed.MsgPostPriceResponse) |  | |

 <!-- end services -->



<a name="ununifidist/ununifidist.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ununifidist/ununifidist.proto



<a name="ununifi.ununifidist.Params"></a>

### Params



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `active` | [bool](#bool) |  |  |
| `periods` | [Period](#ununifi.ununifidist.Period) | repeated |  |






<a name="ununifi.ununifidist.Period"></a>

### Period



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `start` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `end` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `inflation` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="ununifidist/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ununifidist/genesis.proto



<a name="ununifi.ununifidist.GenesisState"></a>

### GenesisState
GenesisState defines the ununifidist module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#ununifi.ununifidist.Params) |  |  |
| `previous_block_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `gov_denom` | [string](#string) |  | this line is used by starport scaffolding # genesis/proto/state |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="ununifidist/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ununifidist/query.proto



<a name="ununifi.ununifidist.QueryGetBalancesRequest"></a>

### QueryGetBalancesRequest







<a name="ununifi.ununifidist.QueryGetBalancesResponse"></a>

### QueryGetBalancesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `balances` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="ununifi.ununifidist.QueryParamsRequest"></a>

### QueryParamsRequest







<a name="ununifi.ununifidist.QueryParamsResponse"></a>

### QueryParamsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#ununifi.ununifidist.Params) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="ununifi.ununifidist.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#ununifi.ununifidist.QueryParamsRequest) | [QueryParamsResponse](#ununifi.ununifidist.QueryParamsResponse) |  | GET|/ununifi/ununifidist/params|
| `Balances` | [QueryGetBalancesRequest](#ununifi.ununifidist.QueryGetBalancesRequest) | [QueryGetBalancesResponse](#ununifi.ununifidist.QueryGetBalancesResponse) | this line is used by starport scaffolding # 2 | GET|/ununifi/ununifidist/balances|

 <!-- end services -->



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |
