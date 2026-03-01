package processors

import (
	"fmt"
	"math/big"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type BaseConvert struct{}

func (p BaseConvert) Name() string    { return "base-convert" }
func (p BaseConvert) Alias() []string { return []string{"baseconv", "radix"} }
func (p BaseConvert) Transform(data []byte, f ...Flag) (string, error) {
	from := uint(10)
	to := uint(16)
	for _, flag := range f {
		switch flag.Short {
		case "f":
			if v, ok := flag.Value.(uint); ok && v >= 2 && v <= 36 {
				from = v
			}
		case "t":
			if v, ok := flag.Value.(uint); ok && v >= 2 && v <= 36 {
				to = v
			}
		}
	}
	input := strings.TrimSpace(string(data))
	input = strings.TrimPrefix(input, "0x")
	input = strings.TrimPrefix(input, "0X")
	input = strings.TrimPrefix(input, "0b")
	input = strings.TrimPrefix(input, "0B")
	input = strings.TrimPrefix(input, "0o")
	input = strings.TrimPrefix(input, "0O")
	n := new(big.Int)
	if _, ok := n.SetString(input, int(from)); !ok {
		return "", fmt.Errorf("invalid number '%s' for base %d", input, from)
	}
	result := n.Text(int(to))
	if to <= 16 {
		result = strings.ToLower(result)
	}
	return result, nil
}
func (p BaseConvert) Flags() []Flag {
	return []Flag{
		{Name: "from", Short: "f", Desc: "Source base (2-36)", Value: uint(10), Type: FlagUint},
		{Name: "to", Short: "t", Desc: "Target base (2-36)", Value: uint(16), Type: FlagUint},
	}
}
func (p BaseConvert) Title() string {
	return fmt.Sprintf("%s (%s)", cases.Title(language.Und, cases.NoLower).String(p.Name()), p.Name())
}
func (p BaseConvert) Description() string { return "Convert number between bases (2-36)" }
func (p BaseConvert) FilterValue() string { return p.Title() }
