BINARY=./build/ununifid
CHAIN_DIR=./data
CHAINID_1=test
NODE_HOME=$CHAIN_DIR/$CHAINID_1
BINARY_MAIN_TOKEN=uguu
VAL1=val
FAUCET=faucet
USER1=user1
USER2=user2
USER3=user3
USER4=user4
PRICEFEED=pricefeed

VAL_MNEMONIC_1="figure web rescue rice quantum sustain alert citizen woman cable wasp eyebrow monster teach hockey giant monitor hero oblige picnic ball never lamp distance"
FAUCET_MNEMONIC_1="chimney diesel tone pipe mouse detect vibrant video first jewel vacuum winter grant almost trim crystal similar giraffe dizzy hybrid trigger muffin awake leader"
USER_MNEMONIC_1="supply release type ostrich rib inflict increase bench wealth course enter pond spare neutral exact retire thing update inquiry atom health number lava taste"
USER_MNEMONIC_2="canyon second appear story film people resist slam waste again race rifle among room hip icon marriage sea quality prepare only liquid column click"
USER_MNEMONIC_3="among follow tooth egg unhappy city road expire solution visit visa skate allow network tissue slogan rose toddler develop utility negative peasant ostrich toward"
USER_MNEMONIC_4="charge split umbrella day gauge two orphan random human clerk buzz funny cabin purse fluid lecture blouse keen twist loud animal supply hat scare"
PRICEFEED_MNEMONIC="jelly fortune hire delay impose daughter praise amazing patch gesture easy achieve intact genre swamp gossip aisle arrest item seek inherit cradle hover involve"

VAL_ADDRESS_1=ununifi1a8jcsmla6heu99ldtazc27dna4qcd4jygsthx6
FAUCET_ADDRESS_1=ununifi1d6zd6awgjxuwrf4y863c9stz9m0eec4ghfy24c
USER_ADDRESS_1=ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w
USER_ADDRESS_2=ununifi1v0h8j7x7kfys29kj4uwdudcc9y0nx6twwxahla
USER_ADDRESS_3=ununifi1y3t7sp0nfe2nfda7r9gf628g6ym6e7d44evfv6
USER_ADDRESS_4=ununifi1pp2ruuhs0k7ayaxjupwj4k5qmgh0d72wrdyjyu
PRICEFEED_ADDRESS=ununifi1h7ulktk5p2gt7tnxwhqzlq0yegq47hum0fahcr

conf="--home=$NODE_HOME --chain-id=$CHAINID_1 --keyring-backend=test -y --broadcast-mode=sync"
# conf="--home=$NODE_HOME --chain-id=$CHAINID_1 --keyring-backend=test -y --broadcast-mode=sync | grep txhash | awk '{ print $2 }'| xargs -I {} sh -c 'sleep 5; $0 q tx {}' $BINARY"

SSH_PREV_KEY_LOCATION=/your/ssh/key/location
ALPHA_NODE_URL=ununifi-alpha-test.cauchye.net
