package processors

import (
	"fmt"
	"regexp"
	"strings"
)

type RegexMatch struct{}

func (p RegexMatch) Name() string    { return "regex-match" }
func (p RegexMatch) Alias() []string { return []string{"regex-find"} }

func (p RegexMatch) Transform(data []byte, f ...Flag) (string, error) {
	var pattern string
	for _, flag := range f {
		if flag.Short == "p" {
			if s, ok := flag.Value.(string); ok {
				pattern = s
			}
		}
	}
	if pattern == "" {
		return "", fmt.Errorf("--pattern is required")
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		return "", fmt.Errorf("invalid regex pattern: %w", err)
	}
	matches := re.FindAllString(string(data), -1)
	return strings.Join(matches, "\n"), nil
}

func (p RegexMatch) Flags() []Flag {
	return []Flag{{Name: "pattern", Short: "p", Desc: "Regular expression pattern", Value: "", Type: FlagString, Required: true}}
}
func (p RegexMatch) Title() string       { return fmt.Sprintf("Regex Match (%s)", p.Name()) }
func (p RegexMatch) Description() string { return "Extract all regex matches" }
func (p RegexMatch) FilterValue() string { return p.Title() }

type RegexReplace struct{}

func (p RegexReplace) Name() string    { return "regex-replace" }
func (p RegexReplace) Alias() []string { return []string{"regex-sub"} }

func (p RegexReplace) Transform(data []byte, f ...Flag) (string, error) {
	var pattern, replacement string
	for _, flag := range f {
		switch flag.Short {
		case "p":
			if s, ok := flag.Value.(string); ok {
				pattern = s
			}
		case "r":
			if s, ok := flag.Value.(string); ok {
				replacement = s
			}
		}
	}
	if pattern == "" {
		return "", fmt.Errorf("--pattern is required")
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		return "", fmt.Errorf("invalid regex pattern: %w", err)
	}
	return re.ReplaceAllString(string(data), replacement), nil
}

func (p RegexReplace) Flags() []Flag {
	return []Flag{
		{Name: "pattern", Short: "p", Desc: "Regular expression pattern", Value: "", Type: FlagString, Required: true},
		{Name: "replacement", Short: "r", Desc: "Replacement string", Value: "", Type: FlagString},
	}
}
func (p RegexReplace) Title() string       { return fmt.Sprintf("Regex Replace (%s)", p.Name()) }
func (p RegexReplace) Description() string { return "Replace regex matches" }
func (p RegexReplace) FilterValue() string { return p.Title() }
