package tests

import (
	"strings"
	"testing"

	"github.com/OpenSyntaxHQ/tweak/processors"
)

func TestTOTP_Format(t *testing.T) {
	flags := []processors.Flag{
		{Name: "digits", Short: "d", Value: uint(6), Type: processors.FlagUint},
		{Name: "period", Short: "p", Value: uint(30), Type: processors.FlagUint},
	}
	got, err := processors.TOTP{}.Transform([]byte("JBSWY3DPEHPK3PXP"), flags...)
	if err != nil {
		t.Fatalf("totp error: %v", err)
	}
	if len(got) != 6 {
		t.Errorf("expected 6-digit code, got %q (len=%d)", got, len(got))
	}
	for _, c := range got {
		if c < '0' || c > '9' {
			t.Errorf("TOTP code should be numeric, got %q", got)
			break
		}
	}
}

func TestTOTP_InvalidSecret(t *testing.T) {
	_, err := processors.TOTP{}.Transform([]byte("not-valid-base32!!!"))
	if err == nil {
		t.Error("expected error for invalid base32 secret")
	}
	_ = strings.Contains("", "")
}
