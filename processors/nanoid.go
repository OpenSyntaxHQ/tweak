package processors

import (
	"fmt"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type NanoID struct{}

func (p NanoID) Name() string      { return "nanoid" }
func (p NanoID) Alias() []string   { return []string{"gen-nanoid"} }
func (p NanoID) IsGenerator() bool { return true }
func (p NanoID) Transform(_ []byte, f ...Flag) (string, error) {
	length := uint(21)
	alphabet := ""
	for _, flag := range f {
		switch flag.Short {
		case "l":
			if l, ok := flag.Value.(uint); ok && l > 0 {
				length = l
			}
		case "a":
			if a, ok := flag.Value.(string); ok {
				alphabet = a
			}
		}
	}
	if alphabet != "" {
		return gonanoid.Generate(alphabet, int(length))
	}
	return gonanoid.New(int(length))
}
func (p NanoID) Flags() []Flag {
	return []Flag{
		{Name: "length", Short: "l", Desc: "ID length", Value: uint(21), Type: FlagUint},
		{Name: "alphabet", Short: "a", Desc: "Custom alphabet", Value: "", Type: FlagString},
	}
}
func (p NanoID) Title() string {
	return fmt.Sprintf("%s (%s)", cases.Title(language.Und, cases.NoLower).String(p.Name()), p.Name())
}
func (p NanoID) Description() string { return "Generate a NanoID" }
func (p NanoID) FilterValue() string { return p.Title() }
