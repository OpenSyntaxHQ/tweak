package processors

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var List = []list.Item{
	Adler32{},
	AESDecrypt{},
	AESEncrypt{},
	Argon2Hash{},
	ASCII85Decoding{},
	ASCII85Encoding{},
	Base32Decode{},
	Base32Encoding{},
	BaseConvert{},
	Base58Decode{},
	Base58Encode{},
	Base62Decode{},
	Base62Encode{},
	Base64Decode{},
	Base64Encode{},
	Base64URLDecode{},
	Base64URLEncode{},
	Bcrypt{},
	BinaryDecode{},
	BinaryEncode{},
	BLAKE2b{},
	BLAKE2s{},
	CaesarDecode{},
	CaesarEncode{},
	Camel{},
	CharFrequency{},
	ChecksumVerify{},
	Column{},
	CountCharacters{},
	CountLines{},
	CountWords{},
	CRC32{},
	CrockfordBase32Decode{},
	CrockfordBase32Encode{},
	CSVToJSON{},
	Detect{},
	Epoch{},
	EscapeQuotes{},
	ExtractEmails{},
	ExtractURLs{},
	ExtractIPs{},
	FormatJSON{},
	Grep{},
	HexDecode{},
	HexEncode{},
	HexToRGB{},
	HMACSHA256{},
	HMACSHA512{},
	HSLToHex{},
	HTMLDecode{},
	HTMLEncode{},
	JSONEscape{},
	JSONMinify{},
	JSONLToJSON{},
	JSONToCSV{},
	JSONToMSGPACK{},
	JSONToTOML{},
	JSONToXML{},
	JSONToYAML{},
	JSONUnescape{},
	JWTDecode{},
	JWTEncode{},
	Kebab{},
	Lorem{},
	Lower{},
	Markdown{},
	MorseCodeEncode{},
	MorseCodeDecode{},
	MD5{},
	MSGPACKToJSON{},
	NanoID{},
	Now{},
	NumberLines{},
	PadLeft{},
	PadRight{},
	Pascal{},
	PasswordGen{},
	QRCode{},
	RegexMatch{},
	RegexReplace{},
	RemoveNewLines{},
	RemoveSpaces{},
	Repeat{},
	ReplaceText{},
	Reverse{},
	ReverseLines{},
	RGBToHex{},
	ROT13{},
	SHA1{},
	SHA224{},
	SHA256{},
	SHA384{},
	SHA512{},
	ShuffleLines{},
	Slug{},
	Snake{},
	SortLines{},
	Title{},
	TOMLToJSON{},
	TOTP{},
	Trim{},
	ULID{},
	UniqueLines{},
	Upper{},
	URLDecode{},
	URLEncode{},
	UUID{},
	ValidateEmail{},
	ValidateJSON{},
	ValidateURL{},
	WordFrequency{},
	Wrap{},
	XMLToJSON{},
	XXH32{},
	XXH64{},
	XXH128{},
	YAMLToJSON{},
	Zeropad{},
}

type Processor interface {
	Name() string
	Alias() []string
	Transform(data []byte, opts ...Flag) (string, error)
	Flags() []Flag
}

type Generator interface {
	IsGenerator() bool
}

func IsGenerator(p Processor) bool {
	if g, ok := p.(Generator); ok {
		return g.IsGenerator()
	}
	return false
}

type StreamingConfig struct {
	ChunkSize    int
	BufferOutput bool
	LineByLine   bool
}

type ConfigurableStreamingProcessor interface {
	Processor
	GetStreamingConfig() StreamingConfig
}

type FlagType string

func (f FlagType) String() string  { return string(f) }
func (f FlagType) IsString() bool  { return f == FlagString }

const (
	FlagInt    FlagType = "Int"
	FlagUint   FlagType = "Uint"
	FlagBool   FlagType = "Bool"
	FlagString FlagType = "String"
)

type Flag struct {
	Name  string
	Short string
	Desc  string
	Type  FlagType
	Value any
}

