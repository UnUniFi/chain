# Messages

**NOTE: This is early draft.**

### Frontend Register
In case that it's convenient to set map type in argument:
```go
type MsgFrontendRegister struct {
  frontend_name string
  map[string]string  
}
```

In case it's not:
```go
type MsgFrontendRegister struct {
  frontend_name string
  address []string
  weight []string
}
```
or possibly take json file

### WithdrawAllFrontendReward
A message to withdraw all accumulated rewards across all denoms.
```go
type MsgWithdrawAllFrontendReward struct {
  sender string
}
```


### WithdrawSpecificFrontendReward
A message to withdraw accumulated reward of specified denom.
```go
type MsgWithdrawSpecificFrontendReward struct {
  sender string
  denom string
}
```
