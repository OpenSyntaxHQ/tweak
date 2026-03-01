package processors

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type RGBToHex struct{}

func (p RGBToHex) Name() string    { return "rgb-hex" }
func (p RGBToHex) Alias() []string { return []string{"rgb-to-hex"} }
func (p RGBToHex) Transform(data []byte, _ ...Flag) (string, error) {
	input := strings.TrimSpace(string(data))
	input = strings.TrimPrefix(input, "rgb(")
	input = strings.TrimPrefix(input, "RGB(")
	input = strings.TrimSuffix(input, ")")
	parts := strings.FieldsFunc(input, func(r rune) bool { return r == ',' || r == ' ' })
	cleaned := make([]string, 0, 3)
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			cleaned = append(cleaned, p)
		}
	}
	if len(cleaned) != 3 {
		return "", fmt.Errorf("expected 3 values (R,G,B), got %d", len(cleaned))
	}
	var rgb [3]int
	for i, s := range cleaned {
		v, err := strconv.Atoi(s)
		if err != nil || v < 0 || v > 255 {
			return "", fmt.Errorf("invalid value '%s': must be 0-255", s)
		}
		rgb[i] = v
	}
	return fmt.Sprintf("#%02x%02x%02x", rgb[0], rgb[1], rgb[2]), nil
}
func (p RGBToHex) Flags() []Flag       { return nil }
func (p RGBToHex) Title() string {
	return fmt.Sprintf("%s (%s)", cases.Title(language.Und, cases.NoLower).String(p.Name()), p.Name())
}
func (p RGBToHex) Description() string { return "Convert RGB to hex color" }
func (p RGBToHex) FilterValue() string { return p.Title() }

type HSLToHex struct{}

func (p HSLToHex) Name() string    { return "hsl-hex" }
func (p HSLToHex) Alias() []string { return []string{"hsl-to-hex"} }
func (p HSLToHex) Transform(data []byte, _ ...Flag) (string, error) {
	input := strings.TrimSpace(string(data))
	input = strings.TrimPrefix(input, "hsl(")
	input = strings.TrimPrefix(input, "HSL(")
	input = strings.TrimSuffix(input, ")")
	parts := strings.FieldsFunc(input, func(r rune) bool { return r == ',' || r == ' ' })
	cleaned := make([]string, 0, 3)
	for _, p := range parts {
		p = strings.TrimSpace(strings.TrimSuffix(strings.TrimSuffix(p, "%"), "°"))
		if p != "" {
			cleaned = append(cleaned, p)
		}
	}
	if len(cleaned) != 3 {
		return "", fmt.Errorf("expected 3 values (H,S,L), got %d", len(cleaned))
	}
	h, err := strconv.ParseFloat(cleaned[0], 64)
	if err != nil {
		return "", fmt.Errorf("invalid hue: %s", cleaned[0])
	}
	s, err := strconv.ParseFloat(cleaned[1], 64)
	if err != nil {
		return "", fmt.Errorf("invalid saturation: %s", cleaned[1])
	}
	l, err := strconv.ParseFloat(cleaned[2], 64)
	if err != nil {
		return "", fmt.Errorf("invalid lightness: %s", cleaned[2])
	}
	if s > 1 {
		s /= 100
	}
	if l > 1 {
		l /= 100
	}
	h = math.Mod(h, 360)
	if h < 0 {
		h += 360
	}
	r, g, b := hslToRGB(h, s, l)
	return fmt.Sprintf("#%02x%02x%02x", r, g, b), nil
}

func hslToRGB(h, s, l float64) (int, int, int) {
	if s == 0 {
		v := int(math.Round(l * 255))
		return v, v, v
	}
	var q float64
	if l < 0.5 {
		q = l * (1 + s)
	} else {
		q = l + s - l*s
	}
	p := 2*l - q
	h /= 360
	r := hueToRGB(p, q, h+1.0/3.0)
	g := hueToRGB(p, q, h)
	b := hueToRGB(p, q, h-1.0/3.0)
	return int(math.Round(r * 255)), int(math.Round(g * 255)), int(math.Round(b * 255))
}

func hueToRGB(p, q, t float64) float64 {
	if t < 0 {
		t++
	}
	if t > 1 {
		t--
	}
	switch {
	case t < 1.0/6.0:
		return p + (q-p)*6*t
	case t < 1.0/2.0:
		return q
	case t < 2.0/3.0:
		return p + (q-p)*(2.0/3.0-t)*6
	default:
		return p
	}
}

func (p HSLToHex) Flags() []Flag       { return nil }
func (p HSLToHex) Title() string {
	return fmt.Sprintf("%s (%s)", cases.Title(language.Und, cases.NoLower).String(p.Name()), p.Name())
}
func (p HSLToHex) Description() string { return "Convert HSL to hex color" }
func (p HSLToHex) FilterValue() string { return p.Title() }
