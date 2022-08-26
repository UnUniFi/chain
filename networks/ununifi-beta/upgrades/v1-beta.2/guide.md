# v1-beta.1 to v1-beta.2 Upgrade Guide

All validators nodes should upgrades to `v1.0.0-beta.2`. The `v1.0.0-beta.2` binary is state machine compatible with `v1.0.0-beta.1` until block 14181300. At 12:00AM UTC on August 30th, 2022, we will have a coordinated re-start of the network. 
All validator nodes have to do is set the binary of `v1.0.0-beta.2` in the appropriate location before 14181300 block height.   
At 14181300, if you use cosmovisor, the system automatically upgrades the binary and block 14181300 will be mined with over 67% voting power.   

## Go Requirement

You will need to be running go1.17 for this, same as the previsou version. You can check with this command:

```shell
go version
```

## Setup

If you use cosmovisor which we highly recommend, create the required folder, make the build, and copy the daemon over to that folder with the appropriate name. NOTE: Don't forget check the file owner of v1-beta.2 binary.

```shell
mkdir -p $DAEMON_HOME/cosmovisor/upgrades/v1-beta.2/bin
cd $HOME/<ununifi-repo>
git pull
git checkout v1-beta.2
make build -B
cp ./build/ununifid $DAEMON_HOME/cosmovisor/upgrades/v1-beta.2/bin
```

Even though the cosmovisor's `DAEMON_ALLOW_DOWNLOAD_BINARIES` variable is set `true`, the cosmovisor uses binary which locates in $DAEMON_HOME/cosmovisor/upgrades/v1-beta.2/bin if exists.   
And you don't have to reboot cosmovisor when to do upgrade. So, after locating the binary into the appropriate place, you don't need to anything.

## Futher Help

If you need more help, please go to our discord https://discord.gg/GXTx9wjA.
