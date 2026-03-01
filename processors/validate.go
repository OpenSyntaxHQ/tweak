package processors

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/mcnijman/go-emailaddress"
	"mvdan.cc/xurls/v2"
)

type ValidateJSON struct{}

func (p ValidateJSON) Name() string    { return "validate-json" }
func (p ValidateJSON) Alias() []string { return []string{"is-json"} }
func (p ValidateJSON) Transform(data []byte, _ ...Flag) (string, error) {
	if len(strings.TrimSpace(string(data))) == 0 {
		return "❌ Invalid JSON: empty input", nil
	}
	if json.Valid(data) {
		return "✅ Valid JSON", nil
	}
	var js json.RawMessage
	err := json.Unmarshal(data, &js)
	return fmt.Sprintf("❌ Invalid JSON: %v", err), nil
}
func (p ValidateJSON) Flags() []Flag       { return nil }
func (p ValidateJSON) Title() string       { return fmt.Sprintf("Validate JSON (%s)", p.Name()) }
func (p ValidateJSON) Description() string { return "Check if input is valid JSON" }
func (p ValidateJSON) FilterValue() string { return p.Title() }

type ValidateEmail struct{}

func (p ValidateEmail) Name() string    { return "validate-email" }
func (p ValidateEmail) Alias() []string { return []string{"is-email"} }
func (p ValidateEmail) Transform(data []byte, _ ...Flag) (string, error) {
	input := strings.TrimSpace(string(data))
	if len(input) == 0 {
		return "❌ Invalid Email: empty input", nil
	}
	parsed, err := emailaddress.Parse(input)
	if err == nil && parsed.String() != "" && parsed.String() == input {
		return fmt.Sprintf("✅ Valid Email\n  Local:  %s\n  Domain: %s", parsed.LocalPart, parsed.Domain), nil
	}
	return "❌ Invalid Email format", nil
}
func (p ValidateEmail) Flags() []Flag       { return nil }
func (p ValidateEmail) Title() string       { return fmt.Sprintf("Validate Email (%s)", p.Name()) }
func (p ValidateEmail) Description() string { return "Check if input is a valid email address" }
func (p ValidateEmail) FilterValue() string { return p.Title() }

type ValidateURL struct{}

func (p ValidateURL) Name() string    { return "validate-url" }
func (p ValidateURL) Alias() []string { return []string{"is-url"} }
func (p ValidateURL) Transform(data []byte, _ ...Flag) (string, error) {
	input := strings.TrimSpace(string(data))
	if len(input) == 0 {
		return "❌ Invalid URL: empty input", nil
	}
	rxStrict := xurls.Strict()
	match := rxStrict.FindString(input)

	if match != "" && match == input {
		return "✅ Valid URL", nil
	}
	return "❌ Invalid URL format (must include scheme like http:// or https://)", nil
}
func (p ValidateURL) Flags() []Flag       { return nil }
func (p ValidateURL) Title() string       { return fmt.Sprintf("Validate URL (%s)", p.Name()) }
func (p ValidateURL) Description() string { return "Check if input is a valid URL" }
func (p ValidateURL) FilterValue() string { return p.Title() }
