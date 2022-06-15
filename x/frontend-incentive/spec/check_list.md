# Check List

**NOTE: You can remove this file once it's done.**

## Needed to be inspected

The check list to achieve all requirements of this module.

- [ ] At EndBlock, is it possible to distinguish transactions from message type?
- [ ] And it id possible to get transactioins of the specific type of message?
- [ ] Is it possible to extract arguments of that message?
- [ ] Is it possible to extract memo field data of that message?
- [ ] Best way to get subject address and its weight via CLI (json file or map?)
- [ ] The way to contain reward information for each addresses and denoms

## Logic ideas which aren't decided

### The way to achieve distribution

1. Actually sending corresponding coin for reward in a process using SendCoinFromModuleAccount
1. In a process, mint corresponding coin for reward for the subject address and just subtract corresponding coin from the subject module account
