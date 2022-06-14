# Term list

The definition of the specific terms which are related to this module.

`frontend_tag`   
The data object that contains all information for the entry of frontend-incentive protocol.   
Specifically, the elements are `name` and `subject_weight_map`.   
e.g. 
```go
type FrontendTag {   
    name string,   
    subject_weight_map map[address]int   
}

example := FrontendTag{"example", {"ununifi1~": 10000~}}
```
Maybe user input is shaped as a json file.

`name`   
Name is the unique identifier of the `frontend_tag`. Hence, it can't be duplicated.    

`subject`    
The UnUniFI account to have a right to receive frontend-incentive.

`weight`   
The ratio of the reward distribution in a     `frontend_tag` unit. `frontend_tag` can contain several `subject`s and ratio for each. 

`reward_rate`   
The rate to determine the percentage for the `frontend-incentive` reward out of the total trading fee.

`denom`   
The token ticker like GUU, BTC etc.   
In UnUniFi NFT market, some tokens other than GUU (native UnUniFi token) can be used to purchase NFTs. So the rewards are accumuleted in some denoms.

