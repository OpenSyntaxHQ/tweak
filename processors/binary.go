package processors

import (
	"fmt"
	"strings"
)

type BinaryEncode struct{}

func (p BinaryEncode) Name() string    { return "binary-encode" }
func (p BinaryEncode) Alias() []string { return []string{"bin-enc"} }

func (p BinaryEncode) Transform(data []byte, _ ...Flag) (string, error) {
	parts := make([]string, len(data))
	for i, b := range data {
		parts[i] = fmt.Sprintf("%08b", b)
	}
	return strings.Join(parts, " "), nil
}

func (p BinaryEncode) Flags() []Flag       { return nil }
func (p BinaryEncode) Title() string       { return fmt.Sprintf("Binary Encode (%s)", p.Name()) }
func (p BinaryEncode) Description() string { return "Encode your text to binary representation" }
func (p BinaryEncode) FilterValue() string { return p.Title() }

type BinaryDecode struct{}

func (p BinaryDecode) Name() string    { return "binary-decode" }
func (p BinaryDecode) Alias() []string { return []string{"bin-dec"} }

func (p BinaryDecode) Transform(data []byte, _ ...Flag) (string, error) {
	parts := strings.Fields(string(data))
	result := make([]byte, 0, len(parts))
	for _, part := range parts {
		var b byte
		for _, bit := range part {
			b = b<<1 | byte(bit-'0')
		}
		result = append(result, b)
	}
	return string(result), nil
}

func (p BinaryDecode) Flags() []Flag       { return nil }
func (p BinaryDecode) Title() string       { return fmt.Sprintf("Binary Decode (%s)", p.Name()) }
func (p BinaryDecode) Description() string { return "Decode binary representation to text" }
func (p BinaryDecode) FilterValue() string { return p.Title() }

type CaesarEncode struct{}

func (p CaesarEncode) Name() string    { return "caesar-encode" }
func (p CaesarEncode) Alias() []string { return []string{"caesar-enc"} }

func (p CaesarEncode) Transform(data []byte, f ...Flag) (string, error) {
	shift := 3
	for _, flag := range f {
		if flag.Short == "s" {
			if s, ok := flag.Value.(int); ok {
				shift = s
			}
		}
	}
	return caesarShift(string(data), shift), nil
}

func (p CaesarEncode) Flags() []Flag {
	return []Flag{{Name: "shift", Short: "s", Desc: "Number of positions to shift", Value: 3, Type: FlagInt}}
}
func (p CaesarEncode) Title() string       { return fmt.Sprintf("Caesar Cipher Encode (%s)", p.Name()) }
func (p CaesarEncode) Description() string { return "Encode text with Caesar cipher" }
func (p CaesarEncode) FilterValue() string { return p.Title() }

type CaesarDecode struct{}

func (p CaesarDecode) Name() string    { return "caesar-decode" }
func (p CaesarDecode) Alias() []string { return []string{"caesar-dec"} }

func (p CaesarDecode) Transform(data []byte, f ...Flag) (string, error) {
	shift := 3
	for _, flag := range f {
		if flag.Short == "s" {
			if s, ok := flag.Value.(int); ok {
				shift = s
			}
		}
	}
	return caesarShift(string(data), -shift), nil
}

func (p CaesarDecode) Flags() []Flag {
	return []Flag{{Name: "shift", Short: "s", Desc: "Number of positions to shift", Value: 3, Type: FlagInt}}
}
func (p CaesarDecode) Title() string       { return fmt.Sprintf("Caesar Cipher Decode (%s)", p.Name()) }
func (p CaesarDecode) Description() string { return "Decode text from Caesar cipher" }
func (p CaesarDecode) FilterValue() string { return p.Title() }

func caesarShift(text string, shift int) string {
	shift = ((shift % 26) + 26) % 26
	return strings.Map(func(r rune) rune {
		switch {
		case r >= 'a' && r <= 'z':
			return rune((int(r-'a')+shift)%26) + 'a'
		case r >= 'A' && r <= 'Z':
			return rune((int(r-'A')+shift)%26) + 'A'
		default:
			return r
		}
	}, text)
}
