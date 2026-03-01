package tests

import (
	"testing"

	"github.com/OpenSyntaxHQ/tweak/processors"
)

func TestPadLeft_Transform(t *testing.T) {
	tests := []struct {
		name  string
		input string
		width uint
		char  string
		want  string
	}{
		{"basic", "hi", 6, "x", "xxxxhi"},
		{"already wide", "hello world", 5, " ", "hello world"},
		{"dash pad", "5", 4, "0", "0005"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flags := []processors.Flag{
				{Name: "width", Short: "w", Value: tt.width, Type: processors.FlagUint},
				{Name: "char", Short: "c", Value: tt.char, Type: processors.FlagString},
			}
			assertTransform(t, processors.PadLeft{}, tt.input, flags, tt.want, false)
		})
	}
}

func TestPadRight_Transform(t *testing.T) {
	tests := []struct {
		name  string
		input string
		width uint
		char  string
		want  string
	}{
		{"basic", "hi", 6, "-", "hi----"},
		{"already wide", "hello world", 5, " ", "hello world"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flags := []processors.Flag{
				{Name: "width", Short: "w", Value: tt.width, Type: processors.FlagUint},
				{Name: "char", Short: "c", Value: tt.char, Type: processors.FlagString},
			}
			assertTransform(t, processors.PadRight{}, tt.input, flags, tt.want, false)
		})
	}
}
