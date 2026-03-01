package processors

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"math"
	"strings"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type TOTP struct{}

func (p TOTP) Name() string    { return "totp" }
func (p TOTP) Alias() []string { return []string{"2fa", "otp"} }
func (p TOTP) Transform(data []byte, f ...Flag) (string, error) {
	digits := uint(6)
	period := uint(30)
	for _, flag := range f {
		switch flag.Short {
		case "d":
			if d, ok := flag.Value.(uint); ok && d > 0 {
				digits = d
			}
		case "p":
			if p, ok := flag.Value.(uint); ok && p > 0 {
				period = p
			}
		}
	}
	secret := strings.TrimSpace(string(data))
	secret = strings.ToUpper(strings.ReplaceAll(secret, " ", ""))
	key, err := base32.StdEncoding.WithPadding(base32.NoPadding).DecodeString(secret)
	if err != nil {
		return "", fmt.Errorf("invalid base32 secret: %w", err)
	}
	counter := uint64(time.Now().Unix()) / uint64(period)
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, counter)
	mac := hmac.New(sha1.New, key)
	mac.Write(buf)
	hash := mac.Sum(nil)
	offset := hash[len(hash)-1] & 0xf
	code := binary.BigEndian.Uint32(hash[offset:offset+4]) & 0x7fffffff
	mod := uint32(math.Pow10(int(digits)))
	otp := code % mod
	return fmt.Sprintf("%0*d", digits, otp), nil
}
func (p TOTP) Flags() []Flag {
	return []Flag{
		{Name: "digits", Short: "d", Desc: "Number of digits", Value: uint(6), Type: FlagUint},
		{Name: "period", Short: "p", Desc: "Time step in seconds", Value: uint(30), Type: FlagUint},
	}
}
func (p TOTP) Title() string {
	return fmt.Sprintf("%s (%s)", cases.Title(language.Und, cases.NoLower).String(p.Name()), p.Name())
}
func (p TOTP) Description() string { return "Generate TOTP code from base32 secret" }
func (p TOTP) FilterValue() string { return p.Title() }
