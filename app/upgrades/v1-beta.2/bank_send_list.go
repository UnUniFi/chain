package v1_beta2

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
      "toAddress": "ununifi1az7xg9wyfruf94teg9u2dr8dec7sjf6esultlc",
      "amount": 2794117647058,
      "denom": "uguu",
      "vesting_starts": 1661126400,
      "vesting_ends": 1724284800
    },
    {
      "toAddress": "ununifi199j70q49338yz9r83xsmfescn74gxachwfp9cs",
      "amount": 2794117647058,
      "denom": "uguu",
      "vesting_starts": 1661126400,
      "vesting_ends": 1724284800
    },
    {
      "toAddress": "ununifi1w92q2r53jgvyvcuuwmlggrvum7c6j72y403puk",
      "amount": 20000000,
      "denom": "uguu",
      "vesting_starts": 1661126400,
      "vesting_ends": 1724284800
    },
    {
      "toAddress": "ununifi19f0w9svr905fhefusyx4z8sf83j6et0gkfnrpz",
      "amount": 20000000,
      "denom": "uguu",
      "vesting_starts": 1661126400,
      "vesting_ends": 1724284800
    }
	],
  "airdropCommunityRewardModerator": [
    {
      "toAddress": "ununifi1vntg7ecgu7f3vydjmzlrtnav5pqlr7wezcw38h",
      "amount": 1199314676,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1rdg72fz6qzztufeg85pqkcpz44vgfvj6y970nz",
      "amount": 856653340,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi12p0q22mjxz7ywpjjpkctvy7yv8h852nnvtfz76",
      "amount": 171330668,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1u5lwnej2gwxtjm2ys7fa6jnlwz8q8e0a9afyap",
      "amount": 1713306681,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1lyx90ps90t6dl92jmm30lzwlw8klmeul6v07gq",
      "amount": 513992004,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1mahxnymyfxszxk7c54veha9dwwvkt42m5h6uu9",
      "amount": 856653340,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1u52lf749s4sz8k6y4tgcx4m62cyqsxdrywuzt2",
      "amount": 1713306681,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1rkglywylt8hdnm3undmv0at2prchy0s5m67hal",
      "amount": 1713306681,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1swdlz05ms6nm4xj6xwlq4m7r6320ncz02r40ud",
      "amount": 1713306681,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1rqeu8c0ljagwulc9nf0acccv2y6ds4mhsptmes",
      "amount": 1713306681,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1pesrpts0a4rpeg4pmfsee5sr579urk4xawp9gp",
      "amount": 856653340,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi13xvqw96lg3f8ek26g6p0hs8x5lwhyz66w2uzry",
      "amount": 856653340,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1y4j0qsj76tgzx4se7657aannrwwtk9tq8glnzl",
      "amount": 513992004,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1clay68h9xf6x49umn0turw294hu5vls64xscmx",
      "amount": 513992004,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1xsrup2xwl3h02vc7gu5spgawm08zhwt0upnggj",
      "amount": 1713306681,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1f39ftt6nl3txq4yl5vnt4vc5ky9slqk7wf7q4m",
      "amount": 1713306681,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1km4tdzat0yw4a8r9yh3q6hggferwy93sp9mzdv",
      "amount": 513992004,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1prvfte89hd7w08kl6n5rgjknuyandh49zg6g4z",
      "amount": 513992004,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1lj0ptujepjfxdketv4h5edyrtd68e5c9m5lv8u",
      "amount": 1199314676,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1jj5erf7kszg5u74q9h8cfqzaxkynf3cg2vhnsd",
      "amount": 513992004,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1t9qvrz8jd20jjx5kz0ruuhnk9c3ctkh8cxp2hj",
      "amount": 513992004,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1zllhxputanq6nr2grgv42xv93d6ykusttx2gpe",
      "amount": 513992004,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi19846ajj4g7cldprrjjtcpkde4c24fhqymlzg67",
      "amount": 1713306681,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1sg5v5ty92slcl39yezzssz4sg8ves3ssy6ywpc",
      "amount": 513992004,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1wtlv63qa4zw4d4pyqjnp6tagchzrkprtyeyq69",
      "amount": 513992004,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1ce8rkj9hhptk7fjzkzq8nzeqgcd0cxcgft3pxn",
      "amount": 513992004,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1qzhyw5gn4yw896pem0gn2nm8hwglm8ftsragxy",
      "amount": 856653340,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1huqml6scl6qzvetkc3jce6knmwhtk4x0av5wgg",
      "amount": 513992004,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1533zjekeeqsgk4l6je9nudaycpup3c3egpcvq8",
      "amount": 1713306681,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1xe0ckssfggcxwv65nr0r59fls5gpe6fthey2x4",
      "amount": 1199314676,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1l24zlfegct4kzxvlddmv3v9zykxazjeksxwpjq",
      "amount": 513992004,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi13ufs9fz4d9svh7ay5eqdcqhnrq369nu0pdr8rp",
      "amount": 171330668,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1kddhvwmw85tpugwldg08usqsvphdsml60rhryc",
      "amount": 856653340,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1r57prxdz5afzr0dy6pdpf9zx4wr5xy3ca9evdh",
      "amount": 1199314676,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi15ff2v6tgltl4sgqqgt2cuxnpe9k996km86hjjc",
      "amount": 171330668,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi12j6hcqnlannlfx4vcpjcezkaay55un3z2mcv5a",
      "amount": 1713306681,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi19s8p8hjyjwh7tl2yrp80fm2tnnn5rf2r0kranf",
      "amount": 856653340,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi156qhxjg924wfpahhgzmat8ssm89k5ndtyq3k7v",
      "amount": 171330668,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1jv08s99dhc0n8t6ya4yqtzgrwerpan3k572c4h",
      "amount": 856653340,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi12v8a7tj50lzds3zmjwe388ss7sz9etjcxewp5f",
      "amount": 513992004,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1nuucmhpay3rkjjz0qwp73c4qc9vjlcmefhw92q",
      "amount": 1713306681,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1tcrzpx3cg94pa87a9nk6m2984960lkatg9lfrt",
      "amount": 1713306681,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1l9qn05tjgrk8rptgzz4nztq72sm6hmxxl0valj",
      "amount": 1713306681,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1fd7rpspd7pwcl57gedzau547jcewcxxwlmqc05",
      "amount": 171330668,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1lcspaj06pwdz9qskt0gxxh5y9l6altm0j4s9lc",
      "amount": 513992004,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1c65c9nlggd2tmn3pxwz4ql5q7dgfjqxyc4pp5h",
      "amount": 171330668,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1t0yyu79gyf0ynuagrxmrgu6evgeckdhtn5eeth",
      "amount": 1713306681,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1h0vrvjkmvzvja32u3ljwlxerv9t75qucr06vpd",
      "amount": 856653340,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1n258cafkx6phkau6uqhyjwpdtsez93nr8cy3r7",
      "amount": 1713306681,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi17yn38x8hjkn9sed9vxkn7858f2t52uk4g8tnyk",
      "amount": 856653340,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1f02ck2lrqmn0cfsphu3wx8x2gkltdmcqt6ctm0",
      "amount": 513992004,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1q84vhdwwhkxtn4pud7js8jfz9l505dyntqusvj",
      "amount": 1713306681,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi15ld4z0vn4psrrwalq035k4mywgps4wc3ht92jx",
      "amount": 513992004,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi15w6g5dugd30csgev2gfnvat4mdtgkveznnujhk",
      "amount": 171330668,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi179pwy6eruf588436g98xseaefcntsgxelh4u8m",
      "amount": 1713306681,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1rs4qaq0crj7v0m9hmznva06drpljht0ehn683v",
      "amount": 856653340,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1wt3tpc7en3rvl3f608auzdlx7w4rhvfz6v4024",
      "amount": 1713306681,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1yxt9vq4svlk8j9dysmlhdx3u8lc08h2wfynkxg",
      "amount": 856653340,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1n5jux49rxn2h05znv92rvqecj8ljmhv2nztuw8",
      "amount": 171330668,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi16y0jgcslxj9f7rxq4v4uy6900afwq28d8r2vlq",
      "amount": 171330668,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi13vypjk7wy3xucmtgaar2y4j404t027a6l7au4a",
      "amount": 171330668,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1hz07sq3deezmp5j32z2xs7sfzccm9sx6820krl",
      "amount": 513992004,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1g87rxhxtmeeqgch235puq73cmfxxk37gl8yxem",
      "amount": 513992004,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi18cameakt4chmpvvh4km3ry6l4dmgpjzvxg96wl",
      "amount": 1713306681,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi145s06fr32tyyd0594lkrulp8n36uau4t5ypgvd",
      "amount": 1713306681,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi17r5as3h2zp3079vcx9ejg36myczl5pl0nywql7",
      "amount": 856653340,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1lt837jpxja5ldnrax8wvadqnl5wvleesmg73zy",
      "amount": 171330668,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1k4mns7svuhrn0vuy23d866xqz882mkpv235czg",
      "amount": 171330668,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1tr5pfv38mtwww2qcuvf7k3cm2yzh6ey08z2jjr",
      "amount": 171330668,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1g74tg3el62spyqfewr9szalxzdwre2wxl36um2",
      "amount": 171330668,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1uue96qs9tdye9mjw2hvta97xkjmy7dmwgwplxx",
      "amount": 1713306681,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1cavga4cpvq3jc4l328877ta8hjjxl4g4yf5820",
      "amount": 513992004,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1juf0xjh8tmjlvt3xa4jcv009vq3nsc4wm8tfay",
      "amount": 171330668,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1u5rq7um22yu38hggfup9j4kv70sngfaelkck0w",
      "amount": 513992004,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1qmswnrm87xy8v9wr644kng3qhjjqpt0sndu579",
      "amount": 513992004,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1sxa5xjlyz9y72smjruc25qeh05q7y2jgwfaun0",
      "amount": 856653340,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1f090d34fnczvlugzt5txeqw29cq9hym5t9uy9r",
      "amount": 1199314676,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1cj5t0r7at6f55nl98t5hz9657z8ca00w2lyd9d",
      "amount": 171330668,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1ale9tx38hm45keu4swvnkc34yc9ln3jz8p6488",
      "amount": 1713306681,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi159ckcjzm8qm4ad64yaenzyhwp3ttpz7q8zgvk5",
      "amount": 1713306681,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1zhs5w4cgwyf5j78my9kp5g9ppxs7d5xlphcvd6",
      "amount": 171330668,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi12u9pxq2snxhft228u0xcs883l5sl0acjdwywqv",
      "amount": 1713306681,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi17naryjvzweytz73ntrqclwwjwtqp5hasty2wum",
      "amount": 1713306681,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi18h77jhystdgptu867actg2wzwkm9xug9ukv5kc",
      "amount": 856653340,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1wpy4gj47ggr8tgzjreqgankylk59z2hzreatjz",
      "amount": 1713306681,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1kgdegkl8s2pf0gr4f409acccrvpujfsf4w4ppr",
      "amount": 171330668,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1qkdppsqhmfy99fu68ud8cqyxtn9j583gg0uamj",
      "amount": 856653340,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1epc2vczt28nud03jk049wpjcrtn4gregr779h4",
      "amount": 1713306681,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1f2pln0kl2vh4rzhwquuky2knkmh8ltwxhuazef",
      "amount": 856653340,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1dtautusrks550sfj7f5mprknq3a50u60tx637z",
      "amount": 1713306681,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1sfyyj50djrd0kju4kvemyrf8pg7l5zt779qpx3",
      "amount": 513992004,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1653py0583z9xudlpxgwuh97p3gjwmfnefwqvyd",
      "amount": 513992004,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1wj77f4mr9jsk5rdezjady7l2fntcgmjtf002gq",
      "amount": 513992004,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi15afu60dw33y65shzg6836wdjckxezx7rerh5nu",
      "amount": 513992004,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1qv8q6jf7sju85s06v6lnuhpyfw7t56w9squs8e",
      "amount": 513992004,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1r9z7kcelx4ylypxa5ey9a0la3vj0wjg4k60agx",
      "amount": 1713306681,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1qnkajfgqhkmxquetadxqn5jyp59wpgsfzfft3p",
      "amount": 513992004,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1fluhhthnqhyzprg849n0qthyc7mzmglhfxvulg",
      "amount": 513992004,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1j45emtyqdqxp6f2umngy3nd7d80xknsxw4g62h",
      "amount": 1713306681,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1y9sf62wxx7gkuw0jlwn909l7q465luwfzfxqdy",
      "amount": 1713306681,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1vm2tnldcenja0ctlqpcqusdpp5jv8k5n7cf4wz",
      "amount": 1713306681,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1sgtjkhzal4mgmd8hmxwxzkzdczwnrfsfack409",
      "amount": 171330668,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1kz643fqvtdc8h7gs4x6k3c9732er0v04rtmap9",
      "amount": 1713306681,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1eljt8f0z49xvmhtuderhd3dvrpyh5x2yd42wm4",
      "amount": 2227298685,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1fmczw6xvue5l658nzfyv8gl0xwxeq2guyg4cf0",
      "amount": 171330668,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1rzd3awqzl6h67u8ltdjnzfm7rqv5jse9zshhtr",
      "amount": 8755422688,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1nwahmkt5lc8ajr35qjfuc9p9gsxr867yt6s9zw",
      "amount": 171330668,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1zqx2lm3vmw9dcps333kce9wxeujqdhtlwp26w7",
      "amount": 171330668,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi19c7m6yksgs9dafur6xw788jxeu54vjzul65usf",
      "amount": 171330668,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi123p0wjzvu4kampc3hn3x9v5cn75zzudaquvld3",
      "amount": 171330668,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1xurwjuy8dvucl6dnrzld0nvw6tpqp30rk0rafs",
      "amount": 171330668,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1wqpdt48849gwr6qfslnu0vvrp7k40662w8nt3f",
      "amount": 1370645344,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1cgjhlgdnzlxp94j9msdl6gxf6qd8eh7wmmdqcn",
      "amount": 1370645344,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1vgjxz84exe55pmvmhrjgd50kuf8nteam9ql3lv",
      "amount": 1370645344,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1cxthcxhsks5hfdd5zfmz2t6gkjapvunte4en4q",
      "amount": 1370645344,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi16m5j7e232uuwgf0qa5ujc9g8rfels40e7jkhxr",
      "amount": 1370645344,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1zjee39sw23unvvnw0vxl3fmz33k3s3mfp4ea92",
      "amount": 1370645344,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1vkecuxvetuv86tgnjfkalmtcgy29l37hl74yuk",
      "amount": 1370645344,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1atmdkat7mqh35u5cpr29njs0qfv33ewn2hdrrv",
      "amount": 1370645344,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi19570z7f7r2x69056hc5ez3tk9gtau3vyg2zlt7",
      "amount": 1370645344,
      "denom": "uguu",
      "vesting_starts": 1664582400,
      "vesting_ends": 1696032000
    },
    {
      "toAddress": "ununifi1s37k9ryxegcfp857r9pm79azc4x72l7vzjjg4j",
      "amount": 37895834536,
      "denom": "uguu",
      "vesting_starts": 1656633600,
      "vesting_ends": 1688169600
    },
    {
      "toAddress": "ununifi1pfkuyknmfhp4uhhd7tqr4yeqvs9qdttvfgfsfm",
      "amount": 16699688467,
      "denom": "uguu",
      "vesting_starts": 1659312000,
      "vesting_ends": 1690848000
    },
    {
      "toAddress": "ununifi1ydtts8a9r5jr0qmls9cy60p2j9ewvg6mrnkh4z",
      "amount": 20763358505,
      "denom": "uguu",
      "vesting_starts": 1659312000,
      "vesting_ends": 1690848000
    },
    {
      "toAddress": "ununifi1qxfymvcz8h2ahnapezz2ed4awh29qswygt2prx",
      "amount": 7706509689,
      "denom": "uguu",
      "vesting_starts": 1659312000,
      "vesting_ends": 1690848000
    },
    {
      "toAddress": "ununifi17c7l66wgut2trha0qv996tqtg4scwusp4wm649",
      "amount": 1900237529,
      "denom": "uguu",
      "vesting_starts": 1656633600,
      "vesting_ends": 1688169600
    },
    {
      "toAddress": "ununifi1p2rkc23ayawfzl8g4kneaymhp3nte7x3t39jx5",
      "amount": 1900237529,
      "denom": "uguu",
      "vesting_starts": 1656633600,
      "vesting_ends": 1688169600
    },
    {
      "toAddress": "ununifi1ddpzczgwt55wtrdjyxfkvymltv7wzjhvv8p34h",
      "amount": 4750593823,
      "denom": "uguu",
      "vesting_starts": 1656633600,
      "vesting_ends": 1688169600
    },
    {
      "toAddress": "ununifi1t2cky7r2rmrp50n2svk59752u4z3349hnj0t58",
      "amount": 1425178147,
      "denom": "uguu",
      "vesting_starts": 1656633600,
      "vesting_ends": 1688169600
    },
    {
      "toAddress": "ununifi1vxg66xlzymyc5fd9hwk602xa0tg09dms8dg509",
      "amount": 10747826085,
      "denom": "uguu",
      "vesting_starts": 1659312000,
      "vesting_ends": 1690848000
    },
    {
      "toAddress": "ununifi1t3syzcj74az3kzta8u5jpkxn6ft7qdumwdyqc0",
      "amount": 2040579710,
      "denom": "uguu",
      "vesting_starts": 1659312000,
      "vesting_ends": 1690848000
    },
    {
      "toAddress": "ununifi18g8870c64sr3t4v6802rj2h4m2xwj6udeuuj6h",
      "amount": 1681159420,
      "denom": "uguu",
      "vesting_starts": 1659312000,
      "vesting_ends": 1690848000
    },
    {
      "toAddress": "ununifi1k8f9hre2szp7p7mjwp70kaydz5hl2lunuzvp78",
      "amount": 1669565217,
      "denom": "uguu",
      "vesting_starts": 1659312000,
      "vesting_ends": 1690848000
    },
    {
      "toAddress": "ununifi163j85dudtqpt40wwuzhn2r3kd5n3hufuacw3pu",
      "amount": 1553623188,
      "denom": "uguu",
      "vesting_starts": 1659312000,
      "vesting_ends": 1690848000
    },
    {
      "toAddress": "ununifi180zz0e8q4c3zyd82e8tz7nu4k4zj0c73ep420z",
      "amount": 1530434782,
      "denom": "uguu",
      "vesting_starts": 1659312000,
      "vesting_ends": 1690848000
    },
    {
      "toAddress": "ununifi104xee9d8gquaz7vyzwzurxjxk0sutqqf3emht6",
      "amount": 5591588761,
      "denom": "uguu",
      "vesting_starts": 1659312000,
      "vesting_ends": 1690848000
    },
    {
      "toAddress": "ununifi1u9nlag5wdtx9wlnqu7h7gnadnjuqqa7tsmk0tn",
      "amount": 1414492753,
      "denom": "uguu",
      "vesting_starts": 1659312000,
      "vesting_ends": 1690848000
    },
    {
      "toAddress": "ununifi17tzgev3du8kttl98xpehd9vf2tvv55asd0kklw",
      "amount": 1391304347,
      "denom": "uguu",
      "vesting_starts": 1659312000,
      "vesting_ends": 1690848000
    },
    {
      "toAddress": "ununifi1nhgprhqn3sl84wqvh2j6raezkach9gyrgvanf6",
      "amount": 1391304347,
      "denom": "uguu",
      "vesting_starts": 1659312000,
      "vesting_ends": 1690848000
    },
    {
      "toAddress": "ununifi1nl4v9kcumsuc6uejqhmqrqsa9a2ju0jmf7gwur",
      "amount": 1310144927,
      "denom": "uguu",
      "vesting_starts": 1659312000,
      "vesting_ends": 1690848000
    },
    {
      "toAddress": "ununifi1md5pmzpq5rzr5hj82tnt3ansgef95x2qepah9a",
      "amount": 1298550724,
      "denom": "uguu",
      "vesting_starts": 1659312000,
      "vesting_ends": 1690848000
    },
    {
      "toAddress": "ununifi1rkh95rzym7cfcthwlc3ee00jcw6h07vey67p7c",
      "amount": 1286956521,
      "denom": "uguu",
      "vesting_starts": 1659312000,
      "vesting_ends": 1690848000
    },
    {
      "toAddress": "ununifi16e6agl7qsans0qjdzuqcq3ukzz8g2ztdnw8xgn",
      "amount": 1286956521,
      "denom": "uguu",
      "vesting_starts": 1659312000,
      "vesting_ends": 1690848000
    },
    {
      "toAddress": "ununifi1xekcjker8npq28qcp5nxr8y0pddlap0x4jgaze",
      "amount": 1286956521,
      "denom": "uguu",
      "vesting_starts": 1659312000,
      "vesting_ends": 1690848000
    },
    {
      "toAddress": "ununifi132srs0hmcwcvl5xee7xmmsnsfu8gjf32kkt7rm",
      "amount": 1275362318,
      "denom": "uguu",
      "vesting_starts": 1659312000,
      "vesting_ends": 1690848000
    },
    {
      "toAddress": "ununifi104y2xw9ys02qlzlgntpefzwtqxk8eksrr0mkkk",
      "amount": 1263768115,
      "denom": "uguu",
      "vesting_starts": 1659312000,
      "vesting_ends": 1690848000
    },
    {
      "toAddress": "ununifi1gqmhnxd5w9zkhhp4llvul0v4ck9cn8w7np4cp3",
      "amount": 1228985507,
      "denom": "uguu",
      "vesting_starts": 1659312000,
      "vesting_ends": 1690848000
    },
    {
      "toAddress": "ununifi1avd4jplu89t6ds66udfspsv9nlkwn2ru03jy9j",
      "amount": 1205797101,
      "denom": "uguu",
      "vesting_starts": 1659312000,
      "vesting_ends": 1690848000
    },
    {
      "toAddress": "ununifi1v45254hgxun3vnydjcl6dhw6sz82ulftu6t53s",
      "amount": 19967751002,
      "denom": "uguu",
      "vesting_starts": 1659312000,
      "vesting_ends": 1690848000
    },
    {
      "toAddress": "ununifi1paj00usj7k69f9vvpknazcppr48pwnyqxg4lqj",
      "amount": 2040858592,
      "denom": "uguu",
      "vesting_starts": 1659312000,
      "vesting_ends": 1690848000
    },
    {
      "toAddress": "ununifi1eq9star05sxl9heurjtcpdmhrcscdur9g85ttq",
      "amount": 4295921537,
      "denom": "uguu",
      "vesting_starts": 1659312000,
      "vesting_ends": 1690848000
    },
    {
      "toAddress": "ununifi1skex8qg79h667zdvnsgt09pum9f7yyr5zv0dar",
      "amount": 23154188671,
      "denom": "uguu",
      "vesting_starts": 1659312000,
      "vesting_ends": 1690848000
    },
    {
      "toAddress": "ununifi1qhxnlhh53y2t73urgdqqxn7ttkseegwz62aq4j",
      "amount": 2401218952,
      "denom": "uguu",
      "vesting_starts": 1659312000,
      "vesting_ends": 1690848000
    }
  ],
	"airdropForfeit": [
			"ununifi1dc8nfka4dgw90rfejrfncsp5rmgtj5ppkfma5v",
			"ununifi1z9affl2d4yfahqn6nvm43tmxuxj084m36e7yjt",
			"ununifi16f32hrzas8ks5pwjdurkm47g8dqhclqn7323fc",
			"ununifi1qmsyp2f9f9tvzk9ncrp55a5p0gwheva8d4z2wv",
			"ununifi196fstf8eqdsenc89v7r30vsu6xf3uj5y6eja0s",
			"ununifi1nup4mfu7wtdychppcnsd76nmy7zuylws8k85le",
			"ununifi1yc7qneet922fxcvgjzm5g7ka26hqfrxx9fs5jk",
			"ununifi1flrwl30prh6p0ysq6jvumes24y4lvmzzlrpupe",
			"ununifi1khscuq3ku9gf9mpgvg897lacug9tn8mhjhldv4",
			"ununifi1cr3drk8decw8u25dqt3xrghu9nswhjjm0m26lc",
			"ununifi1gwpgu55ug72q30sn95fgvupzc7m2frjcxez7es",
			"ununifi1uz90c4u5jzwdcryhsaqskfrtn6mudt2mr93tdj",
			"ununifi19vu4j2wytjru5s779wr5dg0l8gskn5md08h0zy",
			"ununifi1fm2eyrtt3ut2d9g9ksdc3rx2eh77wxzl9y2ln8",
			"ununifi12h7qg749a43e5x7wdrafk3wldlayn84pgmghq0",
			"ununifi1cym27cxuc9zn74md3lq4jam2hs54xflvytl9p8",
			"ununifi1syyklr60ykhjwhwnkf7s344xmwdu2kzxlt34ts",
			"ununifi1j77ze43j79njy3sd62e3umpuglgwqn3ddezqmn"
	]
}`
