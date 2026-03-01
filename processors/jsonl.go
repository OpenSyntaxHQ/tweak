package processors

import (
	"encoding/json"
	"fmt"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type JSONLToJSON struct{}

func (p JSONLToJSON) Name() string    { return "jsonl-json" }
func (p JSONLToJSON) Alias() []string { return []string{"ndjson-json", "jsonl-to-json"} }
func (p JSONLToJSON) Transform(data []byte, _ ...Flag) (string, error) {
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	result := make([]any, 0, len(lines))
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		var obj any
		if err := json.Unmarshal([]byte(line), &obj); err != nil {
			return "", fmt.Errorf("line %d: invalid JSON: %w", i+1, err)
		}
		result = append(result, obj)
	}
	out, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return "", err
	}
	return string(out), nil
}
func (p JSONLToJSON) Flags() []Flag       { return nil }
func (p JSONLToJSON) Title() string {
	return fmt.Sprintf("%s (%s)", cases.Title(language.Und, cases.NoLower).String(p.Name()), p.Name())
}
func (p JSONLToJSON) Description() string { return "Convert JSONL/NDJSON to JSON array" }
func (p JSONLToJSON) FilterValue() string { return p.Title() }
