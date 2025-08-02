package ui

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// GetUserSelection prompts the user to select a file to keep from a list of duplicates.
func GetUserSelection(paths []string) (int, error) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Select which file to keep (enter the number):")
		for i, path := range paths {
			fmt.Printf("  %d: %s\n", i+1, path)
		}
		fmt.Print("> ")

		input, err := reader.ReadString('\n')
		if err != nil {
			return -1, err
		}

		choice, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil || choice < 1 || choice > len(paths) {
			fmt.Println("Invalid selection. Please try again.")
			continue
		}
		return choice - 1, nil
	}
}
