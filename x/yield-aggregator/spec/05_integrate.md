# Integrate

## Yield Farming Contract

```typescript
interface YieldFarmingContract {
  stake(amount: Coin);
  unstake(amount: Coin);
  amount();
  apr();
  interest_fee_rate();
}
```

Rust-formatted interface will be written later.
