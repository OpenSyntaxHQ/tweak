package processors

import (
	"encoding/json"
	"fmt"

	"github.com/clbanning/mxj/v2"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func init() {
	mxj.PrependAttrWithHyphen(false)
}

type XMLToJSON struct{}

func (p XMLToJSON) Name() string    { return "xml-json" }
func (p XMLToJSON) Alias() []string { return []string{"xml-to-json"} }
func (p XMLToJSON) Transform(data []byte, f ...Flag) (string, error) {
	mv, err := mxj.NewMapXml(data, true)
	if err != nil {
		return "", fmt.Errorf("invalid XML: %w", err)
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
	if indent {
		out, err = json.MarshalIndent(map[string]any(mv), "", "  ")
	} else {
		out, err = json.Marshal(map[string]any(mv))
	}
	if err != nil {
		return "", err
	}
	return string(out), nil
}
func (p XMLToJSON) Flags() []Flag {
	return []Flag{{Name: "indent", Short: "i", Desc: "Pretty-print JSON output", Value: false, Type: FlagBool}}
}
func (p XMLToJSON) Title() string {
	return fmt.Sprintf("%s (%s)", cases.Title(language.Und, cases.NoLower).String(p.Name()), p.Name())
}
func (p XMLToJSON) Description() string { return "Convert XML to JSON" }
func (p XMLToJSON) FilterValue() string { return p.Title() }

type JSONToXML struct{}

func (p JSONToXML) Name() string    { return "json-xml" }
func (p JSONToXML) Alias() []string { return []string{"json-to-xml"} }
func (p JSONToXML) Transform(data []byte, f ...Flag) (string, error) {
	indent := false
	root := "root"
	for _, flag := range f {
		switch flag.Short {
		case "i":
			if b, ok := flag.Value.(bool); ok {
				indent = b
			}
		case "r":
			if s, ok := flag.Value.(string); ok && s != "" {
				root = s
			}
		}
	}
	var obj map[string]any
	if err := json.Unmarshal(data, &obj); err != nil {
		return "", fmt.Errorf("invalid JSON: %w", err)
	}
	mv := mxj.Map(obj)
	var out []byte
	var err error
	if indent {
		out, err = mv.XmlIndent("", "  ", root)
	} else {
		out, err = mv.Xml(root)
	}
	if err != nil {
		return "", err
	}
	return string(out), nil
}
func (p JSONToXML) Flags() []Flag {
	return []Flag{
		{Name: "indent", Short: "i", Desc: "Pretty-print XML output", Value: false, Type: FlagBool},
		{Name: "root", Short: "r", Desc: "Root element name", Value: "root", Type: FlagString},
	}
}
func (p JSONToXML) Title() string {
	return fmt.Sprintf("%s (%s)", cases.Title(language.Und, cases.NoLower).String(p.Name()), p.Name())
}
func (p JSONToXML) Description() string { return "Convert JSON to XML" }
func (p JSONToXML) FilterValue() string { return p.Title() }
