package filehasher

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCalculateHashes(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir, err := os.MkdirTemp("", "testdir")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create test files
	file1Path := createTestFile(t, tmpDir, "file1.txt", "hello")
	file2Path := createTestFile(t, tmpDir, "file2.txt", "world")

	// Run the function
	paths := []string{file1Path, file2Path}
	results := CalculateHashes(paths)

	// Check the results
	count := 0
	for result := range results {
		if result.Err != nil {
			t.Errorf("CalculateHashes failed for %s: %v", result.Path, result.Err)
		}
		count++
	}

	if count != 2 {
		t.Errorf("Expected 2 results, but got %d", count)
	}
}

func createTestFile(t *testing.T, dir, name, content string) string {
	t.Helper()
	filePath := filepath.Join(dir, name)
	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to create test file %s: %v", name, err)
	}
	return filePath
}
