package processors

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/argon2"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Argon2Hash struct{}

func (p Argon2Hash) Name() string    { return "argon2" }
func (p Argon2Hash) Alias() []string { return []string{"argon2id"} }
func (p Argon2Hash) Transform(data []byte, f ...Flag) (string, error) {
	var saltStr string
	timeParam := uint(3)
	memParam := uint(65536)
	for _, flag := range f {
		switch flag.Short {
		case "s":
			if s, ok := flag.Value.(string); ok {
				saltStr = s
			}
		case "t":
			if t, ok := flag.Value.(uint); ok && t > 0 {
				timeParam = t
			}
		case "m":
			if m, ok := flag.Value.(uint); ok && m > 0 {
				memParam = m
			}
		}
	}
	var salt []byte
	if saltStr != "" {
		salt = []byte(saltStr)
	} else {
		salt = make([]byte, 16)
		if _, err := rand.Read(salt); err != nil {
			return "", err
		}
	}
	hash := argon2.IDKey(data, salt, uint32(timeParam), uint32(memParam), 4, 32)
	return fmt.Sprintf("$argon2id$v=19$m=%d,t=%d,p=4$%s$%s",
		memParam, timeParam,
		base64.RawStdEncoding.EncodeToString(salt),
		base64.RawStdEncoding.EncodeToString(hash)), nil
}
func (p Argon2Hash) Flags() []Flag {
	return []Flag{
		{Name: "salt", Short: "s", Desc: "Salt string (random if omitted)", Value: "", Type: FlagString},
		{Name: "time", Short: "t", Desc: "Time parameter (iterations)", Value: uint(3), Type: FlagUint},
		{Name: "memory", Short: "m", Desc: "Memory parameter in KiB", Value: uint(65536), Type: FlagUint},
	}
}
func (p Argon2Hash) Title() string {
	return fmt.Sprintf("%s (%s)", cases.Title(language.Und, cases.NoLower).String(p.Name()), p.Name())
}
func (p Argon2Hash) Description() string { return "Argon2id password hash" }
func (p Argon2Hash) FilterValue() string { return p.Title() }
