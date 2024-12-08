package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// ════════════════════════════════════ Functions ═════════════════════════════════════
func main() {
	filePath, wrap, isSmart := parseArgs()
	text := loadFile(*filePath)
	fmt.Println(wrapText(text, *wrap, *isSmart))
}

func loadFile(path string) string {
	dat, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	return string(dat)
}

func parseArgs() (*string, *int, *bool) {
	filePathPtr := flag.String("f", "none", "Input file path")
	wrapPtr := flag.Int("w", 80, "Wrap text at column N")
	isSmartPtr := flag.Bool("s", true, "Sets if the tool should recognise words")

	flag.Parse()
	if *filePathPtr == "none" {
		log.Fatal("Please provide a file path")
	}
	return filePathPtr, wrapPtr, isSmartPtr
}

func wrapText(text string, wrap int, isSmart bool) string {
	text = strings.Replace(text, "\n", " ", -1)
	wrapStr := strconv.Itoa(wrap)
	return text + wrapStr
}
