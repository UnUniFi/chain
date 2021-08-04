#!/bin/bash

# delete old gentx folder
rm -rf ~/.jpyx/config/gentx

# lcnem-a
docker exec -it $(docker ps -qf "name=jpyx") gentx main 5000000000ujsmn --chain-id="jpyx-2" --from="lcnem-a" --ip="a.jpyx.lcnem.net" --moniker="lcnem-a" --identity="lcnem-a" --node-id="8ee6060a7be74cddd3bd620a1206cf3b1a2da0cb" --pubkey="jpyxvalconspub1zcjduepqfqafaz96awchqfyrtzddhuj5t7rug4nllk8yexuvl0yzmee5dt2srkkesq"

# lcnem-b
docker exec -it $(docker ps -qf "name=jpyx") gentx main 5000000000ujsmn --chain-id="jpyx-2" --from="lcnem-b" --ip="b.jpyx.lcnem.net" --moniker="lcnem-b" --identity="lcnem-b" --node-id="1d0651e816b8619bd97d892327a0609f829460a7" --pubkey="jpyxvalconspub1zcjduepq73we20e2l69450m8acplrdlq0w0hshn5pcwzgh8gnvgxey07552svk2lj2"

# mano-san genio
docker exec -it $(docker ps -qf "name=jpyx") gentx mano-san 5000000000ujsmn --chain-id="jpyx-2" --from="mano-san-1" --ip="jpyx.mainnet.validator.genio.neukind.network" --moniker="genio" --identity="genio" --node-id="a4d9b5360a0a8e3887f9ea50807a56d96a6174e4" --pubkey="jpyxvalconspub1zcjduepqc0qh49vf7gypaj4evzn50dzala5ysh0w0tevgdqldv3nm95d8v6sqvx8fg"

# niikura-san-1 zofuku-japan
docker exec -it $(docker ps -qf "name=jpyx") gentx niikura-kun-1 5000000000ujsmn --chain-id="jpyx-2" --from="niikura-san-1" --ip="jpyx.mainnet.validator.zofuku-japan.neukind.network" --moniker="zofuku-japan" --identity="zofuku_japan" --node-id="e572c2ed09f5c996ba6aef1ca01308d465fe5d3b" --pubkey="jpyxvalconspub1zcjduepqw5qee4h96ferfjfcu4gxkwdtq4zq3tdw5gat4sgvjrs6m8rld2gsm3a59h"

# niikura-san-2 zofuku-tokyo
docker exec -it $(docker ps -qf "name=jpyx") gentx niikura-kun-2 5000000000ujsmn --chain-id="jpyx-2" --from="niikura-san-2" --ip="jpyx.mainnet.validator.zofuku-tokyo.neukind.network" --moniker="zofuku-tokyo" --identity="zofuku_tokyo" --node-id="9bb3dd0db9f6c27294f0aa4ac65a6f3db81c33b6" --pubkey="jpyxvalconspub1zcjduepqd06a5ae7343tv662728nugjl2rhd8gml3upal79n38tp4js299qswr6qhh"

# chikako-kurita-1 chikako0903
docker exec -it $(docker ps -qf "name=jpyx") gentx chikako-kurata-san 5000000000ujsmn --chain-id="jpyx-2" --from="chikako-kurita-1" --ip="jpyx.mainnet.validator.chikako0930.neukind.network" --moniker="chikako0903" --identity="chikako0903" --node-id="eab5fcd221e9970735eaa3aae8397bf70bdd04b4" --pubkey="jpyxvalconspub1zcjduepqcx6fccqggvqhyfpx00gkthv49we40pdglae3hkrwhy0q8drv7qjqtlhjfv"

# yoichiro-kurita-1 kurata0211
docker exec -it $(docker ps -qf "name=jpyx") gentx yoichiro-kurata-san 5000000000ujsmn --chain-id="jpyx-2" --from="yoichiro-kurita-1" --ip="jpyx.mainnet.validator.kurata0211.neukind.network" --moniker="kurata0211" --identity="kurata0211" --node-id="e2f9b5e01e115dc74c285ee20dbba606a9263e96" --pubkey="jpyxvalconspub1zcjduepqsfdtdq2fqedlwkd93pr25al6a2gva0tzlwu60y8nz7hgwmzrah0qqt324h"

