# Concepts

Any frontend creator of UnUniFi NFT market and NFTFi service are the subjects to recieve Frontend Incentive from the NFT trading fee in many denoms which are used in NFT market.

### Joining Frontend Incentive Reward

Any subjects can send a register message with the `frontend_name` and `subject_weight_map`.  

### Getting Frontend Incentive Reward 

Once the `frontend_name` is registered, they insert that `frontend_name` in the target message's memo field precisely to get the reward.
Current considering target message is `MsgPayAuctionFee`.

### Withdrawing Frontend Incentive Reward

Any registered subjects can withdraw thier reward by sending a withdrawal message if they are there.   
They can withdraw all rewards across all denoms by sending `MsgWithdrawAllFrontendReward`.   
In other way, they can withdraw specific denom reward by sending `MsgWithdrawSpecificFrontendReward`.

### The Reward Mechanism

All the reward comes from the NFT trading fee which is defined in protocol as glocal parameter. 
There is nothing inflational effect or depletion.

