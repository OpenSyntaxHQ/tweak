package utils

import (
	"bufio"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/iancoleman/strcase"
)

func StdinHasData() bool {
	fi, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	return (fi.Mode() & os.ModeCharDevice) == 0
}

func ReadPipedInput() ([]byte, error) {
	return io.ReadAll(os.Stdin)
}

func ReadMultilineInput() string {
	lines := make([]string, 0, 32)
	scanner := bufio.NewScanner(os.Stdin)
	empty := 0

	for scanner.Scan() {
		text := scanner.Text()
		lines = append(lines, text)

		if text == "" {
			empty++
			if empty == 2 {
				break
			}
		} else {
			empty = 0
		}
	}

	if len(lines) >= 2 {
		lines = lines[:len(lines)-2]
	}
	return strings.Join(lines, "\n")
}

var whitespaceRe = regexp.MustCompile(`\s+`)

func ToKebabCase(input []byte) string {
	str := whitespaceRe.ReplaceAllString(string(input), " ")
	return strcase.ToKebab(str)
}

func ToLowerCamelCase(input []byte) string {
	str := whitespaceRe.ReplaceAllString(string(input), " ")
	return strcase.ToLowerCamel(str)
}
