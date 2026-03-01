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
