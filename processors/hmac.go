package processors

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type HMACSHA256 struct{}

func (p HMACSHA256) Name() string    { return "hmac-sha256" }
func (p HMACSHA256) Alias() []string { return []string{"hmac256"} }
func (p HMACSHA256) Transform(data []byte, f ...Flag) (string, error) {
	var key string
	for _, flag := range f {
		if flag.Short == "k" {
			if s, ok := flag.Value.(string); ok {
				key = s
			}
		}
	}
	if key == "" {
		return "", fmt.Errorf("--key is required")
	}
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write(data)
	return hex.EncodeToString(mac.Sum(nil)), nil
}
func (p HMACSHA256) Flags() []Flag {
	return []Flag{{Name: "key", Short: "k", Desc: "HMAC secret key", Value: "", Type: FlagString, Required: true, Sensitive: true}}
}
func (p HMACSHA256) Title() string {
	return fmt.Sprintf("%s (%s)", cases.Title(language.Und, cases.NoLower).String(p.Name()), p.Name())
}
func (p HMACSHA256) Description() string { return "Generate HMAC-SHA256 keyed hash" }
func (p HMACSHA256) FilterValue() string { return p.Title() }

type HMACSHA512 struct{}

func (p HMACSHA512) Name() string    { return "hmac-sha512" }
func (p HMACSHA512) Alias() []string { return []string{"hmac512"} }
func (p HMACSHA512) Transform(data []byte, f ...Flag) (string, error) {
	var key string
	for _, flag := range f {
		if flag.Short == "k" {
			if s, ok := flag.Value.(string); ok {
				key = s
			}
		}
	}
	if key == "" {
		return "", fmt.Errorf("--key is required")
	}
	mac := hmac.New(sha512.New, []byte(key))
	mac.Write(data)
	return hex.EncodeToString(mac.Sum(nil)), nil
}
func (p HMACSHA512) Flags() []Flag {
	return []Flag{{Name: "key", Short: "k", Desc: "HMAC secret key", Value: "", Type: FlagString, Required: true, Sensitive: true}}
}
func (p HMACSHA512) Title() string {
	return fmt.Sprintf("%s (%s)", cases.Title(language.Und, cases.NoLower).String(p.Name()), p.Name())
}
func (p HMACSHA512) Description() string { return "Generate HMAC-SHA512 keyed hash" }
func (p HMACSHA512) FilterValue() string { return p.Title() }
