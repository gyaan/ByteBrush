package fileutils

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFindFilesBySize(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir, err := os.MkdirTemp("", "testdir")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create test files
	createTestFile(t, tmpDir, "file1.txt", "hello")
	createTestFile(t, tmpDir, "file2.txt", "world")
	createTestFile(t, tmpDir, "file3.txt", "hello")
	createTestFile(t, tmpDir, "file4.txt", "different")

	// Create a subdirectory and a file in it
	subDir := filepath.Join(tmpDir, "subdir")
	os.Mkdir(subDir, 0755)
	createTestFile(t, subDir, "file5.txt", "hello")

	// Run the function
	filesBySize, err := FindFilesBySize(tmpDir, nil, nil, 0, 0)
	if err != nil {
		t.Fatalf("FindFilesBySize failed: %v", err)
	}

	// Check the results
	if len(filesBySize[5]) != 3 {
		t.Errorf("Expected 3 files of size 5, but got %d", len(filesBySize[5]))
	}
	if len(filesBySize[9]) != 1 {
		t.Errorf("Expected 1 file of size 9, but got %d", len(filesBySize[9]))
	}
}

func createTestFile(t *testing.T, dir, name, content string) {
	t.Helper()
	filePath := filepath.Join(dir, name)
	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to create test file %s: %v", name, err)
	}
}
