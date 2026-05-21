package util

import (
	"bytes"
	"compress/zlib"
	"io"
)

func DecompressBytes(compressedData []byte) ([]byte, error) {

	if len(compressedData) == 0 {
		return nil, io.ErrUnexpectedEOF
	}

	reader := bytes.NewReader(compressedData)
	zr, err := zlib.NewReader(reader)
	if err != nil {
		return nil, err
	}
	defer zr.Close()

	// 3. Les den utpakkede dataen over i en buffer
	var uncompressedData bytes.Buffer
	io.Copy(&uncompressedData, reader)

	// Returnerer det originale innholdet (header + filinnhold)
	return uncompressedData.Bytes(), nil

}
