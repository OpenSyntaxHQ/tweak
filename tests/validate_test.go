package tests

import (
	"strings"
	"testing"

	"github.com/OpenSyntaxHQ/tweak/processors"
)

func TestValidateJSON_Valid(t *testing.T) {
	inputs := []string{
		`{"key":"value"}`,
		`[1,2,3]`,
		`null`,
		`true`,
		`42`,
	}
	for _, in := range inputs {
		got, err := processors.ValidateJSON{}.Transform([]byte(in))
		if err != nil {
			t.Errorf("ValidateJSON(%q) error: %v", in, err)
		}
		if !strings.Contains(got, "Valid") {
			t.Errorf("ValidateJSON(%q) = %q, want Valid", in, got)
		}
	}
}

func TestValidateJSON_Invalid(t *testing.T) {
	inputs := []string{`not json`, `{bad}`, `[1,2,`}
	for _, in := range inputs {
		_, err := processors.ValidateJSON{}.Transform([]byte(in))
		if err == nil {
			t.Errorf("ValidateJSON(%q) expected an error for invalid input", in)
		}
	}
}

func TestValidateEmail_Valid(t *testing.T) {
	inputs := []string{"user@example.com", "a@b.io", "test+tag@domain.co.uk"}
	for _, in := range inputs {
		got, err := processors.ValidateEmail{}.Transform([]byte(in))
		if err != nil {
			t.Errorf("ValidateEmail(%q) error: %v", in, err)
		}
		if !strings.Contains(got, "Valid") {
			t.Errorf("ValidateEmail(%q) = %q, want Valid", in, got)
		}
	}
}

func TestValidateEmail_Invalid(t *testing.T) {
	inputs := []string{"not-an-email", "@no-local.com", "no-at-sign"}
	for _, in := range inputs {
		_, err := processors.ValidateEmail{}.Transform([]byte(in))
		if err == nil {
			t.Errorf("ValidateEmail(%q) expected an error for invalid input", in)
		}
	}
}

func TestValidateURL_Valid(t *testing.T) {
	inputs := []string{"https://github.com", "http://example.com/path?q=1"}
	for _, in := range inputs {
		got, err := processors.ValidateURL{}.Transform([]byte(in))
		if err != nil {
			t.Errorf("ValidateURL(%q) error: %v", in, err)
		}
		if !strings.Contains(got, "Valid") {
			t.Errorf("ValidateURL(%q) = %q, want Valid", in, got)
		}
	}
}

func TestValidateURL_Invalid(t *testing.T) {
	inputs := []string{"not a url", "ftp://", "just text"}
	for _, in := range inputs {
		_, err := processors.ValidateURL{}.Transform([]byte(in))
		if err == nil {
			t.Errorf("ValidateURL(%q) expected an error for invalid input", in)
		}
	}
}
