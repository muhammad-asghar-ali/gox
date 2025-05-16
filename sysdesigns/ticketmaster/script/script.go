// Support only uuid.UUID
// FIXME -
//  1. Now on only on flag updating all the filed with same name
//  2. Not formatting the code after update

package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

const (
	out     = "internal/entities"
	queries = "internal/db/queries"
)

var (
	opgex = regexp.MustCompile(`--\s*@optional\s+(\w+)`)
	sfgex = regexp.MustCompile(`(\s+)(\w+)\s+uuid\.UUID(\s+` + "`[^`]*`" + `)?`)
)

func main() {
	files, err := os.ReadDir(out)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	opfs := parse(queries)

	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".go") {
			continue
		}

		filePath := out + "/" + file.Name()
		modify(filePath, opfs)
	}
}

func parse(queries string) map[string]bool {
	opfs := make(map[string]bool)

	files, err := os.ReadDir(queries)
	if err != nil {
		fmt.Println("Error reading SQL queries:", err)
		return opfs
	}

	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".sql") {
			continue
		}

		content, err := os.ReadFile(queries + "/" + file.Name())
		if err != nil {
			fmt.Println("Error reading file:", err)
			continue
		}

		matches := opgex.FindAllStringSubmatch(string(content), -1)
		for _, match := range matches {
			if len(match) == 2 {
				normalized := strings.ToLower(strings.ReplaceAll(match[1], "_", ""))
				opfs[normalized] = true
			}
		}
	}

	return opfs
}

func modify(filePath string, opfs map[string]bool) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading Go file:", err)
		return
	}

	updated := sfgex.ReplaceAllStringFunc(string(content), func(match string) string {
		matches := sfgex.FindStringSubmatch(match)
		if len(matches) >= 3 {
			f_name := matches[2]
			normalized := strings.ToLower(strings.ReplaceAll(f_name, "_", ""))

			if opfs[normalized] {
				return matches[1] + f_name + " *uuid.UUID" + matches[3]
			}
		}
		return match
	})

	if updated != string(content) {
		err := os.WriteFile(filePath, []byte(updated), 0644)
		if err != nil {
			fmt.Println("Error writing updated file:", err)
		}
	}
}
