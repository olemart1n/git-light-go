package object

import (
	"bytes"
	"encoding/hex"
	"errors"
)

func ParseTree(data []byte) ([]TreeEntry, error) {

	_, body, found := bytes.Cut(data, []byte{0})
	if !found {
		return nil, errors.New("could not seperate header and body, null-byte not found")
	}

	var entries []TreeEntry
	pointer := 0

	for pointer < len(body) {
		// SJEKK DETTE: Hvis det bare er tomme bytes eller et linjeskift igjen på slutten, stopp løkken.
		if pointer >= len(body) || len(bytes.TrimSpace(body[pointer:])) == 0 {
			break
		}

		// |<Mode>' '<Name>x\00<Hash>|

		// 1. Finn slutten på Mode (første mellomrom fra der vi står)
		spaceIndex := bytes.IndexByte(body[pointer:], ' ')

		if spaceIndex == -1 {
			return nil, errors.New("ugyldig format: fant ikke mellomrom etter mode")
		}

		mode := string(body[pointer : pointer+spaceIndex])

		// Flytter pekeren forbi mellomrom ' '
		pointer = pointer + spaceIndex + 1

		// 2. Finn slutten på Name (første null-byte fra der pointer er nå)
		nullIndex := bytes.IndexByte(body[pointer:], 0)

		if nullIndex == -1 {
			return nil, errors.New("ugyldig format: fant ikke null-byte etter Name")
		}
		name := string(body[pointer : pointer+nullIndex])

		// Flytter pekeren forbi null-byte
		pointer = pointer + nullIndex + 1

		// 3. Finn slutten på Hash. Hash ER de neste 20 bytene.
		if pointer+20 > len(body) {
			return nil, errors.New("ugyldig format: ufullstendig sha1 hash")
		}
		hashSlice := body[pointer : pointer+20]
		hashHex := hex.EncodeToString(hashSlice)

		// Flytter pekeren 20 bytes frem.
		pointer += 20

		entry := TreeEntry{
			Mode: mode,
			Name: name,
			Hash: hashHex,
		}
		entries = append(entries, entry)
	}

	return entries, nil
}
