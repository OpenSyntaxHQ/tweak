package tests

import (
	"strings"
	"testing"

	"github.com/OpenSyntaxHQ/tweak/processors"
)

func TestJSONLToJSON_Transform(t *testing.T) {
	input := "{\"id\":1}\n{\"id\":2}\n{\"id\":3}"
	got, err := processors.JSONLToJSON{}.Transform([]byte(input))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(got, "[") {
		t.Errorf("expected JSON array output, got: %q", got)
	}
	if !strings.Contains(got, `"id"`) {
		t.Errorf("expected 'id' field in output, got: %q", got)
	}
}

func TestJSONLToJSON_Empty(t *testing.T) {
	_, err := processors.JSONLToJSON{}.Transform([]byte(""))
	_ = err
}
