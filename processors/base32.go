package processors

import (
	"encoding/base32"
	"fmt"
)

type Base32Encoding struct{}

func (p Base32Encoding) Name() string    { return "base32-encode" }
func (p Base32Encoding) Alias() []string { return []string{"b32-enc", "b32-encode"} }

func (p Base32Encoding) Transform(data []byte, _ ...Flag) (string, error) {
	return base32.StdEncoding.EncodeToString(data), nil
}

func (p Base32Encoding) Flags() []Flag       { return nil }
func (p Base32Encoding) Title() string       { return fmt.Sprintf("Base32 Encoding (%s)", p.Name()) }
func (p Base32Encoding) Description() string { return "Encode your text to Base32" }
func (p Base32Encoding) FilterValue() string { return p.Title() }

type Base32Decode struct{}

func (p Base32Decode) Name() string    { return "base32-decode" }
func (p Base32Decode) Alias() []string { return []string{"b32-dec", "b32-decode"} }

func (p Base32Decode) Transform(data []byte, _ ...Flag) (string, error) {
	decoded, err := base32.StdEncoding.DecodeString(string(data))
	return string(decoded), err
}

func (p Base32Decode) Flags() []Flag       { return nil }
func (p Base32Decode) Title() string       { return fmt.Sprintf("Base32 Decode (%s)", p.Name()) }
func (p Base32Decode) Description() string { return "Decode your Base32 text" }
func (p Base32Decode) FilterValue() string { return p.Title() }
