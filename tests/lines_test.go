package tests

import (
	"strings"
	"testing"

	"github.com/OpenSyntaxHQ/tweak/processors"
)

func TestSortLines_Transform(t *testing.T) {
	assertTransform(t, processors.SortLines{}, "banana\napple\ncherry", nil, "apple\nbanana\ncherry", false)
}

func TestReverseLines_Transform(t *testing.T) {
	assertTransform(t, processors.ReverseLines{}, "a\nb\nc", nil, "c\nb\na", false)
}

func TestUniqueLines_Transform(t *testing.T) {
	got, err := (processors.UniqueLines{}).Transform([]byte("a\nb\na\nc\nb"))
	if err != nil {
		t.Fatal(err)
	}
	lines := strings.Split(got, "\n")
	seen := map[string]bool{}
	for _, l := range lines {
		if seen[l] {
			t.Errorf("duplicate line found: %q", l)
		}
		seen[l] = true
	}
}

func TestNumberLines_Transform(t *testing.T) {
	got, err := (processors.NumberLines{}).Transform([]byte("a\nb\nc"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got, "1") || !strings.Contains(got, "3") {
		t.Errorf("expected line numbers 1 and 3, got: %q", got)
	}
}

func TestCountLines_Transform(t *testing.T) {
	assertTransform(t, processors.CountLines{}, "a\nb\nc", nil, "3", false)
}

func TestShuffleLines_Transform(t *testing.T) {
	input := "apple\nbanana\ncherry\ndate\nelderberry"
	got, err := (processors.ShuffleLines{}).Transform([]byte(input))
	if err != nil {
		t.Fatal(err)
	}
	lines := strings.Split(got, "\n")
	if len(lines) != 5 {
		t.Errorf("shuffle should preserve line count, got %d lines", len(lines))
	}
}

func TestCountLines_EdgeCases(t *testing.T) {
	gotEmpty, err := (processors.CountLines{}).Transform([]byte(""))
	if err != nil {
		t.Fatal(err)
	}
	if gotEmpty != "0" {
		t.Fatalf("empty input lines = %q, want 0", gotEmpty)
	}

	gotTrailing, err := (processors.CountLines{}).Transform([]byte("a\nb\n"))
	if err != nil {
		t.Fatal(err)
	}
	if gotTrailing != "2" {
		t.Fatalf("trailing newline count = %q, want 2", gotTrailing)
	}
}

func TestSortLines_DropsTrailingEmptyLine(t *testing.T) {
	got, err := (processors.SortLines{}).Transform([]byte("b\na\n"))
	if err != nil {
		t.Fatal(err)
	}
	if got != "a\nb" {
		t.Fatalf("unexpected sorted output: %q", got)
	}
}
