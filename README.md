# JPYX

## Install

### Environment setup

This is an example for Ubuntu.

```bash
sudo apt install docker.io -y
sudo curl -L "https://github.com/docker/compose/releases/download/1.28.4/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose

sudo gpasswd -a $(whoami) docker
sudo chgrp docker /var/run/docker.sock
sudo systemctl enable docker
sudo systemctl restart docker
```

```bash
vi /etc/docker/daemon.json
```

```json
{
  "ipv6": true,
  "fixed-cidr-v6": "2001:db8:1::/64"
}
```

### Join network

```bash
git clone https://github.com/lcnem/jpyx.git
cd jpyx
docker run -v ~/.jpyx:/root/.jpyx lcnem/jpyx:next [moniker] --chain-id [chain-id]
cp launch/[chain-id]/genesis.json ~/.jpyx/config/genesis.json
docker-compose up -d
```

## Deprecated way of Installation

### Environment setup

This is an example for Ubuntu.

```bash
sudo apt update
sudo apt install build-essential
cd ~
wget https://dl.google.com/go/go1.15.6.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.15.6.linux-amd64.tar.gz
echo export PATH='$PATH:/usr/local/go/bin:$HOME/go/bin' >> ~/.bashrc
source ~/.bashrc
```

### Clone

```bash
git clone https://github.com/lcnem/jpyx.git
cd jpyx
make install
```

### Config daemon

```bash
jpyxd init [moniker] --chain-id [chain-id]
cp launch/[chain-id]/genesis.json ~/.jpyx/config/genesis.json
```

### Register daemon service

```bash
vi /etc/systemd/system/jpyxd.service
```

```txt
[Unit]
Description=JPYX Node
After=network-online.target

[Service]
User=root
ExecStart=/root/go/bin/jpyxd start
Restart=always
RestartSec=3
LimitNOFILE=4096

[Install]
WantedBy=multi-user.target
```

```bash
systemctl enable jpyxd
```

## License

Forked from [Kava](github.com/Kava-Labs/kava).
Thanks Kava Team.

Copyright © LCNEM, Inc. All rights reserved.

Licensed under the [Apache v2 License](LICENSE.md).
