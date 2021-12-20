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

```bash
docker run -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid init [moniker] --chain-id [chain-id]
mkdir ununifi
cd ununifi
curl -L https://raw.githubusercontent.com/UnUniFi/chain/main/launch/[chain-id]/genesis.json -o ~/.ununifi/config/genesis.json
curl -O https://raw.githubusercontent.com/UnUniFi/chain/main/docker-compose.yml
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

### Clone

```bash
git clone https://github.com/UnUniFi/chain.git
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

Copyright © CauchyE, Inc. All rights reserved.

Licensed under the [Apache v2 License](LICENSE.md).
