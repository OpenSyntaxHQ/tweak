package tests

import (
	"strings"
	"testing"

	"github.com/OpenSyntaxHQ/tweak/processors"
)

func TestMD5_Transform(t *testing.T) {
	assertTransform(t, processors.MD5{}, "hello", nil, "5d41402abc4b2a76b9719d911017c592", false)
}

func TestSHA1_Transform(t *testing.T) {
	assertTransform(t, processors.SHA1{}, "hello", nil, "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d", false)
}

func TestSHA224_Transform(t *testing.T) {
	got, err := processors.SHA224{}.Transform([]byte("hello"))
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 56 {
		t.Errorf("SHA224 want 56 hex chars, got %d: %q", len(got), got)
	}
}

func TestSHA256_Transform(t *testing.T) {
	assertTransform(t, processors.SHA256{}, "hello", nil, "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824", false)
}

func TestSHA384_Transform(t *testing.T) {
	got, err := processors.SHA384{}.Transform([]byte("hello"))
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 96 {
		t.Errorf("SHA384 want 96 hex chars, got %d: %q", len(got), got)
	}
}

func TestSHA512_Transform(t *testing.T) {
	got, err := processors.SHA512{}.Transform([]byte("hello"))
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 128 {
		t.Errorf("SHA512 want 128 hex chars, got %d", len(got))
	}
}

func TestBcrypt_Transform(t *testing.T) {
	got, err := processors.Bcrypt{}.Transform([]byte("password"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(got, "$2a$") {
		t.Errorf("expected bcrypt prefix, got %q", got)
	}
}

func TestAdler32_Transform(t *testing.T) {
	assertTransform(t, processors.Adler32{}, "hello", nil, "062c0215", false)
}
