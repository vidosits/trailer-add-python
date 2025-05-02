package test

import (
	"fmt"
	"os"
	"path/filepath"
)

func test() {
	// Define the root directory to start walking from.
	// "." represents the current directory.
	rootDir := "."

	// Use filepath.Walk to recursively traverse the directory tree.
	// It takes the root path and a WalkFunc.
	// The WalkFunc is called for each file or directory visited.
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		// If an error occurs while accessing a file or directory,
		// print the error and continue walking (or return the error
		// to stop the walk entirely). Here, we just print and continue.
		if err != nil {
			fmt.Printf("Error accessing path %q: %v\n", path, err)
			return err // Return the error to stop the walk on this branch
			// return nil // Or return nil to ignore the error and continue
		}

		// Print the path of the current file or directory.
		// info.Name() would print just the name, path prints the full path from rootDir.
		fmt.Println(path)

		// Return nil to continue the walk.
		// To skip a directory and its contents, return filepath.SkipDir.
		return nil
	})

	// Check if there was an error during the walk itself (e.g., permission denied on rootDir)
	if err != nil {
		fmt.Printf("Error walking the directory %q: %v\n", rootDir, err)
		// os.Exit(1) // Exit with a non-zero status code to indicate an error
	}
}
