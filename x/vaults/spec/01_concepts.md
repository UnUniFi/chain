# Concepts

Modules to hold NFTs and balancesModules to hold NFTs and balances.  
NFTs and tokens can be locked in the event of a security incident.  

# requirement

## basic

### deposit
1. トークンを預けることができる
1. NFTを預けることができる

### withdrow 
1. トークンを引き出す事ができる
1. NFTを引き出す事ができる

### lock
1. genesis.jsonのパラメータが有効かつ、XからのMsgだった場合vaultsはロックすることができる
1. ロックされている間は、預入、引き出しをすることができない

### unlock
1. genesis.jsonのパラメータが有効かつ、XからのMsgだった場合vaultsはアンロックすることができる
