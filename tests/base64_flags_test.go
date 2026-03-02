package tests

import (
	"strings"
	"testing"

	"github.com/OpenSyntaxHQ/tweak/processors"
)

func TestBase64RawFlag_ByShortAndName(t *testing.T) {
	input := "hello"

	encodedRawByShort, err := (processors.Base64Encode{}).Transform(
		[]byte(input),
		processors.Flag{Name: "raw", Short: "r", Value: true, Type: processors.FlagBool},
	)
	if err != nil {
		t.Fatalf("base64 raw encode (short) error: %v", err)
	}
	if strings.HasSuffix(encodedRawByShort, "=") {
		t.Fatalf("raw base64 output should be unpadded, got %q", encodedRawByShort)
	}

	decodedRawByName, err := (processors.Base64Decode{}).Transform(
		[]byte(encodedRawByShort),
		processors.Flag{Name: "raw", Value: true, Type: processors.FlagBool},
	)
	if err != nil {
		t.Fatalf("base64 raw decode (name) error: %v", err)
	}
	if decodedRawByName != input {
		t.Fatalf("base64 raw roundtrip mismatch: got %q want %q", decodedRawByName, input)
	}
}

func TestBase64URLRawFlag_RoundTripAndDecodeError(t *testing.T) {
	input := "Hello/World?"
	rawFlag := processors.Flag{Name: "raw", Short: "r", Value: true, Type: processors.FlagBool}

	encodedRaw, err := (processors.Base64URLEncode{}).Transform([]byte(input), rawFlag)
	if err != nil {
		t.Fatalf("base64url raw encode error: %v", err)
	}
	if strings.Contains(encodedRaw, "=") {
		t.Fatalf("raw base64url output should be unpadded, got %q", encodedRaw)
	}

	decodedRaw, err := (processors.Base64URLDecode{}).Transform([]byte(encodedRaw), rawFlag)
	if err != nil {
		t.Fatalf("base64url raw decode error: %v", err)
	}
	if decodedRaw != input {
		t.Fatalf("base64url raw roundtrip mismatch: got %q want %q", decodedRaw, input)
	}

	if _, err := (processors.Base64URLDecode{}).Transform([]byte("%%%"), rawFlag); err == nil {
		t.Fatal("expected base64url raw decode error for invalid payload")
	}
}
