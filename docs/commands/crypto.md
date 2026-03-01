# Cryptography & Security

| Command | Description |
|---------|-------------|
| `hmac-sha256` | HMAC-SHA256 (requires `--key`) |
| `hmac-sha512` | HMAC-SHA512 (requires `--key`) |
| `aes-encrypt` | AES-256-GCM encryption (requires `--key`) |
| `aes-decrypt` | AES-256-GCM decryption (requires `--key`) |
| `argon2` | Argon2id password hash |
| `jwt-encode` | Encode JWT (requires `--secret`) |
| `jwt-decode` | Decode JWT payload |
| `password-gen` | Secure random password |
| `totp` | TOTP code (RFC 6238) |
| `checksum-verify` | Verify checksum against hash |

## Examples

```shell
echo "hello" | tweak hmac-sha256 --key "secret"
echo "my text" | tweak aes-encrypt --key "mysecretpassword"
tweak password-gen --length 32
tweak totp JBSWY3DPEHPK3PXP
```
