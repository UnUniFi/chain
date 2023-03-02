<!-- This file is auto-generated. Please do not modify it yourself. -->

# Protobuf Documentation

<a name="top"></a>

## Table of Contents

<<<<<<< HEAD
<<<<<<< HEAD
=======

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
  - [ListedClass](#ununifi.nftmarket.ListedClass)
  - [ListedNft](#ununifi.nftmarket.ListedNft)
  - [Loan](#ununifi.nftmarket.Loan)
  - [NftBid](#ununifi.nftmarket.NftBid)
  - [NftIdentifier](#ununifi.nftmarket.NftIdentifier)
  - [NftListing](#ununifi.nftmarket.NftListing)
  - [Params](#ununifi.nftmarket.Params)
  - [PaymentStatus](#ununifi.nftmarket.PaymentStatus)

  - [ListingState](#ununifi.nftmarket.ListingState)
  - [ListingType](#ununifi.nftmarket.ListingType)

- [nftmarket/genesis.proto](#nftmarket/genesis.proto)
  - [GenesisState](#ununifi.nftmarket.GenesisState)
- [nftmarket/query.proto](#nftmarket/query.proto)

  - [QueryBidderBidsRequest](#ununifi.nftmarket.QueryBidderBidsRequest)
  - [QueryBidderBidsResponse](#ununifi.nftmarket.QueryBidderBidsResponse)
  - [QueryCDPsListRequest](#ununifi.nftmarket.QueryCDPsListRequest)
  - [QueryCDPsListResponse](#ununifi.nftmarket.QueryCDPsListResponse)
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
  - [MsgBurnStableCoin](#ununifi.nftmarket.MsgBurnStableCoin)
  - [MsgBurnStableCoinResponse](#ununifi.nftmarket.MsgBurnStableCoinResponse)
  - [MsgCancelBid](#ununifi.nftmarket.MsgCancelBid)
  - [MsgCancelBidResponse](#ununifi.nftmarket.MsgCancelBidResponse)
  - [MsgCancelNftListing](#ununifi.nftmarket.MsgCancelNftListing)
  - [MsgCancelNftListingResponse](#ununifi.nftmarket.MsgCancelNftListingResponse)
  - [MsgEndNftListing](#ununifi.nftmarket.MsgEndNftListing)
  - [MsgEndNftListingResponse](#ununifi.nftmarket.MsgEndNftListingResponse)
  - [MsgExpandListingPeriod](#ununifi.nftmarket.MsgExpandListingPeriod)
  - [MsgExpandListingPeriodResponse](#ununifi.nftmarket.MsgExpandListingPeriodResponse)
  - [MsgLiquidate](#ununifi.nftmarket.MsgLiquidate)
  - [MsgLiquidateResponse](#ununifi.nftmarket.MsgLiquidateResponse)
  - [MsgListNft](#ununifi.nftmarket.MsgListNft)
  - [MsgListNftResponse](#ununifi.nftmarket.MsgListNftResponse)
  - [MsgMintNft](#ununifi.nftmarket.MsgMintNft)
  - [MsgMintNftResponse](#ununifi.nftmarket.MsgMintNftResponse)
  - [MsgMintStableCoin](#ununifi.nftmarket.MsgMintStableCoin)
  - [MsgMintStableCoinResponse](#ununifi.nftmarket.MsgMintStableCoinResponse)
  - [MsgPayFullBid](#ununifi.nftmarket.MsgPayFullBid)
  - [MsgPayFullBidResponse](#ununifi.nftmarket.MsgPayFullBidResponse)
  - [MsgPlaceBid](#ununifi.nftmarket.MsgPlaceBid)
  - [MsgPlaceBidResponse](#ununifi.nftmarket.MsgPlaceBidResponse)
  - [MsgRepay](#ununifi.nftmarket.MsgRepay)
  - [MsgRepayResponse](#ununifi.nftmarket.MsgRepayResponse)
  - [MsgSellingDecision](#ununifi.nftmarket.MsgSellingDecision)
  - [MsgSellingDecisionResponse](#ununifi.nftmarket.MsgSellingDecisionResponse)

  - [Msg](#ununifi.nftmarket.Msg)

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

- [yield-aggregator/params.proto](#yield-aggregator/params.proto)
  - [Params](#ununifi.chain.yieldaggregator.Params)
  - [Params.Vault](#ununifi.chain.yieldaggregator.Params.Vault)
- [yield-aggregator/yield-aggregator.proto](#yield-aggregator/yield-aggregator.proto)
  - [Strategy](#ununifi.chain.yieldaggregator.Strategy)
  - [StrategyMetrics](#ununifi.chain.yieldaggregator.StrategyMetrics)
  - [Vault](#ununifi.chain.yieldaggregator.Vault)
  - [VaultMetrics](#ununifi.chain.yieldaggregator.VaultMetrics)
- [yield-aggregator/genesis.proto](#yield-aggregator/genesis.proto)
  - [GenesisState](#ununifi.chain.yieldaggregator.GenesisState)
- [yield-aggregator/query.proto](#yield-aggregator/query.proto)

  - [QueryAllStrategyRequest](#ununifi.chain.yieldaggregator.QueryAllStrategyRequest)
  - [QueryAllStrategyResponse](#ununifi.chain.yieldaggregator.QueryAllStrategyResponse)
  - [QueryAllVaultRequest](#ununifi.chain.yieldaggregator.QueryAllVaultRequest)
  - [QueryAllVaultResponse](#ununifi.chain.yieldaggregator.QueryAllVaultResponse)
  - [QueryGetStrategyRequest](#ununifi.chain.yieldaggregator.QueryGetStrategyRequest)
  - [QueryGetStrategyResponse](#ununifi.chain.yieldaggregator.QueryGetStrategyResponse)
  - [QueryGetVaultRequest](#ununifi.chain.yieldaggregator.QueryGetVaultRequest)
  - [QueryGetVaultResponse](#ununifi.chain.yieldaggregator.QueryGetVaultResponse)
  - [QueryParamsRequest](#ununifi.chain.yieldaggregator.QueryParamsRequest)
  - [QueryParamsResponse](#ununifi.chain.yieldaggregator.QueryParamsResponse)

  - [Query](#ununifi.chain.yieldaggregator.Query)

- [yield-aggregator/tx.proto](#yield-aggregator/tx.proto)

  - [MsgDepositToVault](#ununifi.chain.yieldaggregator.MsgDepositToVault)
  - [MsgDepositToVaultResponse](#ununifi.chain.yieldaggregator.MsgDepositToVaultResponse)
  - [MsgWithdrawFromVault](#ununifi.chain.yieldaggregator.MsgWithdrawFromVault)
  - [MsgWithdrawFromVaultResponse](#ununifi.chain.yieldaggregator.MsgWithdrawFromVaultResponse)

  - [Msg](#ununifi.chain.yieldaggregator.Msg)

- [Scalar Value Types](#scalar-value-types)

  > > > > > > > bb3bfd2c (feat: iya v2)

- [Protobuf Documentation](#protobuf-documentation)
  - [Table of Contents](#table-of-contents)
- [\<\<\<\<\<\<\< HEAD](#-head)
  - [auction/auction.proto](#auctionauctionproto)
    - [BaseAuction](#baseauction)
    - [CollateralAuction](#collateralauction)
    - [DebtAuction](#debtauction)
    - [Params](#params)
    - [SurplusAuction](#surplusauction)
    - [WeightedAddress](#weightedaddress)
  - [auction/genesis.proto](#auctiongenesisproto)
    - [GenesisState](#genesisstate)
  - [auction/query.proto](#auctionqueryproto)
    - [QueryAllAuctionRequest](#queryallauctionrequest)
    - [QueryAllAuctionResponse](#queryallauctionresponse)
    - [QueryGetAuctionRequest](#querygetauctionrequest)
    - [QueryGetAuctionResponse](#querygetauctionresponse)
    - [QueryParamsRequest](#queryparamsrequest)
    - [QueryParamsResponse](#queryparamsresponse)
    - [Query](#query)
  - [auction/tx.proto](#auctiontxproto)
    - [MsgPlaceBid](#msgplacebid)
    - [MsgPlaceBidResponse](#msgplacebidresponse)
    - [Msg](#msg)
  - [cdp/cdp.proto](#cdpcdpproto)
    - [AugmentedCdp](#augmentedcdp)
    - [Cdp](#cdp)
    - [CollateralParam](#collateralparam)
    - [DebtParam](#debtparam)
    - [Deposit](#deposit)
    - [Params](#params-1)
  - [cdp/genesis.proto](#cdpgenesisproto)
    - [GenesisAccumulationTime](#genesisaccumulationtime)
    - [GenesisState](#genesisstate-1)
    - [GenesisTotalPrincipal](#genesistotalprincipal)
  - [cdp/query.proto](#cdpqueryproto)
    - [QueryAllAccountRequest](#queryallaccountrequest)
    - [QueryAllAccountResponse](#queryallaccountresponse)
    - [QueryAllCdpRequest](#queryallcdprequest)
    - [QueryAllCdpResponse](#queryallcdpresponse)
    - [QueryAllDepositRequest](#queryalldepositrequest)
    - [QueryAllDepositResponse](#queryalldepositresponse)
    - [QueryGetCdpRequest](#querygetcdprequest)
    - [QueryGetCdpResponse](#querygetcdpresponse)
    - [QueryParamsRequest](#queryparamsrequest-1)
    - [QueryParamsResponse](#queryparamsresponse-1)
    - [Query](#query-1)
  - [cdp/tx.proto](#cdptxproto)
    - [MsgCreateCdp](#msgcreatecdp)
    - [MsgCreateCdpResponse](#msgcreatecdpresponse)
    - [MsgDeposit](#msgdeposit)
    - [MsgDepositResponse](#msgdepositresponse)
    - [MsgDrawDebt](#msgdrawdebt)
    - [MsgDrawDebtResponse](#msgdrawdebtresponse)
    - [MsgLiquidate](#msgliquidate)
    - [MsgLiquidateResponse](#msgliquidateresponse)
    - [MsgRepayDebt](#msgrepaydebt)
    - [MsgRepayDebtResponse](#msgrepaydebtresponse)
    - [MsgWithdraw](#msgwithdraw)
    - [MsgWithdrawResponse](#msgwithdrawresponse)
    - [Msg](#msg-1)
  - [incentive/incentive.proto](#incentiveincentiveproto)
    - [BaseClaim](#baseclaim)
    - [BaseMultiClaim](#basemulticlaim)
    - [CdpMintingClaim](#cdpmintingclaim)
    - [Multiplier](#multiplier)
    - [Params](#params-2)
    - [RewardIndex](#rewardindex)
    - [RewardPeriod](#rewardperiod)
  - [incentive/genesis.proto](#incentivegenesisproto)
    - [GenesisAccumulationTime](#genesisaccumulationtime-1)
    - [GenesisDenoms](#genesisdenoms)
    - [GenesisState](#genesisstate-2)
  - [incentive/query.proto](#incentivequeryproto)
    - [QueryParamsRequest](#queryparamsrequest-2)
    - [QueryParamsResponse](#queryparamsresponse-2)
    - [Query](#query-2)
  - [incentive/tx.proto](#incentivetxproto)
    - [MsgClaimCdpMintingReward](#msgclaimcdpmintingreward)
    - [MsgClaimCdpMintingRewardResponse](#msgclaimcdpmintingrewardresponse)
    - [Msg](#msg-2)
  - [pricefeed/pricefeed.proto](#pricefeedpricefeedproto)
    - [CurrentPrice](#currentprice)
    - [Market](#market)
    - [Params](#params-3)
    - [PostedPrice](#postedprice)
  - [pricefeed/genesis.proto](#pricefeedgenesisproto)
    - [GenesisState](#genesisstate-3)
  - [pricefeed/query.proto](#pricefeedqueryproto)
    - [QueryAllMarketRequest](#queryallmarketrequest)
    - [QueryAllMarketResponse](#queryallmarketresponse)
    - [QueryAllOracleRequest](#queryalloraclerequest)
    - [QueryAllOracleResponse](#queryalloracleresponse)
    - [QueryAllPriceRequest](#queryallpricerequest)
    - [QueryAllPriceResponse](#queryallpriceresponse)
    - [QueryAllRawPriceRequest](#queryallrawpricerequest)
    - [QueryAllRawPriceResponse](#queryallrawpriceresponse)
    - [QueryGetPriceRequest](#querygetpricerequest)
    - [QueryGetPriceResponse](#querygetpriceresponse)
    - [QueryParamsRequest](#queryparamsrequest-3)
    - [QueryParamsResponse](#queryparamsresponse-3)
    - [Query](#query-3)
  - [pricefeed/tx.proto](#pricefeedtxproto)
    - [MsgPostPrice](#msgpostprice)
    - [MsgPostPriceResponse](#msgpostpriceresponse)
    - [Msg](#msg-3)
  - [ununifidist/ununifidist.proto](#ununifidistununifidistproto)
    - [Params](#params-4)
    - [Period](#period)
  - [ununifidist/genesis.proto](#ununifidistgenesisproto)
    - [GenesisState](#genesisstate-4)
  - [ununifidist/query.proto](#ununifidistqueryproto)
    - [QueryGetBalancesRequest](#querygetbalancesrequest)
    - [QueryGetBalancesResponse](#querygetbalancesresponse)
    - [QueryParamsRequest](#queryparamsrequest-4)
    - [QueryParamsResponse](#queryparamsresponse-4)
    - [Query](#query-4)
- [\<\<\<\<\<\<\< HEAD](#-head-1)
  - [yield-aggregator/params.proto](#yield-aggregatorparamsproto)
    - [Params](#params-5)
    - [Params.Vault](#paramsvault)
  - [yield-aggregator/yield-aggregator.proto](#yield-aggregatoryield-aggregatorproto)
    - [Strategy](#strategy)
- [| `metrics` | StrategyMetrics | | |](#-metrics-----------strategymetrics---------------------)
  - [StrategyMetrics](#strategymetrics)
  - [Vault](#vault)
  - [VaultMetrics](#vaultmetrics)
  - [yield-aggregator/genesis.proto](#yield-aggregatorgenesisproto)
    - [GenesisState](#genesisstate-5)
  - [yield-aggregator/query.proto](#yield-aggregatorqueryproto)
- [\<\<\<\<\<\<\< HEAD](#-head-2)
  - [QueryAllStrategyRequest](#queryallstrategyrequest)
  - [QueryAllStrategyResponse](#queryallstrategyresponse)
  - [QueryAllVaultRequest](#queryallvaultrequest)
  - [QueryAllVaultResponse](#queryallvaultresponse)
  - [QueryGetStrategyRequest](#querygetstrategyrequest)
  - [QueryGetStrategyResponse](#querygetstrategyresponse)
  - [QueryGetVaultRequest](#querygetvaultrequest)
  - [QueryGetVaultResponse](#querygetvaultresponse)
  - [QueryParamsRequest](#queryparamsrequest-5)
  - [QueryParamsResponse](#queryparamsresponse-5)
  - [QueryStrategiesRequest](#querystrategiesrequest)
  - [QueryStrategiesResponse](#querystrategiesresponse)
  - [QueryVaultRequest](#queryvaultrequest)
  - [QueryVaultResponse](#queryvaultresponse)
  - [QueryVaultsRequest](#queryvaultsrequest)
  - [QueryVaultsResponse](#queryvaultsresponse)
- [| `pagination` | cosmos.base.query.v1beta1.PageRequest | | |](#-pagination--cosmosbasequeryv1beta1pagerequest------------------------)
  - [Query](#query-5)
- [| `Strategies` | QueryStrategiesRequest | QueryStrategiesResponse | | GET | /UnUniFi/chain/yield-aggregator/vaults/{denom}/strategies |](#-strategies--querystrategiesrequest--querystrategiesresponse----------------------------------------------------get--------ununifichainyield-aggregatorvaultsdenomstrategies-)
  - [yield-aggregator/tx.proto](#yield-aggregatortxproto)
    - [MsgDepositToVault](#msgdeposittovault)
    - [MsgDepositToVaultResponse](#msgdeposittovaultresponse)
    - [MsgWithdrawFromVault](#msgwithdrawfromvault)
- [| `lp_token_amount` | string | | |](#-lp_token_amount--string---------------------)
  - [MsgWithdrawFromVaultResponse](#msgwithdrawfromvaultresponse)
  - [Msg](#msg-4)
  - [Scalar Value Types](#scalar-value-types)

> > > > > > > b9a58125 (feat: keeper)

<a name="auction/auction.proto"></a>

<p align="right"><a href="#top">Top</a></p>

## auction/auction.proto

<a name="ununifi.auction.BaseAuction"></a>

### BaseAuction

| Field               | Type                                                    | Label | Description |
| ------------------- | ------------------------------------------------------- | ----- | ----------- |
| `id`                | [uint64](#uint64)                                       |       |             |
| `initiator`         | [string](#string)                                       |       |             |
| `lot`               | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin)   |       |             |
| `bidder`            | [string](#string)                                       |       |             |
| `bid`               | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin)   |       |             |
| `has_received_bids` | [bool](#bool)                                           |       |             |
| `end_time`          | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |       |             |
| `max_end_time`      | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |       |             |

<a name="ununifi.auction.CollateralAuction"></a>

### CollateralAuction

| Field                | Type                                                  | Label    | Description |
| -------------------- | ----------------------------------------------------- | -------- | ----------- |
| `base_auction`       | [BaseAuction](#ununifi.auction.BaseAuction)           |          |             |
| `corresponding_debt` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |          |             |
| `max_bid`            | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |          |             |
| `lot_returns`        | [WeightedAddress](#ununifi.auction.WeightedAddress)   | repeated |             |

<a name="ununifi.auction.DebtAuction"></a>

### DebtAuction

| Field                | Type                                                  | Label | Description |
| -------------------- | ----------------------------------------------------- | ----- | ----------- |
| `base_auction`       | [BaseAuction](#ununifi.auction.BaseAuction)           |       |             |
| `corresponding_debt` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |       |             |

<a name="ununifi.auction.Params"></a>

### Params

| Field                  | Type                                                  | Label | Description |
| ---------------------- | ----------------------------------------------------- | ----- | ----------- |
| `max_auction_duration` | [google.protobuf.Duration](#google.protobuf.Duration) |       |             |
| `bid_duration`         | [google.protobuf.Duration](#google.protobuf.Duration) |       |             |
| `increment_surplus`    | [string](#string)                                     |       |             |
| `increment_debt`       | [string](#string)                                     |       |             |
| `increment_collateral` | [string](#string)                                     |       |             |

<a name="ununifi.auction.SurplusAuction"></a>

### SurplusAuction

| Field          | Type                                        | Label | Description |
| -------------- | ------------------------------------------- | ----- | ----------- |
| `base_auction` | [BaseAuction](#ununifi.auction.BaseAuction) |       |             |

<a name="ununifi.auction.WeightedAddress"></a>

### WeightedAddress

| Field     | Type              | Label | Description |
| --------- | ----------------- | ----- | ----------- |
| `address` | [string](#string) |       |             |
| `weight`  | [string](#string) |       |             |

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

| Field             | Type                                        | Label    | Description                                                     |
| ----------------- | ------------------------------------------- | -------- | --------------------------------------------------------------- |
| `next_auction_id` | [uint64](#uint64)                           |          |                                                                 |
| `params`          | [Params](#ununifi.auction.Params)           |          |                                                                 |
| `auctions`        | [google.protobuf.Any](#google.protobuf.Any) | repeated | this line is used by starport scaffolding # genesis/proto/state |

 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->

<a name="auction/query.proto"></a>

<p align="right"><a href="#top">Top</a></p>

## auction/query.proto

<a name="ununifi.auction.QueryAllAuctionRequest"></a>

### QueryAllAuctionRequest

| Field        | Type                                                                            | Label | Description |
| ------------ | ------------------------------------------------------------------------------- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |       |             |

<a name="ununifi.auction.QueryAllAuctionResponse"></a>

### QueryAllAuctionResponse

| Field        | Type                                                                              | Label    | Description |
| ------------ | --------------------------------------------------------------------------------- | -------- | ----------- |
| `auctions`   | [google.protobuf.Any](#google.protobuf.Any)                                       | repeated |             |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |          |             |

<a name="ununifi.auction.QueryGetAuctionRequest"></a>

### QueryGetAuctionRequest

this line is used by starport scaffolding # 3

| Field | Type              | Label | Description |
| ----- | ----------------- | ----- | ----------- |
| `id`  | [uint64](#uint64) |       |             |

<a name="ununifi.auction.QueryGetAuctionResponse"></a>

### QueryGetAuctionResponse

| Field     | Type                                        | Label | Description |
| --------- | ------------------------------------------- | ----- | ----------- |
| `auction` | [google.protobuf.Any](#google.protobuf.Any) |       |             |

<a name="ununifi.auction.QueryParamsRequest"></a>

### QueryParamsRequest

<a name="ununifi.auction.QueryParamsResponse"></a>

### QueryParamsResponse

| Field    | Type                              | Label | Description |
| -------- | --------------------------------- | ----- | ----------- |
| `params` | [Params](#ununifi.auction.Params) |       |             |

 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

<a name="ununifi.auction.Query"></a>

### Query

Query defines the gRPC querier service.

| Method Name  | Request Type                                                      | Response Type                                                       | Description                                   | HTTP Verb | Endpoint                       |
| ------------ | ----------------------------------------------------------------- | ------------------------------------------------------------------- | --------------------------------------------- | --------- | ------------------------------ |
| `Params`     | [QueryParamsRequest](#ununifi.auction.QueryParamsRequest)         | [QueryParamsResponse](#ununifi.auction.QueryParamsResponse)         |                                               | GET       | /ununifi/auction/params        |
| `Auction`    | [QueryGetAuctionRequest](#ununifi.auction.QueryGetAuctionRequest) | [QueryGetAuctionResponse](#ununifi.auction.QueryGetAuctionResponse) | this line is used by starport scaffolding # 2 | GET       | /ununifi/auction/auctions/{id} |
| `AuctionAll` | [QueryAllAuctionRequest](#ununifi.auction.QueryAllAuctionRequest) | [QueryAllAuctionResponse](#ununifi.auction.QueryAllAuctionResponse) |                                               | GET       | /ununifi/auction/auctions      |

 <!-- end services -->

<a name="auction/tx.proto"></a>

<p align="right"><a href="#top">Top</a></p>

## auction/tx.proto

<a name="ununifi.auction.MsgPlaceBid"></a>

### MsgPlaceBid

| Field        | Type                                                  | Label | Description |
| ------------ | ----------------------------------------------------- | ----- | ----------- |
| `auction_id` | [uint64](#uint64)                                     |       |             |
| `bidder`     | [string](#string)                                     |       |             |
| `amount`     | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |       |             |

<a name="ununifi.auction.MsgPlaceBidResponse"></a>

### MsgPlaceBidResponse

 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

<a name="ununifi.auction.Msg"></a>

### Msg

| Method Name | Request Type                                | Response Type                                               | Description | HTTP Verb | Endpoint |
| ----------- | ------------------------------------------- | ----------------------------------------------------------- | ----------- | --------- | -------- |
| `PlaceBid`  | [MsgPlaceBid](#ununifi.auction.MsgPlaceBid) | [MsgPlaceBidResponse](#ununifi.auction.MsgPlaceBidResponse) |             |           |

 <!-- end services -->

<a name="cdp/cdp.proto"></a>

<p align="right"><a href="#top">Top</a></p>

## cdp/cdp.proto

<a name="ununifi.cdp.AugmentedCdp"></a>

### AugmentedCdp

| Field                     | Type                                                  | Label | Description |
| ------------------------- | ----------------------------------------------------- | ----- | ----------- |
| `cdp`                     | [Cdp](#ununifi.cdp.Cdp)                               |       |             |
| `collateral_value`        | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |       |             |
| `collateralization_ratio` | [string](#string)                                     |       |             |

<a name="ununifi.cdp.Cdp"></a>

### Cdp

| Field              | Type                                                    | Label | Description |
| ------------------ | ------------------------------------------------------- | ----- | ----------- |
| `id`               | [uint64](#uint64)                                       |       |             |
| `owner`            | [string](#string)                                       |       |             |
| `type`             | [string](#string)                                       |       |             |
| `collateral`       | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin)   |       |             |
| `principal`        | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin)   |       |             |
| `accumulated_fees` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin)   |       |             |
| `fees_updated`     | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |       |             |
| `interest_factor`  | [string](#string)                                       |       |             |

<a name="ununifi.cdp.CollateralParam"></a>

### CollateralParam

| Field                                 | Type                                                  | Label | Description |
| ------------------------------------- | ----------------------------------------------------- | ----- | ----------- |
| `denom`                               | [string](#string)                                     |       |             |
| `type`                                | [string](#string)                                     |       |             |
| `liquidation_ratio`                   | [string](#string)                                     |       |             |
| `debt_limit`                          | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |       |             |
| `stability_fee`                       | [string](#string)                                     |       |             |
| `auction_size`                        | [string](#string)                                     |       |             |
| `liquidation_penalty`                 | [string](#string)                                     |       |             |
| `prefix`                              | [uint32](#uint32)                                     |       |             |
| `spot_market_id`                      | [string](#string)                                     |       |             |
| `liquidation_market_id`               | [string](#string)                                     |       |             |
| `keeper_reward_percentage`            | [string](#string)                                     |       |             |
| `check_collateralization_index_count` | [string](#string)                                     |       |             |
| `conversion_factor`                   | [string](#string)                                     |       |             |

<a name="ununifi.cdp.DebtParam"></a>

### DebtParam

| Field                       | Type                                                  | Label | Description |
| --------------------------- | ----------------------------------------------------- | ----- | ----------- |
| `denom`                     | [string](#string)                                     |       |             |
| `reference_asset`           | [string](#string)                                     |       |             |
| `conversion_factor`         | [string](#string)                                     |       |             |
| `debt_floor`                | [string](#string)                                     |       |             |
| `global_debt_limit`         | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |       |             |
| `debt_denom`                | [string](#string)                                     |       |             |
| `surplus_auction_threshold` | [string](#string)                                     |       |             |
| `surplus_auction_lot`       | [string](#string)                                     |       |             |
| `debt_auction_threshold`    | [string](#string)                                     |       |             |
| `debt_auction_lot`          | [string](#string)                                     |       |             |
| `circuit_breaker`           | [bool](#bool)                                         |       |             |

<a name="ununifi.cdp.Deposit"></a>

### Deposit

| Field       | Type                                                  | Label | Description |
| ----------- | ----------------------------------------------------- | ----- | ----------- |
| `cdp_id`    | [uint64](#uint64)                                     |       |             |
| `depositor` | [string](#string)                                     |       |             |
| `amount`    | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |       |             |

<a name="ununifi.cdp.Params"></a>

### Params

| Field               | Type                                            | Label    | Description |
| ------------------- | ----------------------------------------------- | -------- | ----------- |
| `collateral_params` | [CollateralParam](#ununifi.cdp.CollateralParam) | repeated |             |
| `debt_params`       | [DebtParam](#ununifi.cdp.DebtParam)             | repeated |             |

 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->

<a name="cdp/genesis.proto"></a>

<p align="right"><a href="#top">Top</a></p>

## cdp/genesis.proto

<a name="ununifi.cdp.GenesisAccumulationTime"></a>

### GenesisAccumulationTime

| Field                        | Type                                                    | Label | Description |
| ---------------------------- | ------------------------------------------------------- | ----- | ----------- |
| `collateral_type`            | [string](#string)                                       |       |             |
| `previous_accumulation_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |       |             |
| `interest_factor`            | [string](#string)                                       |       |             |

<a name="ununifi.cdp.GenesisState"></a>

### GenesisState

GenesisState defines the cdp module's genesis state.

| Field                         | Type                                                            | Label    | Description                                                     |
| ----------------------------- | --------------------------------------------------------------- | -------- | --------------------------------------------------------------- |
| `params`                      | [Params](#ununifi.cdp.Params)                                   |          |                                                                 |
| `cdps`                        | [Cdp](#ununifi.cdp.Cdp)                                         | repeated |                                                                 |
| `deposits`                    | [Deposit](#ununifi.cdp.Deposit)                                 | repeated |                                                                 |
| `starting_cdp_id`             | [uint64](#uint64)                                               |          |                                                                 |
| `gov_denom`                   | [string](#string)                                               |          |                                                                 |
| `previous_accumulation_times` | [GenesisAccumulationTime](#ununifi.cdp.GenesisAccumulationTime) | repeated |                                                                 |
| `total_principals`            | [GenesisTotalPrincipal](#ununifi.cdp.GenesisTotalPrincipal)     | repeated | this line is used by starport scaffolding # genesis/proto/state |

<a name="ununifi.cdp.GenesisTotalPrincipal"></a>

### GenesisTotalPrincipal

| Field             | Type              | Label | Description |
| ----------------- | ----------------- | ----- | ----------- |
| `collateral_type` | [string](#string) |       |             |
| `total_principal` | [string](#string) |       |             |

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

| Field      | Type                                        | Label    | Description |
| ---------- | ------------------------------------------- | -------- | ----------- |
| `accounts` | [google.protobuf.Any](#google.protobuf.Any) | repeated |             |

<a name="ununifi.cdp.QueryAllCdpRequest"></a>

### QueryAllCdpRequest

| Field        | Type                                                                            | Label | Description |
| ------------ | ------------------------------------------------------------------------------- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |       |             |

<a name="ununifi.cdp.QueryAllCdpResponse"></a>

### QueryAllCdpResponse

| Field        | Type                                                                              | Label    | Description |
| ------------ | --------------------------------------------------------------------------------- | -------- | ----------- |
| `cdp`        | [AugmentedCdp](#ununifi.cdp.AugmentedCdp)                                         | repeated |             |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |          |             |

<a name="ununifi.cdp.QueryAllDepositRequest"></a>

### QueryAllDepositRequest

| Field             | Type              | Label | Description |
| ----------------- | ----------------- | ----- | ----------- |
| `owner`           | [string](#string) |       |             |
| `collateral_type` | [string](#string) |       |             |

<a name="ununifi.cdp.QueryAllDepositResponse"></a>

### QueryAllDepositResponse

| Field      | Type                            | Label    | Description |
| ---------- | ------------------------------- | -------- | ----------- |
| `deposits` | [Deposit](#ununifi.cdp.Deposit) | repeated |             |

<a name="ununifi.cdp.QueryGetCdpRequest"></a>

### QueryGetCdpRequest

this line is used by starport scaffolding # 3

| Field             | Type              | Label | Description |
| ----------------- | ----------------- | ----- | ----------- |
| `owner`           | [string](#string) |       |             |
| `collateral_type` | [string](#string) |       |             |

<a name="ununifi.cdp.QueryGetCdpResponse"></a>

### QueryGetCdpResponse

| Field | Type                                      | Label | Description |
| ----- | ----------------------------------------- | ----- | ----------- |
| `cdp` | [AugmentedCdp](#ununifi.cdp.AugmentedCdp) |       |             |

<a name="ununifi.cdp.QueryParamsRequest"></a>

### QueryParamsRequest

<a name="ununifi.cdp.QueryParamsResponse"></a>

### QueryParamsResponse

| Field    | Type                          | Label | Description |
| -------- | ----------------------------- | ----- | ----------- |
| `params` | [Params](#ununifi.cdp.Params) |       |             |

 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

<a name="ununifi.cdp.Query"></a>

### Query

Query defines the gRPC querier service.

| Method Name  | Request Type                                                  | Response Type                                                   | Description                                   | HTTP Verb | Endpoint                                                                |
| ------------ | ------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------- | --------- | ----------------------------------------------------------------------- |
| `Params`     | [QueryParamsRequest](#ununifi.cdp.QueryParamsRequest)         | [QueryParamsResponse](#ununifi.cdp.QueryParamsResponse)         |                                               | GET       | /ununifi/cdp/params                                                     |
| `Cdp`        | [QueryGetCdpRequest](#ununifi.cdp.QueryGetCdpRequest)         | [QueryGetCdpResponse](#ununifi.cdp.QueryGetCdpResponse)         | this line is used by starport scaffolding # 2 | GET       | /ununifi/cdp/cdps/owners/{owner}/collateral-types/{collateral_type}/cdp |
| `CdpAll`     | [QueryAllCdpRequest](#ununifi.cdp.QueryAllCdpRequest)         | [QueryAllCdpResponse](#ununifi.cdp.QueryAllCdpResponse)         |                                               | GET       | /ununifi/cdp/cdps                                                       |
| `AccountAll` | [QueryAllAccountRequest](#ununifi.cdp.QueryAllAccountRequest) | [QueryAllAccountResponse](#ununifi.cdp.QueryAllAccountResponse) |                                               | GET       | /ununifi/cdp/accounts                                                   |
| `DepositAll` | [QueryAllDepositRequest](#ununifi.cdp.QueryAllDepositRequest) | [QueryAllDepositResponse](#ununifi.cdp.QueryAllDepositResponse) |                                               | GET       | /ununifi/cdp/deposits/owners/{owner}/collateral-types/{collateral_type} |

 <!-- end services -->

<a name="cdp/tx.proto"></a>

<p align="right"><a href="#top">Top</a></p>

## cdp/tx.proto

<a name="ununifi.cdp.MsgCreateCdp"></a>

### MsgCreateCdp

| Field             | Type                                                  | Label | Description |
| ----------------- | ----------------------------------------------------- | ----- | ----------- |
| `sender`          | [string](#string)                                     |       |             |
| `collateral`      | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |       |             |
| `principal`       | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |       |             |
| `collateral_type` | [string](#string)                                     |       |             |

<a name="ununifi.cdp.MsgCreateCdpResponse"></a>

### MsgCreateCdpResponse

<a name="ununifi.cdp.MsgDeposit"></a>

### MsgDeposit

| Field             | Type                                                  | Label | Description |
| ----------------- | ----------------------------------------------------- | ----- | ----------- |
| `depositor`       | [string](#string)                                     |       |             |
| `owner`           | [string](#string)                                     |       |             |
| `collateral`      | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |       |             |
| `collateral_type` | [string](#string)                                     |       |             |

<a name="ununifi.cdp.MsgDepositResponse"></a>

### MsgDepositResponse

<a name="ununifi.cdp.MsgDrawDebt"></a>

### MsgDrawDebt

| Field             | Type                                                  | Label | Description |
| ----------------- | ----------------------------------------------------- | ----- | ----------- |
| `sender`          | [string](#string)                                     |       |             |
| `collateral_type` | [string](#string)                                     |       |             |
| `principal`       | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |       |             |

<a name="ununifi.cdp.MsgDrawDebtResponse"></a>

### MsgDrawDebtResponse

<a name="ununifi.cdp.MsgLiquidate"></a>

### MsgLiquidate

| Field             | Type              | Label | Description |
| ----------------- | ----------------- | ----- | ----------- |
| `keeper`          | [string](#string) |       |             |
| `borrower`        | [string](#string) |       |             |
| `collateral_type` | [string](#string) |       |             |

<a name="ununifi.cdp.MsgLiquidateResponse"></a>

### MsgLiquidateResponse

<a name="ununifi.cdp.MsgRepayDebt"></a>

### MsgRepayDebt

| Field             | Type                                                  | Label | Description |
| ----------------- | ----------------------------------------------------- | ----- | ----------- |
| `sender`          | [string](#string)                                     |       |             |
| `collateral_type` | [string](#string)                                     |       |             |
| `payment`         | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |       |             |

<a name="ununifi.cdp.MsgRepayDebtResponse"></a>

### MsgRepayDebtResponse

<a name="ununifi.cdp.MsgWithdraw"></a>

### MsgWithdraw

| Field             | Type                                                  | Label | Description |
| ----------------- | ----------------------------------------------------- | ----- | ----------- |
| `depositor`       | [string](#string)                                     |       |             |
| `owner`           | [string](#string)                                     |       |             |
| `collateral`      | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |       |             |
| `collateral_type` | [string](#string)                                     |       |             |

<a name="ununifi.cdp.MsgWithdrawResponse"></a>

### MsgWithdrawResponse

 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

<a name="ununifi.cdp.Msg"></a>

### Msg

| Method Name | Request Type                              | Response Type                                             | Description | HTTP Verb | Endpoint |
| ----------- | ----------------------------------------- | --------------------------------------------------------- | ----------- | --------- | -------- |
| `CreateCdp` | [MsgCreateCdp](#ununifi.cdp.MsgCreateCdp) | [MsgCreateCdpResponse](#ununifi.cdp.MsgCreateCdpResponse) |             |           |
| `Deposit`   | [MsgDeposit](#ununifi.cdp.MsgDeposit)     | [MsgDepositResponse](#ununifi.cdp.MsgDepositResponse)     |             |           |
| `Withdraw`  | [MsgWithdraw](#ununifi.cdp.MsgWithdraw)   | [MsgWithdrawResponse](#ununifi.cdp.MsgWithdrawResponse)   |             |           |
| `DrawDebt`  | [MsgDrawDebt](#ununifi.cdp.MsgDrawDebt)   | [MsgDrawDebtResponse](#ununifi.cdp.MsgDrawDebtResponse)   |             |           |
| `RepayDebt` | [MsgRepayDebt](#ununifi.cdp.MsgRepayDebt) | [MsgRepayDebtResponse](#ununifi.cdp.MsgRepayDebtResponse) |             |           |
| `Liquidate` | [MsgLiquidate](#ununifi.cdp.MsgLiquidate) | [MsgLiquidateResponse](#ununifi.cdp.MsgLiquidateResponse) |             |           |

 <!-- end services -->

<a name="incentive/incentive.proto"></a>

<p align="right"><a href="#top">Top</a></p>

## incentive/incentive.proto

<a name="ununifi.incentive.BaseClaim"></a>

### BaseClaim

| Field    | Type                                                  | Label | Description |
| -------- | ----------------------------------------------------- | ----- | ----------- |
| `owner`  | [string](#string)                                     |       |             |
| `reward` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |       |             |

<a name="ununifi.incentive.BaseMultiClaim"></a>

### BaseMultiClaim

| Field    | Type                                                  | Label    | Description |
| -------- | ----------------------------------------------------- | -------- | ----------- |
| `owner`  | [string](#string)                                     |          |             |
| `reward` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |             |

<a name="ununifi.incentive.CdpMintingClaim"></a>

### CdpMintingClaim

| Field            | Type                                          | Label    | Description |
| ---------------- | --------------------------------------------- | -------- | ----------- |
| `base_claim`     | [BaseClaim](#ununifi.incentive.BaseClaim)     |          |             |
| `reward_indexes` | [RewardIndex](#ununifi.incentive.RewardIndex) | repeated |             |

<a name="ununifi.incentive.Multiplier"></a>

### Multiplier

| Field           | Type              | Label | Description |
| --------------- | ----------------- | ----- | ----------- |
| `name`          | [string](#string) |       |             |
| `months_lockup` | [int64](#int64)   |       |             |
| `factor`        | [string](#string) |       |             |

<a name="ununifi.incentive.Params"></a>

### Params

| Field                        | Type                                                    | Label    | Description |
| ---------------------------- | ------------------------------------------------------- | -------- | ----------- |
| `cdp_minting_reward_periods` | [RewardPeriod](#ununifi.incentive.RewardPeriod)         | repeated |             |
| `claim_multipliers`          | [Multiplier](#ununifi.incentive.Multiplier)             | repeated |             |
| `claim_end`                  | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |          |             |

<a name="ununifi.incentive.RewardIndex"></a>

### RewardIndex

| Field             | Type              | Label | Description |
| ----------------- | ----------------- | ----- | ----------- |
| `collateral_type` | [string](#string) |       |             |
| `reward_factor`   | [string](#string) |       |             |

<a name="ununifi.incentive.RewardPeriod"></a>

### RewardPeriod

| Field                | Type                                                    | Label | Description |
| -------------------- | ------------------------------------------------------- | ----- | ----------- |
| `active`             | [bool](#bool)                                           |       |             |
| `collateral_type`    | [string](#string)                                       |       |             |
| `start`              | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |       |             |
| `end`                | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |       |             |
| `rewards_per_second` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin)   |       |             |

 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->

<a name="incentive/genesis.proto"></a>

<p align="right"><a href="#top">Top</a></p>

## incentive/genesis.proto

<a name="ununifi.incentive.GenesisAccumulationTime"></a>

### GenesisAccumulationTime

| Field                        | Type                                                    | Label | Description |
| ---------------------------- | ------------------------------------------------------- | ----- | ----------- |
| `collateral_type`            | [string](#string)                                       |       |             |
| `previous_accumulation_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |       |             |

<a name="ununifi.incentive.GenesisDenoms"></a>

### GenesisDenoms

| Field                      | Type              | Label | Description |
| -------------------------- | ----------------- | ----- | ----------- |
| `principal_denom`          | [string](#string) |       |             |
| `cdp_minting_reward_denom` | [string](#string) |       |             |

<a name="ununifi.incentive.GenesisState"></a>

### GenesisState

GenesisState defines the incentive module's genesis state.

| Field                    | Type                                                                  | Label    | Description                                                     |
| ------------------------ | --------------------------------------------------------------------- | -------- | --------------------------------------------------------------- |
| `params`                 | [Params](#ununifi.incentive.Params)                                   |          |                                                                 |
| `cdp_accumulation_times` | [GenesisAccumulationTime](#ununifi.incentive.GenesisAccumulationTime) | repeated |                                                                 |
| `cdp_minting_claims`     | [CdpMintingClaim](#ununifi.incentive.CdpMintingClaim)                 | repeated |                                                                 |
| `denoms`                 | [GenesisDenoms](#ununifi.incentive.GenesisDenoms)                     |          | this line is used by starport scaffolding # genesis/proto/state |

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

| Field    | Type                                | Label | Description |
| -------- | ----------------------------------- | ----- | ----------- |
| `params` | [Params](#ununifi.incentive.Params) |       |             |

 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

<a name="ununifi.incentive.Query"></a>

### Query

Query defines the gRPC querier service.

| Method Name | Request Type                                                | Response Type                                                 | Description                                   | HTTP Verb | Endpoint                  |
| ----------- | ----------------------------------------------------------- | ------------------------------------------------------------- | --------------------------------------------- | --------- | ------------------------- |
| `Params`    | [QueryParamsRequest](#ununifi.incentive.QueryParamsRequest) | [QueryParamsResponse](#ununifi.incentive.QueryParamsResponse) | this line is used by starport scaffolding # 2 | GET       | /ununifi/incentive/params |

 <!-- end services -->

<a name="incentive/tx.proto"></a>

<p align="right"><a href="#top">Top</a></p>

## incentive/tx.proto

<a name="ununifi.incentive.MsgClaimCdpMintingReward"></a>

### MsgClaimCdpMintingReward

| Field             | Type              | Label | Description |
| ----------------- | ----------------- | ----- | ----------- |
| `sender`          | [string](#string) |       |             |
| `multiplier_name` | [string](#string) |       |             |

<a name="ununifi.incentive.MsgClaimCdpMintingRewardResponse"></a>

### MsgClaimCdpMintingRewardResponse

 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

<a name="ununifi.incentive.Msg"></a>

### Msg

| Method Name             | Request Type                                                            | Response Type                                                                           | Description | HTTP Verb | Endpoint |
| ----------------------- | ----------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | ----------- | --------- | -------- |
| `ClaimCdpMintingReward` | [MsgClaimCdpMintingReward](#ununifi.incentive.MsgClaimCdpMintingReward) | [MsgClaimCdpMintingRewardResponse](#ununifi.incentive.MsgClaimCdpMintingRewardResponse) |             |           |

 <!-- end services -->

<a name="pricefeed/pricefeed.proto"></a>

<p align="right"><a href="#top">Top</a></p>

## pricefeed/pricefeed.proto

<a name="ununifi.pricefeed.CurrentPrice"></a>

### CurrentPrice

| Field       | Type              | Label | Description |
| ----------- | ----------------- | ----- | ----------- |
| `market_id` | [string](#string) |       |             |
| `price`     | [string](#string) |       |             |

<a name="ununifi.pricefeed.Market"></a>

### Market

| Field         | Type              | Label    | Description |
| ------------- | ----------------- | -------- | ----------- |
| `market_id`   | [string](#string) |          |             |
| `base_asset`  | [string](#string) |          |             |
| `quote_asset` | [string](#string) |          |             |
| `oracles`     | [string](#string) | repeated |             |
| `active`      | [bool](#bool)     |          |             |

<a name="ununifi.pricefeed.Params"></a>

### Params

| Field     | Type                                | Label    | Description |
| --------- | ----------------------------------- | -------- | ----------- |
| `markets` | [Market](#ununifi.pricefeed.Market) | repeated |             |

<a name="ununifi.pricefeed.PostedPrice"></a>

### PostedPrice

| Field            | Type                                                    | Label | Description |
| ---------------- | ------------------------------------------------------- | ----- | ----------- |
| `market_id`      | [string](#string)                                       |       |             |
| `oracle_address` | [string](#string)                                       |       |             |
| `price`          | [string](#string)                                       |       |             |
| `expiry`         | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |       |             |

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

| Field           | Type                                          | Label    | Description                                                     |
| --------------- | --------------------------------------------- | -------- | --------------------------------------------------------------- |
| `params`        | [Params](#ununifi.pricefeed.Params)           |          |                                                                 |
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

| Field        | Type                                                                            | Label | Description |
| ------------ | ------------------------------------------------------------------------------- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |       |             |

<a name="ununifi.pricefeed.QueryAllMarketResponse"></a>

### QueryAllMarketResponse

| Field        | Type                                                                              | Label    | Description |
| ------------ | --------------------------------------------------------------------------------- | -------- | ----------- |
| `markets`    | [Market](#ununifi.pricefeed.Market)                                               | repeated |             |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |          |             |

<a name="ununifi.pricefeed.QueryAllOracleRequest"></a>

### QueryAllOracleRequest

| Field        | Type                                                                            | Label | Description |
| ------------ | ------------------------------------------------------------------------------- | ----- | ----------- |
| `market_id`  | [string](#string)                                                               |       |             |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |       |             |

<a name="ununifi.pricefeed.QueryAllOracleResponse"></a>

### QueryAllOracleResponse

| Field        | Type                                                                              | Label    | Description |
| ------------ | --------------------------------------------------------------------------------- | -------- | ----------- |
| `oracles`    | [string](#string)                                                                 | repeated |             |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |          |             |

<a name="ununifi.pricefeed.QueryAllPriceRequest"></a>

### QueryAllPriceRequest

| Field        | Type                                                                            | Label | Description |
| ------------ | ------------------------------------------------------------------------------- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |       |             |

<a name="ununifi.pricefeed.QueryAllPriceResponse"></a>

### QueryAllPriceResponse

| Field        | Type                                                                              | Label    | Description |
| ------------ | --------------------------------------------------------------------------------- | -------- | ----------- |
| `prices`     | [CurrentPrice](#ununifi.pricefeed.CurrentPrice)                                   | repeated |             |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |          |             |

<a name="ununifi.pricefeed.QueryAllRawPriceRequest"></a>

### QueryAllRawPriceRequest

| Field        | Type                                                                            | Label | Description |
| ------------ | ------------------------------------------------------------------------------- | ----- | ----------- |
| `market_id`  | [string](#string)                                                               |       |             |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |       |             |

<a name="ununifi.pricefeed.QueryAllRawPriceResponse"></a>

### QueryAllRawPriceResponse

| Field        | Type                                                                              | Label    | Description |
| ------------ | --------------------------------------------------------------------------------- | -------- | ----------- |
| `prices`     | [PostedPrice](#ununifi.pricefeed.PostedPrice)                                     | repeated |             |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |          |             |

<a name="ununifi.pricefeed.QueryGetPriceRequest"></a>

### QueryGetPriceRequest

| Field       | Type              | Label | Description |
| ----------- | ----------------- | ----- | ----------- |
| `market_id` | [string](#string) |       |             |

<a name="ununifi.pricefeed.QueryGetPriceResponse"></a>

### QueryGetPriceResponse

| Field   | Type                                            | Label | Description |
| ------- | ----------------------------------------------- | ----- | ----------- |
| `price` | [CurrentPrice](#ununifi.pricefeed.CurrentPrice) |       |             |

<a name="ununifi.pricefeed.QueryParamsRequest"></a>

### QueryParamsRequest

<a name="ununifi.pricefeed.QueryParamsResponse"></a>

### QueryParamsResponse

| Field    | Type                                | Label | Description |
| -------- | ----------------------------------- | ----- | ----------- |
| `params` | [Params](#ununifi.pricefeed.Params) |       |             |

 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

<a name="ununifi.pricefeed.Query"></a>

### Query

Query defines the gRPC querier service.

| Method Name   | Request Type                                                          | Response Type                                                           | Description                                   | HTTP Verb | Endpoint                                          |
| ------------- | --------------------------------------------------------------------- | ----------------------------------------------------------------------- | --------------------------------------------- | --------- | ------------------------------------------------- |
| `Params`      | [QueryParamsRequest](#ununifi.pricefeed.QueryParamsRequest)           | [QueryParamsResponse](#ununifi.pricefeed.QueryParamsResponse)           |                                               | GET       | /ununifi/pricefeed/params                         |
| `MarketAll`   | [QueryAllMarketRequest](#ununifi.pricefeed.QueryAllMarketRequest)     | [QueryAllMarketResponse](#ununifi.pricefeed.QueryAllMarketResponse)     | this line is used by starport scaffolding # 2 | GET       | /ununifi/pricefeed/markets                        |
| `OracleAll`   | [QueryAllOracleRequest](#ununifi.pricefeed.QueryAllOracleRequest)     | [QueryAllOracleResponse](#ununifi.pricefeed.QueryAllOracleResponse)     |                                               | GET       | /ununifi/pricefeed/markets/{market_id}/oracles    |
| `Price`       | [QueryGetPriceRequest](#ununifi.pricefeed.QueryGetPriceRequest)       | [QueryGetPriceResponse](#ununifi.pricefeed.QueryGetPriceResponse)       |                                               | GET       | /ununifi/pricefeed/markets/{market_id}/price      |
| `PriceAll`    | [QueryAllPriceRequest](#ununifi.pricefeed.QueryAllPriceRequest)       | [QueryAllPriceResponse](#ununifi.pricefeed.QueryAllPriceResponse)       |                                               | GET       | /ununifi/pricefeed/prices                         |
| `RawPriceAll` | [QueryAllRawPriceRequest](#ununifi.pricefeed.QueryAllRawPriceRequest) | [QueryAllRawPriceResponse](#ununifi.pricefeed.QueryAllRawPriceResponse) |                                               | GET       | /ununifi/pricefeed/markets/{market_id}/raw_prices |

 <!-- end services -->

<a name="pricefeed/tx.proto"></a>

<p align="right"><a href="#top">Top</a></p>

## pricefeed/tx.proto

<a name="ununifi.pricefeed.MsgPostPrice"></a>

### MsgPostPrice

| Field       | Type                                                    | Label | Description |
| ----------- | ------------------------------------------------------- | ----- | ----------- |
| `from`      | [string](#string)                                       |       |             |
| `market_id` | [string](#string)                                       |       |             |
| `price`     | [string](#string)                                       |       |             |
| `expiry`    | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |       |             |

<a name="ununifi.pricefeed.MsgPostPriceResponse"></a>

### MsgPostPriceResponse

 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

<a name="ununifi.pricefeed.Msg"></a>

### Msg

| Method Name | Request Type                                    | Response Type                                                   | Description | HTTP Verb | Endpoint |
| ----------- | ----------------------------------------------- | --------------------------------------------------------------- | ----------- | --------- | -------- |
| `PostPrice` | [MsgPostPrice](#ununifi.pricefeed.MsgPostPrice) | [MsgPostPriceResponse](#ununifi.pricefeed.MsgPostPriceResponse) |             |           |

 <!-- end services -->

<a name="ununifidist/ununifidist.proto"></a>

<p align="right"><a href="#top">Top</a></p>

## ununifidist/ununifidist.proto

<a name="ununifi.ununifidist.Params"></a>

### Params

| Field     | Type                                  | Label    | Description |
| --------- | ------------------------------------- | -------- | ----------- |
| `active`  | [bool](#bool)                         |          |             |
| `periods` | [Period](#ununifi.ununifidist.Period) | repeated |             |

<a name="ununifi.ununifidist.Period"></a>

### Period

| Field       | Type                                                    | Label | Description |
| ----------- | ------------------------------------------------------- | ----- | ----------- |
| `start`     | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |       |             |
| `end`       | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |       |             |
| `inflation` | [string](#string)                                       |       |             |

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

| Field                 | Type                                                    | Label | Description                                                     |
| --------------------- | ------------------------------------------------------- | ----- | --------------------------------------------------------------- |
| `params`              | [Params](#ununifi.ununifidist.Params)                   |       |                                                                 |
| `previous_block_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |       |                                                                 |
| `gov_denom`           | [string](#string)                                       |       | this line is used by starport scaffolding # genesis/proto/state |

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

| Field      | Type                                                  | Label    | Description |
| ---------- | ----------------------------------------------------- | -------- | ----------- |
| `balances` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |             |

<a name="ununifi.ununifidist.QueryParamsRequest"></a>

### QueryParamsRequest

<a name="ununifi.ununifidist.QueryParamsResponse"></a>

### QueryParamsResponse

| Field    | Type                                  | Label | Description |
| -------- | ------------------------------------- | ----- | ----------- |
| `params` | [Params](#ununifi.ununifidist.Params) |       |             |

 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

<a name="ununifi.ununifidist.Query"></a>

### Query

Query defines the gRPC querier service.

| Method Name | Request Type                                                            | Response Type                                                             | Description                                   | HTTP Verb | Endpoint                      |
| ----------- | ----------------------------------------------------------------------- | ------------------------------------------------------------------------- | --------------------------------------------- | --------- | ----------------------------- |
| `Params`    | [QueryParamsRequest](#ununifi.ununifidist.QueryParamsRequest)           | [QueryParamsResponse](#ununifi.ununifidist.QueryParamsResponse)           |                                               | GET       | /ununifi/ununifidist/params   |
| `Balances`  | [QueryGetBalancesRequest](#ununifi.ununifidist.QueryGetBalancesRequest) | [QueryGetBalancesResponse](#ununifi.ununifidist.QueryGetBalancesResponse) | this line is used by starport scaffolding # 2 | GET       | /ununifi/ununifidist/balances |

 <!-- end services -->

# <<<<<<< HEAD

<a name="yield-aggregator/params.proto"></a>

<p align="right"><a href="#top">Top</a></p>

## yield-aggregator/params.proto

<a name="ununifi.chain.yieldaggregator.Params"></a>

### Params

Params defines the parameters for the module.

| Field    | Type                                                        | Label    | Description |
| -------- | ----------------------------------------------------------- | -------- | ----------- |
| `vaults` | [Params.Vault](#ununifi.chain.yieldaggregator.Params.Vault) | repeated |             |

<a name="ununifi.chain.yieldaggregator.Params.Vault"></a>

### Params.Vault

| Field                  | Type              | Label | Description |
| ---------------------- | ----------------- | ----- | ----------- |
| `denom`                | [string](#string) |       |             |
| `performance_fee_rate` | [string](#string) |       |             |

 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->

<a name="yield-aggregator/yield-aggregator.proto"></a>

<p align="right"><a href="#top">Top</a></p>

## yield-aggregator/yield-aggregator.proto

<a name="ununifi.chain.yieldaggregator.Strategy"></a>

### Strategy

<<<<<<< HEAD
| Field | Type | Label | Description |
| ------------------ | ----------------------------------------------------------------- | ----- | ----------- |
| `vault_denom` | [string](#string) | | |
| `id` | [string](#string) | | |
| `contract_address` | [string](#string) | | |
| `name` | [string](#string) | | |
| `weight` | [string](#string) | | |
| `metrics` | [StrategyMetrics](#ununifi.chain.yieldaggregator.StrategyMetrics) | | |
=======

| Field              | Type                                                              | Label | Description |
| ------------------ | ----------------------------------------------------------------- | ----- | ----------- |
| `vault_denom`      | [string](#string)                                                 |       |             |
| `id`               | [uint64](#uint64)                                                 |       |             |
| `contract_address` | [string](#string)                                                 |       |             |
| `name`             | [string](#string)                                                 |       |             |
| `weight`           | [string](#string)                                                 |       |             |
| `metrics`          | [StrategyMetrics](#ununifi.chain.yieldaggregator.StrategyMetrics) |       |             |

> > > > > > > bb3bfd2c (feat: iya v2)

<a name="ununifi.chain.yieldaggregator.StrategyMetrics"></a>

### StrategyMetrics

| Field | Type              | Label | Description |
| ----- | ----------------- | ----- | ----------- |
| `apr` | [string](#string) |       |             |
| `tvl` | [string](#string) |       |             |

<a name="ununifi.chain.yieldaggregator.Vault"></a>

### Vault

| Field     | Type                                                        | Label | Description |
| --------- | ----------------------------------------------------------- | ----- | ----------- |
| `denom`   | [string](#string)                                           |       |             |
| `metrics` | [VaultMetrics](#ununifi.chain.yieldaggregator.VaultMetrics) |       |             |

<a name="ununifi.chain.yieldaggregator.VaultMetrics"></a>

### VaultMetrics

| Field | Type              | Label | Description |
| ----- | ----------------- | ----- | ----------- |
| `apy` | [string](#string) |       |             |
| `tvl` | [string](#string) |       |             |

 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->

<a name="yield-aggregator/genesis.proto"></a>

<p align="right"><a href="#top">Top</a></p>

## yield-aggregator/genesis.proto

<a name="ununifi.chain.yieldaggregator.GenesisState"></a>

### GenesisState

GenesisState defines the yieldaggregator module's genesis state.

| Field        | Type                                                | Label    | Description                                                     |
| ------------ | --------------------------------------------------- | -------- | --------------------------------------------------------------- |
| `params`     | [Params](#ununifi.chain.yieldaggregator.Params)     |          |                                                                 |
| `vaults`     | [Vault](#ununifi.chain.yieldaggregator.Vault)       | repeated | this line is used by starport scaffolding # genesis/proto/state |
| `strategies` | [Strategy](#ununifi.chain.yieldaggregator.Strategy) | repeated |                                                                 |

 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->

<a name="yield-aggregator/query.proto"></a>

<p align="right"><a href="#top">Top</a></p>

## yield-aggregator/query.proto

# <<<<<<< HEAD

<a name="ununifi.chain.yieldaggregator.QueryAllStrategyRequest"></a>

### QueryAllStrategyRequest

| Field        | Type                                                                            | Label | Description |
| ------------ | ------------------------------------------------------------------------------- | ----- | ----------- |
| `denom`      | [string](#string)                                                               |       |             |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |       |             |

<a name="ununifi.chain.yieldaggregator.QueryAllStrategyResponse"></a>

### QueryAllStrategyResponse

| Field        | Type                                                                              | Label    | Description |
| ------------ | --------------------------------------------------------------------------------- | -------- | ----------- |
| `strategies` | [Strategy](#ununifi.chain.yieldaggregator.Strategy)                               | repeated |             |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |          |             |

<a name="ununifi.chain.yieldaggregator.QueryAllVaultRequest"></a>

### QueryAllVaultRequest

| Field        | Type                                                                            | Label | Description |
| ------------ | ------------------------------------------------------------------------------- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |       |             |

<a name="ununifi.chain.yieldaggregator.QueryAllVaultResponse"></a>

### QueryAllVaultResponse

| Field        | Type                                                                              | Label    | Description |
| ------------ | --------------------------------------------------------------------------------- | -------- | ----------- |
| `vaults`     | [Vault](#ununifi.chain.yieldaggregator.Vault)                                     | repeated |             |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |          |             |

<a name="ununifi.chain.yieldaggregator.QueryGetStrategyRequest"></a>

### QueryGetStrategyRequest

| Field   | Type              | Label | Description |
| ------- | ----------------- | ----- | ----------- |
| `denom` | [string](#string) |       |             |
| `id`    | [uint64](#uint64) |       |             |

<a name="ununifi.chain.yieldaggregator.QueryGetStrategyResponse"></a>

### QueryGetStrategyResponse

| Field      | Type                                                | Label | Description |
| ---------- | --------------------------------------------------- | ----- | ----------- |
| `strategy` | [Strategy](#ununifi.chain.yieldaggregator.Strategy) |       |             |

<a name="ununifi.chain.yieldaggregator.QueryGetVaultRequest"></a>

### QueryGetVaultRequest

| Field   | Type              | Label | Description |
| ------- | ----------------- | ----- | ----------- |
| `denom` | [string](#string) |       |             |

<a name="ununifi.chain.yieldaggregator.QueryGetVaultResponse"></a>

### QueryGetVaultResponse

| Field   | Type                                          | Label | Description |
| ------- | --------------------------------------------- | ----- | ----------- |
| `vault` | [Vault](#ununifi.chain.yieldaggregator.Vault) |       |             |

> > > > > > > bb3bfd2c (feat: iya v2)
> > > > > > > <a name="ununifi.chain.yieldaggregator.QueryParamsRequest"></a>

### QueryParamsRequest

QueryParamsRequest is request type for the Query/Params RPC method.

<a name="ununifi.chain.yieldaggregator.QueryParamsResponse"></a>

### QueryParamsResponse

QueryParamsResponse is response type for the Query/Params RPC method.

<<<<<<< HEAD
| Field | Type | Label | Description |
| -------- | ----------------------------------------------- | ----- | ----------------------------------------------- |
| `params` | [Params](#ununifi.chain.yieldaggregator.Params) | | params holds all the parameters of this module. |

<a name="ununifi.chain.yieldaggregator.QueryStrategiesRequest"></a>

### QueryStrategiesRequest

| Field   | Type              | Label | Description |
| ------- | ----------------- | ----- | ----------- |
| `denom` | [string](#string) |       |             |

<a name="ununifi.chain.yieldaggregator.QueryStrategiesResponse"></a>

### QueryStrategiesResponse

| Field        | Type                                                                            | Label    | Description |
| ------------ | ------------------------------------------------------------------------------- | -------- | ----------- |
| `strategies` | [Strategy](#ununifi.chain.yieldaggregator.Strategy)                             | repeated |             |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |          |             |

<a name="ununifi.chain.yieldaggregator.QueryVaultRequest"></a>

### QueryVaultRequest

| Field   | Type              | Label | Description |
| ------- | ----------------- | ----- | ----------- |
| `denom` | [string](#string) |       |             |

<a name="ununifi.chain.yieldaggregator.QueryVaultResponse"></a>

### QueryVaultResponse

| Field   | Type                                          | Label    | Description |
| ------- | --------------------------------------------- | -------- | ----------- |
| `vault` | [Vault](#ununifi.chain.yieldaggregator.Vault) | repeated |             |

<a name="ununifi.chain.yieldaggregator.QueryVaultsRequest"></a>

### QueryVaultsRequest

<a name="ununifi.chain.yieldaggregator.QueryVaultsResponse"></a>

### QueryVaultsResponse

| Field        | Type                                                                            | Label    | Description |
| ------------ | ------------------------------------------------------------------------------- | -------- | ----------- |
| `vaults`     | [Vault](#ununifi.chain.yieldaggregator.Vault)                                   | repeated |             |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |          |             |

=======

| Field    | Type                                            | Label | Description                                     |
| -------- | ----------------------------------------------- | ----- | ----------------------------------------------- |
| `params` | [Params](#ununifi.chain.yieldaggregator.Params) |       | params holds all the parameters of this module. |

> > > > > > > bb3bfd2c (feat: iya v2)

 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

<a name="ununifi.chain.yieldaggregator.Query"></a>

### Query

Query defines the gRPC querier service.

<<<<<<< HEAD
| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ------------ | ------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | ------------------------------------------------ | --------- | --------------------------------------------------------- |
| `Params` | [QueryParamsRequest](#ununifi.chain.yieldaggregator.QueryParamsRequest) | [QueryParamsResponse](#ununifi.chain.yieldaggregator.QueryParamsResponse) | Parameters queries the parameters of the module. | GET | /UnUniFi/chain/yield-aggregator/params |
| `Vaults` | [QueryVaultsRequest](#ununifi.chain.yieldaggregator.QueryVaultsRequest) | [QueryVaultsResponse](#ununifi.chain.yieldaggregator.QueryVaultsResponse) | this line is used by starport scaffolding # 2 | GET | /UnUniFi/chain/yield-aggregator/vaults |
| `Vault` | [QueryVaultRequest](#ununifi.chain.yieldaggregator.QueryVaultRequest) | [QueryVaultResponse](#ununifi.chain.yieldaggregator.QueryVaultResponse) | | GET | /UnUniFi/chain/yield-aggregator/vaults/{denom} |
| `Strategies` | [QueryStrategiesRequest](#ununifi.chain.yieldaggregator.QueryStrategiesRequest) | [QueryStrategiesResponse](#ununifi.chain.yieldaggregator.QueryStrategiesResponse) | | GET | /UnUniFi/chain/yield-aggregator/vaults/{denom}/strategies |
=======
| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#ununifi.chain.yieldaggregator.QueryParamsRequest) | [QueryParamsResponse](#ununifi.chain.yieldaggregator.QueryParamsResponse) | Parameters queries the parameters of the module. | GET|/UnUniFi/chain/yield-aggregator/params|
| `VaultAll` | [QueryAllVaultRequest](#ununifi.chain.yieldaggregator.QueryAllVaultRequest) | [QueryAllVaultResponse](#ununifi.chain.yieldaggregator.QueryAllVaultResponse) | this line is used by starport scaffolding # 2 | GET|/UnUniFi/chain/yield-aggregator/vaults|
| `Vault` | [QueryGetVaultRequest](#ununifi.chain.yieldaggregator.QueryGetVaultRequest) | [QueryGetVaultResponse](#ununifi.chain.yieldaggregator.QueryGetVaultResponse) | | GET|/UnUniFi/chain/yield-aggregator/vaults/{denom}|
| `StrategyAll` | [QueryAllStrategyRequest](#ununifi.chain.yieldaggregator.QueryAllStrategyRequest) | [QueryAllStrategyResponse](#ununifi.chain.yieldaggregator.QueryAllStrategyResponse) | | GET|/UnUniFi/chain/yield-aggregator/vaults/{denom}/strategies|
| `Strategy` | [QueryGetStrategyRequest](#ununifi.chain.yieldaggregator.QueryGetStrategyRequest) | [QueryGetStrategyResponse](#ununifi.chain.yieldaggregator.QueryGetStrategyResponse) | | GET|/UnUniFi/chain/yield-aggregator/vaults/{denom}/strategies/{id}|

> > > > > > > bb3bfd2c (feat: iya v2)

 <!-- end services -->

<a name="yield-aggregator/tx.proto"></a>

<p align="right"><a href="#top">Top</a></p>

## yield-aggregator/tx.proto

<a name="ununifi.chain.yieldaggregator.MsgDepositToVault"></a>

### MsgDepositToVault

this line is used by starport scaffolding # proto/tx/message

| Field    | Type                                                  | Label | Description |
| -------- | ----------------------------------------------------- | ----- | ----------- |
| `sender` | [string](#string)                                     |       |             |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |       |             |

<a name="ununifi.chain.yieldaggregator.MsgDepositToVaultResponse"></a>

### MsgDepositToVaultResponse

<a name="ununifi.chain.yieldaggregator.MsgWithdrawFromVault"></a>

### MsgWithdrawFromVault

<<<<<<< HEAD
| Field | Type | Label | Description |
| ----------------- | ----------------- | ----- | ----------- |
| `sender` | [string](#string) | | |
| `principal_denom` | [string](#string) | | |
| `lp_token_amount` | [string](#string) | | |
=======

| Field             | Type              | Label | Description |
| ----------------- | ----------------- | ----- | ----------- |
| `sender`          | [string](#string) |       |             |
| `vault_denom`     | [string](#string) |       |             |
| `lp_token_amount` | [string](#string) |       |             |

> > > > > > > bb3bfd2c (feat: iya v2)

<a name="ununifi.chain.yieldaggregator.MsgWithdrawFromVaultResponse"></a>

### MsgWithdrawFromVaultResponse

 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

<a name="ununifi.chain.yieldaggregator.Msg"></a>

### Msg

Msg defines the Msg service.

| Method Name         | Request Type                                                                | Response Type                                                                               | Description                                              | HTTP Verb | Endpoint |
| ------------------- | --------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | -------------------------------------------------------- | --------- | -------- |
| `DepositToVault`    | [MsgDepositToVault](#ununifi.chain.yieldaggregator.MsgDepositToVault)       | [MsgDepositToVaultResponse](#ununifi.chain.yieldaggregator.MsgDepositToVaultResponse)       | this line is used by starport scaffolding # proto/tx/rpc |           |
| `WithdrawFromVault` | [MsgWithdrawFromVault](#ununifi.chain.yieldaggregator.MsgWithdrawFromVault) | [MsgWithdrawFromVaultResponse](#ununifi.chain.yieldaggregator.MsgWithdrawFromVaultResponse) |                                                          |           |

 <!-- end services -->

> > > > > > > b9a58125 (feat: keeper)

## Scalar Value Types

| .proto Type                    | Notes                                                                                                                                           | C++    | Java       | Python      | Go      | C#         | PHP            | Ruby                           |
| ------------------------------ | ----------------------------------------------------------------------------------------------------------------------------------------------- | ------ | ---------- | ----------- | ------- | ---------- | -------------- | ------------------------------ |
| <a name="double" /> double     |                                                                                                                                                 | double | double     | float       | float64 | double     | float          | Float                          |
| <a name="float" /> float       |                                                                                                                                                 | float  | float      | float       | float32 | float      | float          | Float                          |
| <a name="int32" /> int32       | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32  | int        | int         | int32   | int        | integer        | Bignum or Fixnum (as required) |
| <a name="int64" /> int64       | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64  | long       | int/long    | int64   | long       | integer/string | Bignum                         |
| <a name="uint32" /> uint32     | Uses variable-length encoding.                                                                                                                  | uint32 | int        | int/long    | uint32  | uint       | integer        | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64     | Uses variable-length encoding.                                                                                                                  | uint64 | long       | int/long    | uint64  | ulong      | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32     | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s.                            | int32  | int        | int         | int32   | int        | integer        | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64     | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s.                            | int64  | long       | int/long    | int64   | long       | integer/string | Bignum                         |
| <a name="fixed32" /> fixed32   | Always four bytes. More efficient than uint32 if values are often greater than 2^28.                                                            | uint32 | int        | int         | uint32  | uint       | integer        | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64   | Always eight bytes. More efficient than uint64 if values are often greater than 2^56.                                                           | uint64 | long       | int/long    | uint64  | ulong      | integer/string | Bignum                         |
| <a name="sfixed32" /> sfixed32 | Always four bytes.                                                                                                                              | int32  | int        | int         | int32   | int        | integer        | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes.                                                                                                                             | int64  | long       | int/long    | int64   | long       | integer/string | Bignum                         |
| <a name="bool" /> bool         |                                                                                                                                                 | bool   | boolean    | boolean     | bool    | bool       | boolean        | TrueClass/FalseClass           |
| <a name="string" /> string     | A string must always contain UTF-8 encoded or 7-bit ASCII text.                                                                                 | string | String     | str/unicode | string  | string     | string         | String (UTF-8)                 |
| <a name="bytes" /> bytes       | May contain any arbitrary sequence of bytes.                                                                                                    | string | ByteString | str         | []byte  | ByteString | string         | String (ASCII-8BIT)            |
