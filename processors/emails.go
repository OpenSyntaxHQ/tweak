package processors

import (
	"fmt"
	"strings"

	"github.com/mcnijman/go-emailaddress"
)

type ExtractEmails struct{}

func (p ExtractEmails) Name() string    { return "extract-emails" }
func (p ExtractEmails) Alias() []string { return []string{"find-emails", "find-email", "extract-email"} }
func (p ExtractEmails) Transform(data []byte, f ...Flag) (string, error) {
	var emails []string
	extracted := emailaddress.FindWithIcannSuffix(data, false)
	for _, e := range extracted {
		emails = append(emails, e.String())
	}
	separator := "\n"
	for _, flag := range f {
		if flag.Short == "s" {
			if x, ok := flag.Value.(string); ok && x != "" {
				separator = x
			}
		}
	}
	return strings.Join(emails, separator), nil
}
func (p ExtractEmails) Flags() []Flag {
	return []Flag{{Name: "separator", Short: "s", Desc: "Separator between emails", Value: "", Type: FlagString}}
}
func (p ExtractEmails) Title() string       { return fmt.Sprintf("Extract Emails (%s)", p.Name()) }
func (p ExtractEmails) Description() string { return "Extract emails from given text" }
func (p ExtractEmails) FilterValue() string { return p.Title() }
