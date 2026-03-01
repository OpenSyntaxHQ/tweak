package processors

import (
	"fmt"
	"net"
	"regexp"
	"strings"
)

type ExtractIPs struct{}

func (p ExtractIPs) Name() string    { return "extract-ip" }
func (p ExtractIPs) Alias() []string { return []string{"find-ips", "find-ip", "extract-ips"} }

func (p ExtractIPs) Transform(data []byte, _ ...Flag) (string, error) {
	text := string(data)

	ipv4Re := regexp.MustCompile(`([0-9]{1,3}\.){3}[0-9]{1,3}`)
	ipv6Re := regexp.MustCompile(`(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:(:[0-9a-fA-F]{1,4}){1,6}|:((:[0-9a-fA-F]{1,4}){1,7}|:))`)

	var candidates []string
	candidates = append(candidates, ipv4Re.FindAllString(text, -1)...)
	candidates = append(candidates, ipv6Re.FindAllString(text, -1)...)

	var validIPs []string
	for _, c := range candidates {
		if ip := net.ParseIP(c); ip != nil {
			validIPs = append(validIPs, ip.String())
		}
	}
	return strings.Join(validIPs, "\n"), nil
}

func (p ExtractIPs) Flags() []Flag       { return nil }
func (p ExtractIPs) Title() string       { return fmt.Sprintf("Extract IPs (%s)", p.Name()) }
func (p ExtractIPs) Description() string { return "Extract IPv4 and IPv6 from your text" }
func (p ExtractIPs) FilterValue() string { return p.Title() }
