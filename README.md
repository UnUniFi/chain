# UnUniFi

The Docker image will be automatically built by GitHub Container Registry when releases are created.

## Install

### Environment setup

This is an example for Ubuntu.

```bash
sudo apt install docker.io -y
sudo curl -L "https://github.com/docker/compose/releases/download/1.28.6/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose

sudo gpasswd -a $(whoami) docker
sudo chgrp docker /var/run/docker.sock
sudo systemctl enable docker
sudo systemctl restart docker
```

### Join network

Get resources.

```bash
docker run -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid init [moniker] --chain-id [chain-id]
sudo chown -c -R $USER:docker ~/.ununifi
mkdir ~/ununifi
cd ~/ununifi
curl -L https://raw.githubusercontent.com/UnUniFi/chain/main/launch/[chain-id]/genesis.json -o ~/.ununifi/config/genesis.json
curl -O https://raw.githubusercontent.com/UnUniFi/chain/main/docker-compose.yml
```

Note

- `moniker` is like a nickname of the node.
- `chain-id` is the id of the blockchain network. (ex. In the current public testnet, the chain-id is `ununifi-6-test`.)

Edit config files.

- `~/.ununifi/config/app.toml`
  - `minimum-gas-prices` ... [https://docs.cosmos.network/v0.42/modules/auth/01_concepts.html#gas-fees](https://docs.cosmos.network/v0.42/modules/auth/01_concepts.html#gas-fees)
  - `pruning`
  - Enable defines if the API server should be enabled. `enable = true`
  - EnableUnsafeCORS defines if CORS should be enabled (unsafe - use it at your own risk). `enabled-unsafe-cors = true`
- `.ununifi/config/config.toml` ex. in the case of chain-id="ununifi-6-test", the possible settings are as follows.
  - `persistent-peers = "f0e7dc092e1565ec5aa60d1341c5b6820e5a6c14@a.test.ununifi.cauchye.net:26656,411160f7963c316a83da803daa09914986618531@b.test.ununifi.cauchye.net:26656,1357ac5cd92b215b05253b25d78cf485dd899d55@ununifi.testnet.validator.tokyo-0.neukind.network:26656,25006d6b85daeac2234bcb94dafaa73861b43ee3@ununifi.testnet.validator.tokyo-1.neukind.network:26656"`

Start node.

```bash
docker-compose up -d
```

## Deprecated way of Installation

### Environment setup

This is an example for Ubuntu.

```bash
sudo apt update
sudo apt install build-essential
cd ~
wget https://dl.google.com/go/go1.16.4.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.16.4.linux-amd64.tar.gz
echo export PATH=\$PATH:/usr/local/go/bin:\$HOME/go/bin >> ~/.bashrc
source ~/.bashrc
```

### Install ununifid

```bash
git clone https://github.com/UnUniFi/chain.git UnUniFi
cd UnUniFi
make install
```

### Config daemon

```bash
ununifid init [moniker] --chain-id [chain-id]
cp launch/[chain-id]/genesis.json ~/.ununifi/config/genesis.json
```

### Register daemon service

```bash
vi /etc/systemd/system/ununifid.service
```

```txt
[Unit]
Description=UnUniFi Node
After=network-online.target

[Service]
User=root
ExecStart=/root/go/bin/ununifid start
Restart=always
RestartSec=3
LimitNOFILE=4096

[Install]
WantedBy=multi-user.target
```

```bash
systemctl enable ununifid
```

## License

Forked from [Kava](github.com/Kava-Labs/kava).
Thanks Kava Team.

Copyright © UnUniFi development team. All rights reserved.

Licensed under the [Apache v2 License](LICENSE.md).
