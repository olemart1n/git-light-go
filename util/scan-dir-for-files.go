package util

import (
	"io/fs"
	"path/filepath"
)

// ScanDirForFiles traverses the directory and return an array with filenames, skipping directory names.
// It ignores ".git-light", ".git".
func ScanDirForFiles(dir string) ([]string, error) {

	var files []string

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {

		if err != nil {
			return err
		}
		if path == ".git-light" || path == ".git" {
			return filepath.SkipDir
		}
		if d.IsDir() {
			return nil
		}

		files = append(files, path)

		return nil

	})

	if err != nil {
		return nil, err
	}

	return files, nil

}
