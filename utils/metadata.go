package utils

import (
	"github.com/ipfs/go-cid"
)

type ActorsMetadata struct {
	Network          string
	Version          uint
	ManifestCid      string
	ActorsNameCidMap map[string]cid.Cid
}

type ActorsMetadataMap map[uint]ActorsMetadata

func (a ActorsMetadataMap) GetActorCid(version uint, actorName string) cid.Cid {
	if meta, ok := a[version]; ok {
		if val, ok := meta.ActorsNameCidMap[actorName]; ok {
			return val
		}
	}
	return cid.Cid{}
}

func (a ActorsMetadataMap) GetActorName(version uint, cid cid.Cid) (bool, string) {
	if meta, ok := a[version]; ok {
		for actorName, actorCID := range meta.ActorsNameCidMap {
			if actorCID == cid {
				return true, actorName
			}
		}
	}

	return false, ""
}
