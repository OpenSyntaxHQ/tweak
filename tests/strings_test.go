package tests

import (
	"strings"
	"testing"

	"github.com/OpenSyntaxHQ/tweak/processors"
)

func TestLower_Command(t *testing.T) {
	assertCommandMeta(t, processors.Lower{}, "lower", "To Lower case (lower)", "Transform your text to lower case")
}

func TestLower_Transform(t *testing.T) {
	p := processors.Lower{}
	tests := []struct{ name, in, want string }{
		{"normal", "HELLO WORLD", "hello world"},
		{"already lower", "hello", "hello"},
		{"mixed", "HeLLo WoRLd", "hello world"},
		{"emoji", "😃🇮🇳", "😃🇮🇳"},
		{"multiline", "A\nB\nC", "a\nb\nc"},
		{"numbers", "ABC123", "abc123"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertTransform(t, p, tt.in, nil, tt.want, false)
		})
	}
}

func TestUpper_Command(t *testing.T) {
	assertCommandMeta(t, processors.Upper{}, "upper", "To Upper case (upper)", "Transform your text to UPPER CASE")
}

func TestUpper_Transform(t *testing.T) {
	p := processors.Upper{}
	tests := []struct{ name, in, want string }{
		{"normal", "hello world", "HELLO WORLD"},
		{"already upper", "HELLO", "HELLO"},
		{"multiline", "a\nb", "A\nB"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertTransform(t, p, tt.in, nil, tt.want, false)
		})
	}
}

func TestTitle_Transform(t *testing.T) {
	p := processors.Title{}
	assertTransform(t, p, "hello world", nil, "Hello World", false)
}

func TestSnake_Transform(t *testing.T) {
	p := processors.Snake{}
	assertTransform(t, p, "hello world", nil, "hello_world", false)
	assertTransform(t, p, "Hello World", nil, "hello_world", false)
}

func TestKebab_Transform(t *testing.T) {
	p := processors.Kebab{}
	assertTransform(t, p, "hello world", nil, "hello-world", false)
}

func TestCamel_Transform(t *testing.T) {
	p := processors.Camel{}
	assertTransform(t, p, "hello world", nil, "helloWorld", false)
}

func TestPascal_Transform(t *testing.T) {
	p := processors.Pascal{}
	assertTransform(t, p, "hello world", nil, "HelloWorld", false)
}

func TestSlug_Transform(t *testing.T) {
	p := processors.Slug{}
	assertTransform(t, p, "Hello World!", nil, "hello-world", false)
}

func TestReverse_Transform(t *testing.T) {
	p := processors.Reverse{}
	assertTransform(t, p, "hello", nil, "olleh", false)
	assertTransform(t, p, "abcd", nil, "dcba", false)
}

func TestTrim_Transform(t *testing.T) {
	p := processors.Trim{}
	assertTransform(t, p, "  hello  ", nil, "hello", false)
	assertTransform(t, p, "\t\nhello\n", nil, "hello", false)
}

func TestRepeat_Transform(t *testing.T) {
	p := processors.Repeat{}
	tests := []struct {
		name  string
		in    string
		flags []processors.Flag
		want  string
	}{
		{"default 2x", "ab", nil, "abab"},
		{"3x", "x", []processors.Flag{{Name: "count", Short: "c", Value: uint(3), Type: processors.FlagUint}}, "xxx"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertTransform(t, p, tt.in, tt.flags, tt.want, false)
		})
	}
}

func TestWrap_Transform(t *testing.T) {
	p := processors.Wrap{}
	long := strings.Repeat("word ", 10)
	got, err := p.Transform([]byte(strings.TrimSpace(long)), processors.Flag{Name: "width", Short: "w", Value: uint(20), Type: processors.FlagUint})
	if err != nil {
		t.Fatal(err)
	}
	for _, line := range strings.Split(got, "\n") {
		if len(line) > 20 {
			t.Errorf("line %q exceeds width 20", line)
		}
	}
}

func TestReplaceText_Transform(t *testing.T) {
	p := processors.ReplaceText{}
	flags := []processors.Flag{
		{Name: "find", Short: "f", Value: "foo", Type: processors.FlagString},
		{Name: "with", Short: "w", Value: "bar", Type: processors.FlagString},
	}
	assertTransform(t, p, "foo baz foo", flags, "bar baz bar", false)
}

