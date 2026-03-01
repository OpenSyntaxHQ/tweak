package tests

import (
	"testing"

	"github.com/OpenSyntaxHQ/tweak/processors"
)

func TestAES_RoundTrip(t *testing.T) {
	key := "mysecretpassword"
	plaintext := "Hello, World!"

	encFlags := []processors.Flag{{Name: "key", Short: "k", Value: key, Type: processors.FlagString}}
	encrypted, err := processors.AESEncrypt{}.Transform([]byte(plaintext), encFlags...)
	if err != nil {
		t.Fatalf("AES encrypt error: %v", err)
	}
	if encrypted == plaintext {
		t.Error("encrypted should differ from plaintext")
	}

	decFlags := []processors.Flag{{Name: "key", Short: "k", Value: key, Type: processors.FlagString}}
	decrypted, err := processors.AESDecrypt{}.Transform([]byte(encrypted), decFlags...)
	if err != nil {
		t.Fatalf("AES decrypt error: %v", err)
	}
	if decrypted != plaintext {
		t.Errorf("roundtrip: got %q, want %q", decrypted, plaintext)
	}
}

func TestAESDecrypt_BadKey(t *testing.T) {
	key := "mysecretpassword"
	encFlags := []processors.Flag{{Name: "key", Short: "k", Value: key, Type: processors.FlagString}}
	encrypted, err := processors.AESEncrypt{}.Transform([]byte("secret"), encFlags...)
	if err != nil {
		t.Fatal(err)
	}

	badFlags := []processors.Flag{{Name: "key", Short: "k", Value: "wrongkey", Type: processors.FlagString}}
	_, err = processors.AESDecrypt{}.Transform([]byte(encrypted), badFlags...)
	if err == nil {
		t.Error("expected error decrypting with wrong key")
	}
}
