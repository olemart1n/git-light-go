package cmd

import (
	"fmt"
	"git-light/object"
	"git-light/util"
	"path"
)

func Checkout(hash string) {

	path := path.Join(".git-light", "objects", hash[:2], hash[2:])
	compressedData, err := util.ReadAll(path)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	decompressed, err := util.DecompressBytes(compressedData)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	entries, err := object.ParseTree(decompressed)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	fmt.Println("LENGT OF entries: ", len(entries))

	for _, file := range entries {
		fmt.Println(string(file.Mode), string(file.Name), string(file.Hash))
	}

}
