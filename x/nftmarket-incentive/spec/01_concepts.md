# Concepts

**NOTE: This is early draft.**

Any frontend creator of UnUniFi NFT market and NFTFi service are the subjects to recieve Frontend Incentive reward from the NFT trading fee in many denoms which are used in NFT market.

### Joining Frontend Incentive

Any subjects can send a register message with the `incentive_id` and `subject_weight_map`.   
Or you can simply contrains the address in the target transaction's memo field.   

### Getting Frontend Incentive Reward 

Once the `incentive_id` is registered, they insert that `incentive_id` in the target message's memo field precisely to get the reward.
Current target message is `MsgPayFullBid`.   
Or you can simply contrains the address in the target transaction's memo field.   
The difference of them is whether you can set the weight for each address.   
e.g. In memo field of MsgPayFullBid,   
{"nftmarket-incentive": "`incentive_id`"}  
or   
{"nftmarket-incentive": ["address1", "address2"]}

### Withdrawing Frontend Incentive Reward

Any registered subjects can withdraw thier reward by sending a withdrawal message if they are there.   
They can withdraw all rewards across all denoms by sending `MsgWithdrawAllFrontendReward`.   
In other way, they can withdraw specific denom reward by sending `MsgWithdrawSpecificFrontendReward`.

### The Reward Mechanism

All the reward comes from the NFT trading fee which is defined in protocol as glocal parameter. 
There is nothing inflational effect or depletion.
