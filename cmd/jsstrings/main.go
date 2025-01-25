package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	jsstrings "github.com/bebiksior/jsstrings/strings"
)

func main() {
	showFiles := flag.Bool("show-files", false, "Show which file each string comes from")
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		fmt.Println("Please provide JavaScript file(s) as arguments")
		os.Exit(1)
	}

	seenStrings := make(map[string]bool)

	for _, pattern := range args {
		matches, err := filepath.Glob(pattern)
		if err != nil {
			fmt.Printf("Error globbing pattern %s: %v\n", pattern, err)
			continue
		}

		for _, file := range matches {
			content, err := os.ReadFile(file)
			if err != nil {
				fmt.Printf("Error reading file %s: %v\n", file, err)
				continue
			}

			strings, err := jsstrings.ExtractStrings(string(content), file)
			if err != nil {
				fmt.Printf("Error extracting strings from %s: %v\n", file, err)
				continue
			}

			for _, s := range strings {
				if !seenStrings[s.Value] {
					if *showFiles {
						fmt.Printf("%s: %s\n", file, s.Value)
					} else {
						fmt.Println(s.Value)
					}
					seenStrings[s.Value] = true
				}
			}
		}
	}
}
