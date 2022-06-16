package utils

import (
	"github.com/ipfs/go-cid"
	"testing"
)

func Test_GetActorCID(t *testing.T) {
	actorsMap := make(ActorsMetadataMap)

	ok, meta := GetMetadataForNetwork(ActorsV8, NetworkDevnet)
	if !ok {
		t.Fatal()
	}

	actorsMap[ActorsV8] = meta
	msigCIDV8 := actorsMap.GetActorCid(ActorsV8, "multisig")

	if msigCIDV8.Equals(cid.Cid{}) {
		t.Fatal()
	}

	ok, msigNameV8 := actorsMap.GetActorName(ActorsV8, msigCIDV8)
	if !ok {
		t.Fatal()
	}

	if msigNameV8 != "multisig" {
		t.Fatal()
	}
}
