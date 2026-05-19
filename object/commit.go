// Package object contains types for all git-light objects
package object

import (
	"bytes"
	"strconv"
	"time"
)

type Commit struct {
	Message    string
	Author     string
	ParentHash string
	TreeHash   string
}

// HUSK AT COMMIT IKKE INNEHOLDER NOE KODE RELATERT TIL SELVE PROSJEKTET

func (c Commit) Body() string {
	// 1. Tid i UNIX-format
	now := time.Now()
	timestamp := now.Unix()
	timezone := now.Format("-0700")

	// 2. COMMIT-BODY
	var body bytes.Buffer
	body.WriteString("tree" + " " + c.TreeHash + "\n")
	body.WriteString("\n")
	if c.ParentHash != "" {
		body.WriteString("parent")
		body.WriteString(" ")
		body.WriteString(c.ParentHash)
		body.WriteString("\n")
	}
	body.WriteString("author")
	body.WriteString(" ")
	body.WriteString(c.Author)
	body.WriteString(" ")
	body.WriteString(strconv.FormatInt(timestamp, 10))
	body.WriteString(" ")
	body.WriteString(timezone)
	body.WriteString("\n")

	body.WriteString("committer")
	body.WriteString(" ")
	body.WriteString(c.Author)
	body.WriteString(" ")
	body.WriteString(strconv.FormatInt(timestamp, 10))
	body.WriteString(" ")
	body.WriteString(timezone)
	body.WriteString("\n")
	// TOM LINJE FØR COMMIT-MELDING
	body.WriteString("\n")

	body.WriteString(c.Message)

	return body.String()

}

func (c Commit) Type() string {
	return "commit"
}
