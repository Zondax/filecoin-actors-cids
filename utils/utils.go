package utils

import (
	"fmt"
	lotusBuildLatest "github.com/filecoin-project/lotus/build"
	builtinV8 "github.com/filecoin-project/specs-actors/v8/actors/builtin"
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
	ActorsV8 uint = 8
	ActorsV9 uint = 9
)

func GetFullMetadata(version uint) []ActorsMetadata {
	switch version {
	case ActorsV9:
		return getLatestMetadata()
	case ActorsV8:
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
	const previousVersion = 8
	var meta []ActorsMetadata
	var networks = []string{
		NetworkCalibration,
		NetworkButterfly,
		NetworkCaterpillar,
		NetworkDevnet,
		NetworkMainnet,
	}

	cids := map[string]cid.Cid{
		"account":          builtinV8.AccountActorCodeID,
		"cron":             builtinV8.CronActorCodeID,
		"init":             builtinV8.InitActorCodeID,
		"storagemarket":    builtinV8.StorageMarketActorCodeID,
		"storageminer":     builtinV8.StorageMinerActorCodeID,
		"multisig":         builtinV8.MultisigActorCodeID,
		"paymentchannel":   builtinV8.PaymentChannelActorCodeID,
		"storagepower":     builtinV8.StoragePowerActorCodeID,
		"reward":           builtinV8.RewardActorCodeID,
		"system":           builtinV8.SystemActorCodeID,
		"verifiedregistry": builtinV8.VerifiedRegistryActorCodeID,
	}

	for _, network := range networks {
		meta = append(meta, ActorsMetadata{
			Network:          network,
			Version:          previousVersion,
			ActorsNameCidMap: cids,
		})
	}

	return meta
}
