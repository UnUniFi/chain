<!--
order: 0
title: "BEP3 Overview"
parent:
  title: "bep3"
-->

# `bep3`

<!-- TOC -->
1. **[Concepts](01_concepts.md)**
2. **[State](02_state.md)**
3. **[Messages](03_messages.md)**
4. **[Events](04_events.md)**
5. **[Params](05_params.md)**
6. **[BeginBlock](06_begin_block.md)**

## Abstract

`x/bep3` is a module that handles cross-chain atomic swaps between Stake and blockchains that implement the BEP3 protocol. Atomic swaps are created, then either claimed before their expiration block or refunded after they've expired.

Several user interfaces support Stake BEP3 swaps:
- [Trust Wallet](https://trustwallet.com/)
- [Cosmostation](https://wallet.cosmostation.io/?network=kava)
- [Frontier Wallet](https://frontierwallet.com/)

Swaps can also be created, claimed, and refunded using Stake's [Javascript SDK](https://github.com/Stake-Labs/javascript-sdk) or CLI.
