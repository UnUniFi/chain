# v1-beta.1 to v1-beta.3 Upgrade Guide

## Purpose of this upgrade

This proposal aims to do upgrade to UnUniFi `v1.0.0-beta.3`. The main function is the same after the upgrade. But, we distribute token for the community program winners, moderators and airdrop salgated accounts under the upgrade operation. The
reason why we do by upgrade is we currently make the feature to send tokens disable.

## Brief guide

All validators nodes should upgrades to `v1.0.0-beta.3`. The `v1.0.0-beta.3` binary is state machine compatible with `v1.0.0-beta.1` until **block 1597000. At 01:00 UTC on September 6th, 2022**, we will have a coordinated re-start of the network. 
All validator(full) nodes have to do is set the binary of `v1.0.0-beta.3` in the appropriate location before 1597000 block height.   
At 1597000, if you use cosmovisor, the system automatically upgrades the binary and block 1597000 will be mined with over 67% voting power.   

## Go Requirement

You will need to be running go1.17 for this, same as the previsou version. You can check with this command:

```shell
go version
```

## Setup

If you use cosmovisor which we highly recommend, create the required folder, make the build, and copy the daemon over to that folder with the appropriate name. NOTE: Don't forget check the file owner of v1-beta.3 binary.

```shell
mkdir -p $DAEMON_HOME/cosmovisor/upgrades/v1-beta.3/bin
cd $HOME/<ununifi-repo>
git pull
git checkout v1.0.0-beta.3
make build -B
# you can take checksum in case (but, note that the value can be changed easily)
# md5: 9340e63cf6a04530218b151cb9e554b4
cp ./build/ununifid $DAEMON_HOME/cosmovisor/upgrades/v1-beta.3/bin
```

Even though the cosmovisor's `DAEMON_ALLOW_DOWNLOAD_BINARIES` variable is set `true`, the cosmovisor uses binary which locates in $DAEMON_HOME/cosmovisor/upgrades/v1-beta.3/bin if exists.   
And you don't have to reboot cosmovisor when to do upgrade. So, after locating the binary into the appropriate place, you don't need to anything.

## Futher Help

If you need more help, please go to our discord https://discord.gg/GXTx9wjA.
