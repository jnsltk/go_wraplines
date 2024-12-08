package main

import (
	"flag"
	"fmt"
	"log"
	"os"
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
	isSmartPtr := flag.Bool("smart", true, "Sets if the tool should recognise words")

	flag.Parse()
	if *filePathPtr == "none" {
		log.Fatal("Please provide a file path")
	}
	return filePathPtr, wrapPtr, isSmartPtr
}

func wrapText(text string, wrap int, isSmart bool) string {
	// Remove all existing end line characters
	text = strings.Replace(text, "\n", " ", -1)
	if isSmart {
		for i := 0; i < len(text); i++ {
			if i%wrap == 0 && i > 0 {
				j := i
				for text[j] != ' ' {
					j--
				}
				text = replaceAtIndex(text, '\n', j)
			}
		}
	} else {
		for i := 0; i < len(text); i++ {
			if i%wrap == 0 && i > 0 {
				text = text[:i] + "\n" + text[i:]
			}
		}
	}

	return text
}

func replaceAtIndex(input string, char rune, i int) string {
	out := []rune(input)
	out[i] = char
	return string(out)
}
