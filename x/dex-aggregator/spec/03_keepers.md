# Keepers

## SwapRoute

`SwapRoute` has a validation function to double-check the denom

```go
func (route *SwapRoute) validate(amount sdk.Coin) bool {}
```

## SwapRoutes

`SwapRoute` has a validation function to double-check the denom

```go
type SwapRoutes []SwapRoute

func (routes SwapRoutes) validate(amount sdk.Coin) bool {}
```

```go
type Keeper interface {
  Swap(route SwapRoute, amount sdk.Coin)
  SwapRelay(routes SwapRoutes, amount sdk.Coin)
}
```
