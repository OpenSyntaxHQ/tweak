package tests

import (
	"strings"
	"testing"

	"github.com/OpenSyntaxHQ/tweak/processors"
)

func TestXMLToJSON_Transform(t *testing.T) {
	input := `<root><name>tweak</name><version>1</version></root>`
	got, err := (processors.XMLToJSON{}).Transform([]byte(input))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got, "tweak") {
		t.Errorf("expected 'tweak' in output, got: %q", got)
	}
}

func TestJSONToXML_Transform(t *testing.T) {
	input := `{"name":"tweak","version":1}`
	flags := []processors.Flag{
		{Name: "root", Short: "r", Value: "root", Type: processors.FlagString},
	}
	got, err := (processors.JSONToXML{}).Transform([]byte(input), flags...)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got, "tweak") {
		t.Errorf("expected 'tweak' in XML output, got: %q", got)
	}
	if !strings.Contains(got, "<") {
		t.Errorf("expected XML tags in output, got: %q", got)
	}
}

func TestXMLToJSON_InvalidXML(t *testing.T) {
	if _, err := (processors.XMLToJSON{}).Transform([]byte(`<root><broken>`)); err == nil {
		t.Fatal("expected xml-json error for malformed XML")
	}
}

func TestJSONToXML_IndentAndInvalidJSON(t *testing.T) {
	flags := []processors.Flag{
		{Name: "indent", Short: "i", Value: true, Type: processors.FlagBool},
		{Name: "root", Short: "r", Value: "payload", Type: processors.FlagString},
	}
	got, err := (processors.JSONToXML{}).Transform([]byte(`{"a":1}`), flags...)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got, "\n") {
		t.Fatalf("expected indented XML output, got %q", got)
	}

	if _, err := (processors.JSONToXML{}).Transform([]byte(`{"a":`)); err == nil {
		t.Fatal("expected json-xml error for malformed JSON")
	}
}
