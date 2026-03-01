package processors

import (
	"fmt"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type PadLeft struct{}

func (p PadLeft) Name() string    { return "pad-left" }
func (p PadLeft) Alias() []string { return []string{"lpad"} }
func (p PadLeft) Transform(data []byte, f ...Flag) (string, error) {
	width := uint(10)
	char := " "
	for _, flag := range f {
		switch flag.Short {
		case "w":
			if v, ok := flag.Value.(uint); ok {
				width = v
			}
		case "c":
			if s, ok := flag.Value.(string); ok && len(s) > 0 {
				char = s
			}
		}
	}
	input := string(data)
	for len(input) < int(width) {
		input = char + input
	}
	return input, nil
}
func (p PadLeft) Flags() []Flag {
	return []Flag{
		{Name: "width", Short: "w", Desc: "Target width", Value: uint(10), Type: FlagUint},
		{Name: "char", Short: "c", Desc: "Pad character", Value: " ", Type: FlagString},
	}
}
func (p PadLeft) Title() string {
	return fmt.Sprintf("%s (%s)", cases.Title(language.Und, cases.NoLower).String(p.Name()), p.Name())
}
func (p PadLeft) Description() string { return "Pad text on the left" }
func (p PadLeft) FilterValue() string { return p.Title() }

type PadRight struct{}

func (p PadRight) Name() string    { return "pad-right" }
func (p PadRight) Alias() []string { return []string{"rpad"} }
func (p PadRight) Transform(data []byte, f ...Flag) (string, error) {
	width := uint(10)
	char := " "
	for _, flag := range f {
		switch flag.Short {
		case "w":
			if v, ok := flag.Value.(uint); ok {
				width = v
			}
		case "c":
			if s, ok := flag.Value.(string); ok && len(s) > 0 {
				char = s
			}
		}
	}
	input := string(data)
	for len(input) < int(width) {
		input = input + char
	}
	return input, nil
}
func (p PadRight) Flags() []Flag {
	return []Flag{
		{Name: "width", Short: "w", Desc: "Target width", Value: uint(10), Type: FlagUint},
		{Name: "char", Short: "c", Desc: "Pad character", Value: " ", Type: FlagString},
	}
}
func (p PadRight) Title() string {
	return fmt.Sprintf("%s (%s)", cases.Title(language.Und, cases.NoLower).String(p.Name()), p.Name())
}
func (p PadRight) Description() string { return "Pad text on the right" }
func (p PadRight) FilterValue() string { return p.Title() }


