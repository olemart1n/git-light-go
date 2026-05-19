package repo

import (
	"bytes"
	"fmt"
	"strings"
)

type Commit struct {
	TreeHash      string
	ParentHash    string
	Message       string
	AuthorData    string
	CommitterData string
}

func ParseCommit(data []byte) Commit {
	_, after, found := bytes.Cut(data, []byte{0})
	if !found {
		fmt.Println("empty byte not found")
		return Commit{}
	}

	text := string(after)
	lines := strings.Split(text, "\n")

	c := Commit{}

	inMessage := false

	for i, line := range lines {
		if inMessage {
			c.Message += line
			continue
		}

		// GIT FORMATET HAR EN TOM LINE RETT FØR COMMIT-MELDINGEN
		if line == "" && i > 3 {
			inMessage = true
			continue
		}

		if treeHash, found := strings.CutPrefix(line, "tree "); found {
			c.TreeHash = treeHash
			continue
		}
		if parentHash, found := strings.CutPrefix(line, "parent "); found {
			c.ParentHash = parentHash
			continue
		}
		if authorData, found := strings.CutPrefix(line, "author "); found {
			c.AuthorData = authorData
			continue
		}
		if committerData, found := strings.CutPrefix(line, "committer "); found {
			c.CommitterData = committerData
		}

	}

	return c

}
