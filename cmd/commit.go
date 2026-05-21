// Package cmd contains function for each command line argument
package cmd

import (
	"git-light/object"
	"git-light/repo"
	"git-light/util"
	"log"
)

func Commit(rootDir string, message string) error {
	parentHash, err := repo.ReadHEAD()
	if err != nil {
		return err
	}

	files, err := util.ScanDirForFiles(rootDir)

	if err != nil {
		log.Fatal(err)
	}

	entries := make([]object.TreeEntry, 0, len(files)) // LISTEN MED FILER SOM SKAL I TREET

	for _, file := range files {
		content, err := util.ReadAll(file)
		if err != nil {
			return err
		}

		// Et nytt blob objekt for hver fil i working directory
		blob := object.Blob{Data: content}
		hash, err := object.Write(blob)
		if err != nil {
			return err
		}
		entries = append(entries, object.TreeEntry{
			Mode: "100644",
			Name: file,
			Hash: hash})
	}

	// Tree objektet
	tree := object.Tree{
		Entries: entries,
	}
	treeHash, err := object.Write(tree)
	if err != nil {
		return err
	}

	commit := object.Commit{
		Message:    message,
		ParentHash: parentHash,
		Author:     "Name <name@mail.com>",
		TreeHash:   treeHash,
		Committer:  "some committer name",
	}
	commitHash, err := object.Write(commit)
	if err != nil {
		return err
	}

	repo.UpdateHEAD(commitHash)
	return nil
}
