package fileutils

import (
	"os"
	"path/filepath"
)

// FindFilesBySize walks a directory and returns a map of file sizes to a list of file paths.
func FindFilesBySize(dir string, exclude, include []string, minSize, maxSize int64) (map[int64][]string, error) {
	filesBySize := make(map[int64][]string)
	excludeMap := make(map[string]bool)
	for _, item := range exclude {
		excludeMap[item] = true
	}

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Exclude directories
		if info.IsDir() {
			if excludeMap[info.Name()] {
				return filepath.SkipDir
			}
			return nil
		}

		// Exclude files based on patterns
		for _, pattern := range exclude {
			matched, _ := filepath.Match(pattern, info.Name())
			if matched {
				return nil
			}
		}

		// Include files based on patterns
		if len(include) > 0 && include[0] != "" {
			included := false
			for _, pattern := range include {
				matched, _ := filepath.Match(pattern, info.Name())
				if matched {
					included = true
					break
				}
			}
			if !included {
				return nil
			}
		}

		// Filter by size
		if minSize > 0 && info.Size() < minSize {
			return nil
		}
		if maxSize > 0 && info.Size() > maxSize {
			return nil
		}

		if info.Mode().IsRegular() {
			filesBySize[info.Size()] = append(filesBySize[info.Size()], path)
		}
		return nil
	})
	return filesBySize, err
}
