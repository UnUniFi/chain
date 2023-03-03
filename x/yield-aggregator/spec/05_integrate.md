# Integrate

## Yield Farming Contract

```typescript
interface YieldFarmingContract {
  stake(amount: Coin);
  unstake(amount: Coin);
  get_apr();
  get_interest_fee_rate();
}
```

Rust-formatted interface will be written later.
