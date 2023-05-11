# Messages

## MsgNftLocked(execute only validators)

Mint wrapped NFTs

```protobuf
message MsgNftLocked {
  string sender = 1;
  string toAddress = 2;
  string nftId = 3;
  string nftName = 4;
}
```

## MsgNftUnLocked(execute only validators)

Burn wrapped NFTs

```protobuf
message MsgNftLocked {
  string sender = 1;
  string toAddress = 2;
  string nftId = 3;
}
```

## MsgTransferRequst

deposit wrapped NFTs

```protobuf
message MsgNftLocked {
  string sender = 1;
  string ethAddress = 2;
  string nftId = 3;
}
```

## MsgTransferred(execute only validators)

Burn wrapped NFTs

```protobuf
message MsgNftLocked {
  string sender = 1;
  string nftId = 2;
}
```

## MsgRejectTransfer(execute only validators)

withdraw wrapped NFTs

```protobuf
message MsgNftLocked {
  string sender = 1;
  string nftId = 2;
}
```
