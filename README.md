# ByteBrush: Duplicate File Finder

A command-line tool written in Go to find and manage duplicate files in a specified directory.

## Table of Contents

- [Features](#features)
- [Downloads](#downloads)
- [Usage](#usage)
  - [Using Pre-built Binaries](#using-pre-built-binaries)
  - [Building from Source](#building-from-source)
- [Flags](#flags)
- [Project Structure](#project-structure)
- [Code Quality](#code-quality)
- [Contributing](#contributing)
- [License](#license)

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

## Downloads

You can download the latest pre-built binaries for Windows, macOS, and Linux from the [GitHub Releases page](https://github.com/gyaan/ByteBrush/releases).

## Usage

### Using Pre-built Binaries

1.  Download the appropriate binary for your operating system from the [releases page](https://github.com/gyaan/ByteBrush/releases).
2.  Extract the archive and run the executable from your terminal.

For example, on macOS or Linux:

```bash
./bytebrush_darwin_amd64 --dir /path/to/search
```

On Windows:

```bash
./bytebrush_windows_amd64.exe --dir C:\path\to\search
```

### Building from Source

#### Prerequisites

*   Go (1.23 or later)

#### Installation

1.  Clone the repository:
    ```bash
    git clone https://github.com/gyaan/ByteBrush.git
    ```
2.  Navigate to the project directory:
    ```bash
    cd ByteBrush
    ```

#### Build Commands

To build the binaries for different operating systems, use the following commands:

*   **Linux:**
    ```bash
    GOOS=linux GOARCH=amd64 go build -o bin/bytebrush_linux_amd64 cmd/duplicate-files/main.go
    ```
*   **Windows:**
    ```bash
    GOOS=windows GOARCH=amd64 go build -o bin/bytebrush_windows_amd64.exe cmd/duplicate-files/main.go
    ```
*   **macOS (Intel):**
    ```bash
    GOOS=darwin GOARCH=amd64 go build -o bin/bytebrush_darwin_amd64 cmd/duplicate-files/main.go
    ```
*   **macOS (Apple Silicon):**
    ```bash
    GOOS=darwin GOARCH=arm64 go build -o bin/bytebrush_darwin_arm64 cmd/duplicate-files/main.go
    ```

## Flags

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

## Code Quality

This project uses Go's built-in testing framework. To run the tests, use the following command:

```bash
go test ./...
```

The code is formatted using `gofmt`.

## Contributing

1.  Fork the repository.
2.  Create your feature branch (`git checkout -b feature/amazing-feature`).
3.  Commit your changes (`git commit -m 'Add some amazing feature'`).
4.  Push to the branch (`git push origin feature/amazing-feature`).
5.  Open a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
