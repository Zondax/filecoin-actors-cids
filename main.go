package main

import (
	"encoding/json"
	"fmt"
	lotusBuild "github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/actors"
	"io/ioutil"
)

const outputFileName = "gen/builtin-actors-cids.json"

type BuiltinActorsMetadata struct {
	Network     string
	Version     actors.Version
	ManifestCid string
	Actors      map[string]string
}

func main() {
	if len(lotusBuild.EmbeddedBuiltinActorsMetadata) == 0 {
		fmt.Println("error: EmbeddedBuiltinActorsMetadata is empty!")
		return
	}

	var meta []BuiltinActorsMetadata
	for _, d := range lotusBuild.EmbeddedBuiltinActorsMetadata {
		cids := make(map[string]string)
		for name, cid := range d.Actors {
			cids[name] = cid.String()
		}

		meta = append(meta, BuiltinActorsMetadata{
			Network:     d.Network,
			Version:     d.Version,
			ManifestCid: d.ManifestCid.String(),
			Actors:      cids,
		})
	}

	j, err := json.MarshalIndent(meta, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ioutil.WriteFile(outputFileName, j, 0644)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Downloaded actors cids version %s", meta[0].Version)
}
