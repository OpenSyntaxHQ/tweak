package processors

import (
	"fmt"
	"regexp"
	"strings"
)

type RemoveNewLines struct{}

func (p RemoveNewLines) Name() string    { return "remove-newlines" }
func (p RemoveNewLines) Alias() []string { return []string{"remove-new-lines", "trim-newlines", "trim-new-lines"} }
func (p RemoveNewLines) Transform(data []byte, f ...Flag) (string, error) {
	separator := " "
	for _, flag := range f {
		if flag.Short == "s" {
			if x, ok := flag.Value.(string); ok {
				separator = x
			}
		}
	}
	return regexp.MustCompile(`[\r\n]+`).ReplaceAllString(strings.TrimSpace(string(data)), separator), nil
}
func (p RemoveNewLines) Flags() []Flag {
	return []Flag{{Name: "separator", Short: "s", Desc: "Separator to replace newlines", Value: "", Type: FlagString}}
}
func (p RemoveNewLines) Title() string       { return fmt.Sprintf("Remove all new lines (%s)", p.Name()) }
func (p RemoveNewLines) Description() string { return "Remove all new lines" }
func (p RemoveNewLines) FilterValue() string { return p.Title() }

type RemoveSpaces struct{}

func (p RemoveSpaces) Name() string    { return "remove-spaces" }
func (p RemoveSpaces) Alias() []string { return []string{"remove-space", "trim-spaces", "trim-space"} }
func (p RemoveSpaces) Transform(data []byte, f ...Flag) (string, error) {
	separator := ""
	for _, flag := range f {
		if flag.Short == "s" {
			if x, ok := flag.Value.(string); ok {
				separator = x
			}
		}
	}
	return regexp.MustCompile(`[\s\r\n]+`).ReplaceAllString(strings.TrimSpace(string(data)), separator), nil
}
func (p RemoveSpaces) Flags() []Flag {
	return []Flag{{Name: "separator", Short: "s", Desc: "Separator to replace spaces", Value: "", Type: FlagString}}
}
func (p RemoveSpaces) Title() string       { return fmt.Sprintf("Remove all spaces + new lines (%s)", p.Name()) }
func (p RemoveSpaces) Description() string { return "Remove all spaces + new lines" }
func (p RemoveSpaces) FilterValue() string { return p.Title() }
