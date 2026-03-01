package tests

import (
	"regexp"
	"strings"
	"testing"

	"github.com/OpenSyntaxHQ/tweak/processors"
)

func TestEpoch_UnixToDate(t *testing.T) {
	got, err := processors.Epoch{}.Transform([]byte("0"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got, "1970") {
		t.Errorf("epoch 0 should contain 1970, got: %q", got)
	}
}

func TestEpoch_DateToUnix(t *testing.T) {
	got, err := processors.Epoch{}.Transform([]byte("2023-01-01 00:00:00"))
	if err != nil {
		t.Fatal(err)
	}
	matched, _ := regexp.MatchString(`\d+`, got)
	if !matched {
		t.Errorf("expected unix timestamp in output, got: %q", got)
	}
}

func TestNow_Format(t *testing.T) {
	got, err := processors.Now{}.Transform([]byte(""))
	if err != nil {
		t.Fatal(err)
	}
	if len(got) == 0 {
		t.Error("now should produce non-empty output")
	}
}
