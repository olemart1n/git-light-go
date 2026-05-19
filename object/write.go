package object

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"os"
	"path/filepath"
	"strconv"
)

// Write creates a /objects/**/hash file and writes the content to it.
// func Write(gitDirectory string, hash string, object []byte) (string, error) {
func Write(obj GitObject) (string, error) {

	// 1. HEADER

	body := obj.Body()
	objType := obj.Type()
	var header bytes.Buffer
	header.WriteString(objType)
	header.WriteString(" ")
	header.WriteString(strconv.Itoa(len(body)))
	header.WriteByte(0)
	header.Write([]byte(body))
	//
	// 4. HASH COMMITEN
	sum := sha1.Sum(header.Bytes())             // Array [20]byte
	encodedObject := hex.EncodeToString(sum[:]) // [:] gjør om array til []byte

	// BRUK DE TO FØRSTE KARAKTERENE I HASHEN SOM MAPPE-NAVN
	dir := filepath.Join(".git-light", "objects", encodedObject[:2])
	// BRUK DE RESTERENDE KARAKTERENE SOM FILNAVN
	fp := filepath.Join(dir, encodedObject[2:])

	// SØRG FOR AT MAPPEN FINNES
	if err := os.MkdirAll(dir, 0775); err != nil {
		return "", err
	}

	// SKRIV OBJEKTET TIL FILEN
	if err := os.WriteFile(fp, header.Bytes(), 0444); err != nil { // 0044 GJØR FILEN LESBAR FOR ALLE
		return "", err
	}

	return encodedObject, nil
}
