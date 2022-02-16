#!/bin/bash

date

docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid:test ununifid gentx cauchye-a-private-test-temp 200000000000uguu --chain-id="ununifi-8-private-test" --from="cauchye-a-private-test-temp" --ip="a.private-test.ununifi.cauchye.net" --moniker="cauchye-a-private-test" --identity="cauchye-a-private-test" --website="https://cauchye.com" --node-id="26b87ecbd58732af7dd332786c8659327708eb11" --pubkey="{\"@type\":\"/cosmos.crypto.ed25519.PubKey\",\"key\":\"76XR0s9sp5F2YZ8iQXGM/oNJgbvrI01oEWGvQF0HLIU=\"}"
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid:test ununifid gentx cauchye-b-private-test-temp 200000000000uguu --chain-id="ununifi-8-private-test" --from="cauchye-b-private-test-temp" --ip="b.private-test.ununifi.cauchye.net" --moniker="cauchye-b-private-test" --identity="cauchye-b-private-test" --website="https://cauchye.com" --node-id="a295e623b8722b620df0196500bb419292a3fe76" --pubkey="{\"@type\":\"/cosmos.crypto.ed25519.PubKey\",\"key\":\"P905SJStdg2tPlLcABb5jH8tkk35jWQ0IptqZUT+6as=\"}"
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid:test ununifid gentx cauchye-c-private-test-temp 200000000000uguu --chain-id="ununifi-8-private-test" --from="cauchye-c-private-test-temp" --ip="c.private-test.ununifi.cauchye.net" --moniker="cauchye-c-private-test" --identity="cauchye-c-private-test" --website="https://cauchye.com" --node-id="8317caab7e01954839c3b28c866b62fe42212161" --pubkey="{\"@type\":\"/cosmos.crypto.ed25519.PubKey\",\"key\":\"Pgrc3zr8AgNjKPvKUg+f4eqa3LskrgQyP979ZN2h5EM=\"}"
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid:test ununifid gentx cauchye-d-private-test-temp 200000000000uguu --chain-id="ununifi-8-private-test" --from="cauchye-d-private-test-temp" --ip="d.private-test.ununifi.cauchye.net" --moniker="cauchye-d-private-test" --identity="cauchye-d-private-test" --website="https://cauchye.com" --node-id="d111c086b9eb51b29a144bd7a95b83015279a2d4" --pubkey="{\"@type\":\"/cosmos.crypto.ed25519.PubKey\",\"key\":\"RmN+DShaUkliEFkc6Yi2GCRdfVRXtBqHJlVgKup7VC0=\"}"

sudo chown -c -R $USER:docker ~/.ununifi

date
