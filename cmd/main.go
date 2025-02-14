package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	organizer "github.com/gusevgrishaem1/file-organizer/internal"
)

func main() {
	// Define CLI flag
	path := flag.String("path", ".", "Directory to organize")
	flag.Parse()

	// Check if path exists
	if _, err := os.Stat(*path); os.IsNotExist(err) {
		log.Fatalf("❌ Path does not exist: %s", *path)
	}

	fmt.Println("🔄 Organizing files in:", *path)
	err := organizer.OrganizeFiles(*path)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
