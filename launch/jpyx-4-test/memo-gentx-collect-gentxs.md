# Memo of gentx, collect-gentxs, deploy new network procedures

## 1st gentx

In each node,

```shell
# lcnem-a
cd ~/jpyx
docker-compose down
docker pull lcnem/jpyx:latest

cd ~/.jpyx/config
rm -rf ~/.jpyx/config/gentx
rm ~/.jpyx/config/genesis.json
curl -O https://raw.githubusercontent.com/lcnem/jpyx/main/launch/jpyx-3-test/genesis.json

cd ~/jpyx

docker run -v ~/.jpyx:/root/.jpyx lcnem/jpyx jpyxd gentx main 250000000000ujcbn --chain-id="jpyx-3-test" --from="lcnem-test-a" --ip="a.test.jpyx.lcnem.net" --moniker="lcnem-test-a" --node-id="6fc4109698025d8ee20c09a78fd6b7fb858c26e2" --pubkey="jpyxvalconspub1zcjduepq5xq2kf5n7kl9snq2rh9dhq2ltq8y5ylzj7hrjprajyxjfe743whq52rzqf"

# lcnem-b
cd ~/jpyx
docker-compose down
docker pull lcnem/jpyx:latest

cd ~/.jpyx/config
rm -rf ~/.jpyx/config/gentx
rm ~/.jpyx/config/genesis.json
curl -O https://raw.githubusercontent.com/lcnem/jpyx/main/launch/jpyx-3-test/genesis.json

cd ~/jpyx

docker run -v ~/.jpyx:/root/.jpyx lcnem/jpyx jpyxd gentx main 250000000000ujcbn --chain-id="jpyx-3-test" --from="lcnem-test-b" --ip="b.test.jpyx.lcnem.net" --moniker="lcnem-test-b" --node-id="7fc846bfde6ccd378dd736ac9976f63ed2afdd52" --pubkey="jpyxvalconspub1zcjduepqtj9srjwgcuctaqmm9k3qql4nag6u67472ve73z5grpzth5u2ldpqsh0tqu"

# lcnem-c
cd ~/jpyx
docker-compose down
docker pull lcnem/jpyx:latest

cd ~/.jpyx/config
rm -rf ~/.jpyx/config/gentx
rm ~/.jpyx/config/genesis.json
curl -O https://raw.githubusercontent.com/lcnem/jpyx/main/launch/jpyx-3-test/genesis.json

cd ~/jpyx
docker run -v ~/.jpyx:/root/.jpyx lcnem/jpyx jpyxd gentx main 250000000000ujcbn --chain-id="jpyx-3-test" --from="lcnem-test-c" --ip="c.test.jpyx.lcnem.net" --moniker="lcnem-test-c" --node-id="e676f54115ea73f893a1cc5ba4c334f5524f57ca" --pubkey="jpyxvalconspub1zcjduepqz0nqw0rcxyftlu7tlemvw4pjfgg5h87j60yc46vcmhjm5m6c7x9sez8khz"

# lcnem-d
cd ~/jpyx
docker-compose down
docker pull lcnem/jpyx:latest

cd ~/.jpyx/config
rm -rf ~/.jpyx/config/gentx
rm ~/.jpyx/config/genesis.json
curl -O https://raw.githubusercontent.com/lcnem/jpyx/main/launch/jpyx-3-test/genesis.json

cd ~/jpyx

docker run -v ~/.jpyx:/root/.jpyx lcnem/jpyx jpyxd gentx main 250000000000ujcbn --chain-id="jpyx-3-test" --from="lcnem-test-d" --ip="d.test.jpyx.lcnem.net" --moniker="lcnem-test-d" --node-id="d823c464320c4cdd1f9ae1d8595e3451cd97d520" --pubkey="jpyxvalconspub1zcjduepq9anualsu09egnk7j4ec6fcla63460zk9yvqeslksxqyyskxsd0gqjhxvkm"

```

## 2nd

Copy files generated with 1 procedure to some node's `~/.jpyx/config/gentx` directory,

## 3rd collect-gentxs

```shell
docker run -v ~/.jpyx:/root/.jpyx lcnem/jpyx jpyxd collect-gentxs
```

## 4th backup genesis.json(before unprettify)

Copy genesis.json file on node server to your local working space genesis-pretty.json.

## 5th unprettigy genesis.json

```shell
cd ~/.jpyx/config
curl -O https://raw.githubusercontent.com/lcnem/jpyx/main/launch/unprettify.js
node unprettify.js
rm unprettify.js
```

Now, genesis.json is completed!

## Reset old network data and start new network

```shell
jpyxd-temporal unsafe-reset-all
docker-compose up -d
```

That's all!
