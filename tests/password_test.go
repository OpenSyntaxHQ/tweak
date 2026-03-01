package tests

import (
	"regexp"
	"testing"

	"github.com/OpenSyntaxHQ/tweak/processors"
)

func TestPasswordGen_Default(t *testing.T) {
	got, err := processors.PasswordGen{}.Transform([]byte(""))
	if err != nil {
		t.Fatal(err)
	}
	if len(got) < 12 {
		t.Errorf("password too short: %q (len=%d)", got, len(got))
	}
}

func TestPasswordGen_Length(t *testing.T) {
	flags := []processors.Flag{
		{Name: "length", Short: "l", Value: uint(24), Type: processors.FlagUint},
	}
	got, err := processors.PasswordGen{}.Transform([]byte(""), flags...)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 24 {
		t.Errorf("expected length 24, got %d: %q", len(got), got)
	}
}

func TestPasswordGen_NoSymbols(t *testing.T) {
	flags := []processors.Flag{
		{Name: "length", Short: "l", Value: uint(32), Type: processors.FlagUint},
		{Name: "no-symbols", Short: "n", Value: true, Type: processors.FlagBool},
	}
	got, err := processors.PasswordGen{}.Transform([]byte(""), flags...)
	if err != nil {
		t.Fatal(err)
	}
	matched, _ := regexp.MatchString(`^[a-zA-Z0-9]+$`, got)
	if !matched {
		t.Errorf("--no-symbols password contains non-alphanumeric: %q", got)
	}
}
