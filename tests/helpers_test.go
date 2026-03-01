// Package tests provides black-box integration tests for all tweak processors.
// All tests import the processors package and test only the public API.
package tests

import (
	"testing"

	"github.com/OpenSyntaxHQ/tweak/processors"
)

// assertTransform is a helper that runs a processor's Transform and checks output.
func assertTransform(t *testing.T, p interface {
	Transform([]byte, ...processors.Flag) (string, error)
}, input string, flags []processors.Flag, want string, wantErr bool) {
	t.Helper()
	got, err := p.Transform([]byte(input), flags...)
	if (err != nil) != wantErr {
		t.Fatalf("error = %v, wantErr %v", err, wantErr)
	}
	if !wantErr && got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

// assertCommandMeta checks the metadata methods of a processor.
func assertCommandMeta(t *testing.T, p interface {
	Name() string
	Title() string
	Description() string
	FilterValue() string
}, name, title, description string) {
	t.Helper()
	if got := p.Name(); got != name {
		t.Errorf("Name() = %q, want %q", got, name)
	}
	if got := p.Title(); got != title {
		t.Errorf("Title() = %q, want %q", got, title)
	}
	if got := p.Description(); got != description {
		t.Errorf("Description() = %q, want %q", got, description)
	}
	if got := p.FilterValue(); got != title {
		t.Errorf("FilterValue() = %q, want %q (title)", got, title)
	}
}
