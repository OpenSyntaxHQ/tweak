package tests

import (
	"testing"

	"github.com/OpenSyntaxHQ/tweak/processors"
)

func TestXXH64_Transform(t *testing.T) {
	got, err := processors.XXH64{}.Transform([]byte("hello"))
	if err != nil {
		t.Fatal(err)
	}
	if len(got) == 0 {
		t.Error("xxh64 should return non-empty hash")
	}
}

func TestXXH32_Transform(t *testing.T) {
	got, err := processors.XXH32{}.Transform([]byte("hello"))
	if err != nil {
		t.Fatal(err)
	}
	if len(got) == 0 {
		t.Error("xxh32 should return non-empty hash")
	}
}

func TestXXH128_Transform(t *testing.T) {
	got, err := processors.XXH128{}.Transform([]byte("hello"))
	if err != nil {
		t.Fatal(err)
	}
	if len(got) == 0 {
		t.Error("xxh128 should return non-empty hash")
	}
}

func TestBLAKE2b_Transform(t *testing.T) {
	got, err := processors.BLAKE2b{}.Transform([]byte("hello"))
	if err != nil {
		t.Fatal(err)
	}
	if len(got) == 0 {
		t.Error("BLAKE2b should return non-empty hash")
	}
}

func TestBLAKE2s_Transform(t *testing.T) {
	got, err := processors.BLAKE2s{}.Transform([]byte("hello"))
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
