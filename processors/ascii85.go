package processors

import (
	"bytes"
	"encoding/ascii85"
	"fmt"
	"io"
)

type ASCII85Encoding struct{}

func (p ASCII85Encoding) Name() string    { return "ascii85-encode" }
func (p ASCII85Encoding) Alias() []string { return []string{"ascii85-encoding", "base85-encode", "b85-encode"} }

func (p ASCII85Encoding) Transform(data []byte, _ ...Flag) (string, error) {
	buf := &bytes.Buffer{}
	encoder := ascii85.NewEncoder(buf)
	if _, err := encoder.Write(data); err != nil {
		return "", err
	}
	if err := encoder.Close(); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (p ASCII85Encoding) Flags() []Flag       { return nil }
func (p ASCII85Encoding) Title() string       { return fmt.Sprintf("Ascii85 / Base85 Encoding (%s)", p.Name()) }
func (p ASCII85Encoding) Description() string { return "Encode your text to Ascii85 ( Base85 )" }
func (p ASCII85Encoding) FilterValue() string { return p.Title() }

type ASCII85Decoding struct{}

func (p ASCII85Decoding) Name() string    { return "ascii85-decode" }
func (p ASCII85Decoding) Alias() []string { return []string{"ascii85-decoding", "base85-decode", "b85-decode"} }

func (p ASCII85Decoding) Transform(data []byte, _ ...Flag) (string, error) {
	decoder := ascii85.NewDecoder(bytes.NewReader(data))
	buf, err := io.ReadAll(decoder)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

func (p ASCII85Decoding) Flags() []Flag       { return nil }
func (p ASCII85Decoding) Title() string       { return fmt.Sprintf("Ascii85 / Base85 Decoding (%s)", p.Name()) }
func (p ASCII85Decoding) Description() string { return "Decode your Ascii85 ( Base85 ) text" }
func (p ASCII85Decoding) FilterValue() string { return p.Title() }
