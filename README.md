# jsstrings
Simple tool that uses my fork of go-fAST to read strings from a JS file


## Usage
```go
import (
  "fmt"
  "github.com/bebiksior/jsstrings"
)

// JavaScript code to parse
content := `const str = "hello world";`

// Extract strings, passing the JS content and a URL for reference
strings, err := jsstrings.ExtractStrings(content, "https://example.com/script.js")
if err != nil {
	fmt.Printf("Error extracting strings: %v\n", err)
	return
}

fmt.Println(strings)
```
