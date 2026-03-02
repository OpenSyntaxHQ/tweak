package tests

import (
	"strings"
	"testing"

	"github.com/OpenSyntaxHQ/tweak/processors"
)

func TestFormatJSON_Transform(t *testing.T) {
	flags := []processors.Flag{
		{Name: "indent", Short: "i", Value: true, Type: processors.FlagBool},
	}
	got, err := (processors.FormatJSON{}).Transform([]byte(`{"a":1,"b":2}`), flags...)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got, "\n") {
		t.Errorf("formatted JSON should be multi-line, got: %q", got)
	}
}

func TestJSONMinify_Transform(t *testing.T) {
	input := "{\n  \"a\": 1,\n  \"b\": 2\n}"
	got, err := (processors.JSONMinify{}).Transform([]byte(input))
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(got, "\n") || strings.Contains(got, "  ") {
		t.Errorf("minified JSON should have no whitespace, got: %q", got)
	}
}

func TestJSONToYAML_Transform(t *testing.T) {
	got, err := (processors.JSONToYAML{}).Transform([]byte(`{"name":"tweak","version":1}`))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got, "name: tweak") {
		t.Errorf("expected 'name: tweak' in YAML, got: %q", got)
	}
}

func TestYAMLToJSON_Transform(t *testing.T) {
	input := "name: tweak\nversion: 1\n"
	got, err := (processors.YAMLToJSON{}).Transform([]byte(input))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got, "tweak") {
		t.Errorf("expected 'tweak' in JSON output, got: %q", got)
	}
}

func TestJSONEscape_Transform(t *testing.T) {
	got, err := (processors.JSONEscape{}).Transform([]byte(`{"key":"hello world"}`))
	if err != nil {
		t.Fatal(err)
	}
	if len(got) == 0 {
		t.Error("JSONEscape should produce non-empty output")
	}
}

func TestJSONUnescape_Transform(t *testing.T) {
	input := `{\"key\":\"value\"}`
	got, err := (processors.JSONUnescape{}).Transform([]byte(input))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got, "key") {
		t.Errorf("JSONUnescape should restore keys, got: %q", got)
	}
}

func TestJSONToCSV_Transform(t *testing.T) {
	input := `[{"name":"alice","age":30},{"name":"bob","age":25}]`
	got, err := (processors.JSONToCSV{}).Transform([]byte(input))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got, "alice") {
		t.Errorf("expected 'alice' in CSV output, got: %q", got)
	}
}

func TestCSVToJSON_Transform(t *testing.T) {
	input := "name,age\nalice,30\nbob,25"
	got, err := (processors.CSVToJSON{}).Transform([]byte(input))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got, "alice") {
		t.Errorf("expected 'alice' in JSON output, got: %q", got)
	}
}

func TestJSONToMSGPACK_RoundTrip(t *testing.T) {
	input := `{"name":"tweak"}`
	packed, err := (processors.JSONToMSGPACK{}).Transform([]byte(input))
	if err != nil {
		t.Fatal(err)
	}
	got, err := (processors.MSGPACKToJSON{}).Transform([]byte(packed))
	if err != nil {
		t.Fatalf("msgpack→json error: %v", err)
	}
	if !strings.Contains(got, "tweak") {
		t.Errorf("msgpack roundtrip lost 'tweak': %q", got)
	}
}

func TestFormatJSON_InvalidInput(t *testing.T) {
	if _, err := (processors.FormatJSON{}).Transform([]byte(`{"a":`)); err == nil {
		t.Fatal("expected format-json error for malformed input")
	}
}

func TestJSONEscapeAndUnescape_InvalidInputs(t *testing.T) {
	if _, err := (processors.JSONEscape{}).Transform([]byte(`not-json`)); err == nil {
		t.Fatal("expected json-escape error for invalid JSON")
	}
	if _, err := (processors.JSONUnescape{}).Transform([]byte(`\x`)); err == nil {
		t.Fatal("expected json-unescape error for invalid escaped payload")
	}
}

func TestJSONMinify_InvalidInput(t *testing.T) {
	if _, err := (processors.JSONMinify{}).Transform([]byte(`{`)); err == nil {
		t.Fatal("expected json-minify error for malformed JSON")
	}
}

func TestYAMLToJSON_IndentedOutput(t *testing.T) {
	flags := []processors.Flag{{Name: "indent", Short: "i", Value: true, Type: processors.FlagBool}}
	got, err := (processors.YAMLToJSON{}).Transform([]byte("name: tweak\ncount: 2"), flags...)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got, "\n") {
		t.Fatalf("expected indented JSON output, got %q", got)
	}
}

func TestJSONMSGPACK_InvalidInputBranches(t *testing.T) {
	if _, err := (processors.JSONToMSGPACK{}).Transform([]byte(`{"name":`)); err == nil {
		t.Fatal("expected json-msgpack error for invalid JSON")
	}
	if _, err := (processors.MSGPACKToJSON{}).Transform([]byte{0xc1}); err == nil {
		t.Fatal("expected msgpack-json error for invalid msgpack data")
	}
}

func TestCSVJSON_EdgeBranches(t *testing.T) {
	got, err := (processors.CSVToJSON{}).Transform([]byte("name,age"))
	if err != nil {
		t.Fatal(err)
	}
	if got != "[]" {
		t.Fatalf("expected [] for header-only csv, got %q", got)
	}

	if _, err := (processors.CSVToJSON{}).Transform([]byte("\"unterminated")); err == nil {
		t.Fatal("expected csv-json error for malformed CSV")
	}

	csv, err := (processors.JSONToCSV{}).Transform([]byte(`[]`))
	if err != nil {
		t.Fatal(err)
	}
	if csv != "" {
		t.Fatalf("expected empty output for empty JSON array, got %q", csv)
	}

	if _, err := (processors.JSONToCSV{}).Transform([]byte(`{}`)); err == nil {
		t.Fatal("expected json-csv error for non-array JSON input")
	}
}
