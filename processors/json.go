package processors

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/vmihailenco/msgpack/v5"
)

func unmarshalJSON(data []byte) (any, error) {
	var obj map[string]any
	if err := json.Unmarshal(data, &obj); err == nil {
		return obj, nil
	}
	var arr []any
	if err := json.Unmarshal(data, &arr); err == nil {
		return arr, nil
	}
	return nil, fmt.Errorf("invalid JSON")
}

type FormatJSON struct{}

func (p FormatJSON) Name() string    { return "json" }
func (p FormatJSON) Alias() []string { return nil }

func (p FormatJSON) Transform(data []byte, f ...Flag) (string, error) {
	obj, err := unmarshalJSON(data)
	if err != nil {
		return "", err
	}
	var indent bool
	for _, flag := range f {
		if flag.Short == "i" {
			if b, ok := flag.Value.(bool); ok {
				indent = b
			}
		}
	}
	var out []byte
	if indent {
		out, err = json.MarshalIndent(obj, "", "  ")
	} else {
		out, err = json.Marshal(obj)
	}
	return string(out), err
}

func (p FormatJSON) Flags() []Flag {
	return []Flag{{Name: "indent", Short: "i", Desc: "Indent the output (prettyprint)", Value: false, Type: FlagBool}}
}
func (p FormatJSON) Title() string       { return fmt.Sprintf("Format JSON (%s)", p.Name()) }
func (p FormatJSON) Description() string { return "Format your text as JSON" }
func (p FormatJSON) FilterValue() string { return p.Title() }

type JSONEscape struct{}

func (p JSONEscape) Name() string    { return "json-escape" }
func (p JSONEscape) Alias() []string { return []string{"json-esc"} }
func (p JSONEscape) Transform(data []byte, _ ...Flag) (string, error) {
	obj, err := unmarshalJSON(data)
	if err != nil {
		return "", err
	}
	out, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	output := strconv.Quote(string(out))
	output = strings.TrimPrefix(output, `"`)
	output = strings.TrimSuffix(output, `"`)
	return output, nil
}
func (p JSONEscape) Flags() []Flag       { return nil }
func (p JSONEscape) Title() string       { return fmt.Sprintf("JSON Escape (%s)", p.Name()) }
func (p JSONEscape) Description() string { return "JSON Escape" }
func (p JSONEscape) FilterValue() string { return p.Title() }

type JSONUnescape struct{}

func (p JSONUnescape) Name() string    { return "json-unescape" }
func (p JSONUnescape) Alias() []string { return []string{"json-unesc"} }
func (p JSONUnescape) Transform(data []byte, f ...Flag) (string, error) {
	s, err := strconv.Unquote(`"` + strings.TrimSpace(string(data)) + `"`)
	if err != nil {
		return "", err
	}
	obj, err := unmarshalJSON([]byte(s))
	if err != nil {
		return "", err
	}
	var indent bool
	for _, flag := range f {
		if flag.Short == "i" {
			if b, ok := flag.Value.(bool); ok {
				indent = b
			}
		}
	}
	var out []byte
	if indent {
		out, err = json.MarshalIndent(obj, "", "  ")
	} else {
		out, err = json.Marshal(obj)
	}
	return string(out), err
}
func (p JSONUnescape) Flags() []Flag {
	return []Flag{{Name: "indent", Short: "i", Desc: "Indent the output (prettyprint)", Value: false, Type: FlagBool}}
}
func (p JSONUnescape) Title() string       { return fmt.Sprintf("JSON Unescape (%s)", p.Name()) }
func (p JSONUnescape) Description() string { return "JSON Unescape" }
func (p JSONUnescape) FilterValue() string { return p.Title() }

type JSONToYAML struct{}

func (p JSONToYAML) Name() string    { return "json-yaml" }
func (p JSONToYAML) Alias() []string { return []string{"json-yml"} }
func (p JSONToYAML) Transform(data []byte, _ ...Flag) (string, error) {
	y, err := yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(y), nil
}
func (p JSONToYAML) Flags() []Flag       { return nil }
func (p JSONToYAML) Title() string       { return fmt.Sprintf("JSON to YAML (%s)", p.Name()) }
func (p JSONToYAML) Description() string { return "Convert JSON to YAML text" }
func (p JSONToYAML) FilterValue() string { return p.Title() }

type YAMLToJSON struct{}

func (p YAMLToJSON) Name() string    { return "yaml-json" }
func (p YAMLToJSON) Alias() []string { return []string{"yml-json"} }
func (p YAMLToJSON) Transform(data []byte, f ...Flag) (string, error) {
	y, err := yaml.YAMLToJSON(data)
	if err != nil {
		return "", err
	}
	j := FormatJSON{}
	return j.Transform(y, f...)
}
func (p YAMLToJSON) Flags() []Flag {
	return []Flag{{Name: "indent", Short: "i", Desc: "Indent the output (prettyprint)", Value: false, Type: FlagBool}}
}
func (p YAMLToJSON) Title() string       { return fmt.Sprintf("YAML To JSON (%s)", p.Name()) }
func (p YAMLToJSON) Description() string { return "Convert YAML to JSON text" }
func (p YAMLToJSON) FilterValue() string { return p.Title() }

type JSONToMSGPACK struct{}

func (p JSONToMSGPACK) Name() string    { return "json-msgpack" }
func (p JSONToMSGPACK) Alias() []string { return nil }
func (p JSONToMSGPACK) Transform(data []byte, _ ...Flag) (string, error) {
	var raw any
	if err := json.Unmarshal(data, &raw); err != nil {
		return "", err
	}
	m, err := msgpack.Marshal(raw)
	if err != nil {
		return "", err
	}
	return string(m), nil
}
func (p JSONToMSGPACK) Flags() []Flag       { return nil }
func (p JSONToMSGPACK) Title() string       { return fmt.Sprintf("JSON To MSGPACK (%s)", p.Name()) }
func (p JSONToMSGPACK) Description() string { return "Convert JSON to MSGPACK" }
func (p JSONToMSGPACK) FilterValue() string { return p.Title() }

type MSGPACKToJSON struct{}

func (p MSGPACKToJSON) Name() string    { return "msgpack-json" }
func (p MSGPACKToJSON) Alias() []string { return nil }
func (p MSGPACKToJSON) Transform(data []byte, _ ...Flag) (string, error) {
	var raw any
	if err := msgpack.Unmarshal(data, &raw); err != nil {
		return "", err
	}
	m, err := json.Marshal(raw)
	if err != nil {
		return "", err
	}
	return string(m), nil
}
func (p MSGPACKToJSON) Flags() []Flag       { return nil }
func (p MSGPACKToJSON) Title() string       { return fmt.Sprintf("MSGPACK to JSON (%s)", p.Name()) }
func (p MSGPACKToJSON) Description() string { return "Convert MSGPACK to JSON" }
func (p MSGPACKToJSON) FilterValue() string { return p.Title() }

type JSONMinify struct{}

func (p JSONMinify) Name() string    { return "json-minify" }
func (p JSONMinify) Alias() []string { return []string{"json-min"} }
func (p JSONMinify) Transform(data []byte, _ ...Flag) (string, error) {
	var buf bytes.Buffer
	if err := json.Compact(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}
func (p JSONMinify) Flags() []Flag       { return nil }
func (p JSONMinify) Title() string       { return fmt.Sprintf("JSON Minify (%s)", p.Name()) }
func (p JSONMinify) Description() string { return "Minify JSON (remove whitespace)" }
func (p JSONMinify) FilterValue() string { return p.Title() }
