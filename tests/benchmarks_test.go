package tests

import (
	"strings"
	"testing"

	"github.com/OpenSyntaxHQ/tweak/processors"
)

func BenchmarkSHA256(b *testing.B) {
	payload := []byte(strings.Repeat("abc123", 1024))
	p := processors.SHA256{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := p.Transform(payload); err != nil {
			b.Fatalf("transform failed: %v", err)
		}
	}
}

func BenchmarkBase64Encode(b *testing.B) {
	payload := []byte(strings.Repeat("hello world", 1024))
	p := processors.Base64Encode{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := p.Transform(payload); err != nil {
			b.Fatalf("transform failed: %v", err)
		}
	}
}

func BenchmarkSortLines(b *testing.B) {
	payload := []byte(strings.Repeat("delta\nalpha\ncharlie\nbravo\n", 1024))
	p := processors.SortLines{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := p.Transform(payload); err != nil {
			b.Fatalf("transform failed: %v", err)
		}
	}
}

func BenchmarkFormatJSON(b *testing.B) {
	payload := []byte(`{"z":1,"a":[1,2,3],"nested":{"x":"y"}}`)
	p := processors.FormatJSON{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := p.Transform(payload); err != nil {
			b.Fatalf("transform failed: %v", err)
		}
	}
}
