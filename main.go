package main

import (
	"encoding/json"
	"fmt"
	"github.com/zondax/filecoin-actors-cids/utils"
	"io/ioutil"
)

const outputFileName = "gen/builtin-actors-cids"

func main() {
	exportMetadata(utils.ActorsV7)
	exportMetadata(utils.ActorsV8)
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

	err = ioutil.WriteFile(fmt.Sprintf("%s_V%d.json", outputFileName, version), j, 0600)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("json file created for actors cids with version", version)
}
