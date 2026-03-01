package tests

import (
	"strings"
	"testing"

	"github.com/OpenSyntaxHQ/tweak/processors"
)

func TestChecksumVerify_Match(t *testing.T) {
	tests := []struct {
		algo string
		hash string
		flag string
	}{
		{"sha256", "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824", "sha256"},
		{"md5", "5d41402abc4b2a76b9719d911017c592", "md5"},
		{"sha1", "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d", "sha1"},
	}
	for _, tt := range tests {
		t.Run(tt.algo, func(t *testing.T) {
			flags := []processors.Flag{
				{Name: "hash", Short: "x", Value: tt.hash, Type: processors.FlagString},
				{Name: "algo", Short: "g", Value: tt.flag, Type: processors.FlagString},
			}
			got, err := processors.ChecksumVerify{}.Transform([]byte("hello"), flags...)
			if err != nil {
				t.Fatal(err)
			}
			if !strings.Contains(got, "Match") {
				t.Errorf("expected Match, got: %q", got)
			}
		})
	}
}

func TestChecksumVerify_Mismatch(t *testing.T) {
	flags := []processors.Flag{
		{Name: "hash", Short: "x", Value: "0000000000000000000000000000000000000000000000000000000000000000", Type: processors.FlagString},
		{Name: "algo", Short: "g", Value: "sha256", Type: processors.FlagString},
	}
	got, err := processors.ChecksumVerify{}.Transform([]byte("hello"), flags...)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got, "Mismatch") {
		t.Errorf("expected Mismatch, got: %q", got)
	}
}
