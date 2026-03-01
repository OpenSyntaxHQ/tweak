package processors

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"strings"
)

type CSVToJSON struct{}

func (p CSVToJSON) Name() string    { return "csv-json" }
func (p CSVToJSON) Alias() []string { return nil }

func (p CSVToJSON) Transform(data []byte, _ ...Flag) (string, error) {
	reader := csv.NewReader(bytes.NewReader(data))
	records, err := reader.ReadAll()
	if err != nil {
		return "", err
	}
	if len(records) < 2 {
		return "[]", nil
	}
	headers := records[0]
	var results []map[string]string
	for _, row := range records[1:] {
		obj := make(map[string]string)
		for i, val := range row {
			if i < len(headers) {
				obj[headers[i]] = val
			}
		}
		results = append(results, obj)
	}
	out, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func (p CSVToJSON) Flags() []Flag       { return nil }
func (p CSVToJSON) Title() string       { return fmt.Sprintf("CSV to JSON (%s)", p.Name()) }
func (p CSVToJSON) Description() string { return "Convert CSV to JSON" }
func (p CSVToJSON) FilterValue() string { return p.Title() }

type JSONToCSV struct{}

func (p JSONToCSV) Name() string    { return "json-csv" }
func (p JSONToCSV) Alias() []string { return nil }

func (p JSONToCSV) Transform(data []byte, _ ...Flag) (string, error) {
	var records []map[string]any
	if err := json.Unmarshal(data, &records); err != nil {
		return "", err
	}
	if len(records) == 0 {
		return "", nil
	}

	headerSet := make(map[string]struct{})
	for _, rec := range records {
		for k := range rec {
			headerSet[k] = struct{}{}
		}
	}
	headers := make([]string, 0, len(headerSet))
	for h := range headerSet {
		headers = append(headers, h)
	}

	var buf strings.Builder
	w := csv.NewWriter(&buf)
	_ = w.Write(headers)
	for _, rec := range records {
		row := make([]string, len(headers))
		for i, h := range headers {
			if v, ok := rec[h]; ok {
				row[i] = fmt.Sprintf("%v", v)
			}
		}
		_ = w.Write(row)
	}
	w.Flush()
	return strings.TrimSuffix(buf.String(), "\n"), nil
}

func (p JSONToCSV) Flags() []Flag       { return nil }
func (p JSONToCSV) Title() string       { return fmt.Sprintf("JSON to CSV (%s)", p.Name()) }
func (p JSONToCSV) Description() string { return "Convert JSON array to CSV" }
func (p JSONToCSV) FilterValue() string { return p.Title() }