# takano-san-1 keyplayers01
docker exec -it $(docker ps -qf "name=jpyx") gentx takano-san-1 5000000000ujsmn --chain-id="jpyx-2" --from="takano-san-1" --ip="jpyx.mainnet.validator.keyplayers01.neukind.network" --moniker="keyplayers01" --identity="keyplayers01" --node-id="661fbed5faf614563264fcf91c7ce292fac54bb2" --pubkey="jpyxvalconspub1zcjduepqucxcr7nl5l49dsc04r036ea8vkwhd9grar6xgxxtg4dhz967yh7qdq4p8k"

# takano-san-2 keyplayers02
docker exec -it $(docker ps -qf "name=jpyx") gentx takano-san-2 5000000000ujsmn --chain-id="jpyx-2" --from="takano-san-2" --ip="jpyx.mainnet.validator.keyplayers02.neukind.network" --moniker="keyplayers02" --identity="keyplayers02" --node-id="c29c9cfc1d3d588de78a350f73f31fd4c8dc4cdd" --pubkey="jpyxvalconspub1zcjduepqx86xjj20tw7wt0cra9cpyh4zmuka9cx05255k24ypyt97ysdxsqsm2emyt"

# tokyo-0
docker exec -it $(docker ps -qf "name=jpyx") gentx neukind-tokyo-0 5000000000ujsmn --chain-id="jpyx-2" --from="tokyo-0" --ip="jpyx.mainnet.validator.tokyo-0.neukind.network" --moniker="tokyo-0" --identity="tokyo-0" --node-id="809ca087ef81466d49900aad02ad265b65964605" --pubkey="jpyxvalconspub1zcjduepq06hef840twnjqeemqqx4mhrtcydm4gcvxdedz5txtf9088s5w5qsjx008n"

# tokyo-1
docker exec -it $(docker ps -qf "name=jpyx") gentx neukind-tokyo-1 5000000000ujsmn --chain-id="jpyx-2" --from="tokyo-1" --ip="jpyx.mainnet.validator.tokyo-1.neukind.network" --moniker="tokyo-1" --identity="tokyo-1" --node-id="678bba63e5aa3a8d29ede81769f3568e721756fc" --pubkey="jpyxvalconspub1zcjduepq78mjsrtmgt96uze02j4s43w32fcl5ua0vlm0hx08lw4pm88uyuts5uc00w"

# tokyo-2
docker exec -it $(docker ps -qf "name=jpyx") gentx neukind-tokyo-2 5000000000ujsmn --chain-id="jpyx-2" --from="tokyo-2" --ip="jpyx.mainnet.validator.tokyo-2.neukind.network" --moniker="tokyo-2" --identity="tokyo-2" --node-id="091362955d0079773990ac1a56ee3602b103134b" --pubkey="jpyxvalconspub1zcjduepqa23wcrkccpaa3dflgq3x0rq3s99ptrwvywhdlgyyt86ulmhm6sqq0mhufy"

# tokyo-3
docker exec -it $(docker ps -qf "name=jpyx") gentx neukind-tokyo-3 5000000000ujsmn --chain-id="jpyx-2" --from="tokyo-3" --ip="jpyx.mainnet.validator.tokyo-3.neukind.network" --moniker="tokyo-3" --identity="tokyo-3" --node-id="d3cfbc7eb166b21f879bfbf20457ae4ee2620a24" --pubkey="jpyxvalconspub1zcjduepqlv0z8epq90stsavujfhnjph9ru30eqqsjj6eeeqa3j5ddmkuazeqqz994q"

# tokyo-4
docker exec -it $(docker ps -qf "name=jpyx") gentx neukind-tokyo-4 5000000000ujsmn --chain-id="jpyx-2" --from="tokyo-4" --ip="jpyx.mainnet.validator.tokyo-4.neukind.network" --moniker="tokyo-4" --identity="tokyo-4" --node-id="b9af9c349e7410ecde657e64c6633e3e4cd6764e" --pubkey="jpyxvalconspub1zcjduepqulr7qeykw4rauddf5u5u0vzvc8pxw3ld2e6mcar7yvrtn9nmstpqkfetmh"

# add val iijima-san-1 toko1631
docker exec -it $(docker ps -qf "name=jpyx") gentx iijima-san-1 5000000000ujsmn --chain-id="jpyx-2" --from="iijima-san-1" --ip="jpyx.mainnet.validator.toko1631.neukind.network" --moniker="toko1631" --identity="toko1631" --node-id="d2a672562c275e4651b76ca3359809942afe4886" --pubkey="jpyxvalconspub1zcjduepqg0ug5x8xt8y52ydlxqrfw2de8a22mjwj7an9jj2mln5z6qp2h82sn4yeyh"


# gentx-"hash".json into genesis.json
docker exec -it $(docker ps -qf "name=jpyx") collect-gentxs