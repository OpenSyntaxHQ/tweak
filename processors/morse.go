package processors

import (
	"fmt"
	"strings"
)

var charToMorse = map[rune]string{
	'A': ".-", 'B': "-...", 'C': "-.-.", 'D': "-..", 'E': ".", 'F': "..-.",
	'G': "--.", 'H': "....", 'I': "..", 'J': ".---", 'K': "-.-", 'L': ".-..",
	'M': "--", 'N': "-.", 'O': "---", 'P': ".--.", 'Q': "--.-", 'R': ".-.",
	'S': "...", 'T': "-", 'U': "..-", 'V': "...-", 'W': ".--", 'X': "-..-",
	'Y': "-.--", 'Z': "--..",
	'0': "-----", '1': ".----", '2': "..---", '3': "...--", '4': "....-",
	'5': ".....", '6': "-....", '7': "--...", '8': "---..", '9': "----.",
	'.': ".-.-.-", ',': "--..--", '?': "..--..", '\'': ".----.", '!': "-.-.--",
	'/': "-..-.", '(': "-.--.", ')': "-.--.-", '&': ".-...", ':': "---...",
	';': "-.-.-.", '=': "-...-", '+': ".-.-.", '-': "-....-", '_': "..--.-",
	'"': ".-..-.", '$': "...-..-", '@': ".--.-.", ' ': "/",
}

var morseToChar map[string]rune

func init() {
	morseToChar = make(map[string]rune, len(charToMorse))
	for k, v := range charToMorse {
		morseToChar[v] = k
	}
}

type MorseCodeEncode struct{}

func (p MorseCodeEncode) Name() string    { return "morse-encode" }
func (p MorseCodeEncode) Alias() []string { return nil }

func (p MorseCodeEncode) Transform(data []byte, _ ...Flag) (string, error) {
	var parts []string
	for _, r := range strings.ToUpper(string(data)) {
		if code, ok := charToMorse[r]; ok {
			parts = append(parts, code)
		}
	}
	return strings.Join(parts, " "), nil
}

func (p MorseCodeEncode) Flags() []Flag       { return nil }
func (p MorseCodeEncode) Title() string       { return fmt.Sprintf("Morse Code Encode (%s)", p.Name()) }
func (p MorseCodeEncode) Description() string { return "Encode your text to Morse code" }
func (p MorseCodeEncode) FilterValue() string { return p.Title() }

type MorseCodeDecode struct{}

func (p MorseCodeDecode) Name() string    { return "morse-decode" }
func (p MorseCodeDecode) Alias() []string { return nil }

func (p MorseCodeDecode) Transform(data []byte, _ ...Flag) (string, error) {
	var result strings.Builder
	words := strings.Split(string(data), " / ")
	for wi, word := range words {
		if wi > 0 {
			result.WriteRune(' ')
		}
		codes := strings.Fields(word)
		for _, code := range codes {
			if ch, ok := morseToChar[code]; ok {
				result.WriteRune(ch)
			}
		}
	}
	return result.String(), nil
}

func (p MorseCodeDecode) Flags() []Flag       { return nil }
func (p MorseCodeDecode) Title() string       { return fmt.Sprintf("Morse Code Decode (%s)", p.Name()) }
func (p MorseCodeDecode) Description() string { return "Decode your Morse code" }
func (p MorseCodeDecode) FilterValue() string { return p.Title() }
