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

	// 4. HASH header
	sum := sha1.Sum(header.Bytes())             // Array [20]byte
	encodedObject := hex.EncodeToString(sum[:]) //

	dir := filepath.Join(".git-light", "objects", encodedObject[:2])
	fp := filepath.Join(dir, encodedObject[2:])

	// SØRG FOR AT MAPPEN FINNES
	if err := os.MkdirAll(dir, 0775); err != nil {
		return "", err
	}

	// OBJEKTET ER IMMUTABLE, SÅ SJEKK OM DET ALLEREDE EKSISTER. VISST IKKE HOPP OVER WRITE
	if _, err := os.Stat(fp); err == nil {
		return encodedObject, nil
	}

	// SKRIV OBJEKTET TIL FILEN
	if err := os.WriteFile(fp, header.Bytes(), 0444); err != nil { // 0044 GJØR FILEN LESBAR FOR ALLE
		return "", err
	}

	return encodedObject, nil
}
