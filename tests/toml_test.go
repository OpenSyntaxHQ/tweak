package tests

import (
	"strings"
	"testing"

	"github.com/OpenSyntaxHQ/tweak/processors"
)

func TestTOMLToJSON_Transform(t *testing.T) {
	input := "[server]\nhost = \"localhost\"\nport = 8080"
	got, err := (processors.TOMLToJSON{}).Transform([]byte(input))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got, "localhost") {
		t.Errorf("expected 'localhost' in output, got: %q", got)
	}
	if !strings.Contains(got, "8080") {
		t.Errorf("expected '8080' in output, got: %q", got)
	}
}

func TestJSONToTOML_Transform(t *testing.T) {
	input := `{"server":{"host":"localhost","port":8080}}`
	got, err := (processors.JSONToTOML{}).Transform([]byte(input))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got, "localhost") {
		t.Errorf("expected 'localhost' in output, got: %q", got)
	}
}

func TestTOMLJSON_RoundTrip(t *testing.T) {
	toml := "[db]\nname = \"mydb\""
	json, err := (processors.TOMLToJSON{}).Transform([]byte(toml))
	if err != nil {
		t.Fatal(err)
	}
	back, err := (processors.JSONToTOML{}).Transform([]byte(json))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(back, "mydb") {
		t.Errorf("roundtrip lost 'mydb': %q", back)
	}
}

func TestTOMLJSON_InvalidInputsAndIndent(t *testing.T) {
	if _, err := (processors.TOMLToJSON{}).Transform([]byte("invalid = [")); err == nil {
		t.Fatal("expected toml-json error for malformed TOML")
	}

	flags := []processors.Flag{{Name: "indent", Short: "i", Value: true, Type: processors.FlagBool}}
	got, err := (processors.TOMLToJSON{}).Transform([]byte("a = 1"), flags...)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got, "\n") {
		t.Fatalf("expected indented JSON from toml-json, got %q", got)
	}

	if _, err := (processors.JSONToTOML{}).Transform([]byte(`{"a":`)); err == nil {
		t.Fatal("expected json-toml error for malformed JSON")
	}
}
