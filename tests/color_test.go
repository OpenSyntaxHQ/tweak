package tests

import (
	"testing"

	"github.com/OpenSyntaxHQ/tweak/processors"
)

func TestRGBHex_Transform(t *testing.T) {
	tests := []struct{ name, in, want string }{
		{"pure red", "255, 0, 0", "#ff0000"},
		{"pure green", "0, 255, 0", "#00ff00"},
		{"pure blue", "0, 0, 255", "#0000ff"},
		{"white", "255, 255, 255", "#ffffff"},
		{"black", "0, 0, 0", "#000000"},
		{"mixed", "255, 128, 0", "#ff8000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertTransform(t, processors.RGBToHex{}, tt.in, nil, tt.want, false)
		})
	}
}

func TestHSLHex_Transform(t *testing.T) {
	tests := []struct{ name, in, want string }{
		{"red hsl", "0, 100, 50", "#ff0000"},
		{"white hsl", "0, 0, 100", "#ffffff"},
		{"black hsl", "0, 0, 0", "#000000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := processors.HSLToHex{}.Transform([]byte(tt.in))
			if err != nil {
				t.Fatal(err)
			}
			if got != tt.want {
				t.Errorf("HSLToHex(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
