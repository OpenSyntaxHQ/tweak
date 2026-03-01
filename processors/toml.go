package processors

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/BurntSushi/toml"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type TOMLToJSON struct{}

func (p TOMLToJSON) Name() string    { return "toml-json" }
func (p TOMLToJSON) Alias() []string { return []string{"toml-to-json"} }
func (p TOMLToJSON) Transform(data []byte, f ...Flag) (string, error) {
	var obj any
	if err := toml.Unmarshal(data, &obj); err != nil {
		return "", fmt.Errorf("invalid TOML: %w", err)
	}
	indent := false
	for _, flag := range f {
		if flag.Short == "i" {
			if b, ok := flag.Value.(bool); ok {
				indent = b
			}
		}
	}
	var out []byte
	var err error
	if indent {
		out, err = json.MarshalIndent(obj, "", "  ")
	} else {
		out, err = json.Marshal(obj)
	}
	if err != nil {
		return "", err
	}
	return string(out), nil
}
func (p TOMLToJSON) Flags() []Flag {
	return []Flag{{Name: "indent", Short: "i", Desc: "Pretty-print JSON output", Value: false, Type: FlagBool}}
}
func (p TOMLToJSON) Title() string {
	return fmt.Sprintf("%s (%s)", cases.Title(language.Und, cases.NoLower).String(p.Name()), p.Name())
}
func (p TOMLToJSON) Description() string { return "Convert TOML to JSON" }
func (p TOMLToJSON) FilterValue() string { return p.Title() }

type JSONToTOML struct{}

func (p JSONToTOML) Name() string    { return "json-toml" }
func (p JSONToTOML) Alias() []string { return []string{"json-to-toml"} }
func (p JSONToTOML) Transform(data []byte, _ ...Flag) (string, error) {
	var obj any
	if err := json.Unmarshal(data, &obj); err != nil {
		return "", fmt.Errorf("invalid JSON: %w", err)
	}
	var buf bytes.Buffer
	enc := toml.NewEncoder(&buf)
	if err := enc.Encode(obj); err != nil {
		return "", err
	}
	return buf.String(), nil
}
func (p JSONToTOML) Flags() []Flag       { return nil }
func (p JSONToTOML) Title() string {
	return fmt.Sprintf("%s (%s)", cases.Title(language.Und, cases.NoLower).String(p.Name()), p.Name())
}
func (p JSONToTOML) Description() string { return "Convert JSON to TOML" }
func (p JSONToTOML) FilterValue() string { return p.Title() }
