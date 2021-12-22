#!/bin/bash

date

cd ~/ununifi
docker-compose down
sudo chown -c -R $USER:docker ~/.ununifi

# Todo: Create or Edit `~/.ununifi/config/genesis.json`.
# rm ~/.ununifi/config/genesis.json
# curl -L https://raw.githubusercontent.com/UnUniFi/chain/main/launch/ununifi-6-test/genesis-pretty.json -o ~/.ununifi/config/genesis.json
curl -O https://raw.githubusercontent.com/UnUniFi/chain/main/docker-compose.yml

docker pull ghcr.io/ununifi/ununifid:latest

# gentxs
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid gentx cauchye-a-test-temp 5000000000uguu --chain-id="ununifi-6-test" --from="cauchye-a-test-temp" --ip="a.test.ununifi.cauchye.net" --moniker="cauchye-a-test" --identity="cauchye-a-test" --website="https://cauchye.com" --node-id="f0e7dc092e1565ec5aa60d1341c5b6820e5a6c14" --pubkey="ununifivalconspub1zcjduepq0cywhfd6a9k306u7e8p0ps35c97hgumxmh4573vec5f3hkzkdp8qxykjwg"
sudo chown -c -R $USER:docker ~/.ununifi

docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid gentx cauchye-b-test-temp 5000000000uguu --chain-id="ununifi-6-test" --from="cauchye-b-test-temp" --ip="b.test.ununifi.cauchye.net" --moniker="cauchye-b-test" --identity="cauchye-b-test" --website="https://cauchye.com" --node-id="411160f7963c316a83da803daa09914986618531" --pubkey="ununifivalconspub1zcjduepq2zyearj8c695aqf9e0hzsdk38w787p7ljacnllgaxl2n8whvrnuq38ry6p"
sudo chown -c -R $USER:docker ~/.ununifi

docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid gentx tokyo-0-test-temp 5000000000uguu --chain-id="ununifi-6-test" --from="tokyo-0-test-temp" --ip="ununifi.testnet.validator.tokyo-0.neukind.network" --moniker="tokyo-0-test" --identity="tokyo-0-test" --node-id="1357ac5cd92b215b05253b25d78cf485dd899d55" --pubkey="ununifivalconspub1zcjduepq89wlgjv3ndhe49w695tkkpuv6s9tacnztthdwlved8cwwp83te6qp22dpm"
sudo chown -c -R $USER:docker ~/.ununifi

docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid gentx tokyo-1-test-temp 5000000000uguu --chain-id="ununifi-6-test" --from="tokyo-1-test-temp" --ip="ununifi.testnet.validator.tokyo-1.neukind.network" --moniker="tokyo-1-test" --identity="tokyo-1-test" --node-id="25006d6b85daeac2234bcb94dafaa73861b43ee3" --pubkey="ununifivalconspub1zcjduepq2aqtvv6kef76eh2y36q43m8klf2ku5ld9648fl59q5g9anw5dvssa5cclf"
sudo chown -c -R $USER:docker ~/.ununifi

docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid gentx genio01-test-temp 5000000000uguu --chain-id="ununifi-6-test" --from="genio01-test-temp" --ip="ununifi.testnet.validator.genio01.neukind.network" --moniker="genio01-test" --identity="genio01-test" --website="https://polkadot-coin.jp" --node-id="fe482d2bf89681d5684269e351b23e7ff3ba4b41" --pubkey="ununifivalconspub1zcjduepqgdp6v6pegrfmhj9sqc7d7rt6xhuqkpsr2rs4jk2gqdz9f44ejs8s65gjrr"
sudo chown -c -R $USER:docker ~/.ununifi

docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid gentx zofuku-japan-test-temp 5000000000uguu --chain-id="ununifi-6-test" --from="zofuku-japan-test-temp" --ip="ununifi.testnet.validator.zofuku-japan.neukind.network" --moniker="zofuku-japan-test" --identity="zofuku-japan-test" --website="https://zofuku.com" --node-id="c858ea83ef8ad23db02744d4f0ce7e3f592b4d94" --pubkey="ununifivalconspub1zcjduepqxmgct99g5e80rcf6vdqryncthsy6mu9fxgt22pzqw3mhje06rgzsfc0nhk"
sudo chown -c -R $USER:docker ~/.ununifi

docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid gentx zofuku-tokyo-test-temp 5000000000uguu --chain-id="ununifi-6-test" --from="zofuku-tokyo-test-temp" --ip="ununifi.testnet.validator.zofuku-tokyo.neukind.network" --moniker="zofuku-tokyo-test" --identity="zofuku-tokyo-test" --website="https://zofuku.com" --node-id="85131e04b014f9fee85b805783080c4518521c54" --pubkey="ununifivalconspub1zcjduepqddcpz325ydav2wpweccpdnad7pxs2tghfnpzd6yxu6s0gczqqnfq2astac"
sudo chown -c -R $USER:docker ~/.ununifi

docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid gentx chikako0903-test-temp 5000000000uguu --chain-id="ununifi-6-test" --from="chikako0903-test-temp" --ip="ununifi.testnet.validator.chikako0930.neukind.network" --moniker="chikako0903-test" --identity="chikako0903-test" --node-id="afc260c54043b671a5b084f905b565da1c7595fd" --pubkey="ununifivalconspub1zcjduepqtgeqrvfhw0st7ce7dzmzltlrmf6ddfr29n9z3237mfw4m3lsj6hsprmrac"
sudo chown -c -R $USER:docker ~/.ununifi

docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid gentx kurata0211-test-temp 5000000000uguu --chain-id="ununifi-6-test" --from="kurata0211-test-temp" --ip="ununifi.testnet.validator.kurata0211.neukind.network" --moniker="kurata0211-test" --identity="kurata0211-test" --node-id="7257ef8f772c8e41a2b68fdb0340697b60b9c433" --pubkey="ununifivalconspub1zcjduepqwklnky5mtt32u6vn6s9vpm257csgrnuzxpladzse9mdrcmhuq7wq5z6jrn"
sudo chown -c -R $USER:docker ~/.ununifi

docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid gentx keyplayers01-test-temp 5000000000uguu --chain-id="ununifi-6-test" --from="keyplayers01-test-temp" --ip="ununifi.testnet.validator.keyplayers01.neukind.network" --moniker="keyplayers01-test" --identity="keyplayers01-test" --node-id="320ec43486005c4022c8103b5087d851c4d544fe" --pubkey="ununifivalconspub1zcjduepq9chpalue7dr4cfqs7c8apfcztke6kz6e7e3ngjgr0rksmd3kfeaqh3ztv8"
sudo chown -c -R $USER:docker ~/.ununifi

docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid gentx keyplayers02-test-temp 5000000000uguu --chain-id="ununifi-6-test" --from="keyplayers02-test-temp" --ip="ununifi.testnet.validator.keyplayers02.neukind.network" --moniker="keyplayers02-test" --identity="keyplayers02-test" --node-id="65ae4a6e43abdb4bc4bc3ff830bdf8d4e889b4b6" --pubkey="ununifivalconspub1zcjduepq6r04zanwtv7p2a3undnnnxrf4h0eqere9359dl0rlnccm7nyc8wsr2l358"
sudo chown -c -R $USER:docker ~/.ununifi

docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid gentx tokyo-2-test-temp 5000000000uguu --chain-id="ununifi-6-test" --from="tokyo-2-test-temp" --ip="ununifi.testnet.validator.tokyo-2.neukind.network" --moniker="tokyo-2-test" --identity="tokyo-2-test" --node-id="caf792ed396dd7e737574a030ae8eabe19ecdf5c" --pubkey="ununifivalconspub1zcjduepqvmfcsnzpr4e7cg3xf789ka3ygangsmh2jv6nr46vhwcxlkzpkgxsyqw2kg"
sudo chown -c -R $USER:docker ~/.ununifi

docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid gentx tokyo-3-test-temp 5000000000uguu --chain-id="ununifi-6-test" --from="tokyo-3-test-temp" --ip="ununifi.testnet.validator.tokyo-3.neukind.network" --moniker="tokyo-3-test" --identity="tokyo-3-test" --node-id="796c62bb2af411c140cf24ddc409dff76d9d61cf" --pubkey="ununifivalconspub1zcjduepq6xfx9kw3uypx7690ry38gfn0hcqkddzjkqqxvpwsn9cwxvy4pe2sjf5ftm"
sudo chown -c -R $USER:docker ~/.ununifi

docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid gentx tokyo-4-test-temp 5000000000uguu --chain-id="ununifi-6-test" --from="tokyo-4-test-temp" --ip="ununifi.testnet.validator.tokyo-4.neukind.network" --moniker="tokyo-4-test" --identity="tokyo-4-test" --node-id="cea8d05b6e01188cf6481c55b7d1bc2f31de0eed" --pubkey="ununifivalconspub1zcjduepq6ag9yud9yn8q849dj8qsk870jsgvuev0nmaar408pqrcd4hty2xqc4z5hz"
sudo chown -c -R $USER:docker ~/.ununifi

docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid gentx toko1631-test-temp 5000000000uguu --chain-id="ununifi-6-test" --from="toko1631-test-temp" --ip="ununifi.testnet.validator.toko1631.neukind.network" --moniker="toko1631-test" --identity="toko1631-test" --node-id="062655515d4c12de85018a69adb3fcdb790f7d51" --pubkey="ununifivalconspub1zcjduepql04k05pahjtc0tdq2909gdj80pkrdjdpf6vcmkhqyzy829zk3gqql78sdf"
sudo chown -c -R $USER:docker ~/.ununifi

# collect-gentxs
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid collect-gentxs
sudo chown -c -R $USER:docker ~/.ununifi

# unsafe-reset-all
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid unsafe-reset-all
sudo chown -c -R $USER:docker ~/.ununifi

date
