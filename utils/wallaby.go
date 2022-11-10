package utils

import (
	"github.com/filecoin-project/lotus/build"
	"github.com/ipfs/go-cid"
)

var WallabyBuiltinActorsMetadata = []*build.BuiltinActorsMetadata{
	{
		Network:     "wallaby",
		Version:     8,
		ManifestCid: build.MustParseCid("bafy2bzaceb5g64cj6swfgmmf4hreo6lcvu4bup5ollfrqsq5qpazzkd564pw4"),
		Actors: map[string]cid.Cid{
			"account":          build.MustParseCid("bafk2bzacecy3hive4t2j2ytbq3jrhueu3q3yvtmdny4lq6gwpaaovb5nrr76s"),
			"cron":             build.MustParseCid("bafk2bzaceaqq2p7g5g2gbrwviwi5rerq7gclsryzmr3cgwdym3q6cuwr2m6jy"),
			"eam":              build.MustParseCid("bafk2bzacedwuvyzfaaf6vpxx4lhervvs4qs4ukfqitjxikeemzpec3lbqu5ba"),
			"embryo":           build.MustParseCid("bafk2bzacecau3tohdilfx66pohfqdrngpuqd5oew2j5iv3c7sjlrkcm5npqos"),
			"evm":              build.MustParseCid("bafk2bzacedzg2dsdry6cy5nzfldtqatuopljgdxt5hxdwn2gmuj3fk566bndg"),
			"init":             build.MustParseCid("bafk2bzacebfgbr3jpistckeoomgmub65hplg7mhwnvznqprxxg7ujajkvixoi"),
			"multisig":         build.MustParseCid("bafk2bzacebdwl2ffznau33xzjztvgyiptgbsz6imn2g2pakpvwuguacwigxtg"),
			"paymentchannel":   build.MustParseCid("bafk2bzacedu3vnale73bu5k2dcaiz423orrtyvx4bazte773kyabiwgdvw6ye"),
			"reward":           build.MustParseCid("bafk2bzacec2yqixjcmwnhplqnpfflcn3md3illzd7n4sxs4axbvoux6lkim3w"),
			"storagemarket":    build.MustParseCid("bafk2bzaceddnsy6esfaxzpcczcwo6vjgwd6zqhy7n67mmok5o2zfj3fyaw6ow"),
			"storageminer":     build.MustParseCid("bafk2bzacebc5dnbl2q3u6aevyxxk6x7aoomkg3hbp73i7u5bo4ltc2kjpsuzo"),
			"storagepower":     build.MustParseCid("bafk2bzacebdq4txea7jncypuxoomanpsiz6q2zmz7ejga3rkg4reryjozkcbk"),
			"system":           build.MustParseCid("bafk2bzacebevepzhwqe25epwpkjbekm4std6l4b3umlvdxrzd2x2dyaq6p224"),
			"verifiedregistry": build.MustParseCid("bafk2bzacecfa4bxq2tdkscq4iu2ys5skq6by7x534p3oxziooc6oymsg64sdo"),
		},
	}}
