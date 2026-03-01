package tests

import (
	"testing"

	"github.com/OpenSyntaxHQ/tweak/processors"
)

func TestBaseConvert_Transform(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		fromBase uint
		toBase   uint
		want     string
	}{
		{"dec to hex", "255", 10, 16, "ff"},
		{"bin to dec", "1010", 2, 10, "10"},
		{"dec to bin", "10", 10, 2, "1010"},
		{"hex to dec", "ff", 16, 10, "255"},
		{"dec to oct", "8", 10, 8, "10"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flags := []processors.Flag{
				{Name: "from", Short: "f", Value: tt.fromBase, Type: processors.FlagUint},
				{Name: "to", Short: "t", Value: tt.toBase, Type: processors.FlagUint},
			}
			got, err := processors.BaseConvert{}.Transform([]byte(tt.input), flags...)
			if err != nil {
				t.Fatalf("error: %v", err)
			}
			if got != tt.want {
				t.Errorf("BaseConvert(%q from=%d to=%d) = %q, want %q", tt.input, tt.fromBase, tt.toBase, got, tt.want)
			}
		})
	}
}
