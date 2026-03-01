package processors

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type ChecksumVerify struct{}

func (p ChecksumVerify) Name() string    { return "checksum-verify" }
func (p ChecksumVerify) Alias() []string { return []string{"verify-hash", "check-hash"} }
func (p ChecksumVerify) Transform(data []byte, f ...Flag) (string, error) {
	var expectedHash, algo string
	for _, flag := range f {
		switch flag.Short {
		case "x":
			if s, ok := flag.Value.(string); ok {
				expectedHash = s
			}
		case "g":
			if s, ok := flag.Value.(string); ok {
				algo = s
			}
		}
	}
	if expectedHash == "" {
		return "", fmt.Errorf("--hash is required")
	}
	algo = strings.ToLower(strings.TrimSpace(algo))
	expectedHash = strings.ToLower(strings.TrimSpace(expectedHash))
	var h hash.Hash
	switch algo {
	case "md5":
		h = md5.New()
	case "sha1":
		h = sha1.New()
	case "sha256", "":
		h = sha256.New()
		if algo == "" {
			algo = "sha256"
		}
	case "sha384":
		h = sha512.New384()
	case "sha512":
		h = sha512.New()
	default:
		return "", fmt.Errorf("unsupported algorithm: %s (use md5, sha1, sha256, sha384, sha512)", algo)
	}
	h.Write(data)
	actual := hex.EncodeToString(h.Sum(nil))
	if actual == expectedHash {
		return fmt.Sprintf("✅ Match (%s)", algo), nil
	}
	return fmt.Sprintf("❌ Mismatch (%s)\n  expected: %s\n  actual:   %s", algo, expectedHash, actual), nil
}
func (p ChecksumVerify) Flags() []Flag {
	return []Flag{
		{Name: "hash", Short: "x", Desc: "Expected hash to verify against", Value: "", Type: FlagString},
		{Name: "algo", Short: "g", Desc: "Hash algorithm (md5/sha1/sha256/sha384/sha512)", Value: "sha256", Type: FlagString},
	}
}
func (p ChecksumVerify) Title() string {
	return fmt.Sprintf("%s (%s)", cases.Title(language.Und, cases.NoLower).String(p.Name()), p.Name())
}
func (p ChecksumVerify) Description() string { return "Verify checksum against expected hash" }
func (p ChecksumVerify) FilterValue() string { return p.Title() }
