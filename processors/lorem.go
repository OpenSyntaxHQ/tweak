package processors

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var loremWords = []string{
	"lorem", "ipsum", "dolor", "sit", "amet", "consectetur", "adipiscing", "elit",
	"sed", "do", "eiusmod", "tempor", "incididunt", "ut", "labore", "et", "dolore",
	"magna", "aliqua", "enim", "ad", "minim", "veniam", "quis", "nostrud",
	"exercitation", "ullamco", "laboris", "nisi", "aliquip", "ex", "ea", "commodo",
	"consequat", "duis", "aute", "irure", "in", "reprehenderit", "voluptate",
	"velit", "esse", "cillum", "fugiat", "nulla", "pariatur", "excepteur", "sint",
	"occaecat", "cupidatat", "non", "proident", "sunt", "culpa", "qui", "officia",
	"deserunt", "mollit", "anim", "id", "est", "laborum",
}

var loremTitleCaser = cases.Title(language.Und, cases.NoLower)

type Lorem struct{}

func (p Lorem) Name() string    { return "lorem" }
func (p Lorem) Alias() []string { return []string{"lorem-ipsum"} }

func (p Lorem) Transform(_ []byte, f ...Flag) (string, error) {
	paragraphs := uint(1)
	for _, flag := range f {
		if flag.Short == "p" {
			if n, ok := flag.Value.(uint); ok && n > 0 {
				paragraphs = n
			}
		}
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var pars []string
	for i := uint(0); i < paragraphs; i++ {
		sentences := 4 + r.Intn(4) // 4-7 sentences per paragraph
		var sents []string
		for s := 0; s < sentences; s++ {
			wordCount := 8 + r.Intn(12) // 8-19 words per sentence
			words := make([]string, wordCount)
			for w := 0; w < wordCount; w++ {
				words[w] = loremWords[r.Intn(len(loremWords))]
			}
			words[0] = loremTitleCaser.String(words[0])
			sents = append(sents, strings.Join(words, " ")+".")
		}
		pars = append(pars, strings.Join(sents, " "))
	}
	return strings.Join(pars, "\n\n"), nil
}

func (p Lorem) IsGenerator() bool { return true }
func (p Lorem) Flags() []Flag {
	return []Flag{{Name: "paragraphs", Short: "p", Desc: "Number of paragraphs", Value: uint(1), Type: FlagUint}}
}
func (p Lorem) Title() string       { return fmt.Sprintf("Lorem Ipsum Generator (%s)", p.Name()) }
func (p Lorem) Description() string { return "Generate Lorem Ipsum text" }
func (p Lorem) FilterValue() string { return p.Title() }
