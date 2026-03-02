package tests

import (
	"regexp"
	"strings"
	"testing"

	"github.com/OpenSyntaxHQ/tweak/processors"
)

func TestUUID_Format(t *testing.T) {
	uuidRe := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`)
	for i := 0; i < 5; i++ {
		got, err := (processors.UUID{}).Transform([]byte(""))
		if err != nil {
			t.Fatal(err)
		}
		if !uuidRe.MatchString(got) {
			t.Errorf("UUID %q does not match v4 format", got)
		}
	}
}

func TestROT13_Transform(t *testing.T) {
	assertTransform(t, processors.ROT13{}, "Hello, World!", nil, "Uryyb, Jbeyq!", false)
	got1, _ := (processors.ROT13{}).Transform([]byte("Hello, World!"))
	got2, _ := (processors.ROT13{}).Transform([]byte(got1))
	if got2 != "Hello, World!" {
		t.Errorf("ROT13 applied twice should restore original, got: %q", got2)
	}
}

func TestMorseCodeEncode_Transform(t *testing.T) {
	got, err := (processors.MorseCodeEncode{}).Transform([]byte("SOS"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got, "...") {
		t.Errorf("Morse SOS should contain '...', got: %q", got)
	}
}

func TestMorseCodeDecode_Transform(t *testing.T) {
	encoded, err := (processors.MorseCodeEncode{}).Transform([]byte("HELLO"))
	if err != nil {
		t.Fatal(err)
	}
	got, err := (processors.MorseCodeDecode{}).Transform([]byte(encoded))
	if err != nil {
		t.Fatalf("morse decode error: %v", err)
	}
	if !strings.Contains(strings.ToUpper(got), "HELLO") {
		t.Errorf("morse decode: expected HELLO, got: %q", got)
	}
}

func TestLorem_Transform(t *testing.T) {
	got, err := (processors.Lorem{}).Transform([]byte(""))
	if err != nil {
		t.Fatal(err)
	}
	if len(got) < 10 {
		t.Errorf("Lorem ipsum output too short: %q", got)
	}
}

func TestMarkdown_Transform(t *testing.T) {
	input := "# Hello\n\n**World**"
	got, err := (processors.Markdown{}).Transform([]byte(input))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got, "Hello") {
		t.Errorf("Markdown should preserve text, got: %q", got)
	}
}

func TestQRCode_Transform(t *testing.T) {
	got, err := (processors.QRCode{}).Transform([]byte("https://example.com"))
	if err != nil {
		t.Fatal(err)
	}
	if len(got) == 0 {
		t.Error("QRCode should produce non-empty output")
	}
}

func TestQRCode_FlagBranches(t *testing.T) {
	tests := []struct {
		name  string
		flags []processors.Flag
	}{
		{
			name: "size high + full blocks",
			flags: []processors.Flag{
				{Name: "size", Short: "s", Value: uint(300), Type: processors.FlagUint},
				{Name: "full", Short: "f", Value: true, Type: processors.FlagBool},
			},
		},
		{
			name: "size medium level medium",
			flags: []processors.Flag{
				{Name: "size", Short: "s", Value: uint(220), Type: processors.FlagUint},
				{Name: "level", Short: "l", Value: "medium", Type: processors.FlagString},
			},
		},
		{
			name: "size low level low",
			flags: []processors.Flag{
				{Name: "size", Short: "s", Value: uint(100), Type: processors.FlagUint},
				{Name: "level", Short: "l", Value: "low", Type: processors.FlagString},
			},
		},
		{
			name: "level high",
			flags: []processors.Flag{
				{Name: "level", Short: "l", Value: "H", Type: processors.FlagString},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := (processors.QRCode{}).Transform([]byte("https://example.com"), tc.flags...)
			if err != nil {
				t.Fatal(err)
			}
			if len(strings.TrimSpace(got)) == 0 {
				t.Fatal("expected qr output with provided flags")
			}
		})
	}
}

func TestZeropad_Transform(t *testing.T) {
	tests := []struct {
		in     string
		n      uint
		minLen int
	}{
		{"5", 4, 4},
		{"42", 6, 6},
		{"100", 2, 3},
	}
	for _, tt := range tests {
		flags := []processors.Flag{
			{Name: "n", Short: "n", Value: tt.n, Type: processors.FlagUint},
		}
		got, err := (processors.Zeropad{}).Transform([]byte(tt.in), flags...)
		if err != nil {
			t.Fatalf("Zeropad(%q, n=%d) error: %v", tt.in, tt.n, err)
		}
		if len(got) < tt.minLen {
			t.Errorf("Zeropad(%q, n=%d) = %q, expected len >= %d", tt.in, tt.n, got, tt.minLen)
		}
	}
}

func TestHexToRGB_Transform(t *testing.T) {
	got, err := (processors.HexToRGB{}).Transform([]byte("#ff0000"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got, "255") {
		t.Errorf("HexToRGB(#ff0000) should contain 255, got: %q", got)
	}
}

func TestRemoveSpaces_Transform(t *testing.T) {
	assertTransform(t, processors.RemoveSpaces{}, "hello world", nil, "helloworld", false)
}

func TestRemoveNewLines_Transform(t *testing.T) {
	got, err := (processors.RemoveNewLines{}).Transform([]byte("a\nb\nc"))
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(got, "\n") {
		t.Errorf("RemoveNewLines should remove all newlines, got: %q", got)
	}
}

func TestExtractEmails_Transform(t *testing.T) {
	input := "contact us at hello@example.com or support@test.org"
	got, err := (processors.ExtractEmails{}).Transform([]byte(input))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got, "hello@example.com") {
		t.Errorf("expected email in output, got: %q", got)
	}
}

func TestExtractURLs_Transform(t *testing.T) {
	input := "see https://example.com and http://test.org for more"
	got, err := (processors.ExtractURLs{}).Transform([]byte(input))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got, "example.com") {
		t.Errorf("expected URL in output, got: %q", got)
	}
}

func TestExtractIPs_Transform(t *testing.T) {
	input := "server at 192.168.1.1 and backup at 10.0.0.1"
	got, err := (processors.ExtractIPs{}).Transform([]byte(input))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got, "192.168.1.1") {
		t.Errorf("expected IP in output, got: %q", got)
	}
}

func TestCaesar_RoundTrip(t *testing.T) {
	input := "HELLO"
	flags := []processors.Flag{
		{Name: "shift", Short: "s", Value: uint(3), Type: processors.FlagUint},
	}
	encoded, err := (processors.CaesarEncode{}).Transform([]byte(input), flags...)
	if err != nil {
		t.Fatal(err)
	}
	decoded, err := (processors.CaesarDecode{}).Transform([]byte(encoded), flags...)
	if err != nil {
		t.Fatalf("caesar decode error: %v", err)
	}
	if !strings.EqualFold(decoded, input) {
		t.Errorf("caesar roundtrip: got %q, want %q", decoded, input)
	}
}

func TestCaesar_MixedCaseAndPunctuation(t *testing.T) {
	flags := []processors.Flag{{Name: "shift", Short: "s", Value: 2, Type: processors.FlagInt}}
	got, err := (processors.CaesarEncode{}).Transform([]byte("Abc xyz!"), flags...)
	if err != nil {
		t.Fatal(err)
	}
	if got != "Cde zab!" {
		t.Fatalf("unexpected caesar output: %q", got)
	}

	back, err := (processors.CaesarDecode{}).Transform([]byte(got), flags...)
	if err != nil {
		t.Fatal(err)
	}
	if back != "Abc xyz!" {
		t.Fatalf("unexpected caesar decode output: %q", back)
	}
}

func TestRegexMatch_Transform(t *testing.T) {
	flags := []processors.Flag{
		{Name: "pattern", Short: "p", Value: `\d+`, Type: processors.FlagString},
	}
	got, err := (processors.RegexMatch{}).Transform([]byte("abc 123 def 456"), flags...)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got, "123") {
		t.Errorf("RegexMatch should find '123', got: %q", got)
	}
}

func TestRegexReplace_Transform(t *testing.T) {
	flags := []processors.Flag{
		{Name: "pattern", Short: "p", Value: `\d+`, Type: processors.FlagString},
		{Name: "replace", Short: "r", Value: "NUM", Type: processors.FlagString},
	}
	got, err := (processors.RegexReplace{}).Transform([]byte("abc 123 def 456"), flags...)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(got, "123") {
		t.Errorf("RegexReplace should have replaced '123', got: %q", got)
	}
	if !strings.Contains(got, "NUM") {
		t.Errorf("RegexReplace should contain 'NUM', got: %q", got)
	}
}

func TestCountWords_Transform(t *testing.T) {
	assertTransform(t, processors.CountWords{}, "hello world foo", nil, "3", false)
}

func TestCountCharacters_Transform(t *testing.T) {
	assertTransform(t, processors.CountCharacters{}, "hello", nil, "5", false)
}
