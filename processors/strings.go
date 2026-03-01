package processors

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/OpenSyntaxHQ/tweak/utils"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/iancoleman/strcase"
)

var whitespaceNormRe = regexp.MustCompile(`\s+`)

type Lower struct{}

func (p Lower) GetStreamingConfig() StreamingConfig {
	return StreamingConfig{ChunkSize: 64 * 1024, BufferOutput: false}
}
func (p Lower) Name() string                        { return "lower" }
func (p Lower) Alias() []string                     { return nil }
func (p Lower) Transform(data []byte, _ ...Flag) (string, error) { return strings.ToLower(string(data)), nil }
func (p Lower) Flags() []Flag                       { return nil }
func (p Lower) Title() string                       { return fmt.Sprintf("To Lower case (%s)", p.Name()) }
func (p Lower) Description() string                 { return "Transform your text to lower case" }
func (p Lower) FilterValue() string                 { return p.Title() }

type Upper struct{}

func (p Upper) GetStreamingConfig() StreamingConfig {
	return StreamingConfig{ChunkSize: 64 * 1024, BufferOutput: false}
}
func (p Upper) Name() string                        { return "upper" }
func (p Upper) Alias() []string                     { return nil }
func (p Upper) Transform(data []byte, _ ...Flag) (string, error) { return strings.ToUpper(string(data)), nil }
func (p Upper) Flags() []Flag                       { return nil }
func (p Upper) Title() string                       { return fmt.Sprintf("To Upper case (%s)", p.Name()) }
func (p Upper) Description() string                 { return "Transform your text to UPPER CASE" }
func (p Upper) FilterValue() string                 { return p.Title() }

type Title struct{}

func (p Title) Name() string    { return "title" }
func (p Title) Alias() []string { return nil }
func (p Title) Transform(data []byte, _ ...Flag) (string, error) {
	return cases.Title(language.Und, cases.NoLower).String(string(data)), nil
}
func (p Title) Flags() []Flag       { return nil }
func (p Title) Title() string       { return fmt.Sprintf("To Title Case (%s)", p.Name()) }
func (p Title) Description() string { return "Transform your text to Title Case" }
func (p Title) FilterValue() string { return p.Title() }

type Snake struct{}

func (p Snake) Name() string    { return "snake" }
func (p Snake) Alias() []string { return nil }
func (p Snake) Transform(data []byte, _ ...Flag) (string, error) {
	str := whitespaceNormRe.ReplaceAllString(string(data), " ")
	return strcase.ToSnake(str), nil
}
func (p Snake) Flags() []Flag       { return nil }
func (p Snake) Title() string       { return fmt.Sprintf("To Snake case (%s)", p.Name()) }
func (p Snake) Description() string { return "Transform your text to snake_case" }
func (p Snake) FilterValue() string { return p.Title() }

type Kebab struct{}

func (p Kebab) Name() string    { return "kebab" }
func (p Kebab) Alias() []string { return nil }
func (p Kebab) Transform(data []byte, _ ...Flag) (string, error) {
	return utils.ToKebabCase(data), nil
}
func (p Kebab) Flags() []Flag       { return nil }
func (p Kebab) Title() string       { return fmt.Sprintf("To Kebab case (%s)", p.Name()) }
func (p Kebab) Description() string { return "Transform your text to kebab-case" }
func (p Kebab) FilterValue() string { return p.Title() }

type Camel struct{}

func (p Camel) Name() string    { return "camel" }
func (p Camel) Alias() []string { return nil }
func (p Camel) Transform(data []byte, _ ...Flag) (string, error) {
	str := whitespaceNormRe.ReplaceAllString(string(data), " ")
	return strcase.ToLowerCamel(str), nil
}
func (p Camel) Flags() []Flag       { return nil }
func (p Camel) Title() string       { return fmt.Sprintf("To Camel case (%s)", p.Name()) }
func (p Camel) Description() string { return "Transform your text to camelCase" }
func (p Camel) FilterValue() string { return p.Title() }

type Pascal struct{}

func (p Pascal) Name() string    { return "pascal" }
func (p Pascal) Alias() []string { return nil }
func (p Pascal) Transform(data []byte, _ ...Flag) (string, error) {
	str := whitespaceNormRe.ReplaceAllString(string(data), " ")
	return strcase.ToCamel(str), nil
}
func (p Pascal) Flags() []Flag       { return nil }
func (p Pascal) Title() string       { return fmt.Sprintf("To Pascal case (%s)", p.Name()) }
func (p Pascal) Description() string { return "Transform your text to PascalCase" }
func (p Pascal) FilterValue() string { return p.Title() }

type Slug struct{}

func (p Slug) Name() string    { return "slug" }
func (p Slug) Alias() []string { return nil }
func (p Slug) Transform(data []byte, _ ...Flag) (string, error) {
	re := regexp.MustCompile("[^a-z0-9]+")
	return strings.Trim(re.ReplaceAllString(strings.ToLower(string(data)), "-"), "-"), nil
}
func (p Slug) Flags() []Flag       { return nil }
func (p Slug) Title() string       { return fmt.Sprintf("To Slug case (%s)", p.Name()) }
func (p Slug) Description() string { return "Transform your text to slug-case" }
func (p Slug) FilterValue() string { return p.Title() }

