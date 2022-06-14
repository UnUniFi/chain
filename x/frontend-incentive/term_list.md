# Term list

The definition of the specific terms which are related to this module.

`frontend_store`   
The KVStore that contains all information for the entry of frontend-incentive protocol.   
Specifically, the key is `frontend_name` and value is `subject_weight_map`.   
e.g. 
{ frontend_name: "example", { ununifi1a~: 0.50, ununifi1b~: 0.50 } }

Maybe user input is shaped as a json file.

`frontend_name`   
frontend_name is the unique identifier in the `frontend_store`. Hence, it can't be duplicated.    

`subject`    
The UnUniFI account to have a right to receive frontend-incentive.

`weight`   
The ratio of the reward distribution in a     `frontend_store` unit. `frontend_store` can contain several `subject`s and ratio for each. 

`reward_rate`   
The rate to determine the percentage for the `frontend-incentive` reward out of the total trading fee.

`denom`   
The token ticker like GUU, BTC etc.   
In UnUniFi NFT market, some tokens other than GUU (native UnUniFi token) can be used to purchase NFTs. So the rewards are accumuleted in some denoms.

`reward_payer`   
The account to pay the reward for `frontend-incentive` protocol.   
It's possibly community-pool (distribution module account). But, thisn't final decision.
