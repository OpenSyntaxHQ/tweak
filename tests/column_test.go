package tests

import (
	"testing"

	"github.com/OpenSyntaxHQ/tweak/processors"
)

func TestColumn_Whitespace(t *testing.T) {
	input := "one two three"
	flags := []processors.Flag{
		{Name: "field", Short: "f", Value: uint(2), Type: processors.FlagUint},
	}
	got, err := processors.Column{}.Transform([]byte(input), flags...)
	if err != nil {
		t.Fatal(err)
	}
	if got != "two" {
		t.Errorf("Column field 2 = %q, want %q", got, "two")
	}
}

func TestColumn_CSVDelimiter(t *testing.T) {
	input := "a,b,c,d"
	flags := []processors.Flag{
		{Name: "field", Short: "f", Value: uint(3), Type: processors.FlagUint},
		{Name: "delimiter", Short: "d", Value: ",", Type: processors.FlagString},
	}
	got, err := processors.Column{}.Transform([]byte(input), flags...)
	if err != nil {
		t.Fatal(err)
	}
	if got != "c" {
		t.Errorf("Column(csv, field=3) = %q, want %q", got, "c")
	}
}
