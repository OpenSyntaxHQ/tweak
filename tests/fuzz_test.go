package tests

import (
	"testing"

	"github.com/OpenSyntaxHQ/tweak/processors"
)

func FuzzJSONTransforms(f *testing.F) {
	f.Add(`{"k":"v"}`)
	f.Add(`[1,2,3]`)
	f.Add(`not json`)

	f.Fuzz(func(t *testing.T, input string) {
		_, _ = processors.FormatJSON{}.Transform([]byte(input))
		_, _ = processors.JSONMinify{}.Transform([]byte(input))
		_, _ = processors.JSONToYAML{}.Transform([]byte(input))
	})
}

func FuzzRegexProcessors(f *testing.F) {
	f.Add("hello 123 world", "[a-z]+", "x")
	f.Add("", ".*", "")

	f.Fuzz(func(t *testing.T, input, pattern, replacement string) {
		_, _ = processors.RegexMatch{}.Transform([]byte(input), processors.Flag{Short: "p", Value: pattern})
		_, _ = processors.RegexReplace{}.Transform(
			[]byte(input),
			processors.Flag{Short: "p", Value: pattern},
			processors.Flag{Short: "r", Value: replacement},
		)
	})
}
