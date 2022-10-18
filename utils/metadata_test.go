package utils

import (
	"github.com/ipfs/go-cid"
	"testing"
)

func Test_GetActorCID(t *testing.T) {
	actorsMap := make(ActorsMetadataMap)

	ok, meta := GetMetadataForNetwork(LatestVersion, NetworkDevnet)
	if !ok {
		t.Fatal()
	}

	actorsMap[LatestVersion] = meta
	msigCIDV8 := actorsMap.GetActorCid(LatestVersion, "multisig")

	if msigCIDV8.Equals(cid.Cid{}) {
		t.Fatal()
	}

	ok, msigNameV8 := actorsMap.GetActorName(LatestVersion, msigCIDV8)
	if !ok {
		t.Fatal()
	}

	if msigNameV8 != "multisig" {
		t.Fatal()
	}
}
