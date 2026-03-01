package processors

import (
	"fmt"
	"strconv"
	"strings"
)

type NumberLines struct{}

func (p NumberLines) GetStreamingConfig() StreamingConfig {
	return StreamingConfig{ChunkSize: 64 * 1024, BufferOutput: true}
}
func (p NumberLines) Name() string    { return "number-lines" }
func (p NumberLines) Alias() []string { return []string{"nl", "line-numbers", "line-number", "number-line", "numberlines", "numberline"} }

func (p NumberLines) Transform(data []byte, _ ...Flag) (string, error) {
	lines := strings.Split(string(data), "\n")
	counter := 1
	nec := nonEmptyCount(lines)
	maxDigits := len(strconv.Itoa(nec))
	var result strings.Builder
	for i, line := range lines {
		if line != "" {
			result.WriteString(fmt.Sprintf("%*d. %s", maxDigits, counter, line))
			counter++
		}
		if i < len(lines)-1 {
			result.WriteByte('\n')
		}
	}
	return result.String(), nil
}

func nonEmptyCount(strs []string) int {
	count := 0
	for _, s := range strs {
		if s != "" {
			count++
		}
	}
	return count
}

func (p NumberLines) Flags() []Flag       { return nil }
func (p NumberLines) Title() string       { return "Line numberer" }
func (p NumberLines) Description() string { return "Prepend consecutive number to each input line" }
func (p NumberLines) FilterValue() string { return p.Title() }
