package utils

import (
	"fmt"
	lotusBuildLatest "github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/specs-actors/v8/actors/builtin"
	"github.com/ipfs/go-cid"
)

const (
	NetworkCalibration = "calibrationnet"
	NetworkButterfly   = "butterflynet"
	NetworkCaterpillar = "caterpillarnet"
	NetworkDevnet      = "devnet"
	NetworkMainnet     = "mainnet"
)

const (
	PreviousVersion uint = 8
	LatestVersion   uint = 9
)

func GetFullMetadata(version uint) []ActorsMetadata {
	switch version {
	case LatestVersion:
		return getLatestMetadata()
	case PreviousVersion:
		return getPreviousMetadata()
	default:
		fmt.Println("error actor version unsupported!")
		return nil
	}
}

func GetMetadataForNetwork(version uint, network string) (bool, ActorsMetadata) {
	meta := GetFullMetadata(version)
	for _, metadata := range meta {
		if metadata.Network == network {
			return true, metadata
		}
	}
	return false, ActorsMetadata{}
}

func getLatestMetadata() []ActorsMetadata {
	var meta []ActorsMetadata
	for _, d := range lotusBuildLatest.EmbeddedBuiltinActorsMetadata {
		cids := make(map[string]cid.Cid)
		for name, cid := range d.Actors {
			cids[name] = cid
		}

		meta = append(meta, ActorsMetadata{
			Network:          d.Network,
			Version:          uint(d.Version),
			ManifestCid:      d.ManifestCid.String(),
			ActorsNameCidMap: cids,
		})
	}

	return meta
}

func getPreviousMetadata() []ActorsMetadata {
	var meta []ActorsMetadata
	var networks = []string{
		NetworkCalibration,
		NetworkButterfly,
		NetworkCaterpillar,
		NetworkDevnet,
		NetworkMainnet,
	}

	cids := map[string]cid.Cid{
		"account":          builtin.AccountActorCodeID,
		"cron":             builtin.CronActorCodeID,
		"init":             builtin.InitActorCodeID,
		"storagemarket":    builtin.StorageMarketActorCodeID,
		"storageminer":     builtin.StorageMinerActorCodeID,
		"multisig":         builtin.MultisigActorCodeID,
		"paymentchannel":   builtin.PaymentChannelActorCodeID,
		"storagepower":     builtin.StoragePowerActorCodeID,
		"reward":           builtin.RewardActorCodeID,
		"system":           builtin.SystemActorCodeID,
		"verifiedregistry": builtin.VerifiedRegistryActorCodeID,
	}

	for _, network := range networks {
		meta = append(meta, ActorsMetadata{
			Network:          network,
			Version:          PreviousVersion,
			ActorsNameCidMap: cids,
		})
	}

	return meta
}
