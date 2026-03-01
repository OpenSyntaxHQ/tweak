# JSON & Data Formats

| Command | Description |
|---------|-------------|
| `json` | Format / pretty-print JSON |
| `json-minify` | Minify JSON |
| `json-escape` / `json-unescape` | Escape / unescape |
| `json-yaml` / `yaml-json` | JSON ↔ YAML |
| `json-toml` / `toml-json` | JSON ↔ TOML |
| `json-xml` / `xml-json` | JSON ↔ XML |
| `json-csv` / `csv-json` | JSON ↔ CSV |
| `json-msgpack` / `msgpack-json` | JSON ↔ MessagePack |
| `jsonl-json` | JSONL/NDJSON → JSON array |

## Examples

```shell
cat config.json | tweak json
cat config.yaml | tweak yaml-json
curl https://api.example.com | tweak json-yaml
echo '[{"id":1}' | tweak jsonl-json
```
