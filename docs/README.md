# JPYX documentation

## Environment

This is an example for Ubuntu.

```shell
apt update
apt install build-essential
cd ~
wget https://dl.google.com/go/go1.14.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.14.linux-amd64.tar.gz
echo export PATH='$PATH:/usr/local/go/bin:$HOME/go/bin' >> ~/.bashrc
source ~/.bashrc
```

## Install

```shell
mkdir -p /usr/local/src/github.com/lcnem
cd /usr/local/src/github.com/lcnem
git clone https://github.com/lcnem/jpyx.git
cd jpyx
git checkout v0.1.0
make install
```

## Setup genesis.json

```shell
jpyxd init [moniker] --chain-id jpyx-1
cd /usr/local/src/github.com/lcnem/jpyx
cp launch/genesis.json ~/.jpyxd/config/genesis.json
```

## Setup services

```shell
jpyxcli config chain-id jpyx-1
jpyxcli config trust-node true
```

### Daemon service

```shell
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

```shell
systemctl enable jpyxd
```

### REST service

```shell
vi /etc/systemd/system/jpyxrest.service
```

```txt
[Unit]
Description=JPYX Rest
After=network-online.target

[Service]
User=root
ExecStart=/root/go/bin/jpyxcli rest-server
Restart=always
RestartSec=3
LimitNOFILE=4096

[Install]
WantedBy=multi-user.target
```

```shell
systemctl enable jpyxrest
```
