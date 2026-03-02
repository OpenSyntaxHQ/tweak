# Validation & Detection

| Command | Description |
|---------|-------------|
| `validate-json` | Validate JSON |
| `validate-email` | Validate email address |
| `validate-url` | Validate URL |
| `detect` | Auto-detect encoding and decode |
| `extract-emails` | Extract all emails from text |
| `extract-urls` | Extract all URLs from text |
| `extract-ips` | Extract all IP addresses from text |

## Examples

```shell
echo '{"key":"value"}' | tweak validate-json
echo "user@example.com" | tweak validate-email
echo "not-an-email" | tweak validate-email
echo "SGVsbG8=" | tweak detect
cat log.txt | tweak extract-emails
cat page.html | tweak extract-urls
cat access.log | tweak extract-ips
```

`validate-json`, `validate-email`, and `validate-url` return non-zero exit codes on invalid input.
