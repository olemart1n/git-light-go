package repo

import (
	"fmt"
	"git-light/object"
	"git-light/util"
	"path"
)

func Checkout(hash string) {

	path := path.Join(".git-light", "objects", hash[:2], hash[2:])
	commitData, err := util.ReadAll(path)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	entries, err := object.ParseTree(commitData)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	fmt.Println("LENGT OF entries: ", len(entries))

	for _, file := range entries {
		fmt.Println("Mode: ", string(file.Mode), " Name", string(file.Name), " Hash", string(file.Hash))
	}

}
