package tests

import (
	"strings"
	"testing"

	"github.com/OpenSyntaxHQ/tweak/processors"
)

func TestBase58Check_RoundTrip(t *testing.T) {
	input := "hello tweak"
	checkFlag := processors.Flag{Name: "check", Short: "c", Value: true, Type: processors.FlagBool}

	encoded, err := (processors.Base58Encode{}).Transform([]byte(input), checkFlag)
	if err != nil {
		t.Fatalf("base58check encode error: %v", err)
	}

	decoded, err := (processors.Base58Decode{}).Transform([]byte(encoded), checkFlag)
	if err != nil {
		t.Fatalf("base58check decode error: %v", err)
	}
	if decoded != input {
		t.Fatalf("base58check roundtrip mismatch: got %q want %q", decoded, input)
	}
}

func TestBase58Check_ChecksumFailureAndShortInput(t *testing.T) {
	checkFlag := processors.Flag{Name: "check", Short: "c", Value: true, Type: processors.FlagBool}

	encoded, err := (processors.Base58Encode{}).Transform([]byte("checksum-me"), checkFlag)
	if err != nil {
		t.Fatalf("base58check encode error: %v", err)
	}

	changed := "1"
	if strings.HasSuffix(encoded, "1") {
		changed = "2"
	}
	tampered := encoded[:len(encoded)-1] + changed

	if _, err := (processors.Base58Decode{}).Transform([]byte(tampered), checkFlag); err == nil {
		t.Fatal("expected checksum verification failure for tampered base58check payload")
	}

	if _, err := (processors.Base58Decode{}).Transform([]byte("1"), checkFlag); err == nil {
		t.Fatal("expected short-input error for base58check decode")
	}
}

func TestCrockfordChecksum_VerifyBranches(t *testing.T) {
	checksumFlag := processors.Flag{Name: "checksum", Short: "c", Value: true, Type: processors.FlagBool}
	verifyFlag := processors.Flag{Name: "verify", Short: "v", Value: true, Type: processors.FlagBool}

	encodedWithChecksum, err := (processors.CrockfordBase32Encode{}).Transform([]byte("hello"), checksumFlag)
	if err != nil {
		t.Fatalf("crockford encode error: %v", err)
	}

	decoded, err := (processors.CrockfordBase32Decode{}).Transform([]byte(encodedWithChecksum), verifyFlag)
	if err != nil {
		t.Fatalf("crockford verify decode error: %v", err)
	}
	if decoded != "hello" {
		t.Fatalf("crockford verify decode mismatch: got %q", decoded)
	}

	validReplacement := "A"
	if strings.HasSuffix(encodedWithChecksum, "A") {
		validReplacement = "B"
	}
	wrongChecksum := encodedWithChecksum[:len(encodedWithChecksum)-1] + validReplacement
	if _, err := (processors.CrockfordBase32Decode{}).Transform([]byte(wrongChecksum), verifyFlag); err == nil {
		t.Fatal("expected checksum mismatch error for crockford verify")
	}

	invalidChecksum := encodedWithChecksum[:len(encodedWithChecksum)-1] + "!"
	if _, err := (processors.CrockfordBase32Decode{}).Transform([]byte(invalidChecksum), verifyFlag); err == nil {
		t.Fatal("expected invalid checksum character error for crockford verify")
	}
}

func TestBase62Decode_PrefixAndInvalidCharacter(t *testing.T) {
	prefixFlag := processors.Flag{Name: "prefix", Short: "p", Value: "id", Type: processors.FlagString}

	encoded, err := (processors.Base62Encode{}).Transform([]byte("hello"), prefixFlag)
	if err != nil {
		t.Fatalf("base62 encode error: %v", err)
	}
	if !strings.HasPrefix(encoded, "id_") {
		t.Fatalf("expected prefixed base62 output, got %q", encoded)
	}

	decoded, err := (processors.Base62Decode{}).Transform([]byte(encoded))
	if err != nil {
		t.Fatalf("base62 decode error: %v", err)
	}
	if decoded != "hello" {
		t.Fatalf("base62 decode mismatch: got %q want %q", decoded, "hello")
	}

	if _, err := (processors.Base62Decode{}).Transform([]byte("bad-!")); err == nil {
		t.Fatal("expected invalid-character error for base62 decode")
	}
}
