package utils

import (
	"fmt"
	lotusBuildLatest "github.com/filecoin-project/lotus/build"
	builtinV7 "github.com/filecoin-project/specs-actors/v7/actors/builtin"
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
	ActorsV7 uint = 7
	ActorsV8 uint = 8
)

func GetFullMetadata(version uint) []ActorsMetadata {
	switch version {
	case ActorsV8:
		return getLatestMetadata()
	case ActorsV7:
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
	const previousVersion = 7
	var meta []ActorsMetadata
	var networks = []string{
		NetworkCalibration,
		NetworkButterfly,
		NetworkCaterpillar,
		NetworkDevnet,
		NetworkMainnet,
	}

	cids := map[string]cid.Cid{
		"account":          builtinV7.AccountActorCodeID,
		"cron":             builtinV7.CronActorCodeID,
		"init":             builtinV7.InitActorCodeID,
		"storagemarket":    builtinV7.StorageMarketActorCodeID,
		"storageminer":     builtinV7.StorageMinerActorCodeID,
		"multisig":         builtinV7.MultisigActorCodeID,
		"paymentchannel":   builtinV7.PaymentChannelActorCodeID,
		"storagepower":     builtinV7.StoragePowerActorCodeID,
		"reward":           builtinV7.RewardActorCodeID,
		"system":           builtinV7.SystemActorCodeID,
		"verifiedregistry": builtinV7.VerifiedRegistryActorCodeID,
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
