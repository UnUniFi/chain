# v1.0.0-beta.4 to v2 Upgrade Guide

## Purpose of this upgrade

This proposal aims to do upgrade to UnUniFi `v2.0.0`. 
This upgrade will upgrade you to cosmos sdk v47 and allow you to use the nft, wasmd, and ibc modules.

## Brief guide

### The following text will be updated soon.

~~ All validators nodes should upgrades to `v2.0.0`. The `v2.0.0` binary is state machine compatible with `v1.0.0-beta.4` until **block 3452990. At 01:00 UTC on January 10th, 2023**, we will have a coordinated re-start of the network.  ~~ 
~~ All validator(full) nodes have to do is set the binary of `v3.0.0` in the appropriate location before 3452990 block height.   ~~
~~ At 3452990, if you use cosmovisor, the system automatically upgrades the binary and block 3452990 will be mined with over 67% voting power.   ~~

## Go Requirement

You will need to be running go1.19 for this, same as the previsou version. You can check with this command:

```shell
go version
```

## Setup

If you use cosmovisor which we highly recommend, create the required folder, make the build, and copy the daemon over to that folder with the appropriate name. NOTE: Don't forget check the file owner of v2 binary.

```shell
mkdir -p $DAEMON_HOME/cosmovisor/upgrades/v2/bin
cd $HOME/<ununifi-repo>
git pull
git checkout v2.0.0
make build -B
## to make sure the build status
# ./build/ununifid version
cp ./build/ununifid $DAEMON_HOME/cosmovisor/upgrades/v2/bin
```

Even though the cosmovisor's `DAEMON_ALLOW_DOWNLOAD_BINARIES` variable is set `true`, the cosmovisor uses binary which locates in $DAEMON_HOME/cosmovisor/upgrades/v2/bin if exists.   
And you don't have to reboot cosmovisor when to do upgrade. So, after locating the binary into the appropriate place, you don't need to anything.

## Futher Help

If you need more help, please go to our discord https://discord.com/channels/762953633230356480/762953916593340416.
