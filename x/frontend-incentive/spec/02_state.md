# State

**NOTE: This is early draft.**
### frontend_store
The KVStore that contains all information for the entry of frontend-incentive protocol.   
Specifically, the key is `frontend_name` and value is `subject_weight_map`.   
e.g. 
{ frontend_name: "example", { ununifi1a~: 0.50, ununifi1b~: 0.50 } }

Maybe user input is shaped as a json file.

### frontend_name
Name is the unique identifier in the `frontend_store`. Hence, it can't be duplicated.    

### subject
The UnUniFI account to have a right to receive frontend-incentive.

### weight
The ratio of the reward distribution in a     `frontend_store` unit. `frontend_store` can contain several `subject`s and ratio for each. 

### reward_rate
The rate to determine the percentage for the `frontend-incentive` reward out of the total trading fee.
