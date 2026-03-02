package processors

import (
	crypto_rand "crypto/rand"
	"fmt"
	"math/big"
	"math/rand"
	"sort"
	"strings"
	"time"
)

type CountLines struct{}

func (p CountLines) StreamingSpec() StreamingSpec {
	return StreamingSpec{Mode: StreamingModeBuffered, ChunkSize: 64 * 1024}
}
func (p CountLines) Name() string    { return "count-lines" }
func (p CountLines) Alias() []string { return nil }
func (p CountLines) Transform(data []byte, _ ...Flag) (string, error) {
	if len(data) == 0 {
		return "0", nil
	}
	text := string(data)
	lines := strings.Count(text, "\n")
	if !strings.HasSuffix(text, "\n") {
		lines++
	}
	return fmt.Sprintf("%d", lines), nil
}
func (p CountLines) Flags() []Flag       { return nil }
func (p CountLines) Title() string       { return fmt.Sprintf("Count Number of Lines (%s)", p.Name()) }
func (p CountLines) Description() string { return "Count the number of lines in your text" }
func (p CountLines) FilterValue() string { return p.Title() }

type SortLines struct{}

func (p SortLines) StreamingSpec() StreamingSpec {
	return StreamingSpec{Mode: StreamingModeBuffered, ChunkSize: 64 * 1024}
}
func (p SortLines) Name() string    { return "sort-lines" }
func (p SortLines) Alias() []string { return nil }
func (p SortLines) Transform(data []byte, _ ...Flag) (string, error) {
	sorted := strings.Split(string(data), "\n")
	if len(sorted) > 0 && sorted[len(sorted)-1] == "" {
		sorted = sorted[:len(sorted)-1]
	}
	sort.Strings(sorted)
	return strings.Join(sorted, "\n"), nil
}
func (p SortLines) Flags() []Flag       { return nil }
func (p SortLines) Title() string       { return fmt.Sprintf("Sort Lines (%s)", p.Name()) }
func (p SortLines) Description() string { return "Sort lines alphabetically" }
func (p SortLines) FilterValue() string { return p.Title() }

type ShuffleLines struct{}

func (p ShuffleLines) Name() string    { return "shuffle-lines" }
func (p ShuffleLines) Alias() []string { return nil }
func (p ShuffleLines) Transform(data []byte, _ ...Flag) (string, error) {
	seed, err := crypto_rand.Int(crypto_rand.Reader, big.NewInt(int64(time.Now().Nanosecond())))
	if err != nil {
		return "", err
	}
	r := rand.New(rand.NewSource(seed.Int64()))
	lines := strings.Split(string(data), "\n")
	r.Shuffle(len(lines), func(i, j int) { lines[i], lines[j] = lines[j], lines[i] })
	return strings.Join(lines, "\n"), nil
}
func (p ShuffleLines) Flags() []Flag       { return nil }
func (p ShuffleLines) Title() string       { return fmt.Sprintf("Shuffle Lines (%s)", p.Name()) }
func (p ShuffleLines) Description() string { return "Shuffle lines randomly" }
func (p ShuffleLines) FilterValue() string { return p.Title() }

type UniqueLines struct{}

func (p UniqueLines) StreamingSpec() StreamingSpec {
	return StreamingSpec{Mode: StreamingModeBuffered, ChunkSize: 64 * 1024}
}
func (p UniqueLines) Name() string    { return "unique-lines" }
func (p UniqueLines) Alias() []string { return nil }
func (p UniqueLines) Transform(data []byte, _ ...Flag) (string, error) {
	unique := make(map[string]int)
	lines := strings.Split(string(data), "\n")
	for k, v := range lines {
		unique[v] = k
	}
	type kv struct {
		Key   string
		Value int
	}
	pairs := make([]kv, 0, len(unique))
	for k, v := range unique {
		pairs = append(pairs, kv{k, v})
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i].Value < pairs[j].Value })
	output := make([]string, 0, len(pairs))
	for _, p := range pairs {
		output = append(output, p.Key)
	}
	return strings.Join(output, "\n"), nil
}
func (p UniqueLines) Flags() []Flag       { return nil }
func (p UniqueLines) Title() string       { return fmt.Sprintf("Unique Lines (%s)", p.Name()) }
func (p UniqueLines) Description() string { return "Get unique lines from list" }
func (p UniqueLines) FilterValue() string { return p.Title() }

type ReverseLines struct{}

func (p ReverseLines) StreamingSpec() StreamingSpec {
	return StreamingSpec{Mode: StreamingModeBuffered, ChunkSize: 64 * 1024}
}
func (p ReverseLines) Name() string    { return "reverse-lines" }
func (p ReverseLines) Alias() []string { return nil }
func (p ReverseLines) Transform(data []byte, _ ...Flag) (string, error) {
	lines := strings.Split(string(data), "\n")
	for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
		lines[i], lines[j] = lines[j], lines[i]
	}
	return strings.Join(lines, "\n"), nil
}
func (p ReverseLines) Flags() []Flag       { return nil }
func (p ReverseLines) Title() string       { return fmt.Sprintf("Reverse Lines (%s)", p.Name()) }
func (p ReverseLines) Description() string { return "Reverse the order of lines" }
func (p ReverseLines) FilterValue() string { return p.Title() }
