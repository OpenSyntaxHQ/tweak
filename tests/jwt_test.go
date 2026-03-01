package tests

import (
	"strings"
	"testing"

	"github.com/OpenSyntaxHQ/tweak/processors"
)

func TestJWT_RoundTrip(t *testing.T) {
	payload := `{"uid":42,"role":"admin"}`
	secret := "mysupersecret"

	encFlags := []processors.Flag{
		{Name: "secret", Short: "s", Value: secret, Type: processors.FlagString},
	}
	token, err := processors.JWTEncode{}.Transform([]byte(payload), encFlags...)
	if err != nil {
		t.Fatalf("jwt encode error: %v", err)
	}
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		t.Errorf("expected 3 JWT parts, got %d: %q", len(parts), token)
	}

	decoded, err := processors.JWTDecode{}.Transform([]byte(token))
	if err != nil {
		t.Fatalf("jwt decode error: %v", err)
	}
	if !strings.Contains(decoded, "uid") {
		t.Errorf("decoded JWT should contain 'uid', got: %q", decoded)
	}
	if !strings.Contains(decoded, "42") {
		t.Errorf("decoded JWT should contain '42', got: %q", decoded)
	}
}

func TestJWTDecode_InvalidToken(t *testing.T) {
	_, err := processors.JWTDecode{}.Transform([]byte("not.a.jwt"))
	_ = err
}
