package processors

import (
	"crypto/rand"
	"fmt"
	"time"

	"github.com/oklog/ulid/v2"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type ULID struct{}

func (p ULID) Name() string      { return "ulid" }
func (p ULID) Alias() []string   { return []string{"gen-ulid"} }
func (p ULID) IsGenerator() bool { return true }
func (p ULID) Transform(_ []byte, _ ...Flag) (string, error) {
	id, err := ulid.New(ulid.Timestamp(time.Now()), rand.Reader)
	if err != nil {
		return "", err
	}
	return id.String(), nil
}
func (p ULID) Flags() []Flag       { return nil }
func (p ULID) Title() string {
	return fmt.Sprintf("%s (%s)", cases.Title(language.Und, cases.NoLower).String(p.Name()), p.Name())
}
func (p ULID) Description() string { return "Generate a ULID" }
func (p ULID) FilterValue() string { return p.Title() }
