package object

import (
	"encoding/hex"
)

type TreeEntry struct {
	Mode string
	Name string
	Hash string
}

type Tree struct {
	Entries []TreeEntry
}

func (t Tree) Body() []byte {

	body := make([]byte, 0)

	for _, e := range t.Entries {
		// mode + space
		body = append(body, e.Mode...) // ANNEN SYNTAX MEN FUNKER OGSÅ. HVER BYTE (RUNE) FRA STRENGER BLIR PAKKER UT OG APPENDET.
		body = append(body, ' ')

		// navn + nullbyte
		body = append(body, []byte(e.Name)...) // DET SAMME SKJER HER SOM 4 LINJER OVER.
		body = append(body, 0)

		// SHA1 MÅ VÆRE RÅBYTES. OFFISIELLE GIT FORVENTER AT HASHENE I UKOMPRIMERT TREE OBJEKT SKAL VÆRE 20 BYTES, IKKE 40.
		rawHash, _ := hex.DecodeString(e.Hash) // KOMPRIMERER SHA1(40 bytes) til RÅ-BYTES (20 bytes)
		body = append(body, rawHash...)

	}

	return body
}

func (t Tree) Type() string {
	return "tree"
}
