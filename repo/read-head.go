// Package repo contains current info like HEAD
package repo

import (
	"os"
	"path/filepath"
	"strings"
)

func ReadHEAD() (string, error) {
	path := filepath.Join(".git-light", "HEAD")
	data, err := os.ReadFile(path)

	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(data)), nil

}
