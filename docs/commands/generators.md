# Generators

| Command | Description |
|---------|-------------|
| `uuid` | UUID v4 |
| `ulid` | ULID (sortable unique ID) |
| `nanoid` | NanoID (URL-friendly) |
| `lorem` | Lorem ipsum text |
| `now` | Current timestamp |
| `epoch` | Unix ↔ human timestamp |
| `password-gen` | Secure password |
| `totp` | TOTP code (RFC 6238) |
| `qrcode` | QR code in terminal |

## Examples

```shell
tweak uuid
tweak ulid
tweak nanoid --length 10
tweak lorem
tweak now
tweak epoch 0
tweak epoch "2024-01-01"
tweak qrcode "https://opensyntaxhq.github.io/tweak"
```
