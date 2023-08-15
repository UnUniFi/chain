# UnUniFi

## Install

### Environment setup

Install the necessary libraries for the build.

```bash
cd $HOME
sudo apt update -y; sudo apt upgrade -y
sudo apt install -y jq git build-essential
```

Install Go. Use the 20.x series version.

```bash
$ wget https://go.dev/dl/go1.20.linux-amd64.tar.gz
$ sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.20.linux-amd64.tar.gz
$ vim ~/.bashrc
export PATH=$PATH:/usr/local/go/bin:$HOME/go/bin
$ source ~/.bashrc
$ go version
go version go1.20 linux/amd64
```

Clone the UnUniFi blockchain repository, check out the given branch, and build it with `make install` to build binaries.

```bash
git clone https://github.com/UnUniFi/chain chain_repo
cd chain_repo
git checkout v3.2.0
git pull
make install
```

Check the binary version.

```bash
ununifid version
```

### Start node for debugging

Initialize node

```bash
ununifid init [your-moniker] --chain-id [chain-id]
```

- `your-moniker` is like a nickname of your node.
- `chain-id` is the id of the blockchain network. (ex. In the current mainnet beta, the chain-id is `ununifi-beta-v1`.)

Download correct genesis.json

```bash
rm ~/.ununifi/config/genesis.json
curl -L https://raw.githubusercontent.com/UnUniFi/network/main/launch/[chain-id]/genesis.json -o ~/.ununifi/config/genesis.json
```

If you necessary, Edit config files.

