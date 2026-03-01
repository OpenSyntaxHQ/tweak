package tests

import (
	"regexp"
	"testing"

	"github.com/OpenSyntaxHQ/tweak/processors"
)

var ulidRegex = regexp.MustCompile(`^[0-9A-Z]{26}$`)

func TestULID_Format(t *testing.T) {
	for i := 0; i < 5; i++ {
		got, err := processors.ULID{}.Transform([]byte(""))
		if err != nil {
			t.Fatal(err)
		}
		if !ulidRegex.MatchString(got) {
			t.Errorf("ULID %q does not match expected format ^[0-9A-Z]{26}$", got)
		}
	}
}

func TestULID_Unique(t *testing.T) {
	seen := map[string]bool{}
	for i := 0; i < 10; i++ {
		got, _ := processors.ULID{}.Transform([]byte(""))
		if seen[got] {
			t.Errorf("duplicate ULID generated: %q", got)
		}
		seen[got] = true
	}
}
