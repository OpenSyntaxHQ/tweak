package tests

import (
	"strings"
	"testing"

	"github.com/OpenSyntaxHQ/tweak/processors"
)

func TestArgon2_Transform(t *testing.T) {
	flags := []processors.Flag{
		{Name: "salt", Short: "s", Value: "testsalt", Type: processors.FlagString},
		{Name: "time", Short: "t", Value: uint(1), Type: processors.FlagUint},
		{Name: "memory", Short: "m", Value: uint(65536), Type: processors.FlagUint},
	}
	got, err := processors.Argon2Hash{}.Transform([]byte("mypassword"), flags...)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got, "argon2id") {
		t.Errorf("expected argon2id hash, got: %q", got)
	}
}

func TestArgon2_Deterministic(t *testing.T) {
	flags := []processors.Flag{
		{Name: "salt", Short: "s", Value: "fixedsalt", Type: processors.FlagString},
		{Name: "time", Short: "t", Value: uint(1), Type: processors.FlagUint},
		{Name: "memory", Short: "m", Value: uint(65536), Type: processors.FlagUint},
	}
	got1, _ := processors.Argon2Hash{}.Transform([]byte("pass"), flags...)
	got2, _ := processors.Argon2Hash{}.Transform([]byte("pass"), flags...)
	if got1 != got2 {
		t.Error("argon2 with same params should be deterministic")
	}
}
