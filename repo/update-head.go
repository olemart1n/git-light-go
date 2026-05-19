package repo

import (
	"os"
	"path/filepath"
)

func UpdateHEAD(commitHash string) error {
	path := filepath.Join(".git-light", "HEAD")
	content := []byte(commitHash + "\n")
	return os.WriteFile(path, content, 0664)

}