- `~/.ununifi/config/app.toml`
  - `minimum-gas-prices` ... [https://docs.cosmos.network/v0.45/modules/auth/01_concepts.html#gas-fees](https://docs.cosmos.network/v0.45/modules/auth/01_concepts.html#gas-fees)
  - `pruning`
  - Enable defines if the API server should be enabled. `enable = true`
  - EnableUnsafeCORS defines if CORS should be enabled (unsafe - use it at your own risk). `enabled-unsafe-cors = true`
- `.ununifi/config/config.toml` ex. in the case of chain-id="ununifi-beta-v1", the possible settings are as follows.
  - `persistent-peers = "fa38d2a851de43d34d9602956cd907eb3942ae89@a.ununifi.cauchye.net:26656,404ea79bd31b1734caacced7a057d78ae5b60348@b.ununifi.cauchye.net:26656,1357ac5cd92b215b05253b25d78cf485dd899d55@[2600:1f1c:534:8f02:7bf:6b31:3702:2265]:26656,25006d6b85daeac2234bcb94dafaa73861b43ee3@[2600:1f1c:534:8f02:a407:b1c6:e8f5:94b]:26656,caf792ed396dd7e737574a030ae8eabe19ecdf5c@[2600:1f1c:534:8f02:b0a4:dbf6:e50b:d64e]:26656,796c62bb2af411c140cf24ddc409dff76d9d61cf@[2600:1f1c:534:8f02:ca0e:14e9:8e60:989e]:26656,cea8d05b6e01188cf6481c55b7d1bc2f31de0eed@[2600:1f1c:534:8f02:ba43:1f69:e23a:df6b]:26656"`

Start node.

```bash
ununifid start
```

## Keep your node stable

If you wanna keep your node stable, we recommend to use Cosmovisor.

[https://docs.cosmos.network/master/run-node/cosmovisor.html](https://docs.cosmos.network/master/run-node/cosmovisor.html)

### Setup

Set the following environment variables.
Some environment variables must be set to appropriate values for each node and each network.

`~/.profile`

```bash
export CHAIN_REPO=https://github.com/UnUniFi/chain
export CHAIN_REPO_BRANCHE=main
export TARGET=ununifid
export TARGET_HOME=.ununifi

# This value will be different for each node.
export MONIKER=<your-moniker>
export DL_CHAIN_BIN=

# This value is example of mainnet.
export CHAIN_ID=ununifi-beta-v1

# This value is example of mainnet.
export GENESIS_FILE_URL=https://raw.githubusercontent.com/UnUniFi/network/main/launch/ununifi-beta-v1/genesis.json
export SETUP_NODE_CONFIG_ENV=TRUE
export SETUP_NODE_ENV=TRUE
export SETUP_NODE_MASTER=TRUE
export DAEMON_NAME=ununifid

# This value will be different for each node.
export DAEMON_HOME=/home/<your-user>/.ununifi
export DAEMON_ALLOW_DOWNLOAD_BINARIES=true
export DAEMON_LOG_BUFFER_SIZE=512
export DAEMON_RESTART_AFTER_UPGRADE=true
export UNSAFE_SKIP_BACKUP=true
```

### Install Cosmovisor

```bash
cd $HOME
go install github.com/cosmos/cosmos-sdk/cosmovisor/cmd/cosmovisor@v1.0.0
mkdir -p $DAEMON_HOME/cosmovisor
mkdir -p $DAEMON_HOME/cosmovisor/genesis
mkdir -p $DAEMON_HOME/cosmovisor/genesis/bin
mkdir -p $DAEMON_HOME/cosmovisor/upgrades
cp ~/go/bin/$DAEMON_NAME $DAEMON_HOME/cosmovisor/genesis/bin
```

### Register daemon service

Create service file.

`/lib/systemd/system/cosmovisor.service`

```shell
[Unit]
Description=Cosmovisor daemon
After=network-online.target
[Service]
Environment="DAEMON_NAME=ununifid"
Environment="DAEMON_HOME=/home/<your-user>/.ununifi"
Environment="DAEMON_RESTART_AFTER_UPGRADE=true"
Environment="DAEMON_ALLOW_DOWNLOAD_BINARIES=true"
Environment="DAEMON_LOG_BUFFER_SIZE=512"
Environment="UNSAFE_SKIP_BACKUP=true"
User=<your-user>
ExecStart=/home/<your-user>/go/bin/cosmovisor start
Restart=always
RestartSec=3
LimitNOFILE=infinity
LimitNPROC=infinity
[Install]
WantedBy=multi-user.target
```

Set systemctl

```bash
sudo systemctl daemon-reload
sudo systemctl restart systemd-journald
sudo systemctl enable cosmovisor
```

### Start node with cosmovisor

Start node.

```bash
sudo systemctl start cosmovisor
```

You can check your node status with the following command.

```bash
$ sudo systemctl status cosmovisor

# Log sample
● cosmovisor.service - Cosmovisor daemon
      Loaded: loaded (/lib/systemd/system/cosmovisor.service; enabled; vendor preset: enabled)
      Active: active (running) since Mon 2022-05-16 17:57:24 +08; 46s ago
    Main PID: 232015 (cosmovisor)
      Tasks: 29 (limit: 36046)
      Memory: 123.3M
      CGroup: /system.slice/cosmovisor.service
              ├─232015 /home/ununifi/go/bin/cosmovisor start
              └─232029 /home/ununifi/.ununifi/cosmovisor/genesis/bin/ununifid start

May 16 17:57:26 cosmovisor[232029]: 5:57PM INF Completed ABCI Handshake - Tendermint and App are synced appHash= appHeight=0 module=consensus
May 16 17:57:26 cosmovisor[232029]: 5:57PM INF Version info block=11 p2p=8 tendermint_version=v0.34.16
May 16 17:57:26 cosmovisor[232029]: 5:57PM INF This node is not a validator addr=83FD137D6541F5198D7107FE6B75ACDDBCC72329 module=consensus pubKey=B8tjjYkW51s6bFqDNRIhJdZJsTR68Ez>
May 16 17:57:26 cosmovisor[232029]: 5:57PM INF P2P Node ID ID=729318b4ee913b1d56a1fe22b93860aa01bff82a file=/home/ununifi/.ununifi/config/node_key.json module=p2p
May 16 17:57:26 cosmovisor[232029]: 5:57PM INF Adding persistent peers addrs=["fa38d2a851de43d34d9602956cd907eb3942ae89@a.ununifi.cauchye.net:26656","404ea79bd31b1734caacced7a05>
May 16 17:57:26 cosmovisor[232029]: 5:57PM INF Adding unconditional peer ids ids=[] module=p2p
May 16 17:57:26 cosmovisor[232029]: 5:57PM INF Add our address to book addr={"id":"729318b4ee913b1d56a1fe22b93860aa01bff82a","ip":"0.0.0.0","port":26656} book=/home/ununifi/.unu>
May 16 17:57:26 cosmovisor[232029]: 5:57PM INF Starting Node service impl=Node
May 16 17:57:26 cosmovisor[232029]: 5:57PM INF Genesis time is in the future. Sleeping until then... genTime=2022-05-17T03:00:00Z
May 16 17:57:26 cosmovisor[232029]: 5:57PM INF Starting pprof server laddr=localhost:6060
```

## License

Copyright © UnUniFi development team. All rights reserved.

Licensed under the [Apache v2 License](LICENSE.md).
