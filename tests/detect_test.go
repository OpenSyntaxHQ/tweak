package tests

import (
	"strings"
	"testing"

	"github.com/OpenSyntaxHQ/tweak/processors"
)

func TestDetect_Base64(t *testing.T) {
	got, err := (processors.Detect{}).Transform([]byte("SGVsbG8gV29ybGQ="))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got, "Base64") {
		t.Errorf("expected Base64 detection, got: %q", got)
	}
	if !strings.Contains(got, "Hello World") {
		t.Errorf("expected decoded 'Hello World', got: %q", got)
	}
}

func TestDetect_HexEncoded(t *testing.T) {
	got, err := (processors.Detect{}).Transform([]byte("68656c6c6f"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got, "Hex") {
		t.Errorf("expected Hex detection, got: %q", got)
	}
}

func TestDetect_JWT(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiJ9.eyJ1aWQiOjF9.abc123"
	got, err := (processors.Detect{}).Transform([]byte(token))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got, "JWT") {
		t.Errorf("expected JWT detection, got: %q", got)
	}
}

func TestDetect_URLEncoded(t *testing.T) {
	got, err := (processors.Detect{}).Transform([]byte("Hello%20World"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got, "URL") {
		t.Errorf("expected URL-encoded detection, got: %q", got)
	}
}

func TestDetect_BinaryString(t *testing.T) {
	got, err := (processors.Detect{}).Transform([]byte("01001000 01101001"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got, "Binary") {
		t.Fatalf("expected binary detection, got: %q", got)
	}
	if !strings.Contains(got, "Hi") {
		t.Fatalf("expected decoded text in output, got: %q", got)
	}
}

func TestDetect_EmptyInput(t *testing.T) {
	if _, err := (processors.Detect{}).Transform([]byte("   ")); err == nil {
		t.Fatal("expected empty-input error")
	}
}

func TestDetect_UnknownFormat(t *testing.T) {
	got, err := (processors.Detect{}).Transform([]byte("this-is-not-encoded-data"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got, "Could not automatically detect") {
		t.Fatalf("expected unknown-format message, got: %q", got)
	}
}
