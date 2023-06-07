# v2.0.0 to v2.1.0 Upgrade Guide

## Purpose of this upgrade

This proposal aims to do upgrade to UnUniFi `v2.1.0`.
This upgrade will enable everyone to deploy cosmwasm contract & update [ibc-go v7.0.1](https://github.com/cosmos/ibc-go/releases/tag/v7.0.1) to fix for the [huckleberry security advisory](https://forum.cosmos.network/t/ibc-security-advisory-huckleberry/10731).

## Brief guide

### Time

The upgrade height is `5630000`.
About June 4, 2023, 7:30 AM UTC.

## Go Requirement

You will need to be running go1.19 for this, same as the previous version. You can check with this command:

```shell
go version
```

## Setup

If the cosmovisor's `DAEMON_ALLOW_DOWNLOAD_BINARIES` variable is set `true`, no need to do the following steps, it will be downloaded automatically.
But, if `$DAEMON_HOME/cosmovisor/upgrades/v2_1/bin` already exists, the cosmovisor uses it.

If you use cosmovisor which we highly recommend, create the required folder, make the build, and copy the daemon over to that folder with the appropriate name.

```shell
# Download
mkdir -p $DAEMON_HOME/cosmovisor/upgrades/v2_1/bin
wget https://github.com/UnUniFi/chain/releases/download/v2.1.0/ununifid?checksum=md5:870390f317f30995ea0b4457df05c53b -O $DAEMON_HOME/cosmovisor/upgrades/v2_1/bin/ununifid
```

or

```shell
# Build
mkdir -p $DAEMON_HOME/cosmovisor/upgrades/v2_1/bin
cd $HOME/<your-ununifi-repo>
git pull
git checkout v2.1.0
make build -B
## to make sure the build status
# ./build/ununifid version
cp ./build/ununifid $DAEMON_HOME/cosmovisor/upgrades/v2_1/bin
```

NOTE: Don't forget check the file owner of v2 binary.

```shell
chmod 755 $DAEMON_HOME/cosmovisor/upgrades/v2_1/bin/ununifid
```

And you don't have to reboot cosmovisor when to do upgrade. So, after locating the binary into the appropriate place, you don't need to do anything.

## Futher Help

If you need more help, please go to [validator channel](https://discord.com/channels/762953633230356480/762953916593340416) in UnUniFi Discord.