var DefaultStreamingConfig = StreamingConfig{
	ChunkSize:    64 * 1024,
	BufferOutput: false,
	LineByLine:   false,
}

func TransformStream(processor Processor, reader io.Reader, writer io.Writer, opts ...Flag) error {
	config := DefaultStreamingConfig
	if sp, ok := processor.(ConfigurableStreamingProcessor); ok {
		config = sp.GetStreamingConfig()
	}

	switch {
	case config.LineByLine:
		return transformStreamLineByLine(processor, reader, writer, opts...)
	case config.BufferOutput:
		return transformStreamBuffered(processor, reader, writer, opts...)
	default:
		return transformStreamChunked(processor, reader, writer, config.ChunkSize, opts...)
	}
}

func transformStreamLineByLine(p Processor, r io.Reader, w io.Writer, opts ...Flag) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		result, err := p.Transform(scanner.Bytes(), opts...)
		if err != nil {
			return err
		}
		if _, err := fmt.Fprintln(w, result); err != nil {
			return err
		}
	}
	return scanner.Err()
}

func transformStreamBuffered(p Processor, r io.Reader, w io.Writer, opts ...Flag) error {
	data, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	result, err := p.Transform(data, opts...)
	if err != nil {
		return err
	}
	_, err = w.Write([]byte(result))
	return err
}

func transformStreamChunked(p Processor, r io.Reader, w io.Writer, chunkSize int, opts ...Flag) error {
	buf := make([]byte, chunkSize)
	for {
		n, err := r.Read(buf)
		if n > 0 {
			result, tErr := p.Transform(buf[:n], opts...)
			if tErr != nil {
				return tErr
			}
			if _, wErr := w.Write([]byte(result)); wErr != nil {
				return wErr
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func CanStream(_ Processor) bool { return true }

func PreferStream(processor Processor) bool {
	if sp, ok := processor.(ConfigurableStreamingProcessor); ok {
		cfg := sp.GetStreamingConfig()
		if !cfg.BufferOutput || cfg.LineByLine {
			return true
		}
	}

	name := processor.Name()
	friendly := []string{
		"md5", "sha1", "sha224", "sha256", "sha384", "sha512",
		"hex-encode", "hex-decode", "base64-encode", "base64-decode",
		"base32-encode", "base32-decode", "upper", "lower",
	}
	for _, f := range friendly {
		if name == f {
			return true
		}
	}
	return false
}

type Zeropad struct{}

func (p Zeropad) Name() string    { return "zeropad" }
func (p Zeropad) Alias() []string { return nil }

func (p Zeropad) Transform(data []byte, f ...Flag) (string, error) {
	strIn := strings.TrimSpace(string(data))
	neg := ""
	i, err := strconv.ParseFloat(strIn, 64)
	if err != nil {
		return "", fmt.Errorf("number expected: '%s'", data)
	}
	if i < 0 {
		neg = "-"
		data = data[1:]
	}

	var n int
	pre := ""
	for _, flag := range f {
		switch flag.Short {
		case "n":
			if x, ok := flag.Value.(uint); ok {
				n = int(x)
			}
		case "p":
			if x, ok := flag.Value.(string); ok {
				pre = x
			}
		}
	}
	return fmt.Sprintf("%s%s%s%s", pre, neg, strings.Repeat("0", n), data), nil
}

func (p Zeropad) Flags() []Flag {
	return []Flag{
		{Name: "number-of-zeros", Short: "n", Desc: "Number of zeros to pad", Value: uint(5), Type: FlagUint},
		{Name: "prefix", Short: "p", Desc: "Prefix before the number", Value: "", Type: FlagString},
	}
}

func (p Zeropad) Title() string {
	return fmt.Sprintf("%s (%s)", cases.Title(language.Und, cases.NoLower).String(p.Name()), p.Name())
}
func (p Zeropad) Description() string { return "Pad a number with zeros" }
func (p Zeropad) FilterValue() string { return p.Title() }
