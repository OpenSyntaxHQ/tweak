package processors

import (
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/mr-tron/base58"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Detect struct{}

func (p Detect) Name() string    { return "detect" }
func (p Detect) Alias() []string { return []string{"auto-decode", "magic"} }
func (p Detect) Transform(data []byte, _ ...Flag) (string, error) {
	str := strings.TrimSpace(string(data))
	if len(str) == 0 {
		return "", fmt.Errorf("empty input")
	}

	// 1. JWT
	parts := strings.Split(str, ".")
	if len(parts) == 3 {
		if _, err := base64.RawURLEncoding.DecodeString(parts[0]); err == nil {
			if _, err := base64.RawURLEncoding.DecodeString(parts[1]); err == nil {
				out, err := JWTDecode{}.Transform(data)
				if err == nil {
					return fmt.Sprintf("🔍 Detected: JWT Token\n\n%s", out), nil
				}
			}
		}
	}

	// 2. Base64
	if decoded, err := base64.StdEncoding.DecodeString(str); err == nil && isText(decoded) {
		return fmt.Sprintf("🔍 Detected: Base64\n\n%s", string(decoded)), nil
	}
	if decoded, err := base64.URLEncoding.DecodeString(str); err == nil && isText(decoded) {
		return fmt.Sprintf("🔍 Detected: Base64 (URL-safe)\n\n%s", string(decoded)), nil
	}
	if decoded, err := base64.RawStdEncoding.DecodeString(str); err == nil && isText(decoded) {
		return fmt.Sprintf("🔍 Detected: Base64 (unpadded)\n\n%s", string(decoded)), nil
	}
	if decoded, err := base64.RawURLEncoding.DecodeString(str); err == nil && isText(decoded) {
		return fmt.Sprintf("🔍 Detected: Base64 (URL-safe, unpadded)\n\n%s", string(decoded)), nil
	}

	// 3. Hex
	cleanHex := strings.TrimPrefix(strings.ToLower(str), "0x")
	if decoded, err := hex.DecodeString(cleanHex); err == nil && isText(decoded) {
		return fmt.Sprintf("🔍 Detected: Hexadecimal\n\n%s", string(decoded)), nil
	}

	// 4. Base32
	cleanB32 := strings.ToUpper(str)
	if decoded, err := base32.StdEncoding.DecodeString(cleanB32); err == nil && isText(decoded) {
		return fmt.Sprintf("🔍 Detected: Base32\n\n%s", string(decoded)), nil
	}

	// 5. URL Encoded
	if strings.Contains(str, "%") {
		decoded, err := (URLDecode{}).Transform(data)
		if err == nil && decoded != str && isText([]byte(decoded)) {
			return fmt.Sprintf("🔍 Detected: URL Encoded\n\n%s", decoded), nil
		}
	}

	// 6. Base58
	if decoded, err := base58.Decode(str); err == nil && isText(decoded) {
		return fmt.Sprintf("🔍 Detected: Base58\n\n%s", string(decoded)), nil
	}

	// 7. Binary
	if isBinaryString(str) {
		decoded, err := (BinaryDecode{}).Transform([]byte(str))
		if err == nil && isText([]byte(decoded)) {
			return fmt.Sprintf("🔍 Detected: Binary\n\n%s", decoded), nil
		}
	}

	return "❌ Could not automatically detect and decode format. Try using a specific decoder command.", nil
}

func isBinaryString(s string) bool {
	s = strings.ReplaceAll(s, " ", "")
	for _, c := range s {
		if c != '0' && c != '1' {
			return false
		}
	}
	return len(s) >= 8
}

func isText(data []byte) bool {
	if len(data) == 0 {
		return false
	}
	printable := 0
	for _, b := range data {
		if (b >= 32 && b <= 126) || b == '\n' || b == '\r' || b == '\t' {
			printable++
		}
	}
	return float64(printable)/float64(len(data)) > 0.9
}

func (p Detect) Flags() []Flag       { return nil }
func (p Detect) Title() string {
	return fmt.Sprintf("%s (%s)", cases.Title(language.Und, cases.NoLower).String(p.Name()), p.Name())
}
func (p Detect) Description() string { return "Auto-detect and decode base64/hex/JWT/URL-encoded" }
func (p Detect) FilterValue() string { return p.Title() }
