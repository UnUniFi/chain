# Strategy

Decide on an investment strategy.
This is part of FarmingOrder state

## Format

```json
{
 "strategyType": "string",
 "option": {
  "whiteTargetIdList":["string"],
  "blackTargetIdList":["string"],
 },
}
```

example

```json
{
 "strategyType": "recent30DaysHighDPRStrategy",
 "option": {
  "blackTargetIdList":["2","3"],
 },
}
```

```json
{
 "strategyType": "ManualStrategy",
 "option": {
  "whiteTargetIdList":["1"],
 },
}
```

## Params

- `strategyType` is the type of investment strategy

|key                        |Description                                                      |
|---------------------------|-----------------------------------------------------------------|
|recent30DaysHighDPRStrategy|Invest in the best DPR destination in the last 30 days on average|
|recent1DayHighDPRStrategy  |Invest in the best DPR destination in the last average day       |
|notHaveDPRStrategy         |Invest in something that does not have a DPR.                    |
|ManualStrategy             |Manual investment, whiteTargetIdlist required.                   |

- (ManualStrategy only) `whiteTargetIdList` is list of potential investments.Select only one that meets your criteria and invest in it.Preference is given to the top of the list.
- `blackTargetIdList` is the list in which we do not invest.
