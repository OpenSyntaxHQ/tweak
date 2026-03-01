package processors

import (
	"encoding/hex"
	"fmt"
	"html"
	"net/url"
	"strings"

	"mvdan.cc/xurls/v2"
)

type HexEncode struct{}

func (p HexEncode) Name() string    { return "hex-encode" }
func (p HexEncode) Alias() []string { return []string{"hex-enc", "hexadecimal-encode"} }
func (p HexEncode) GetStreamingConfig() StreamingConfig {
	return StreamingConfig{ChunkSize: 64 * 1024, BufferOutput: false, LineByLine: false}
}
func (p HexEncode) Transform(data []byte, _ ...Flag) (string, error) {
	return hex.EncodeToString(data), nil
}
func (p HexEncode) Flags() []Flag       { return nil }
func (p HexEncode) Title() string       { return fmt.Sprintf("Hex Encode (%s)", p.Name()) }
func (p HexEncode) Description() string { return "Encode your text to Hex" }
func (p HexEncode) FilterValue() string { return p.Title() }

type HexDecode struct{}

func (p HexDecode) Name() string    { return "hex-decode" }
func (p HexDecode) Alias() []string { return []string{"hex-dec", "hexadecimal-decode"} }
func (p HexDecode) GetStreamingConfig() StreamingConfig {
	return StreamingConfig{ChunkSize: 64 * 1024, BufferOutput: true, LineByLine: false}
}
func (p HexDecode) Transform(data []byte, _ ...Flag) (string, error) {
	output, err := hex.DecodeString(string(data))
	if err != nil {
		return "", err
	}
	return string(output), nil
}
func (p HexDecode) Flags() []Flag       { return nil }
func (p HexDecode) Title() string       { return fmt.Sprintf("Hex Decode (%s)", p.Name()) }
func (p HexDecode) Description() string { return "Convert Hexadecimal to String" }
func (p HexDecode) FilterValue() string { return p.Title() }

type HTMLEncode struct{}

func (p HTMLEncode) Name() string    { return "html-encode" }
func (p HTMLEncode) Alias() []string { return []string{"html-enc", "html-escape"} }
func (p HTMLEncode) Transform(data []byte, _ ...Flag) (string, error) {
	return html.EscapeString(string(data)), nil
}
func (p HTMLEncode) Flags() []Flag       { return nil }
func (p HTMLEncode) Title() string       { return fmt.Sprintf("HTML Encode (%s)", p.Name()) }
func (p HTMLEncode) Description() string { return "Escape your HTML" }
func (p HTMLEncode) FilterValue() string { return p.Title() }

type HTMLDecode struct{}

func (p HTMLDecode) Name() string    { return "html-decode" }
func (p HTMLDecode) Alias() []string { return []string{"html-dec", "html-unescape"} }
func (p HTMLDecode) Transform(data []byte, _ ...Flag) (string, error) {
	return html.UnescapeString(string(data)), nil
}
func (p HTMLDecode) Flags() []Flag       { return nil }
func (p HTMLDecode) Title() string       { return fmt.Sprintf("HTML Decode (%s)", p.Name()) }
func (p HTMLDecode) Description() string { return "Unescape your HTML" }
func (p HTMLDecode) FilterValue() string { return p.Title() }

type URLEncode struct{}

func (p URLEncode) Name() string    { return "url-encode" }
func (p URLEncode) Alias() []string { return []string{"url-enc"} }
func (p URLEncode) Transform(data []byte, _ ...Flag) (string, error) {
	return url.QueryEscape(string(data)), nil
}
func (p URLEncode) Flags() []Flag       { return nil }
func (p URLEncode) Title() string       { return fmt.Sprintf("URL Encode (%s)", p.Name()) }
func (p URLEncode) Description() string { return "Encode URL entities" }
func (p URLEncode) FilterValue() string { return p.Title() }

type URLDecode struct{}

func (p URLDecode) Name() string    { return "url-decode" }
func (p URLDecode) Alias() []string { return []string{"url-dec"} }
func (p URLDecode) Transform(data []byte, _ ...Flag) (string, error) {
	res, _ := url.QueryUnescape(string(data))
	return res, nil
}
func (p URLDecode) Flags() []Flag       { return nil }
func (p URLDecode) Title() string       { return fmt.Sprintf("URL Decode (%s)", p.Name()) }
func (p URLDecode) Description() string { return "Decode URL entities" }
func (p URLDecode) FilterValue() string { return p.Title() }

type ExtractURLs struct{}

func (p ExtractURLs) Name() string    { return "extract-url" }
func (p ExtractURLs) Alias() []string { return []string{"url-ext", "extract-urls", "ext-url"} }
func (p ExtractURLs) Transform(data []byte, _ ...Flag) (string, error) {
	rx := xurls.Relaxed()
	urls := rx.FindAllString(string(data), -1)
	return strings.Join(urls, "\n"), nil
}
func (p ExtractURLs) Flags() []Flag       { return nil }
func (p ExtractURLs) Title() string       { return fmt.Sprintf("Extract URLs (%s)", p.Name()) }
func (p ExtractURLs) Description() string { return "Extract URLs from text" }
func (p ExtractURLs) FilterValue() string { return p.Title() }
