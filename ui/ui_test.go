package ui

import (
	"strings"
	"testing"

	"github.com/charmbracelet/bubbles/textinput"

	"github.com/OpenSyntaxHQ/tweak/processors"
)

func TestParseFlagInput_RequiredString(t *testing.T) {
	flag := processors.Flag{Name: "key", Short: "k", Type: processors.FlagString, Required: true}
	if _, err := parseFlagInput(flag, ""); err == nil {
		t.Fatalf("expected required string flag to fail on empty input")
	}
}

func TestParseFlagInput_BoolUint(t *testing.T) {
	boolFlag := processors.Flag{Name: "invert", Short: "v", Type: processors.FlagBool, Value: false}
	parsedBool, err := parseFlagInput(boolFlag, "true")
	if err != nil {
		t.Fatalf("unexpected bool parse error: %v", err)
	}
	if v, ok := parsedBool.Value.(bool); !ok || !v {
		t.Fatalf("expected parsed bool=true, got %#v", parsedBool.Value)
	}

	uintFlag := processors.Flag{Name: "exp", Short: "e", Type: processors.FlagUint, Value: uint(24)}
	parsedUint, err := parseFlagInput(uintFlag, "48")
	if err != nil {
		t.Fatalf("unexpected uint parse error: %v", err)
	}
	if v, ok := parsedUint.Value.(uint); !ok || v != 48 {
		t.Fatalf("expected parsed uint=48, got %#v", parsedUint.Value)
	}
}

func TestPrepareFlagInput_SensitiveEcho(t *testing.T) {
	m := initialModel("hello")
	m.flags = []processors.Flag{{Name: "secret", Short: "s", Type: processors.FlagString, Sensitive: true}}
	m.flagIndex = 0
	m.prepareFlagInput()

	if m.flagInput.EchoMode != textinput.EchoPassword {
		t.Fatalf("expected sensitive flag to use password echo mode")
	}
}

func TestTruncateRuneSafe(t *testing.T) {
	in := "こんにちは世界"
	out := truncate(in, 6)
	if out == "" {
		t.Fatalf("truncate returned empty output")
	}
	if !strings.HasSuffix(out, "...") {
		t.Fatalf("expected truncated output to end with ellipsis, got %q", out)
	}
}
