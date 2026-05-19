// Package cmd contains function for each command line argument
package cmd

import (
	"git-light/object"
	"git-light/repo"
	"git-light/util"
	"log"
)

func Commit(rootDir string, message string) error {

	files, err := util.ScanWorkingDirectory(rootDir)
	if err != nil {
		log.Fatal(err)
	}

	entries := make([]object.TreeEntry, 0, len(files))

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
	treeHash, _ := object.Write(tree)

	parentHash, _ := repo.ReadHEAD()

	commit := object.Commit{
		Message:    message,
		ParentHash: parentHash,
		Author:     "Name <name@mail.com>",
		TreeHash:   treeHash,
	}
	commitHash, err := object.Write(commit)
	if err != nil {
		return err
	}

	repo.UpdateHEAD(commitHash)
	return nil
}
