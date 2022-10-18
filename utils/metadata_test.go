package utils

import (
	"github.com/ipfs/go-cid"
	"testing"
)

func Test_GetActorCID(t *testing.T) {
	actorsMap := make(ActorsMetadataMap)

	ok, meta := GetMetadataForNetwork(ActorsV9, NetworkDevnet)
	if !ok {
		t.Fatal()
	}

	actorsMap[ActorsV9] = meta
	msigCIDV8 := actorsMap.GetActorCid(ActorsV9, "multisig")

	if msigCIDV8.Equals(cid.Cid{}) {
		t.Fatal()
	}

	ok, msigNameV8 := actorsMap.GetActorName(ActorsV9, msigCIDV8)
	if !ok {
		t.Fatal()
	}

	if msigNameV8 != "multisig" {
		t.Fatal()
	}
}
