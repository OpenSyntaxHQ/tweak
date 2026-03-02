package tests

import (
	"testing"

	"github.com/OpenSyntaxHQ/tweak/processors"
)

func TestXXH64_Transform(t *testing.T) {
	got, err := (processors.XXH64{}).Transform([]byte("hello"))
	if err != nil {
		t.Fatal(err)
	}
	if len(got) == 0 {
		t.Error("xxh64 should return non-empty hash")
	}
}

func TestXXH32_Transform(t *testing.T) {
	got, err := (processors.XXH32{}).Transform([]byte("hello"))
	if err != nil {
		t.Fatal(err)
	}
	if len(got) == 0 {
		t.Error("xxh32 should return non-empty hash")
	}
}

func TestXXH128_Transform(t *testing.T) {
	got, err := (processors.XXH128{}).Transform([]byte("hello"))
	if err != nil {
		t.Fatal(err)
	}
	if len(got) == 0 {
		t.Error("xxh128 should return non-empty hash")
	}
}

func TestBLAKE2b_Transform(t *testing.T) {
	got, err := (processors.BLAKE2b{}).Transform([]byte("hello"))
	if err != nil {
		t.Fatal(err)
	}
	if len(got) == 0 {
		t.Error("BLAKE2b should return non-empty hash")
	}
}

func TestBLAKE2b_SizeFlagBranches(t *testing.T) {
	got, err := (processors.BLAKE2b{}).Transform(
		[]byte("hello"),
		processors.Flag{Name: "size", Short: "s", Value: uint(32), Type: processors.FlagUint},
	)
	if err != nil {
		t.Fatalf("blake2b size=32 error: %v", err)
	}
	if len(got) != 64 {
		t.Fatalf("blake2b size=32 hex length = %d, want 64", len(got))
	}

	if _, err := (processors.BLAKE2b{}).Transform(
		[]byte("hello"),
		processors.Flag{Name: "size", Short: "s", Value: uint(0), Type: processors.FlagUint},
	); err == nil {
		t.Fatal("expected blake2b size validation error")
	}
}

func TestBLAKE2s_Transform(t *testing.T) {
	got, err := (processors.BLAKE2s{}).Transform([]byte("hello"))
	if err != nil {
		t.Fatal(err)
	}
	if len(got) == 0 {
		t.Error("BLAKE2s should return non-empty hash")
	}
}

func TestCRC32_Transform(t *testing.T) {
	assertTransform(t, processors.CRC32{}, "hello", nil, "3610a686", false)
}

func TestAdler32Sum_Transform(t *testing.T) {
	assertTransform(t, processors.Adler32{}, "hello", nil, "062c0215", false)
}

func TestCRC32_PolynomialBranches(t *testing.T) {
	tests := []struct {
		name  string
		flags []processors.Flag
	}{
		{
			name: "castagnoli",
			flags: []processors.Flag{
				{Name: "polynomial", Short: "p", Value: "castagnoli", Type: processors.FlagString},
			},
		},
		{
			name: "koopman",
			flags: []processors.Flag{
				{Name: "polynomial", Short: "p", Value: "koopman", Type: processors.FlagString},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := (processors.CRC32{}).Transform([]byte("hello"), tt.flags...)
			if err != nil {
				t.Fatalf("crc32 %s error: %v", tt.name, err)
			}
			if len(got) != 8 {
				t.Fatalf("crc32 %s output length = %d, want 8", tt.name, len(got))
			}
		})
	}
}

func TestCRC32_InvalidPolynomial(t *testing.T) {
	_, err := (processors.CRC32{}).Transform(
		[]byte("hello"),
		processors.Flag{Name: "polynomial", Short: "p", Value: "invalid", Type: processors.FlagString},
	)
	if err == nil {
		t.Fatal("expected unsupported polynomial error")
	}
}