type Reverse struct{}

func (p Reverse) Name() string    { return "reverse" }
func (p Reverse) Alias() []string { return nil }
func (p Reverse) Transform(data []byte, _ ...Flag) (string, error) {
	runes := []rune(string(data))
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes), nil
}
func (p Reverse) Flags() []Flag       { return nil }
func (p Reverse) Title() string       { return fmt.Sprintf("Reverse text (%s)", p.Name()) }
func (p Reverse) Description() string { return "Reverse Text ( txeT esreveR )" }
func (p Reverse) FilterValue() string { return p.Title() }

type EscapeQuotes struct{}

func (p EscapeQuotes) Name() string    { return "escape-quotes" }
func (p EscapeQuotes) Alias() []string { return []string{"esc-quotes"} }

func (p EscapeQuotes) Transform(data []byte, f ...Flag) (string, error) {
	var result strings.Builder
	for _, v := range data {
		for _, flag := range f {
			switch flag.Short {
			case "d":
				if v == '"' {
					result.WriteString("\\")
				}
			case "s":
				if v == '\'' {
					result.WriteString("\\")
				}
			}
		}
		if len(f) == 0 {
			if v == '"' || v == '\'' {
				result.WriteString("\\")
			}
		}
		result.WriteByte(v)
	}
	return result.String(), nil
}

func (p EscapeQuotes) Flags() []Flag {
	return []Flag{
		{Name: "double-quote", Short: "d", Desc: "Escape double quote", Value: true, Type: FlagBool},
		{Name: "single-quote", Short: "s", Desc: "Escape single quote", Value: true, Type: FlagBool},
	}
}
func (p EscapeQuotes) Title() string       { return fmt.Sprintf("Escape Quotes (%s)", p.Name()) }
func (p EscapeQuotes) Description() string { return "Escapes single and double quotes by default" }
func (p EscapeQuotes) FilterValue() string { return p.Title() }

type CountCharacters struct{}

func (p CountCharacters) GetStreamingConfig() StreamingConfig {
	return StreamingConfig{ChunkSize: 64 * 1024, BufferOutput: true}
}
func (p CountCharacters) Name() string    { return "count-chars" }
func (p CountCharacters) Alias() []string { return nil }
func (p CountCharacters) Transform(data []byte, _ ...Flag) (string, error) {
	return fmt.Sprintf("%d", len([]rune(string(data)))), nil
}
func (p CountCharacters) Flags() []Flag       { return nil }
func (p CountCharacters) Title() string       { return fmt.Sprintf("Count Number of Characters (%s)", p.Name()) }
func (p CountCharacters) Description() string { return "Find the length of your text (including spaces)" }
func (p CountCharacters) FilterValue() string { return p.Title() }

type CountWords struct{}

func (p CountWords) GetStreamingConfig() StreamingConfig {
	return StreamingConfig{ChunkSize: 64 * 1024, BufferOutput: true}
}
func (p CountWords) Name() string    { return "count-words" }
func (p CountWords) Alias() []string { return nil }
func (p CountWords) Transform(data []byte, _ ...Flag) (string, error) {
	return fmt.Sprintf("%d", len(strings.Fields(string(data)))), nil
}
func (p CountWords) Flags() []Flag       { return nil }
func (p CountWords) Title() string       { return fmt.Sprintf("Count Number of Words (%s)", p.Name()) }
func (p CountWords) Description() string { return "Count the number of words in your text" }
func (p CountWords) FilterValue() string { return p.Title() }

type Trim struct{}

func (p Trim) Name() string    { return "trim" }
func (p Trim) Alias() []string { return []string{"strip"} }
func (p Trim) Transform(data []byte, _ ...Flag) (string, error) {
	return strings.TrimSpace(string(data)), nil
}
func (p Trim) Flags() []Flag       { return nil }
func (p Trim) Title() string       { return fmt.Sprintf("Trim Whitespace (%s)", p.Name()) }
func (p Trim) Description() string { return "Trim leading and trailing whitespace" }
func (p Trim) FilterValue() string { return p.Title() }

type Repeat struct{}

func (p Repeat) Name() string    { return "repeat" }
func (p Repeat) Alias() []string { return []string{"rep"} }
func (p Repeat) Transform(data []byte, f ...Flag) (string, error) {
	count := uint(2)
	for _, flag := range f {
		if flag.Short == "c" {
			if c, ok := flag.Value.(uint); ok {
				count = c
			}
		}
	}
	return strings.Repeat(string(data), int(count)), nil
}
func (p Repeat) Flags() []Flag {
	return []Flag{{Name: "count", Short: "c", Desc: "Number of repetitions", Value: uint(2), Type: FlagUint}}
}
func (p Repeat) Title() string       { return fmt.Sprintf("Repeat Text (%s)", p.Name()) }
func (p Repeat) Description() string { return "Repeat text N times" }
func (p Repeat) FilterValue() string { return p.Title() }

