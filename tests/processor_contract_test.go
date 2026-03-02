package tests

import (
	"bytes"
	"strings"
	"testing"

	"github.com/OpenSyntaxHQ/tweak/processors"
)

type invalidNativeDeclaration struct{}

func (invalidNativeDeclaration) Name() string { return "invalid-native" }
func (invalidNativeDeclaration) Alias() []string {
	return nil
}
func (invalidNativeDeclaration) Transform(data []byte, _ ...processors.Flag) (string, error) {
	return string(data), nil
}
func (invalidNativeDeclaration) Flags() []processors.Flag {
	return nil
}
func (invalidNativeDeclaration) StreamingSpec() processors.StreamingSpec {
	return processors.StreamingSpec{Mode: processors.StreamingModeNative}
}

type chunkedZeroSize struct{}

func (chunkedZeroSize) Name() string { return "chunked-zero-size" }
func (chunkedZeroSize) Alias() []string {
	return nil
}
func (chunkedZeroSize) Transform(data []byte, _ ...processors.Flag) (string, error) {
	return string(data), nil
}
func (chunkedZeroSize) Flags() []processors.Flag { return nil }
func (chunkedZeroSize) StreamingSpec() processors.StreamingSpec {
	return processors.StreamingSpec{Mode: processors.StreamingModeChunked, ChunkSize: 0}
}

func TestProcessorRegistryContracts(t *testing.T) {
	knownGenerators := map[string]bool{
		"lorem":    true,
		"nanoid":   true,
		"now":      true,
		"password": true,
		"ulid":     true,
		"uuid":     true,
	}
	seenNames := make(map[string]struct{}, len(processors.List))

	for i, item := range processors.List {
		p, ok := item.(processors.Processor)
		if !ok {
			t.Fatalf("registry item %d (%T) does not implement processors.Processor", i, item)
		}

		name := p.Name()
		if name == "" {
			t.Fatalf("registry item %d has empty Name()", i)
		}
		if _, exists := seenNames[name]; exists {
			t.Fatalf("duplicate processor name %q", name)
		}
		seenNames[name] = struct{}{}

		meta, ok := item.(interface {
			Title() string
			Description() string
			FilterValue() string
		})
		if !ok {
			t.Fatalf("processor %q does not implement list metadata methods", name)
		}
		title := meta.Title()
		desc := meta.Description()
		filterValue := meta.FilterValue()
		if title == "" {
			t.Fatalf("processor %q has empty Title()", name)
		}
		if desc == "" {
			t.Fatalf("processor %q has empty Description()", name)
		}
		if filterValue == "" {
			t.Fatalf("processor %q has empty FilterValue()", name)
		}
		if filterValue != title {
			t.Fatalf("processor %q must return FilterValue() == Title()", name)
		}

		aliasSet := make(map[string]struct{})
		for _, alias := range p.Alias() {
			if alias == "" {
				t.Fatalf("processor %q has empty alias", name)
			}
			if alias == name {
				t.Fatalf("processor %q alias duplicates name", name)
			}
			if _, exists := aliasSet[alias]; exists {
				t.Fatalf("processor %q has duplicate alias %q", name, alias)
			}
			aliasSet[alias] = struct{}{}
		}

		flags := p.Flags()
		validateFlags(t, name, flags)

		if hasRequiredFlag(flags) {
			if _, err := p.Transform([]byte("test-input-without-required-flags")); err == nil {
				t.Fatalf("processor %q has required flags but Transform succeeded without them", name)
			}
		}

		validateStreamingSpec(t, p)

		if got, want := processors.IsGenerator(p), knownGenerators[name]; got != want {
			t.Fatalf("processor %q IsGenerator()=%v, want %v", name, got, want)
		}
	}

	if got, want := len(seenNames), len(processors.List); got != want {
		t.Fatalf("registry name count mismatch: got %d, want %d", got, want)
	}
}

