# Messages

## MsgDepositToVault

Call `stake` function of farming contract,

## MsgWithdrawFromVault

Call `unstake` function of farming contract, and then return the principal token with amount such that

$$
\text{amount} = \text{principal\_token\_amount\_in\_vault} \times \frac{\text{lp\_token\_amount\_to\_burn}}{\text{total\_supply\_of\_lp\_token}}
$$
