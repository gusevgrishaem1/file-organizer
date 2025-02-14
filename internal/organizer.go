package organizer

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

// OrganizeFiles moves files into categorized folders
func OrganizeFiles(directory string) error {
	files, err := os.ReadDir(directory)
	if err != nil {
		return fmt.Errorf("failed to read directory: %w", err)
	}

	file, err := os.Open("config.json")
	if err != nil {
		return fmt.Errorf("failed to read config: %w", err)
	}

	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	log.Print(len(data))

	var categories map[string][]string
	json.Unmarshal(data, &categories)

	log.Print(files)

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fileExt := filepath.Ext(file.Name())
		destFolder := getDestination(fileExt, categories)

		if destFolder != "" {
			moveFile(directory, file.Name(), destFolder)
		}
	}

	fmt.Println("✅ Organization complete!")
	return nil
}

// getDestination finds the category for an extension
func getDestination(ext string, categories map[string][]string) string {
	for category, extensions := range categories {
		for _, validExt := range extensions {
			if validExt == ext {
				return category
			}
		}
	}
	return ""
}

// moveFile moves a file into the target folder
func moveFile(directory, fileName, category string) {
	destPath := filepath.Join(directory, category)
	os.MkdirAll(destPath, os.ModePerm)

	oldPath := filepath.Join(directory, fileName)
	newPath := filepath.Join(destPath, fileName)

	err := os.Rename(oldPath, newPath)
	if err != nil {
		fmt.Printf("❌ Error moving %s: %v\n", fileName, err)
	} else {
		fmt.Printf("📂 Moved %s → %s\n", fileName, category)
	}
}
