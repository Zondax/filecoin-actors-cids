package utils

import (
	"testing"
)

var (
	versions = []uint{ActorsV7, ActorsV8}
	networks = []string{NetworkCalibration, NetworkButterfly,
		NetworkCaterpillar, NetworkDevnet, NetworkMainnet}
)

func Test_GetFullMetadata(t *testing.T) {
	for _, version := range versions {
		meta := GetFullMetadata(version)
		if len(meta) == 0 {
			t.Fatalf("metadata is empty for version %d", version)
		}
	}
}

func Test_GetMetadataForNetwork(t *testing.T) {
	for _, network := range networks {
		for _, version := range versions {
			ok, _ := GetMetadataForNetwork(version, network)
			if !ok {
				t.Fatalf("metadata unavailable for network '%s' and version '%d'", network, version)
			}
		}
	}
}
