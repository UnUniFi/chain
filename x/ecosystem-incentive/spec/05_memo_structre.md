## Data structure for the memo field

We use memo field data to know which frontend a lisetd nft used in the case of frontend-incentive model. So we have to use the organized data structure of memo field in a listing tx to distingush it as a legitimate entry or not.

The v1 archtecture is:

```json
{
  "version": "v1",
  "incentive-type": 0,
  "incentive-id": "incentive_id"
}
```

The `incentive_type` 0 means `NFTMARKET_FRONTEND` type here for the record.

NOTE: There's a lot of chances to be changed this structure with the change of the version. Please note it when to use.
