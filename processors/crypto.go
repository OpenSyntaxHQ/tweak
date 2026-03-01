package processors

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type MD5 struct{}

func (p MD5) GetStreamingConfig() StreamingConfig {
	return StreamingConfig{ChunkSize: 64 * 1024, BufferOutput: true, LineByLine: false}
}
func (p MD5) Name() string    { return "md5" }
func (p MD5) Alias() []string { return []string{"md5-sum"} }
func (p MD5) Transform(data []byte, _ ...Flag) (string, error) {
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil)), nil
}
func (p MD5) Flags() []Flag       { return nil }
func (p MD5) Title() string       { return fmt.Sprintf("MD5 Sum (%s)", p.Name()) }
func (p MD5) Description() string { return "Get the MD5 checksum of your text" }
func (p MD5) FilterValue() string { return p.Title() }

type SHA1 struct{}

func (p SHA1) GetStreamingConfig() StreamingConfig {
	return StreamingConfig{ChunkSize: 64 * 1024, BufferOutput: true}
}
func (p SHA1) Name() string    { return "sha1" }
func (p SHA1) Alias() []string { return []string{"sha1-sum"} }
func (p SHA1) Transform(data []byte, _ ...Flag) (string, error) {
	h := sha1.New()
	h.Write(data)
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
func (p SHA1) Flags() []Flag       { return nil }
func (p SHA1) Title() string       { return fmt.Sprintf("SHA-1 Sum (%s)", p.Name()) }
func (p SHA1) Description() string { return "Get the SHA-1 checksum of your text" }
func (p SHA1) FilterValue() string { return p.Title() }

type SHA224 struct{}

func (p SHA224) GetStreamingConfig() StreamingConfig {
	return StreamingConfig{ChunkSize: 64 * 1024, BufferOutput: true}
}
func (p SHA224) Name() string    { return "sha224" }
func (p SHA224) Alias() []string { return []string{"sha224-sum"} }
func (p SHA224) Transform(data []byte, _ ...Flag) (string, error) {
	bs := sha256.Sum224(data)
	return fmt.Sprintf("%x", bs), nil
}
func (p SHA224) Flags() []Flag       { return nil }
func (p SHA224) Title() string       { return fmt.Sprintf("SHA-224 Sum (%s)", p.Name()) }
func (p SHA224) Description() string { return "Get the SHA-224 checksum of your text" }
func (p SHA224) FilterValue() string { return p.Title() }

type SHA256 struct{}

func (p SHA256) GetStreamingConfig() StreamingConfig {
	return StreamingConfig{ChunkSize: 64 * 1024, BufferOutput: true}
}
func (p SHA256) Name() string    { return "sha256" }
func (p SHA256) Alias() []string { return []string{"sha256-sum"} }
func (p SHA256) Transform(data []byte, _ ...Flag) (string, error) {
	h := sha256.New()
	h.Write(data)
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
func (p SHA256) Flags() []Flag       { return nil }
func (p SHA256) Title() string       { return fmt.Sprintf("SHA-256 Sum (%s)", p.Name()) }
func (p SHA256) Description() string { return "Get the SHA-256 checksum of your text" }
func (p SHA256) FilterValue() string { return p.Title() }

type SHA384 struct{}

func (p SHA384) GetStreamingConfig() StreamingConfig {
	return StreamingConfig{ChunkSize: 64 * 1024, BufferOutput: true}
}
func (p SHA384) Name() string    { return "sha384" }
func (p SHA384) Alias() []string { return []string{"sha384-sum"} }
func (p SHA384) Transform(data []byte, _ ...Flag) (string, error) {
	bs := sha512.Sum384(data)
	return fmt.Sprintf("%x", bs), nil
}
func (p SHA384) Flags() []Flag       { return nil }
func (p SHA384) Title() string       { return fmt.Sprintf("SHA-384 Sum (%s)", p.Name()) }
func (p SHA384) Description() string { return "Get the SHA-384 checksum of your text" }
func (p SHA384) FilterValue() string { return p.Title() }

type SHA512 struct{}

func (p SHA512) GetStreamingConfig() StreamingConfig {
	return StreamingConfig{ChunkSize: 64 * 1024, BufferOutput: true}
}
func (p SHA512) Name() string    { return "sha512" }
func (p SHA512) Alias() []string { return []string{"sha512-sum"} }
func (p SHA512) Transform(data []byte, _ ...Flag) (string, error) {
	h := sha512.New()
	h.Write(data)
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
func (p SHA512) Flags() []Flag       { return nil }
func (p SHA512) Title() string       { return fmt.Sprintf("SHA-512 Sum (%s)", p.Name()) }
func (p SHA512) Description() string { return "Get the SHA-512 checksum of your text" }
func (p SHA512) FilterValue() string { return p.Title() }

type Bcrypt struct{}

func (p Bcrypt) Name() string    { return "bcrypt" }
func (p Bcrypt) Alias() []string { return []string{"bcrypt-hash"} }
func (p Bcrypt) Transform(data []byte, f ...Flag) (string, error) {
	var rounds uint = 10
	for _, flag := range f {
		if flag.Short == "r" {
			if r, ok := flag.Value.(uint); ok {
				rounds = r
			}
		}
	}
	hashed, err := bcrypt.GenerateFromPassword(data, int(rounds))
	return string(hashed), err
}
func (p Bcrypt) Flags() []Flag {
	return []Flag{{Name: "number-of-rounds", Short: "r", Desc: "Number of rounds", Value: uint(10), Type: FlagUint}}
}
func (p Bcrypt) Title() string       { return fmt.Sprintf("bcrypt Hash (%s)", p.Name()) }
func (p Bcrypt) Description() string { return "Get the bcrypt hash of your text" }
func (p Bcrypt) FilterValue() string { return p.Title() }
