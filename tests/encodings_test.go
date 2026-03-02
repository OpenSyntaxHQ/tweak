package tests

import (
	"testing"

	"github.com/OpenSyntaxHQ/tweak/processors"
)

func TestBase64_RoundTrip(t *testing.T) {
	input := "Hello, World!"
	enc, err := (processors.Base64Encode{}).Transform([]byte(input))
	if err != nil {
		t.Fatal(err)
	}
	dec, err := (processors.Base64Decode{}).Transform([]byte(enc))
	if err != nil {
		t.Fatal(err)
	}
	if dec != input {
		t.Errorf("base64 roundtrip: got %q, want %q", dec, input)
	}
}

func TestBase32_RoundTrip(t *testing.T) {
	input := "hello"
	enc, err := (processors.Base32Encoding{}).Transform([]byte(input))
	if err != nil {
		t.Fatal(err)
	}
	dec, err := (processors.Base32Decode{}).Transform([]byte(enc))
	if err != nil {
		t.Fatalf("base32 decode error: %v", err)
	}
	if dec != input {
		t.Errorf("base32 roundtrip: got %q, want %q", dec, input)
	}
}

func TestBase64URL_RoundTrip(t *testing.T) {
	input := "Hello, World!"
	enc, err := (processors.Base64URLEncode{}).Transform([]byte(input))
	if err != nil {
		t.Fatal(err)
	}
	dec, err := (processors.Base64URLDecode{}).Transform([]byte(enc))
	if err != nil {
		t.Fatalf("base64url decode error: %v", err)
	}
	if dec != input {
		t.Errorf("base64url roundtrip: got %q, want %q", dec, input)
	}
}

func TestHexEncode_RoundTrip(t *testing.T) {
	input := "hello"
	enc, err := (processors.HexEncode{}).Transform([]byte(input))
	if err != nil {
		t.Fatal(err)
	}
	dec, err := (processors.HexDecode{}).Transform([]byte(enc))
	if err != nil {
		t.Fatalf("hex decode error: %v", err)
	}
	if dec != input {
		t.Errorf("hex roundtrip: got %q, want %q", dec, input)
	}
}

func TestBinary_RoundTrip(t *testing.T) {
	input := "hi"
	enc, err := (processors.BinaryEncode{}).Transform([]byte(input))
	if err != nil {
		t.Fatal(err)
	}
	dec, err := (processors.BinaryDecode{}).Transform([]byte(enc))
	if err != nil {
		t.Fatalf("binary decode error: %v", err)
	}
	if dec != input {
		t.Errorf("binary roundtrip: got %q, want %q", dec, input)
	}
}

func TestBase58_RoundTrip(t *testing.T) {
	input := "hello"
	enc, err := (processors.Base58Encode{}).Transform([]byte(input))
	if err != nil {
		t.Fatal(err)
	}
	dec, err := (processors.Base58Decode{}).Transform([]byte(enc))
	if err != nil {
		t.Fatalf("base58 decode error: %v", err)
	}
	if dec != input {
		t.Errorf("base58 roundtrip: got %q, want %q", dec, input)
	}
}

func TestBase62_RoundTrip(t *testing.T) {
	input := "hello"
	enc, err := (processors.Base62Encode{}).Transform([]byte(input))
	if err != nil {
		t.Fatal(err)
	}
	dec, err := (processors.Base62Decode{}).Transform([]byte(enc))
	if err != nil {
		t.Fatalf("base62 decode error: %v", err)
	}
	if dec != input {
		t.Errorf("base62 roundtrip: got %q, want %q", dec, input)
	}
}

func TestASCII85_RoundTrip(t *testing.T) {
	input := "hello world"
	enc, err := (processors.ASCII85Encoding{}).Transform([]byte(input))
	if err != nil {
		t.Fatal(err)
	}
	dec, err := (processors.ASCII85Decoding{}).Transform([]byte(enc))
	if err != nil {
		t.Fatalf("ascii85 decode error: %v", err)
	}
	if dec != input {
		t.Errorf("ascii85 roundtrip: got %q, want %q", dec, input)
	}
}

func TestCrockfordBase32_RoundTrip(t *testing.T) {
	input := "hello"
	enc, err := (processors.CrockfordBase32Encode{}).Transform([]byte(input))
	if err != nil {
		t.Fatal(err)
	}
	dec, err := (processors.CrockfordBase32Decode{}).Transform([]byte(enc))
	if err != nil {
		t.Fatalf("crockford decode error: %v", err)
	}
	if dec != input {
		t.Errorf("crockford roundtrip: got %q, want %q", dec, input)
	}
}

func TestURL_RoundTrip(t *testing.T) {
	input := "hello world & more"
	enc, err := (processors.URLEncode{}).Transform([]byte(input))
	if err != nil {
		t.Fatal(err)
	}
	dec, err := (processors.URLDecode{}).Transform([]byte(enc))
	if err != nil {
		t.Fatalf("url decode error: %v", err)
	}
	if dec != input {
		t.Errorf("url roundtrip: got %q, want %q", dec, input)
	}
}

func TestHTML_RoundTrip(t *testing.T) {
	input := `<hello> & "world"`
	enc, err := (processors.HTMLEncode{}).Transform([]byte(input))
	if err != nil {
		t.Fatal(err)
	}
	dec, err := (processors.HTMLDecode{}).Transform([]byte(enc))
	if err != nil {
		t.Fatalf("html decode error: %v", err)
	}
	if dec != input {
		t.Errorf("html roundtrip: got %q, want %q", dec, input)
	}
}

func TestHexDecode_InvalidInput(t *testing.T) {
	if _, err := (processors.HexDecode{}).Transform([]byte("zz")); err == nil {
		t.Fatal("expected hex decode error for invalid input")
	}
}

func TestURLDecode_MalformedPercent(t *testing.T) {
	got, err := (processors.URLDecode{}).Transform([]byte("%zz"))
	if err != nil {
		t.Fatalf("url-decode should not return errors, got: %v", err)
	}
	if got != "" {
		t.Fatalf("expected empty output for malformed query escape, got %q", got)
	}
}
