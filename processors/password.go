package processors

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const (
	pwLetters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	pwDigits  = "0123456789"
	pwSymbols = "!@#$%^&*()-_=+[]{}|;:,.<>?"
)

type PasswordGen struct{}

func (p PasswordGen) Name() string      { return "password" }
func (p PasswordGen) Alias() []string   { return []string{"pass-gen", "gen-pass"} }
func (p PasswordGen) IsGenerator() bool { return true }
func (p PasswordGen) Transform(_ []byte, f ...Flag) (string, error) {
	length := uint(20)
	noSymbols := false
	for _, flag := range f {
		switch flag.Short {
		case "l":
			if l, ok := flag.Value.(uint); ok && l > 0 {
				length = l
			}
		case "n":
			if b, ok := flag.Value.(bool); ok {
				noSymbols = b
			}
		}
	}
	charset := pwLetters + pwDigits
	if !noSymbols {
		charset += pwSymbols
	}
	result := make([]byte, length)
	for i := range result {
		idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		result[i] = charset[idx.Int64()]
	}
	return string(result), nil
}
func (p PasswordGen) Flags() []Flag {
	return []Flag{
		{Name: "length", Short: "l", Desc: "Password length", Value: uint(20), Type: FlagUint},
		{Name: "no-symbols", Short: "n", Desc: "Exclude symbols", Value: false, Type: FlagBool},
	}
}
func (p PasswordGen) Title() string {
	return fmt.Sprintf("%s (%s)", cases.Title(language.Und, cases.NoLower).String(p.Name()), p.Name())
}
func (p PasswordGen) Description() string { return "Generate a secure random password" }
func (p PasswordGen) FilterValue() string { return p.Title() }
