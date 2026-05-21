package util

import (
	"bytes"
	"compress/zlib"
)

func CompressBytes(uncompressedData []byte) ([]byte, error) {
	// 1. Opprett en buffer som skal holde på de komprimerte bytes.
	var compressedData bytes.Buffer

	// 2. Opprett en zlib-writer som skriver til bufferen. zlib.DefaultCompression (nivå 6)
	writer := zlib.NewWriter(&compressedData)

	// 3. Skriv de ukomprimerte rå-bytesene inn til zlib writeren.
	_, err := writer.Write(uncompressedData)
	if err != nil {
		return nil, err
	}

	// 4. Writeren må lukkes, men ikke med defer!
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	return compressedData.Bytes(), nil
}
