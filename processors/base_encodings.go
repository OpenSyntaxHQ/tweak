package processors

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"strings"
)

const crockfordAlphabet = "0123456789ABCDEFGHJKMNPQRSTVWXYZ"
const base58Alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
const base62Alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

type CrockfordBase32Encode struct{}

func (p CrockfordBase32Encode) Name() string    { return "crockford-base32-encode" }
func (p CrockfordBase32Encode) Alias() []string { return []string{"crockford-b32-enc", "cb32-encode"} }

func (p CrockfordBase32Encode) Transform(data []byte, f ...Flag) (string, error) {
	var checksum bool
	for _, flag := range f {
		if flag.Short == "c" {
			if c, ok := flag.Value.(bool); ok {
				checksum = c
			}
		}
	}
	encoded := encodeCrockfordBase32(data)
	if checksum {
		cv := calculateCrockfordChecksum(data)
		encoded += string(crockfordAlphabet[cv])
	}
	return encoded, nil
}

func (p CrockfordBase32Encode) Flags() []Flag {
	return []Flag{{Name: "checksum", Short: "c", Desc: "Add Crockford checksum", Value: false, Type: FlagBool}}
}
func (p CrockfordBase32Encode) Title() string       { return fmt.Sprintf("Crockford Base32 Encode (%s)", p.Name()) }
func (p CrockfordBase32Encode) Description() string { return "Encode your text to Crockford Base32" }
func (p CrockfordBase32Encode) FilterValue() string { return p.Title() }

type CrockfordBase32Decode struct{}

func (p CrockfordBase32Decode) Name() string    { return "crockford-base32-decode" }
func (p CrockfordBase32Decode) Alias() []string { return []string{"crockford-b32-dec", "cb32-decode"} }

func (p CrockfordBase32Decode) Transform(data []byte, f ...Flag) (string, error) {
	var verify bool
	for _, flag := range f {
		if flag.Short == "v" {
			if v, ok := flag.Value.(bool); ok {
				verify = v
			}
		}
	}
	input := strings.ToUpper(string(data))
	if verify && len(input) > 0 {
		checksumChar := input[len(input)-1:]
		dataStr := input[:len(input)-1]
		decoded, err := decodeCrockfordBase32(dataStr)
		if err != nil {
			return "", err
		}
		expected := calculateCrockfordChecksum(decoded)
		actual := strings.Index(crockfordAlphabet, checksumChar)
		if actual == -1 || actual != expected {
			return "", fmt.Errorf("checksum verification failed")
		}
		return string(decoded), nil
	}
	decoded, err := decodeCrockfordBase32(input)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}

func (p CrockfordBase32Decode) Flags() []Flag {
	return []Flag{{Name: "verify", Short: "v", Desc: "Verify Crockford checksum", Value: false, Type: FlagBool}}
}
func (p CrockfordBase32Decode) Title() string       { return fmt.Sprintf("Crockford Base32 Decode (%s)", p.Name()) }
func (p CrockfordBase32Decode) Description() string { return "Decode your Crockford Base32 text" }
func (p CrockfordBase32Decode) FilterValue() string { return p.Title() }

type Base58Encode struct{}

func (p Base58Encode) Name() string    { return "base58-encode" }
func (p Base58Encode) Alias() []string { return []string{"b58-enc", "b58-encode"} }

func (p Base58Encode) Transform(data []byte, f ...Flag) (string, error) {
	var check bool
	for _, flag := range f {
		if flag.Short == "c" {
			if c, ok := flag.Value.(bool); ok {
				check = c
			}
		}
	}
	if check {
		return encodeBase58Check(data), nil
	}
	return encodeBase58(data), nil
}

func (p Base58Encode) Flags() []Flag {
	return []Flag{{Name: "check", Short: "c", Desc: "Use Base58Check encoding (with checksum)", Value: false, Type: FlagBool}}
}
func (p Base58Encode) Title() string       { return fmt.Sprintf("Base58 Encode (%s)", p.Name()) }
func (p Base58Encode) Description() string { return "Encode your text to Base58" }
func (p Base58Encode) FilterValue() string { return p.Title() }

type Base58Decode struct{}

func (p Base58Decode) Name() string    { return "base58-decode" }
func (p Base58Decode) Alias() []string { return []string{"b58-dec", "b58-decode"} }

func (p Base58Decode) Transform(data []byte, f ...Flag) (string, error) {
	var check bool
	for _, flag := range f {
		if flag.Short == "c" {
			if c, ok := flag.Value.(bool); ok {
				check = c
			}
		}
	}
	if check {
		decoded, err := decodeBase58Check(string(data))
		if err != nil {
			return "", err
		}
		return string(decoded), nil
	}
	decoded, err := decodeBase58(string(data))
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}

func (p Base58Decode) Flags() []Flag {
	return []Flag{{Name: "check", Short: "c", Desc: "Use Base58Check decoding (verify checksum)", Value: false, Type: FlagBool}}
}
func (p Base58Decode) Title() string       { return fmt.Sprintf("Base58 Decode (%s)", p.Name()) }
func (p Base58Decode) Description() string { return "Decode your Base58 text" }
func (p Base58Decode) FilterValue() string { return p.Title() }

type Base62Encode struct{}

func (p Base62Encode) Name() string    { return "base62-encode" }
func (p Base62Encode) Alias() []string { return []string{"b62-enc", "b62-encode"} }

func (p Base62Encode) Transform(data []byte, f ...Flag) (string, error) {
	var prefix string
	for _, flag := range f {
		if flag.Short == "p" {
			if p, ok := flag.Value.(string); ok {
				prefix = p
			}
		}
	}
	encoded := encodeBase62(data)
	if prefix != "" {
		encoded = prefix + "_" + encoded
	}
	return encoded, nil
}