func validateFlags(t *testing.T, processorName string, flags []processors.Flag) {
	t.Helper()

	nameSet := make(map[string]struct{}, len(flags))
	shortSet := make(map[string]struct{}, len(flags))

	for _, flag := range flags {
		if flag.Name == "" {
			t.Fatalf("processor %q has a flag with empty Name", processorName)
		}
		if flag.Desc == "" {
			t.Fatalf("processor %q flag %q has empty Desc", processorName, flag.Name)
		}
		if flag.Short == "" {
			t.Fatalf("processor %q flag %q has empty Short", processorName, flag.Name)
		}
		if _, exists := nameSet[flag.Name]; exists {
			t.Fatalf("processor %q has duplicate flag name %q", processorName, flag.Name)
		}
		nameSet[flag.Name] = struct{}{}
		if _, exists := shortSet[flag.Short]; exists {
			t.Fatalf("processor %q has duplicate flag short %q", processorName, flag.Short)
		}
		shortSet[flag.Short] = struct{}{}

		switch flag.Type {
		case processors.FlagString:
			if _, ok := flag.Value.(string); !ok {
				t.Fatalf("processor %q flag %q has non-string default for FlagString", processorName, flag.Name)
			}
		case processors.FlagBool:
			if _, ok := flag.Value.(bool); !ok {
				t.Fatalf("processor %q flag %q has non-bool default for FlagBool", processorName, flag.Name)
			}
		case processors.FlagUint:
			if _, ok := flag.Value.(uint); !ok {
				t.Fatalf("processor %q flag %q has non-uint default for FlagUint", processorName, flag.Name)
			}
		case processors.FlagInt:
			if _, ok := flag.Value.(int); !ok {
				t.Fatalf("processor %q flag %q has non-int default for FlagInt", processorName, flag.Name)
			}
		default:
			t.Fatalf("processor %q flag %q has unsupported type %q", processorName, flag.Name, flag.Type)
		}
	}
}

func hasRequiredFlag(flags []processors.Flag) bool {
	for _, flag := range flags {
		if flag.Required {
			return true
		}
	}
	return false
}

func validateStreamingSpec(t *testing.T, p processors.Processor) {
	t.Helper()

	spec := processors.GetStreamingSpec(p)
	switch spec.Mode {
	case processors.StreamingModeNone,
		processors.StreamingModeBuffered,
		processors.StreamingModeChunked,
		processors.StreamingModeLine,
		processors.StreamingModeNative:
	default:
		t.Fatalf("processor %q has invalid StreamingSpec.Mode %q", p.Name(), spec.Mode)
	}

	if spec.Mode == processors.StreamingModeChunked && spec.ChunkSize <= 0 {
		t.Fatalf("processor %q has chunked streaming with non-positive chunk size", p.Name())
	}

	if _, isNative := p.(processors.NativeStreamProcessor); isNative && spec.Mode != processors.StreamingModeNative {
		t.Fatalf("native stream processor %q must use native mode", p.Name())
	}
}

func TestStreamingSpecNormalizationAndInvalidNativeDeclaration(t *testing.T) {
	chunked := processors.GetStreamingSpec(chunkedZeroSize{})
	if chunked.Mode != processors.StreamingModeChunked {
		t.Fatalf("expected chunked mode, got %q", chunked.Mode)
	}
	if chunked.ChunkSize <= 0 {
		t.Fatalf("expected normalized chunk size > 0, got %d", chunked.ChunkSize)
	}

	invalid := invalidNativeDeclaration{}
	spec := processors.GetStreamingSpec(invalid)
	if spec.Mode != processors.StreamingModeNone {
		t.Fatalf("invalid native declaration should normalize to none mode, got %q", spec.Mode)
	}
	if processors.SupportsStreaming(invalid) {
		t.Fatalf("invalid native declaration must not be treated as streamable")
	}

	var out bytes.Buffer
	err := processors.TransformStream(invalid, strings.NewReader("abc"), &out)
	if err == nil {
		t.Fatal("expected TransformStream error for invalid native declaration")
	}
}

func TestContractsHelpers(t *testing.T) {
	if !processors.IsGenerator(processors.Now{}) {
		t.Fatal("Now must satisfy generator contract")
	}
	if processors.IsGenerator(processors.Upper{}) {
		t.Fatal("Upper must not satisfy generator contract")
	}

	if got, want := processors.FlagString.String(), "String"; got != want {
		t.Fatalf("FlagString.String()=%q, want %q", got, want)
	}
	if !processors.FlagString.IsString() {
		t.Fatal("FlagString.IsString() should be true")
	}
	if processors.FlagBool.IsString() {
		t.Fatal("FlagBool.IsString() should be false")
	}

	required := processors.Flag{Name: "key", Required: true}
	if got, want := required.HelpLabel(), "--key (required)"; got != want {
		t.Fatalf("HelpLabel()=%q, want %q", got, want)
	}
	optional := processors.Flag{Name: "indent"}
	if got, want := optional.HelpLabel(), "--indent"; got != want {
		t.Fatalf("HelpLabel()=%q, want %q", got, want)
	}
}
