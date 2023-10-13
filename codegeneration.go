package main

import (
	"encoding/json"
	"fmt"
)

type Program struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Author  string `json:"author"`
	Code    struct {
		Type    string `json:"type"`
		Content string `json:"description"`
	} `json:"code"`
}

func generateProgram(jsonStr string) string {
	var p Program
	if err := json.Unmarshal([]byte(jsonStr), &p); err != nil {
		return fmt.Sprintf("Json problem: %v", err)
	}

	programCode := fmt.Sprintf(
		"# %s (verison %s)\n# Author: %s\n\nfunc main() {\n",
		p.Name, p.Version, p.Author,
	)

	if p.Code.Type == "Go" {
		programCode += fmt.Sprintf("    %s\n}\n", p.Code.Content)
	} else {
		programCode += "    fmt.Println(\"Unknown type of code.\")\n}\n"
	}

	return programCode
}

func main() {
	jsonStr := `
    {
      "name": "My program",
      "version": "1.0",
      "author": "Author name",
      "code": {
        "type": "Go",
        "description": "fmt.Println(\"Hello, world!\")"
      }
    }`

	generatedCode := generateProgram(jsonStr)
	fmt.Println(generatedCode)
}
