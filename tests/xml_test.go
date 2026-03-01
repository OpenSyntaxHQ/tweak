package tests

import (
	"strings"
	"testing"

	"github.com/OpenSyntaxHQ/tweak/processors"
)

func TestXMLToJSON_Transform(t *testing.T) {
	input := `<root><name>tweak</name><version>1</version></root>`
	got, err := processors.XMLToJSON{}.Transform([]byte(input))
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
	got, err := processors.JSONToXML{}.Transform([]byte(input), flags...)
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
