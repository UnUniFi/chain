# EndBlock

**NOTE: This is early draft.**

All rewards accumulation are executed at the EndBlock.   
First, search the trasactino contains the message type to be subject to `frontend-incentive`.   
Second, extract the memo field data and arguments of that message and transaction.   
Third, calculate the actual `frontend-incentive` reward that is earned by that transaction.    
And, update and reflect the subjects' accumulated and claimable reward amount
