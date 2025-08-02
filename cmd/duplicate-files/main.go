package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"duplicate-files/pkg/filehasher"
	"duplicate-files/pkg/fileutils"
	"duplicate-files/pkg/ui"
)

func main() {
	// --- Flags ---
	dir := flag.String("dir", ".", "Directory to search for duplicate files")
	dryRun := flag.Bool("dry-run", false, "Print files that would be deleted without actually deleting them")
	exclude := flag.String("exclude", "", "Comma-separated list of directories or file patterns to exclude")
	include := flag.String("include", "", "Comma-separated list of file patterns to include")
	minSize := flag.Int64("min-size", 0, "Minimum file size in bytes")
	maxSize := flag.Int64("max-size", 0, "Maximum file size in bytes")
	outputFormat := flag.String("format", "text", "Output format (text, json)")
	interactive := flag.Bool("interactive", false, "Prompt for which file to keep when duplicates are found")

	flag.Parse()

	// --- File Discovery ---
	filesBySize, err := fileutils.FindFilesBySize(*dir, strings.Split(*exclude, ","), strings.Split(*include, ","), *minSize, *maxSize)
	if err != nil {
		log.Fatalf("Error finding files: %v", err)
	}

	// --- Hashing ---
	hashes := make(map[[32]byte][]string)
	for _, paths := range filesBySize {
		if len(paths) < 2 {
			continue
		}

		results := filehasher.CalculateHashes(paths)
		for result := range results {
			if result.Err != nil {
				log.Printf("Could not calculate hash for %s: %v", result.Path, result.Err)
				continue
			}
			hashes[result.Hash] = append(hashes[result.Hash], result.Path)
		}
	}

	// --- Output and Deletion ---
	duplicates := getDuplicates(hashes)

	switch *outputFormat {
	case "json":
		printJSON(duplicates)
	case "text":
		printText(duplicates, *dryRun, *interactive)
	default:
		log.Fatalf("Unsupported output format: %s", *outputFormat)
	}
}

func getDuplicates(hashes map[[32]byte][]string) [][]string {
	var duplicates [][]string
	for _, paths := range hashes {
		if len(paths) > 1 {
			duplicates = append(duplicates, paths)
		}
	}
	return duplicates
}

func printJSON(duplicates [][]string) {
	output, err := json.MarshalIndent(duplicates, "", "  ")
	if err != nil {
		log.Fatalf("Error formatting JSON: %v", err)
	}
	fmt.Println(string(output))
}

func printText(duplicates [][]string, dryRun, interactive bool) {
	if len(duplicates) == 0 {
		fmt.Println("No duplicate files found.")
		return
	}

	fmt.Println("Duplicate files found:")
	for _, paths := range duplicates {
		fmt.Println(strings.Join(paths, "\n"))
		fmt.Println()

		if interactive {
			selection, err := ui.GetUserSelection(paths)
			if err != nil {
				log.Printf("Error getting user selection: %v", err)
				continue
			}

			for i, path := range paths {
				if i != selection {
					deleteFile(path, dryRun)
				}
			}
		}
	}
}

func deleteFile(path string, dryRun bool) {
	if dryRun {
		fmt.Printf("[Dry Run] Would delete file: %s\n", path)
		return
	}

	fmt.Printf("Deleting file: %s\n", path)
	err := os.Remove(path)
	if err != nil {
		log.Printf("Could not delete file %s: %v", path, err)
	}
}