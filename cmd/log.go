package cmd

import (
	"fmt"
	"git-light/repo"
	"git-light/util"
	"path/filepath"
)

func Log() error {

	//. START FRA HEAD
	hash, err := repo.ReadHEAD()
	if err != nil {
		fmt.Print("No commits yet \n")
		return err
	}

	for hash != "" {
		subFolder := hash[:2]
		filename := hash[2:]
		path := filepath.Join(".git-light", "objects", subFolder, filename)
		data, err := util.ReadAll(path)
		if err != nil {
			return err
		}
		commit := repo.ParseCommit(data)
		fmt.Println("commit", hash)
		fmt.Println("      ", commit.Message)
		fmt.Println()

		hash = commit.ParentHash

	}
	return nil
}
