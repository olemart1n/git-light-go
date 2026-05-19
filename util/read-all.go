package util

import (
	"io"
	"os"
)

// ReadAll returns all bytes from a file ([]byte)
func ReadAll(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	buf := make([]byte, 4096) // 4kb buffer
	var out []byte
	for {
		n, err := file.Read(buf) // n is all read bytes

		if n > 0 {
			out = append(out, buf[:n]...)
		}

		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

	}

	return out, nil
}
