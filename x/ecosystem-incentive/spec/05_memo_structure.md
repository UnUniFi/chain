# Data structure for the memo field

We use tx memo field data to identify what incentive will be distributed to what `incentive-unit` by putting the correct formatted json data into that.

The v1's formal data archtecture is:

```json
{
  "version": "v1",
  "incentive-unit-id": "incentive_unit-1"
}
```

NOTE: There's a lot of chances to be changed this structure with the change of the version. Please note it when to use.

## Frontends

We use memo field data to know which frontend a lisetd nft used in the case of frontend-incentive model.   
So we have to use the organized data structure of memo field in a listing tx (MsgListNft) to distingush it as a legitimate entry or not.

Even if you put the wrong formatted data in the memo of tx contains MsgListNft, the MsgListNft itself will still succeed. The registration of the information which nft-id relates to what `incentive-unit-id` will just fail.
