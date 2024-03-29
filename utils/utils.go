package utils

import (
	"fmt"
	lotusBuildLatest "github.com/filecoin-project/lotus/build"
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
	builtinActorsMetadata := lotusBuildLatest.EmbeddedBuiltinActorsMetadata
	builtinActorsMetadata = append(builtinActorsMetadata, WallabyBuiltinActorsMetadata...)
	for _, d := range builtinActorsMetadata {
		if uint(d.Version) != LatestVersion {
			continue
		}

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
	builtinActorsMetadata := lotusBuildLatest.EmbeddedBuiltinActorsMetadata
	builtinActorsMetadata = append(builtinActorsMetadata, WallabyBuiltinActorsMetadata...)
	for _, d := range builtinActorsMetadata {
		if uint(d.Version) != PreviousVersion {
			continue
		}

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