func TestCharFrequency_Transform(t *testing.T) {
	p := processors.CharFrequency{}
	got, err := p.Transform([]byte("aab"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got, "a: 2") {
		t.Errorf("expected 'a: 2' in %q", got)
	}
}

func TestWordFrequency_Transform(t *testing.T) {
	p := processors.WordFrequency{}
	got, err := p.Transform([]byte("the cat and the cat"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got, "the: 2") {
		t.Errorf("expected 'the: 2' in %q", got)
	}
}

func TestEscapeQuotes_Transform(t *testing.T) {
	assertTransform(t, processors.EscapeQuotes{}, `say "hello"`, nil, `say \"hello\"`, false)
}

func TestEscapeQuotes_FlagVariants(t *testing.T) {
	p := processors.EscapeQuotes{}

	doubleOnly, err := p.Transform(
		[]byte(`say "hello" and 'bye'`),
		processors.Flag{Name: "double-quote", Short: "d", Value: true, Type: processors.FlagBool},
	)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(doubleOnly, `\"hello\"`) {
		t.Fatalf("expected escaped double quotes, got %q", doubleOnly)
	}
	if strings.Contains(doubleOnly, `\'bye\'`) {
		t.Fatalf("did not expect escaped single quotes in double-only mode, got %q", doubleOnly)
	}

	singleOnly, err := p.Transform(
		[]byte(`say "hello" and 'bye'`),
		processors.Flag{Name: "single-quote", Short: "s", Value: true, Type: processors.FlagBool},
	)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(singleOnly, `\'bye\'`) {
		t.Fatalf("expected escaped single quotes, got %q", singleOnly)
	}
	if strings.Contains(singleOnly, `\"hello\"`) {
		t.Fatalf("did not expect escaped double quotes in single-only mode, got %q", singleOnly)
	}
}

func TestWrap_ZeroWidthAndWhitespaceInput(t *testing.T) {
	p := processors.Wrap{}

	unchanged, err := p.Transform(
		[]byte("this text should stay unchanged"),
		processors.Flag{Name: "width", Short: "w", Value: uint(0), Type: processors.FlagUint},
	)
	if err != nil {
		t.Fatal(err)
	}
	if unchanged != "this text should stay unchanged" {
		t.Fatalf("wrap width=0 should return input unchanged, got %q", unchanged)
	}

	empty, err := p.Transform([]byte("   \t   "))
	if err != nil {
		t.Fatal(err)
	}
	if empty != "" {
		t.Fatalf("wrap on whitespace-only input should be empty, got %q", empty)
	}
}

func TestReplaceText_NoFindFlag(t *testing.T) {
	p := processors.ReplaceText{}
	got, err := p.Transform([]byte("foo bar baz"))
	if err != nil {
		t.Fatal(err)
	}
	if got != "foo bar baz" {
		t.Fatalf("replace without find flag should keep input, got %q", got)
	}
}

func TestCharFrequency_SpecialCharacterLabels(t *testing.T) {
	got, err := (processors.CharFrequency{}).Transform([]byte(" \n\t"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got, "[space]: 1") {
		t.Fatalf("expected [space] label in output, got %q", got)
	}
	if !strings.Contains(got, `\n: 1`) {
		t.Fatalf("expected newline label in output, got %q", got)
	}
	if !strings.Contains(got, `\t: 1`) {
		t.Fatalf("expected tab label in output, got %q", got)
	}
}

func TestWordFrequency_CaseInsensitive(t *testing.T) {
	got, err := (processors.WordFrequency{}).Transform([]byte("Go go GO gO"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got, "go: 4") {
		t.Fatalf("expected case-insensitive aggregation, got %q", got)
	}
}

func TestRepeat_ZeroCount(t *testing.T) {
	got, err := (processors.Repeat{}).Transform(
		[]byte("abc"),
		processors.Flag{Name: "count", Short: "c", Value: uint(0), Type: processors.FlagUint},
	)
	if err != nil {
		t.Fatal(err)
	}
	if got != "" {
		t.Fatalf("repeat count 0 should be empty string, got %q", got)
	}
}
