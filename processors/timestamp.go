package processors

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Epoch struct{}

func (p Epoch) Name() string    { return "epoch" }
func (p Epoch) Alias() []string { return []string{"timestamp", "ts"} }
func (p Epoch) Transform(data []byte, f ...Flag) (string, error) {
	format := "2006-01-02 15:04:05 MST"
	tz := "Local"
	for _, flag := range f {
		switch flag.Short {
		case "f":
			if s, ok := flag.Value.(string); ok && s != "" {
				format = s
			}
		case "z":
			if s, ok := flag.Value.(string); ok && s != "" {
				tz = s
			}
		}
	}
	input := strings.TrimSpace(string(data))
	loc, err := time.LoadLocation(tz)
	if err != nil {
		loc = time.Local
	}
	if epoch, err := strconv.ParseInt(input, 10, 64); err == nil {
		var t time.Time
		switch {
		case epoch > 1e15:
			t = time.UnixMicro(epoch)
		case epoch > 1e12:
			t = time.UnixMilli(epoch)
		default:
			t = time.Unix(epoch, 0)
		}
		return t.In(loc).Format(format), nil
	}
	layouts := []string{
		time.RFC3339, time.RFC3339Nano,
		"2006-01-02 15:04:05", "2006-01-02",
		time.RFC1123, time.RFC1123Z, time.RFC822, time.RFC822Z,
	}
	for _, layout := range layouts {
		if t, err := time.ParseInLocation(layout, input, loc); err == nil {
			return fmt.Sprintf("%d", t.Unix()), nil
		}
	}
	return "", fmt.Errorf("cannot parse '%s' as epoch or date", input)
}
func (p Epoch) Flags() []Flag {
	return []Flag{
		{Name: "format", Short: "f", Desc: "Output date format (Go layout)", Value: "2006-01-02 15:04:05 MST", Type: FlagString},
		{Name: "timezone", Short: "z", Desc: "Timezone (e.g. UTC, America/New_York)", Value: "Local", Type: FlagString},
	}
}
func (p Epoch) Title() string {
	return fmt.Sprintf("%s (%s)", cases.Title(language.Und, cases.NoLower).String(p.Name()), p.Name())
}
func (p Epoch) Description() string { return "Convert epoch ↔ human-readable timestamp" }
func (p Epoch) FilterValue() string { return p.Title() }

type Now struct{}

func (p Now) Name() string      { return "now" }
func (p Now) Alias() []string   { return []string{"time", "date"} }
func (p Now) IsGenerator() bool { return true }
func (p Now) Transform(_ []byte, f ...Flag) (string, error) {
	format := "2006-01-02 15:04:05 MST"
	utc := false
	for _, flag := range f {
		switch flag.Short {
		case "f":
			if s, ok := flag.Value.(string); ok && s != "" {
				format = s
			}
		case "u":
			if b, ok := flag.Value.(bool); ok {
				utc = b
			}
		}
	}
	t := time.Now()
	if utc {
		t = t.UTC()
	}
	return t.Format(format), nil
}
func (p Now) Flags() []Flag {
	return []Flag{
		{Name: "format", Short: "f", Desc: "Output format (Go layout or 'epoch')", Value: "2006-01-02 15:04:05 MST", Type: FlagString},
		{Name: "utc", Short: "u", Desc: "Use UTC timezone", Value: false, Type: FlagBool},
	}
}
func (p Now) Title() string {
	return fmt.Sprintf("%s (%s)", cases.Title(language.Und, cases.NoLower).String(p.Name()), p.Name())
}
func (p Now) Description() string { return "Print current timestamp" }
func (p Now) FilterValue() string { return p.Title() }
