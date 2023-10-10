package main

import (
	"encoding/json"
	"fmt"
)

type Program struct {
	Name    string `json:"название"`
	Version string `json:"версия"`
	Author  string `json:"автор"`
	Code    struct {
		Type    string `json:"тип"`
		Content string `json:"содержание"`
	} `json:"код"`
}

func generateProgram(jsonStr string) string {
	var p Program
	if err := json.Unmarshal([]byte(jsonStr), &p); err != nil {
		return fmt.Sprintf("Ошибка разбора JSON: %v", err)
	}

	programCode := fmt.Sprintf(
		"# %s (версия %s)\n# Автор: %s\n\nfunc main() {\n",
		p.Name, p.Version, p.Author,
	)

	if p.Code.Type == "Go" {
		programCode += fmt.Sprintf("    %s\n}\n", p.Code.Content)
	} else {
		programCode += "    fmt.Println(\"Неизвестный тип кода.\")\n}\n"
	}

	return programCode
}

func main() {
	jsonStr := `
    {
      "название": "My program",
      "версия": "1.0",
      "автор": "Имя автора",
      "код": {
        "тип": "Go",
        "содержание": "fmt.Println(\"Привет, мир!\")"
      }
    }`

	generatedCode := generateProgram(jsonStr)
	fmt.Println(generatedCode)
}
