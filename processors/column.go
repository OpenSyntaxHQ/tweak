package processors

import (
	"fmt"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Column struct{}

func (p Column) Name() string    { return "column" }
func (p Column) Alias() []string { return []string{"col", "field"} }
func (p Column) Transform(data []byte, f ...Flag) (string, error) {
	fieldNum := uint(1)
	delimiter := ""
	for _, flag := range f {
		switch flag.Short {
		case "f":
			if v, ok := flag.Value.(uint); ok && v > 0 {
				fieldNum = v
			}
		case "d":
			if s, ok := flag.Value.(string); ok {
				delimiter = s
			}
		}
	}
	lines := strings.Split(string(data), "\n")
	var result []string
	for _, line := range lines {
		var fields []string
		if delimiter == "" {
			fields = strings.Fields(line)
		} else {
			fields = strings.Split(line, delimiter)
		}
		idx := int(fieldNum) - 1
		if idx < len(fields) {
			result = append(result, strings.TrimSpace(fields[idx]))
		}
	}
	return strings.Join(result, "\n"), nil
}
func (p Column) Flags() []Flag {
	return []Flag{
		{Name: "field", Short: "f", Desc: "Field number (1-based)", Value: uint(1), Type: FlagUint},
		{Name: "delimiter", Short: "d", Desc: "Field delimiter (whitespace if empty)", Value: "", Type: FlagString},
	}
}
func (p Column) Title() string {
	return fmt.Sprintf("%s (%s)", cases.Title(language.Und, cases.NoLower).String(p.Name()), p.Name())
}
func (p Column) Description() string { return "Extract column/field from text" }
func (p Column) FilterValue() string { return p.Title() }
