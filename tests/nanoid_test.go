package tests

import (
	"regexp"
	"testing"

	"github.com/OpenSyntaxHQ/tweak/processors"
)

func TestNanoID_DefaultLength(t *testing.T) {
	got, err := processors.NanoID{}.Transform([]byte(""))
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 21 {
		t.Errorf("default NanoID length should be 21, got %d: %q", len(got), got)
	}
}

func TestNanoID_CustomLength(t *testing.T) {
	flags := []processors.Flag{
		{Name: "length", Short: "l", Value: uint(10), Type: processors.FlagUint},
	}
	got, err := processors.NanoID{}.Transform([]byte(""), flags...)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 10 {
		t.Errorf("expected length 10, got %d: %q", len(got), got)
	}
}

func TestNanoID_Format(t *testing.T) {
	re := regexp.MustCompile(`^[A-Za-z0-9_-]+$`)
	for i := 0; i < 5; i++ {
		got, _ := processors.NanoID{}.Transform([]byte(""))
		if !re.MatchString(got) {
			t.Errorf("NanoID %q contains unexpected characters", got)
		}
	}
}
