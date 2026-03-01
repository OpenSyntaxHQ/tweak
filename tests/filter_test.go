package tests

import (
	"strings"
	"testing"

	"github.com/OpenSyntaxHQ/tweak/processors"
)

func TestGrep_Match(t *testing.T) {
	input := "apple\nbanana\napricot\ncherry"
	flags := []processors.Flag{
		{Name: "pattern", Short: "p", Value: "^a", Type: processors.FlagString},
	}
	got, err := processors.Grep{}.Transform([]byte(input), flags...)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got, "apple") {
		t.Errorf("expected 'apple', got: %q", got)
	}
	if strings.Contains(got, "banana") {
		t.Errorf("'banana' should be filtered out, got: %q", got)
	}
}

func TestGrep_Invert(t *testing.T) {
	input := "apple\nbanana\ncherry"
	flags := []processors.Flag{
		{Name: "pattern", Short: "p", Value: "banana", Type: processors.FlagString},
		{Name: "invert", Short: "v", Value: true, Type: processors.FlagBool},
	}
	got, err := processors.Grep{}.Transform([]byte(input), flags...)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(got, "banana") {
		t.Errorf("inverted grep should exclude 'banana', got: %q", got)
	}
	if !strings.Contains(got, "apple") {
		t.Errorf("inverted grep should include 'apple', got: %q", got)
	}
}

func TestGrep_InvalidRegex(t *testing.T) {
	flags := []processors.Flag{
		{Name: "pattern", Short: "p", Value: "[invalid", Type: processors.FlagString},
	}
	_, err := processors.Grep{}.Transform([]byte("test"), flags...)
	if err == nil {
		t.Error("expected error for invalid regex")
	}
}
