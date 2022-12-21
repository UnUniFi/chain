package v1_beta4

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func BankSendList(ctx sdk.Context) (ResultList, error) {
	ctx.Logger().Info(fmt.Sprintf("bank send list:%s", UpgradeName))

	// Read file and get list
	var result ResultList
	if err := json.Unmarshal([]byte(BANK_SEND_LIST), &result); err != nil {
		panic(err)
	}
	return result, nil
}

const BANK_SEND_LIST string = `{
  "validator": [
    {
      "toAddress": "ununifi1jzr5ux9ydjmrch54c04rjt3fr5pdlrx08ng6p8",
      "amount": 20000000,
      "denom": "uguu",
      "vesting_starts": 1703084400,
      "vesting_ends": 1734620400
    },
    {
      "toAddress": "ununifi1uny0f22uj5mn93sgg5d80nhnl7608fd9ap4ld6",
      "amount": 3000000000000,
      "denom": "uguu",
      "vesting_starts": 1703084400,
      "vesting_ends": 1734620400
    },
    {
      "toAddress": "ununifi1llr0q0ggvnyu7fxh57wxkuvmlpyf6ygmknga26",
      "amount": 3000000000000,
      "denom": "uguu",
      "vesting_starts": 1703084400,
      "vesting_ends": 1734620400
    }
  ],
  "lendValidator": [
    {
      "toAddress": "ununifi1n92vgzkyn32n4wgk0wetq5xw5j5v8zwseetg8f",
      "amount": 511145986071,
      "denom": "uguu",
      "vesting_starts": 1703084400,
      "vesting_ends": 1734620400
    },
    {
      "toAddress": "ununifi1zqt3snrfpq7zlqvwzvm4v926w26pcmrutklsl9",
      "amount": 511426337194,
      "denom": "uguu",
      "vesting_starts": 1703084400,
      "vesting_ends": 1734620400
    },
    {
      "toAddress": "ununifi1p705f2vfnqx6t0k68my4ah6w4gdcujs3yvcwua",
      "amount": 464896876114,
      "denom": "uguu",
      "vesting_starts": 1703084400,
      "vesting_ends": 1734620400
    },
    {
      "toAddress": "ununifi1gwcrz4cnzfurdxfcz7k4uklcrqd6prrme8m66s",
      "amount": 441655905328,
      "denom": "uguu",
      "vesting_starts": 1703084400,
      "vesting_ends": 1734620400
    },
    {
      "toAddress": "ununifi13ua8cusfmx3lwta4n76vjuf6rarr5m90hggnvq",
      "amount": 441591206812,
      "denom": "uguu",
      "vesting_starts": 1703084400,
      "vesting_ends": 1734620400
    },
    {
      "toAddress": "ununifi1ca3k5gk7elrpej8rd74hruu2pm8dn5jzny2qqe",
      "amount": 288264021643,
      "denom": "uguu",
      "vesting_starts": 1703084400,
      "vesting_ends": 1734620400
    },
    {
      "toAddress": "ununifi1w92q2r53jgvyvcuuwmlggrvum7c6j72y403puk",
      "amount": 309526695085,
      "denom": "uguu",
      "vesting_starts": 1703084400,
      "vesting_ends": 1734620400
    },
    {
      "toAddress": "ununifi19f0w9svr905fhefusyx4z8sf83j6et0gkfnrpz",
      "amount": 302587573611,
      "denom": "uguu",
      "vesting_starts": 1703084400,
      "vesting_ends": 1734620400
    }
  ],
  "ecocsytemDevelopment": [
    {
      "toAddress": "ununifi1ydk5djlj53p84gq9c3fhxe7mt3urhk0g5sq0vr",
      "amount": 15652505410,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1k8f9hre2szp7p7mjwp70kaydz5hl2lunuzvp78",
      "amount": 8441578742,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1t3syzcj74az3kzta8u5jpkxn6ft7qdumwdyqc0",
      "amount": 6374629312,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1d5csc65hy3l8yk9s8rges7zxl9fwka88f6waec",
      "amount": 4359974259,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi12x2trndk9p5umj64jjvn2j9kgguxfrujfzks9u",
      "amount": 3537503493,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1qxfymvcz8h2ahnapezz2ed4awh29qswygt2prx",
      "amount": 2137028121,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi182tul2dksmwur370ks4ghel0yr8mmgrf9g45ay",
      "amount": 2003170316,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1aakg8cq30w03xn3gldaf9mh04sr6wume8ytgfc",
      "amount": 13024479227,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi13kaqr9vzfxxpmerplda2me9eg3y8jyhgnc2f9n",
      "amount": 1737803088,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1j4ng55e6snak7gyvzwx9xtsjfph752pez22cj6",
      "amount": 1568719544,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi12fvdlw4ltmexpvj957gfintnhrrr4n4ay7u503",
      "amount": 6792443400,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1mcvgt3306fw8zx95hyacem932jusvu5epqalqy",
      "amount": 2470498444,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1e34uc97tty95r70jx7ek7gkzf8aat84m3u6w5d",
      "amount": 1343274819,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1ed2wt0dgnfrntns3gwsjzfkqhtprr3krjku4l8",
      "amount": 1324487759,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1pvzssder5a5md4wdhjyzdhtxmcrpvp9ctd02jj",
      "amount": 1293958785,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1jj9dysvjwyqvjv86gr8467sr3hjqc3why5vkga",
      "amount": 1997755953,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1ug5x6ldl9xfqkwl70qn0tuxkudwu5jk83tz0fq",
      "amount": 1056772148,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1nl4v9kcumsuc6uejqhmqrqsa9a2ju0jmf7gwur",
      "amount": 1056772148,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1zkcuchahmc0n4mrqfzyhyaalypdzsw2gqrqxr3",
      "amount": 1019198027,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1fsg8xkl0vyhz2ha2kjzzk9ardpa27uczx5kws4",
      "amount": 939353020,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1qdyn7rss0ad6g85wj2ad28a52cvsqr56nrqe5g",
      "amount": 1502964832,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi18g8870c64sr3t4v6802rj2h4m2xwj6udeuuj6h",
      "amount": 901778899,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1vxg66xlzymyc5fd9hwk602xa0tg09dms8dg509",
      "amount": 27476759026,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1y4uqc6tf435saqlqmdeswz5w8chtxhrhp6w6x7",
      "amount": 786708154,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1lul6f99rxw0d9j5rn39mx3djndldu8yt0sstff",
      "amount": 2428006775,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1xgw4zsxtm7mx0zav83an7pcrschkr8rxr5r7j9",
      "amount": 3274054771,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1qzy40tuamwfdf320le6cnfkntfrn63j72quc43",
      "amount": 1773009599,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1874lfx44c2gu47k6hh2nx8jpyh7n570pdcdcuv",
      "amount": 2796194945,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1xekcjker8npq28qcp5nxr8y0pddlap0x4jgaze",
      "amount": 1592320722,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1dn4tsvhuqpcksj4tdnrnkvmhylnzydcvyclf0y",
      "amount": 1581027667,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1ecgmgzttrq52g6gp4d4f687g2xymnjxvdfqsqc",
      "amount": 9719130344,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi15acyslq77wylle9kc3cueme5ml42aacsyz79z2",
      "amount": 2997306289,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1c3dxnlqwfvjfgcyzqhdhtp7h348e78erj6a8lt",
      "amount": 1212121212,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi179xh34gna498hdtglq7mv2hwrxqwzw8tfg43ec",
      "amount": 1110483719,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1j5dtezz2e7mpvqj42x2c5p0qy7glsac7mqfrcl",
      "amount": 1076604554,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi12evsxkfkq59v9s88zfx0cl9tn4zlavs9zuueuj",
      "amount": 1057782796,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1kuhp2ywypwg2ekuz05zw4dlvtm40zswfawf32v",
      "amount": 1054018445,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1j027zxk5wqp8ek2x490mf2awkyfcnztxc3e6zx",
      "amount": 752870318,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1nhgprhqn3sl84wqvh2j6raezkach9gyrgvanf6",
      "amount": 3757306262,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1pfkuyknmfhp4uhhd7tqr4yeqvs9qdttvfgfsfm",
      "amount": 26162280745,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1l832mnjzn24hfqkvjr6se2r00at9p6el9w6rvq",
      "amount": 2347518820,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1ydtts8a9r5jr0qmls9cy60p2j9ewvg6mrnkh4z",
      "amount": 39492343498,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi19xek9465etmvshl5vhet2s7c8uc2lrjvte65ec",
      "amount": 1646950376,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1n7d0pd3jj0r4wq764vth26radlnvcljele0cd4",
      "amount": 737440467,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi18hpy9dlnf9wcmwu9mu778rwezuxnfmjjj3pvwl",
      "amount": 1542479643,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1zvsl6u04wgxd8j9qek6nucycynu856pxcftq0f",
      "amount": 1272084805,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1jz76pjr0zt997c07efpuq9n0g7zpjz64udnrgg",
      "amount": 2393506942,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi132srs0hmcwcvl5xee7xmmsnsfu8gjf32kkt7rm",
      "amount": 1966592211,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1p6naal89l93mlvwnjp4gzjuhm9jrlp06ag26lw",
      "amount": 4000694565,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1aj9y7y4ngaqp7w9t56mr8cw5lrtsu9q0f0f5ce",
      "amount": 2569890606,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1vchphhj9ykdt8sre4a2xwyzmmg4j2xd6zrgydd",
      "amount": 1896162528,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi163j85dudtqpt40wwuzhn2r3kd5n3hufuacw3pu",
      "amount": 1764195172,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1s8tgzyc3tws66a9lpu6hlj0s4th9r3nlx2qkk6",
      "amount": 1666956068,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1awlrha5r7e5jzcceqgsle0vva2d8qrtmskg83t",
      "amount": 1444695259,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1pzjzaftghc7m8k67qxs0p3nl7kkdvlghg28k8r",
      "amount": 1284945303,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1453sp8gyfje8y8kuux9rhald6yzcjk03ylkmxj",
      "amount": 1048793193,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1395u5yh5sx67hx3krhxa2w9gut479m2y2tz76d",
      "amount": 1021010592,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1s82uxe4h3077temj7plxe0ydmhc84g70xkp2fe",
      "amount": 1000000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1az7k4r6qcyw54e5ltvfdva0lz9k7kqhy6a9num",
      "amount": 1000000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1ces9q4hsgn6x7j040z7a8jg5tnudrwtsps2vfa",
      "amount": 1000000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi16g2reacq384f8uqfkehdnt5q6wu6lgpc0xgk8v",
      "amount": 1000000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi15ur7aylv5tu5q4y5zkzmrwdwuxx0rajctgmmkz",
      "amount": 4000000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi14v2jryt8828p0mjqy4jqx29s9wj7n7gus87sdx",
      "amount": 1000000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1qrpd96dapnd87ff92k6hxgl7zvjqlykc8pffh5",
      "amount": 1000000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1rvge28l7le97e6qtzlnzrj9vp7v77pcujq0938",
      "amount": 1000000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1epc2vczt28nud03jk049wpjcrtn4gregr779h4",
      "amount": 1000000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi183hfvu0hnju50s32kqcf8nk6f5d70p54mc82w5",
      "amount": 10000000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1c5w9gv4x0cp3wnxnhk4vr08h4552u0fqe525d0",
      "amount": 1000000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi10eln7ccy0vth6r6hr7xgxfyg6yhfc3g9ck55vc",
      "amount": 1000000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1wuxggfh5verzhsyfjavtn5c4grd8k3jjlqsmxy",
      "amount": 1000000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1gwuu4xe3c759atkwk6uhryj6jwgyth5tde8m5d",
      "amount": 1000000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1cpu56kjrnhj8x04rxmdst4ztu77wz8ncflajd5",
      "amount": 4000000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi129ax5ctrjd0ddstm8pge52kwkvjvtvqjpmfjp2",
      "amount": 1000000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1uc68tmfc40yurpw8xqk30sxggvd5javt9xvny5",
      "amount": 4000000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1dsqsrgt5xkt6pj3nuht375jrt0xfl8p8ev336s",
      "amount": 4000000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1v45254hgxun3vnydjcl6dhw6sz82ulftu6t53s",
      "amount": 25897028517,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1pnndh4ge5svw86myh3ng9cd0s729y4x3eh5gn8",
      "amount": 5000000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1pq9e4epludy78a96y0j89uqt0uq5k4z3205lka",
      "amount": 52756766713,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1a5yqx4xa5v7ggs5xc5mnk4qzgrylqxd9xy46zy",
      "amount": 1335497760,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi17c7l66wgut2trha0qv996tqtg4scwusp4wm649",
      "amount": 25052095925,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1avd4jplu89t6ds66udfspsv9nlkwn2ru03jy9j",
      "amount": 15289146778,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1paj00usj7k69f9vvpknazcppr48pwnyqxg4lqj",
      "amount": 5428568139,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    }
  ],
  "marketing": [
    {
      "toAddress": "ununifi1llr0q0ggvnyu7fxh57wxkuvmlpyf6ygmknga26",
      "amount": 241720000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1uny0f22uj5mn93sgg5d80nhnl7608fd9ap4ld6",
      "amount": 241720000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi1yf5cv54mclv84y2gufafh9cm2yzdne64zasfnk",
      "amount": 181720000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
    {
      "toAddress": "ununifi172rcust8rcuf5ef2j4ywrfccf9lre70dmv8k8v",
      "amount": 241720000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    }
  ],
  "advisors": [
    {
      "toAddress": "ununifi1g94dh4vzqflc4xrjerza9t7ydc5gfpc4zhtal2",
      "amount": 100000000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1687186800
    },
    {
      "toAddress": "ununifi1dn7ta3u4zxlug5npfcy3ue6stjlfw7kscynwtp",
      "amount": 2000000000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1687186800
    }
  ],
	"others": [
		{
			"fromAddress": "ununifi1pa29ejcfrylh69pvntrx3va9xej69tnx7re567",
			"bankSendTarget": {
				"toAddress": "ununifi1q6jfv5un5cc7lh26njttg0tje0jevt93shy9zv",
				"amount": 2205862352941,
				"denom": "uguu"
			}
		},
		{
			"fromAddress": "ununifi19srj7ga7t2pyflz7f50le5fv0wa9kuf7tmdtla",
			"bankSendTarget": {
				"toAddress": "ununifi1q6jfv5un5cc7lh26njttg0tje0jevt93shy9zv",
				"amount": 2050518544482,
				"denom": "uguu"
			}
		}
	]
}`