type Wrap struct{}

func (p Wrap) Name() string    { return "wrap" }
func (p Wrap) Alias() []string { return []string{"word-wrap"} }
func (p Wrap) Transform(data []byte, f ...Flag) (string, error) {
	width := uint(80)
	for _, flag := range f {
		if flag.Short == "w" {
			if w, ok := flag.Value.(uint); ok {
				width = w
			}
		}
	}
	return wordWrap(string(data), int(width)), nil
}
func (p Wrap) Flags() []Flag {
	return []Flag{{Name: "width", Short: "w", Desc: "Maximum line width", Value: uint(80), Type: FlagUint}}
}
func (p Wrap) Title() string       { return fmt.Sprintf("Word Wrap (%s)", p.Name()) }
func (p Wrap) Description() string { return "Wrap text at specified column width" }
func (p Wrap) FilterValue() string { return p.Title() }

func wordWrap(text string, width int) string {
	if width <= 0 {
		return text
	}
	words := strings.Fields(text)
	if len(words) == 0 {
		return ""
	}
	var lines []string
	current := words[0]
	for _, word := range words[1:] {
		if len(current)+1+len(word) > width {
			lines = append(lines, current)
			current = word
		} else {
			current += " " + word
		}
	}
	lines = append(lines, current)
	return strings.Join(lines, "\n")
}

type ReplaceText struct{}

func (p ReplaceText) Name() string    { return "replace" }
func (p ReplaceText) Alias() []string { return []string{"find-replace"} }
func (p ReplaceText) Transform(data []byte, f ...Flag) (string, error) {
	var find, with string
	for _, flag := range f {
		switch flag.Short {
		case "f":
			if s, ok := flag.Value.(string); ok {
				find = s
			}
		case "w":
			if s, ok := flag.Value.(string); ok {
				with = s
			}
		}
	}
	if find == "" {
		return string(data), nil
	}
	return strings.ReplaceAll(string(data), find, with), nil
}
func (p ReplaceText) Flags() []Flag {
	return []Flag{
		{Name: "find", Short: "f", Desc: "String to find", Value: "", Type: FlagString},
		{Name: "with", Short: "w", Desc: "Replacement string", Value: "", Type: FlagString},
	}
}
func (p ReplaceText) Title() string       { return fmt.Sprintf("Find and Replace (%s)", p.Name()) }
func (p ReplaceText) Description() string { return "Find and replace text" }
func (p ReplaceText) FilterValue() string { return p.Title() }

type CharFrequency struct{}

func (p CharFrequency) Name() string    { return "char-freq" }
func (p CharFrequency) Alias() []string { return []string{"char-frequency"} }
func (p CharFrequency) Transform(data []byte, _ ...Flag) (string, error) {
	freq := make(map[rune]int)
	for _, r := range string(data) {
		freq[r]++
	}
	type kv struct {
		Key   rune
		Value int
	}
	pairs := make([]kv, 0, len(freq))
	for k, v := range freq {
		pairs = append(pairs, kv{k, v})
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i].Value > pairs[j].Value })
	var sb strings.Builder
	for _, p := range pairs {
		switch p.Key {
		case '\n':
			fmt.Fprintf(&sb, "\\n: %d\n", p.Value)
		case '\t':
			fmt.Fprintf(&sb, "\\t: %d\n", p.Value)
		case ' ':
			fmt.Fprintf(&sb, "[space]: %d\n", p.Value)
		default:
			fmt.Fprintf(&sb, "%c: %d\n", p.Key, p.Value)
		}
	}
	return strings.TrimSuffix(sb.String(), "\n"), nil
}
func (p CharFrequency) Flags() []Flag       { return nil }
func (p CharFrequency) Title() string       { return fmt.Sprintf("Character Frequency (%s)", p.Name()) }
func (p CharFrequency) Description() string { return "Character frequency analysis" }
func (p CharFrequency) FilterValue() string { return p.Title() }

type WordFrequency struct{}

func (p WordFrequency) Name() string    { return "word-freq" }
func (p WordFrequency) Alias() []string { return []string{"word-frequency"} }
func (p WordFrequency) Transform(data []byte, _ ...Flag) (string, error) {
	freq := make(map[string]int)
	for _, w := range strings.Fields(string(data)) {
		freq[strings.ToLower(w)]++
	}
	type kv struct {
		Key   string
		Value int
	}
	pairs := make([]kv, 0, len(freq))
	for k, v := range freq {
		pairs = append(pairs, kv{k, v})
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i].Value > pairs[j].Value })
	var sb strings.Builder
	for _, p := range pairs {
		sb.WriteString(fmt.Sprintf("%s: %d\n", p.Key, p.Value))
	}
	return strings.TrimSuffix(sb.String(), "\n"), nil
}
func (p WordFrequency) Flags() []Flag       { return nil }
func (p WordFrequency) Title() string       { return fmt.Sprintf("Word Frequency (%s)", p.Name()) }
func (p WordFrequency) Description() string { return "Word frequency analysis" }
func (p WordFrequency) FilterValue() string { return p.Title() }
