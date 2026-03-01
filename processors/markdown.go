package processors

import (
	"bytes"
	"fmt"

	"github.com/yuin/goldmark"
)

type Markdown struct{}

func (p Markdown) Name() string    { return "markdown-html" }
func (p Markdown) Alias() []string { return []string{"md-html"} }
func (p Markdown) Transform(data []byte, _ ...Flag) (string, error) {
	var buf bytes.Buffer
	if err := goldmark.Convert(data, &buf); err != nil {
		return "", err
	}
	return buf.String(), nil
}
func (p Markdown) Flags() []Flag       { return nil }
func (p Markdown) Title() string       { return fmt.Sprintf("Markdown to HTML (%s)", p.Name()) }
func (p Markdown) Description() string { return "Convert Markdown to HTML" }
func (p Markdown) FilterValue() string { return p.Title() }
