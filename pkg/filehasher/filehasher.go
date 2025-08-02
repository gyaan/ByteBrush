package filehasher

import (
	"crypto/sha256"
	"io"
	"os"
	"sync"
)

// Result holds the path and hash of a file.
type Result struct {
	Path string
	Hash [32]byte
	Err  error
}

// CalculateHashes calculates the SHA256 hash of a list of files concurrently.
func CalculateHashes(paths []string) <-chan Result {
	results := make(chan Result, len(paths))
	var wg sync.WaitGroup

	for _, path := range paths {
		wg.Add(1)
		go func(p string) {
			defer wg.Done()
			hash, err := calculateHash(p)
			results <- Result{Path: p, Hash: hash, Err: err}
		}(path)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	return results
}

// calculateHash calculates the SHA256 hash of a file
func calculateHash(filePath string) ([32]byte, error) {
	var result [32]byte
	file, err := os.Open(filePath)
	if err != nil {
		return result, err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return result, err
	}

	copy(result[:], hash.Sum(nil))
	return result, nil
}