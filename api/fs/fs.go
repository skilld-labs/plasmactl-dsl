package fs

import (
	"io/fs"
	"os"
	"path/filepath"
)

func DiscoverFiles(paths []string, ext string) ([]string, error) {
	var files []string

	for _, path := range paths {
		pathFiles, err := listFiles(path, ext)
		if err != nil {
			return nil, err
		}
		files = append(files, pathFiles...)
	}

	return files, nil
}

func listFiles(path string, ext string) ([]string, error) {
	var files []string
	err := filepath.WalkDir(path, func(file string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		match, err := filepath.Match(ext, filepath.Ext(d.Name()))
		if err != nil {
			return err
		}
		if match {
			files = append(files, file)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

func WriteFile(file string, bytes []byte) error {
	err := os.WriteFile("/tmp/dat1", bytes, 0644)
	if err != nil {
		return err
	}
	return nil
}
