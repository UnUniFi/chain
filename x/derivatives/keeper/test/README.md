# NFT Backed Loan Keeper Test

Before run debug `build then launch`

The chain must be initialized before each test is run.
When re-run same script, chain needs to be initialized.

Test DLP exchanges at Pool and token transfers in perpetual-futures positions

## Pool

### 01_pool_all_withdraw

Test token transfers and changes in DLP prices when all DLP are burned.

- If all DLPs are burned, the rate returns to 0 and the initial rate (1:1 regardless of tokens) is applied to the next deposit
- After everything is burned, depositor will get back all tokens but the fee.

### 02_selling_decision_no_paid

Test token transfers and DLP price if 1udlp is left and the other is burned.

- When 1udlp remains, the pool has 1 token, thus maintaining the initial rate of 1:1.

### 03_pool_liquidation

Test token transfers when position is happened Levy Period.

- Traders suffer losses from falling prices.
- Fees and losses go into the pool.

### 04_pool_levy_period

- Traders lose swap fees.
- The reporter gets a portion of the fee. Under the current spec, the module account (Pool).
- The NFT return to lister

### 05_pool_withdraw_opened_position

Test that the Pool balance used to open a position cannot be withdrawn by the depositor.

- Pool balance is locked and the depositor's withdrawal fails.

## Perpetual Futures Position

### 06_position_close_price_up

Test token transfers when closing a position during a price increase

- Traders get profits from price increases.
- Pool loses for that.

### 07_position_close_price_down

Test token transfers when closing a position during a price decrease

- Traders suffer losses from falling prices.
- Pool benefits from that.
