package processors

import (
	"fmt"
	"strings"
)

type ROT13 struct{}

func (p ROT13) Name() string    { return "rot13" }
func (p ROT13) Alias() []string { return []string{"rot13-encode", "rot13-enc"} }

func (p ROT13) Transform(data []byte, _ ...Flag) (string, error) {
	return strings.Map(rot13Rune, string(data)), nil
}

func (p ROT13) Flags() []Flag       { return nil }
func (p ROT13) Title() string       { return fmt.Sprintf("ROT13 Letter Substitution (%s)", p.Name()) }
func (p ROT13) Description() string { return "Cipher/Decipher your text with ROT13 letter substitution" }
func (p ROT13) FilterValue() string { return p.Title() }

func rot13Rune(r rune) rune {
	switch {
	case r >= 'a' && r <= 'z':
		if r >= 'm' {
			return r - 13
		}
		return r + 13
	case r >= 'A' && r <= 'Z':
		if r >= 'M' {
			return r - 13
		}
		return r + 13
	default:
		return r
	}
}
