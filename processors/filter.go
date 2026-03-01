package processors

import (
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Grep struct{}

func (p Grep) Name() string    { return "grep" }
func (p Grep) Alias() []string { return []string{"filter", "filter-lines"} }
func (p Grep) Transform(data []byte, f ...Flag) (string, error) {
	var pattern string
	invert := false
	for _, flag := range f {
		switch flag.Short {
		case "p":
			if s, ok := flag.Value.(string); ok {
				pattern = s
			}
		case "v":
			if b, ok := flag.Value.(bool); ok {
				invert = b
			}
		}
	}
	if pattern == "" {
		return "", fmt.Errorf("--pattern is required")
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		return "", fmt.Errorf("invalid regex: %w", err)
	}
	lines := strings.Split(string(data), "\n")
	var result []string
	for _, line := range lines {
		matches := re.MatchString(line)
		if (matches && !invert) || (!matches && invert) {
			result = append(result, line)
		}
	}
	return strings.Join(result, "\n"), nil
}
func (p Grep) Flags() []Flag {
	return []Flag{
		{Name: "pattern", Short: "p", Desc: "Regex pattern to match", Value: "", Type: FlagString},
		{Name: "invert", Short: "v", Desc: "Invert match (exclude matching lines)", Value: false, Type: FlagBool},
	}
}
func (p Grep) Title() string {
	return fmt.Sprintf("%s (%s)", cases.Title(language.Und, cases.NoLower).String(p.Name()), p.Name())
}
func (p Grep) Description() string { return "Filter lines matching a pattern" }
func (p Grep) FilterValue() string { return p.Title() }
