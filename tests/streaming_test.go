package tests

import (
	"bytes"
	"strings"
	"testing"

	"github.com/OpenSyntaxHQ/tweak/processors"
)

type lineStreamingUpper struct{}

func (lineStreamingUpper) Name() string    { return "line-streaming-upper" }
func (lineStreamingUpper) Alias() []string { return nil }
func (lineStreamingUpper) Flags() []processors.Flag {
	return nil
}
func (lineStreamingUpper) Transform(data []byte, _ ...processors.Flag) (string, error) {
	return strings.ToUpper(string(data)), nil
}
func (lineStreamingUpper) StreamingSpec() processors.StreamingSpec {
	return processors.StreamingSpec{Mode: processors.StreamingModeLine}
}

func TestStreamingParity_Chunked(t *testing.T) {
	input := "Hello, 世界!\nThis is a chunked streaming test."
	p := processors.Lower{}

	direct, err := p.Transform([]byte(input))
	if err != nil {
		t.Fatalf("direct transform failed: %v", err)
	}

	var out bytes.Buffer
	if err := processors.TransformStream(p, strings.NewReader(input), &out); err != nil {
		t.Fatalf("stream transform failed: %v", err)
	}

	if out.String() != direct {
		t.Fatalf("chunked stream mismatch\n got: %q\nwant: %q", out.String(), direct)
	}
}

func TestStreamingParity_Buffered(t *testing.T) {
	input := "line1\nline2\nline3"
	p := processors.SHA256{}

	direct, err := p.Transform([]byte(input))
	if err != nil {
		t.Fatalf("direct transform failed: %v", err)
	}

	var out bytes.Buffer
	if err := processors.TransformStream(p, strings.NewReader(input), &out); err != nil {
		t.Fatalf("stream transform failed: %v", err)
	}

	if out.String() != direct {
		t.Fatalf("buffered stream mismatch\n got: %q\nwant: %q", out.String(), direct)
	}
}

func TestStreamingParity_Native(t *testing.T) {
	input := strings.Repeat("abc123", 1000)
	p := processors.BLAKE2s{}

	direct, err := p.Transform([]byte(input))
	if err != nil {
		t.Fatalf("direct transform failed: %v", err)
	}

	var out bytes.Buffer
	if err := processors.TransformStream(p, strings.NewReader(input), &out); err != nil {
		t.Fatalf("stream transform failed: %v", err)
	}

	if out.String() != direct {
		t.Fatalf("native stream mismatch\n got: %q\nwant: %q", out.String(), direct)
	}
}

func TestShouldStreamRules(t *testing.T) {
	if !processors.ShouldStream(processors.Lower{}, 1, 10*1024*1024) {
		t.Fatalf("expected lower to prefer streaming")
	}

	if processors.ShouldStream(processors.SortLines{}, 100, 10*1024*1024) {
		t.Fatalf("expected sort-lines to avoid small-file streaming")
	}

	if !processors.ShouldStream(processors.SortLines{}, 20*1024*1024, 10*1024*1024) {
		t.Fatalf("expected sort-lines to stream for large file")
	}

	if processors.ShouldStream(processors.ROT13{}, 20*1024*1024, 10*1024*1024) {
		t.Fatalf("expected processors without streaming spec not to stream")
	}
}

func TestTransformStreamRejectsUnsupportedProcessor(t *testing.T) {
	var out bytes.Buffer
	err := processors.TransformStream(processors.ROT13{}, strings.NewReader("abc"), &out)
	if err == nil {
		t.Fatalf("expected error for unsupported streaming processor")
	}
}

func TestStreamingParity_NativeBlake2b(t *testing.T) {
	input := strings.Repeat("blake2b", 4096)
	p := processors.BLAKE2b{}

	direct, err := p.Transform([]byte(input))
	if err != nil {
		t.Fatalf("direct transform failed: %v", err)
	}

	var out bytes.Buffer
	if err := processors.TransformStream(p, strings.NewReader(input), &out); err != nil {
		t.Fatalf("stream transform failed: %v", err)
	}

	if out.String() != direct {
		t.Fatalf("native blake2b stream mismatch\n got: %q\nwant: %q", out.String(), direct)
	}
}

func TestTransformStream_LineMode(t *testing.T) {
	p := lineStreamingUpper{}
	input := "a\nb\nc"

	var out bytes.Buffer
	if err := processors.TransformStream(p, strings.NewReader(input), &out); err != nil {
		t.Fatalf("line streaming transform failed: %v", err)
	}

	if got, want := out.String(), "A\nB\nC\n"; got != want {
		t.Fatalf("line streaming mismatch\n got: %q\nwant: %q", got, want)
	}
}
