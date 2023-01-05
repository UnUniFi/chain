# v1-beta.3 to v1-beta.4 Upgrade Guide

## Purpose of this upgrade

This proposal aims to do upgrade to UnUniFi `v1.0.0-beta.4`. The main function is the same after the upgrade. But, we distribute token for the community program winners, moderators and airdrop salgated accounts under the upgrade operation. The
reason why we do by upgrade is we currently make the feature to send tokens disable.  
In short, we do state-modification in the middle of this upgrade to send some token.

## Brief guide

All validators nodes should upgrades to `v1.0.0-beta.4`. The `v1.0.0-beta.4` binary is state machine compatible with `v1.0.0-beta.3` until **block <todo>. At 01:00 UTC on January 10th, 2023**, we will have a coordinated re-start of the network. 
All validator(full) nodes have to do is set the binary of `v1.0.0-beta.4` in the appropriate location before <todo> block height.   
At <todo>, if you use cosmovisor, the system automatically upgrades the binary and block <todo> will be mined with over 67% voting power.   

## Go Requirement

You will need to be running go1.17 for this, same as the previsou version. You can check with this command:

```shell
go version
```

## Setup

If you use cosmovisor which we highly recommend, create the required folder, make the build, and copy the daemon over to that folder with the appropriate name. NOTE: Don't forget check the file owner of v1-beta.4 binary.

```shell
mkdir -p $DAEMON_HOME/cosmovisor/upgrades/v1-beta.4/bin
cd $HOME/<ununifi-repo>
git pull
git checkout v1.0.0-beta.4
make build -B
## to make sure the build status
# ./build/ununifid version
cp ./build/ununifid $DAEMON_HOME/cosmovisor/upgrades/v1-beta.4/bin
```

Even though the cosmovisor's `DAEMON_ALLOW_DOWNLOAD_BINARIES` variable is set `true`, the cosmovisor uses binary which locates in $DAEMON_HOME/cosmovisor/upgrades/v1-beta.4/bin if exists.   
And you don't have to reboot cosmovisor when to do upgrade. So, after locating the binary into the appropriate place, you don't need to anything.

## Futher Help

If you need more help, please go to our discord https://discord.gg/GXTx9wjA.
