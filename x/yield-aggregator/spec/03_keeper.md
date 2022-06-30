# Keeper

## GetAssetManagementAccounts

## GetAssetManagementTargetsOfAccount

Get `AssetManagementTarget`s with `asset_management_account_id`.

## GetAssetManagementTargetsOfDenom

Get `AssetManagementTarget`s with `denom`.

## GetDepositsOfAddress

## Deposit

## Withdraw

## Handler

BlockHandlerかなんかでそれぞれのコントラクトに情報とりにいく
nft-marketmakerはcosmwasmじゃなくてgolangモジュールなので、２つのやり方
- golangモジュールに情報取りに行く処理も書く
- ngt-marketmakerをラップするcosmwasmコントラクトを書く
