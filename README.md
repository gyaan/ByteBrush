# Go Duplicate File Finder

A command-line tool written in Go to find and manage duplicate files in a specified directory.

## Features

*   Finds duplicate files based on their content hash (SHA-256).
*   Recursively searches through subdirectories.
*   Concurrent processing for faster hashing.
*   Optimized hashing strategy (groups files by size first).
*   Interactive mode to select which files to delete.
*   Dry run mode to preview deletions.
*   Exclude directories and file patterns.
*   Filter by file type and size.
*   Formatted output (text or JSON).

## Getting Started

### Prerequisites

*   Go (1.23 or later)

### Installation

1.  Clone the repository:
    ```bash
    git clone https://github.com/gyaan/ByteBrush.git
    ```
2.  Navigate to the project directory:
    ```bash
    cd ByteBrush
    ```

### Usage

To run the tool, use the following command:

```bash
go run cmd/duplicate-files/main.go [flags]
```

#### Flags

*   `--dir`: Directory to search for duplicate files (default: `.`)
*   `--dry-run`: Print files that would be deleted without actually deleting them.
*   `--exclude`: Comma-separated list of directories or file patterns to exclude.
*   `--include`: Comma-separated list of file patterns to include.
*   `--min-size`: Minimum file size in bytes.
*   `--max-size`: Maximum file size in bytes.
*   `--format`: Output format (text, json) (default: `text`)
*   `--interactive`: Prompt for which file to keep when duplicates are found.

## Project Structure

*   `cmd/duplicate-files/main.go`: The main application entry point.
*   `pkg/filehasher/filehasher.go`: The package responsible for calculating file hashes.
*   `pkg/fileutils/fileutils.go`: The package responsible for finding files and filtering them.
*   `pkg/ui/ui.go`: The package responsible for the interactive user interface.
*   `go.mod`: The Go module definition file.

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.