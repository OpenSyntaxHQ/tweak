package processors

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func deriveAESKey(passphrase string) []byte {
	h := sha256.Sum256([]byte(passphrase))
	return h[:]
}

type AESEncrypt struct{}

func (p AESEncrypt) Name() string    { return "aes-encrypt" }
func (p AESEncrypt) Alias() []string { return []string{"aes-enc", "encrypt"} }
func (p AESEncrypt) Transform(data []byte, f ...Flag) (string, error) {
	var key string
	for _, flag := range f {
		if flag.Short == "k" {
			if s, ok := flag.Value.(string); ok {
				key = s
			}
		}
	}
	if key == "" {
		return "", fmt.Errorf("--key is required")
	}
	keyBytes := deriveAESKey(key)
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	ciphertext := aesGCM.Seal(nonce, nonce, data, nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}
func (p AESEncrypt) Flags() []Flag {
	return []Flag{{Name: "key", Short: "k", Desc: "Encryption passphrase", Value: "", Type: FlagString}}
}
func (p AESEncrypt) Title() string {
	return fmt.Sprintf("%s (%s)", cases.Title(language.Und, cases.NoLower).String(p.Name()), p.Name())
}
func (p AESEncrypt) Description() string { return "AES-256-GCM encrypt" }
func (p AESEncrypt) FilterValue() string { return p.Title() }

type AESDecrypt struct{}

func (p AESDecrypt) Name() string    { return "aes-decrypt" }
func (p AESDecrypt) Alias() []string { return []string{"aes-dec", "decrypt"} }
func (p AESDecrypt) Transform(data []byte, f ...Flag) (string, error) {
	var key string
	for _, flag := range f {
		if flag.Short == "k" {
			if s, ok := flag.Value.(string); ok {
				key = s
			}
		}
	}
	if key == "" {
		return "", fmt.Errorf("--key is required")
	}
	ciphertext, err := base64.StdEncoding.DecodeString(string(data))
	if err != nil {
		return "", fmt.Errorf("invalid base64 input: %w", err)
	}
	keyBytes := deriveAESKey(key)
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonceSize := aesGCM.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", fmt.Errorf("ciphertext too short")
	}
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", fmt.Errorf("decryption failed (wrong key?): %w", err)
	}
	return string(plaintext), nil
}
func (p AESDecrypt) Flags() []Flag {
	return []Flag{{Name: "key", Short: "k", Desc: "Decryption passphrase", Value: "", Type: FlagString}}
}
func (p AESDecrypt) Title() string {
	return fmt.Sprintf("%s (%s)", cases.Title(language.Und, cases.NoLower).String(p.Name()), p.Name())
}
func (p AESDecrypt) Description() string { return "AES-256-GCM decrypt" }
func (p AESDecrypt) FilterValue() string { return p.Title() }
