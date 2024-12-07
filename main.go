package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	executablePath, err := os.Executable()
	if err != nil {
		log.Fatalf("Error getting executable path: %v", err)
	}
	executableDir := filepath.Dir(executablePath)
	fmt.Printf(executableDir)
}