func (p Base62Encode) Flags() []Flag {
	return []Flag{{Name: "prefix", Short: "p", Desc: "Add prefix to encoded string", Value: "", Type: FlagString}}
}
func (p Base62Encode) Title() string       { return fmt.Sprintf("Base62 Encode (%s)", p.Name()) }
func (p Base62Encode) Description() string { return "Encode your text to Base62" }
func (p Base62Encode) FilterValue() string { return p.Title() }

type Base62Decode struct{}

func (p Base62Decode) Name() string    { return "base62-decode" }
func (p Base62Decode) Alias() []string { return []string{"b62-dec", "b62-decode"} }

func (p Base62Decode) Transform(data []byte, _ ...Flag) (string, error) {
	input := string(data)
	if idx := strings.Index(input, "_"); idx != -1 {
		input = input[idx+1:]
	}
	decoded, err := decodeBase62(input)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}

func (p Base62Decode) Flags() []Flag       { return nil }
func (p Base62Decode) Title() string       { return fmt.Sprintf("Base62 Decode (%s)", p.Name()) }
func (p Base62Decode) Description() string { return "Decode your Base62 text" }
func (p Base62Decode) FilterValue() string { return p.Title() }

func encodeCrockfordBase32(data []byte) string {
	if len(data) == 0 {
		return ""
	}
	num := new(big.Int).SetBytes(data)
	if num.Sign() == 0 {
		return "0"
	}
	base := big.NewInt(32)
	rem := new(big.Int)
	var result string
	for num.Sign() > 0 {
		num.DivMod(num, base, rem)
		result = string(crockfordAlphabet[rem.Int64()]) + result
	}
	return result
}

func decodeCrockfordBase32(encoded string) ([]byte, error) {
	if encoded == "" {
		return []byte{}, nil
	}
	encoded = strings.ToUpper(encoded)
	encoded = strings.NewReplacer("O", "0", "I", "1", "L", "1").Replace(encoded)

	num := new(big.Int)
	base := big.NewInt(32)
	for _, ch := range encoded {
		idx := strings.IndexRune(crockfordAlphabet, ch)
		if idx == -1 {
			return nil, fmt.Errorf("invalid character in Crockford Base32: %c", ch)
		}
		num.Mul(num, base)
		num.Add(num, big.NewInt(int64(idx)))
	}
	return num.Bytes(), nil
}

func calculateCrockfordChecksum(data []byte) int {
	sum := 0
	for _, b := range data {
		sum += int(b)
	}
	return sum % 32
}

func encodeBase58(data []byte) string {
	if len(data) == 0 {
		return ""
	}
	leadingZeros := 0
	for _, b := range data {
		if b == 0 {
			leadingZeros++
		} else {
			break
		}
	}
	num := new(big.Int).SetBytes(data)
	if num.Sign() == 0 {
		return strings.Repeat("1", leadingZeros)
	}
	base := big.NewInt(58)
	rem := new(big.Int)
	var result string
	for num.Sign() > 0 {
		num.DivMod(num, base, rem)
		result = string(base58Alphabet[rem.Int64()]) + result
	}
	return strings.Repeat("1", leadingZeros) + result
}

func decodeBase58(encoded string) ([]byte, error) {
	if encoded == "" {
		return []byte{}, nil
	}
	leadingOnes := 0
	for _, ch := range encoded {
		if ch == '1' {
			leadingOnes++
		} else {
			break
		}
	}
	num := new(big.Int)
	base := big.NewInt(58)
	for _, ch := range encoded {
		idx := strings.IndexRune(base58Alphabet, ch)
		if idx == -1 {
			return nil, fmt.Errorf("invalid character in Base58: %c", ch)
		}
		num.Mul(num, base)
		num.Add(num, big.NewInt(int64(idx)))
	}
	b := num.Bytes()
	result := make([]byte, leadingOnes+len(b))
	copy(result[leadingOnes:], b)
	return result, nil
}

func encodeBase58Check(data []byte) string {
	h1 := sha256.Sum256(data)
	h2 := sha256.Sum256(h1[:])
	payload := append(data, h2[:4]...)
	return encodeBase58(payload)
}

func decodeBase58Check(encoded string) ([]byte, error) {
	decoded, err := decodeBase58(encoded)
	if err != nil {
		return nil, err
	}
	if len(decoded) < 4 {
		return nil, fmt.Errorf("Base58Check data too short")
	}
	data := decoded[:len(decoded)-4]
	checksum := decoded[len(decoded)-4:]
	h1 := sha256.Sum256(data)
	h2 := sha256.Sum256(h1[:])
	for i := 0; i < 4; i++ {
		if checksum[i] != h2[i] {
			return nil, fmt.Errorf("Base58Check checksum verification failed")
		}
	}
	return data, nil
}

func encodeBase62(data []byte) string {
	if len(data) == 0 {
		return ""
	}
	num := new(big.Int).SetBytes(data)
	if num.Sign() == 0 {
		return "0"
	}
	base := big.NewInt(62)
	rem := new(big.Int)
	var result string
	for num.Sign() > 0 {
		num.DivMod(num, base, rem)
		result = string(base62Alphabet[rem.Int64()]) + result
	}
	return result
}

func decodeBase62(encoded string) ([]byte, error) {
	if encoded == "" {
		return []byte{}, nil
	}
	num := new(big.Int)
	base := big.NewInt(62)
	for _, ch := range encoded {
		idx := strings.IndexRune(base62Alphabet, ch)
		if idx == -1 {
			return nil, fmt.Errorf("invalid character in Base62: %c", ch)
		}
		num.Mul(num, base)
		num.Add(num, big.NewInt(int64(idx)))
	}
	return num.Bytes(), nil
}
