package processors

import (
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Zeropad struct{}

func (p Zeropad) Name() string    { return "zeropad" }
func (p Zeropad) Alias() []string { return nil }

func (p Zeropad) Transform(data []byte, f ...Flag) (string, error) {
	strIn := strings.TrimSpace(string(data))
	neg := ""
	i, err := strconv.ParseFloat(strIn, 64)
	if err != nil {
		return "", fmt.Errorf("number expected: '%s'", data)
	}
	if i < 0 {
		neg = "-"
		data = data[1:]
	}

	var n int
	pre := ""
	for _, flag := range f {
		switch flag.Short {
		case "n":
			if x, ok := flag.Value.(uint); ok {
				n = int(x)
			}
		case "p":
			if x, ok := flag.Value.(string); ok {
				pre = x
			}
		}
	}
	return fmt.Sprintf("%s%s%s%s", pre, neg, strings.Repeat("0", n), data), nil
}

func (p Zeropad) Flags() []Flag {
	return []Flag{
		{Name: "number-of-zeros", Short: "n", Desc: "Number of zeros to pad", Value: uint(5), Type: FlagUint},
		{Name: "prefix", Short: "p", Desc: "Prefix before the number", Value: "", Type: FlagString},
	}
}

func (p Zeropad) Title() string {
	return fmt.Sprintf("%s (%s)", cases.Title(language.Und, cases.NoLower).String(p.Name()), p.Name())
}
func (p Zeropad) Description() string { return "Pad a number with zeros" }
func (p Zeropad) FilterValue() string { return p.Title() }
