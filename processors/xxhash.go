package processors

import (
	"fmt"

	"github.com/harsh16coder/xxhash"
)

type XXH32 struct{}

func (p XXH32) Name() string    { return "xxh-32" }
func (p XXH32) Alias() []string { return []string{"xxh32", "xxhash32", "xxhash-32"} }
func (p XXH32) Transform(data []byte, _ ...Flag) (string, error) {
	h := xxhash.New32()
	if _, err := h.Write(data); err != nil {
		return "", err
	}
	return fmt.Sprintf("%08x", h.Sum32()), nil
}
func (p XXH32) Flags() []Flag       { return nil }
func (p XXH32) Title() string       { return fmt.Sprintf("xxHash - XXH32 (%s)", p.Name()) }
func (p XXH32) Description() string { return "Get the XXH32 checksum of your text" }
func (p XXH32) FilterValue() string { return p.Title() }

type XXH64 struct{}

func (p XXH64) Name() string    { return "xxh-64" }
func (p XXH64) Alias() []string { return []string{"xxh64", "xxhash64", "xxhash-64"} }
func (p XXH64) Transform(data []byte, _ ...Flag) (string, error) {
	h := xxhash.New64()
	if _, err := h.Write(data); err != nil {
		return "", err
	}
	return fmt.Sprintf("%016x", h.Sum64()), nil
}
func (p XXH64) Flags() []Flag       { return nil }
func (p XXH64) Title() string       { return fmt.Sprintf("xxHash - XXH64 (%s)", p.Name()) }
func (p XXH64) Description() string { return "Get the XXH64 checksum of your text" }
func (p XXH64) FilterValue() string { return p.Title() }

type XXH128 struct{}

func (p XXH128) Name() string    { return "xxh-128" }
func (p XXH128) Alias() []string { return []string{"xxh128", "xxhash128", "xxhash-128"} }
func (p XXH128) Transform(data []byte, _ ...Flag) (string, error) {
	h := xxhash.New128()
	if _, err := h.Write(data); err != nil {
		return "", err
	}
	s := h.Sum128()
	return fmt.Sprintf("%016x%016x", s.Hi, s.Lo), nil
}
func (p XXH128) Flags() []Flag       { return nil }
func (p XXH128) Title() string       { return fmt.Sprintf("xxHash - XXH128 (%s)", p.Name()) }
func (p XXH128) Description() string { return "Get the XXH128 checksum of your text" }
func (p XXH128) FilterValue() string { return p.Title() }
