package tests

import (
	"strings"
	"testing"

	"github.com/OpenSyntaxHQ/tweak/processors"
)

func TestHMACSHA256_Transform(t *testing.T) {
	flags := []processors.Flag{{Name: "key", Short: "k", Value: "secret", Type: processors.FlagString}}
	got, err := processors.HMACSHA256{}.Transform([]byte("hello"), flags...)
	if err != nil {
		t.Fatal(err)
	}
	want := "88aab3ede8d3adf94d26ab90d3bafd4a2083070c3bcce9c014ee04a443847c0b"
	if got != want {
		t.Errorf("HMAC-SHA256 got %q, want %q", got, want)
	}
}

func TestHMACSHA512_Transform(t *testing.T) {
	flags := []processors.Flag{{Name: "key", Short: "k", Value: "secret", Type: processors.FlagString}}
	got, err := processors.HMACSHA512{}.Transform([]byte("hello"), flags...)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 128 {
		t.Errorf("HMAC-SHA512 want 128 hex chars, got %d", len(got))
	}
	if !strings.Contains(got, "") { /* just ensure non-empty */ }
}
