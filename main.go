package main

import (
	"encoding/json"
	"fmt"
	"github.com/zondax/filecoin-actors-cids/utils"
	"io/ioutil"
	"os"
	"path/filepath"
)

const outputFileName = "gen/builtin-actors-cids"

func main() {
	exportMetadata(utils.PreviousVersion)
	exportMetadata(utils.LatestVersion)
}

func exportMetadata(version uint) {
	meta := utils.GetFullMetadata(version)
	if len(meta) == 0 {
		fmt.Printf("Metadata for actors version '%d' is empty", version)
		return
	}

	j, err := json.MarshalIndent(meta, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}

	outPath := filepath.Join(".", "gen")
	err = os.MkdirAll(outPath, os.ModePerm)
	if err != nil {
		fmt.Println("could not create output path")
		return
	}

	err = ioutil.WriteFile(fmt.Sprintf("%s_V%d.json", outputFileName, version), j, 0600)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("json file created for actors cids with version", version)
}
