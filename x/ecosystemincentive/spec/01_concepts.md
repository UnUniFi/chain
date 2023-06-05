# Concepts

**NOTE: This is early draft.**

This module aims to provide the incentive for the parties which especially bring value to our ecosystem like frontend service creator.   
Fucosing on the case for the frontend service creator, any of them who creates UnUniFi NFT market and NFTFi frontend service are the subjects to recieve Ecosystem Incentive reward from the NFT trading fee in many denoms which are used in NFT market.

## Joining Ecosystem Incentive

Any subjects can send a register message `MsgIncentiveRegister` with the `incentive_id` and `subject_weight_map`.   

## Getting Ecosystem Incentive Reward

This model of distribution reward could be applied to many use-cases. But, we write down only about the case for Nftmarket Frontend model here for better explanation of the sense of this module.   
First, the subjects must register to get incentive by sending `MsgIncentiveRegister`.   
Once the `incentive_id` is registered, they insert that `incentive_id` in the target message which is `MsgListNft` memo field precisely to get the reward for the Nftmarket Frontend incentive mode.
Once the `NftIdentifer` on the market is connected with `incentive_id`, `AfterNftPaymentWithCommission` hook function triggers methods to reflect the reward amount for according addresses in `incentive_id`.

## Withdrawing Ecosystem Incentive Reward

Any registered subjects can withdraw thier reward by sending a withdrawal message if they are there.   
They can withdraw all rewards across all denoms by sending `MsgWithdrawAllRewards`.   
In other way, they can withdraw specific denom reward by sending `MsgWithdrawSpecificDenomReward`.

## The Reward Mechanism

All the reward comes from the fees that UnUniFi protocol earned without gas fee which is defined in protocol as glocal parameter.   
There is nothing inflational effect or depletion by rewarding subjects.
