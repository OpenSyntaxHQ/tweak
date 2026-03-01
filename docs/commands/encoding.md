# Encoding & Decoding

| Command | Description |
|---------|-------------|
| `base32-encode` / `base32-decode` | Base32 |
| `base58-encode` / `base58-decode` | Base58 |
| `base62-encode` / `base62-decode` | Base62 |
| `base64-encode` / `base64-decode` | Base64 |
| `base64url-encode` / `base64url-decode` | Base64 URL-safe |
| `ascii85-encode` / `ascii85-decode` | Ascii85 |
| `crockford-encode` / `crockford-decode` | Crockford Base32 |
| `hex-encode` / `hex-decode` | Hexadecimal |
| `binary-encode` / `binary-decode` | Binary |
| `html-encode` / `html-decode` | HTML escape |
| `url-encode` / `url-decode` | URL encoding |

## Examples

```shell
tweak base64-encode "Hello World"
echo "SGVsbG8gV29ybGQ=" | tweak base64-decode
tweak hex-encode "hello"
tweak url-encode "hello world & more"
```
