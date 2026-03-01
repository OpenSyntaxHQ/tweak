package processors

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type JWTDecode struct{}

func (p JWTDecode) Name() string    { return "jwt-decode" }
func (p JWTDecode) Alias() []string { return []string{"jwt-d", "jwt-inspect"} }
func (p JWTDecode) Transform(data []byte, _ ...Flag) (string, error) {
	token := strings.TrimSpace(string(data))
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return "", fmt.Errorf("invalid JWT: expected 3 parts, got %d", len(parts))
	}
	headerJSON, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		return "", fmt.Errorf("invalid JWT header: %w", err)
	}
	payloadJSON, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return "", fmt.Errorf("invalid JWT payload: %w", err)
	}
	var header, payload any
	if err := json.Unmarshal(headerJSON, &header); err != nil {
		return "", fmt.Errorf("invalid JWT header JSON: %w", err)
	}
	if err := json.Unmarshal(payloadJSON, &payload); err != nil {
		return "", fmt.Errorf("invalid JWT payload JSON: %w", err)
	}
	result := map[string]any{
		"header":  header,
		"payload": payload,
	}
	out, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return "", err
	}
	return string(out), nil
}
func (p JWTDecode) Flags() []Flag       { return nil }
func (p JWTDecode) Title() string {
	return fmt.Sprintf("%s (%s)", cases.Title(language.Und, cases.NoLower).String(p.Name()), p.Name())
}
func (p JWTDecode) Description() string { return "Decode JWT token (header + payload)" }
func (p JWTDecode) FilterValue() string { return p.Title() }

type JWTEncode struct{}

func (p JWTEncode) Name() string    { return "jwt-encode" }
func (p JWTEncode) Alias() []string { return []string{"jwt-e", "jwt-sign"} }
func (p JWTEncode) Transform(data []byte, f ...Flag) (string, error) {
	var secret string
	expHours := uint(24)
	for _, flag := range f {
		switch flag.Short {
		case "s":
			if s, ok := flag.Value.(string); ok {
				secret = s
			}
		case "e":
			if e, ok := flag.Value.(uint); ok && e > 0 {
				expHours = e
			}
		}
	}
	if secret == "" {
		return "", fmt.Errorf("--secret is required")
	}
	var payload map[string]any
	if err := json.Unmarshal(data, &payload); err != nil {
		return "", fmt.Errorf("payload must be valid JSON: %w", err)
	}
	if _, ok := payload["iat"]; !ok {
		payload["iat"] = time.Now().Unix()
	}
	if _, ok := payload["exp"]; !ok {
		payload["exp"] = time.Now().Add(time.Duration(expHours) * time.Hour).Unix()
	}
	header := map[string]string{"alg": "HS256", "typ": "JWT"}
	headerJSON, _ := json.Marshal(header)
	payloadJSON, _ := json.Marshal(payload)
	headerB64 := base64.RawURLEncoding.EncodeToString(headerJSON)
	payloadB64 := base64.RawURLEncoding.EncodeToString(payloadJSON)
	signingInput := headerB64 + "." + payloadB64
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(signingInput))
	sig := base64.RawURLEncoding.EncodeToString(mac.Sum(nil))
	return signingInput + "." + sig, nil
}
func (p JWTEncode) Flags() []Flag {
	return []Flag{
		{Name: "secret", Short: "s", Desc: "HMAC signing secret", Value: "", Type: FlagString},
		{Name: "exp", Short: "e", Desc: "Expiry in hours from now", Value: uint(24), Type: FlagUint},
	}
}
func (p JWTEncode) Title() string {
	return fmt.Sprintf("%s (%s)", cases.Title(language.Und, cases.NoLower).String(p.Name()), p.Name())
}
func (p JWTEncode) Description() string { return "Encode and sign JWT (HS256)" }
func (p JWTEncode) FilterValue() string { return p.Title() }
