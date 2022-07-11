# Network Params

`Params` describes global parameters that are maintained by governance.

|  Key                       | Type                     |
|  --------------------------| -------------------------|
|  MaxTokenSupplyLimit       | uint64                   |
|  MinClassNameLen           | uint64                   |
|  MaxClassNameLen           | uint64                   |
|  MaxClassUriLen            | uint64                   |
|  MaxBaseTokenUriLen        | uint64                   |
|  MaxSymbolLen              | uint64                   |
|  MaxDescriptionLen         | uint64                   |

1. **MaxTokenSupply** - The max token supply is the cap of the number of each `Class`'s `NFT`.
1. **MinClassNameLen** - The min class name length is the min string length that `Class.Name` can be put.
1. **MaxClassNameLen** - The max class name length is the max string length that `Class.Name` can be put.
1. **MinUriLen** - The min class uri length is the min string length that `Class.Uri` can be put.
1. **MaxUriLen** - The max class uri length is the max string length that `Class.Uri` can be put.
1. **MaxSymbolLen** - The max symbol length is the max string length that `Class.Symbol` can be put.
1. **MaxDescriptionLen** - The max description length is the max string length that `Class.Description` can be put.

### Default values

```json
{
    "MaxTokenSupplyLimit": 100000,
    "MinClassNameLen": 3,
    "MaxClassNameLen": 128,
    "MinUriLen": 8,
    "MaxUriLen": 512,
    "MaxSymbolLen": 16,
    "MaxDescriptionLen": 1024
}
```
