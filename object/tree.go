package object

import "encoding/hex"

type TreeEntry struct {
	Mode string
	Name string
	Hash string
}

type Tree struct {
	Entries []TreeEntry
}

func (t Tree) Body() string {

	var body []byte

	for _, e := range t.Entries {

		// mode + space
		body = append(body, []byte(e.Mode)...)
		body = append(body, ' ')

		// navn + nullbyte
		body = append(body, []byte(e.Name)...)
		body = append(body, 0)

		// SHA1 MÅ VÆRE RÅBYTES
		rawHash, _ := hex.DecodeString(e.Hash)
		body = append(body, rawHash...)
	}

	return string(body)
}

func (t Tree) Type() string {
	return "tree"
}
