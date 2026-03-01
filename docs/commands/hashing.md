# Hashing & Checksums

| Command | Algorithm |
|---------|-----------|
| `md5` | MD5 (128-bit) |
| `sha1` | SHA-1 (160-bit) |
| `sha224` | SHA-224 |
| `sha256` | SHA-256 |
| `sha384` | SHA-384 |
| `sha512` | SHA-512 (512-bit) |
| `bcrypt` | bcrypt |
| `adler32` | Adler-32 |
| `crc32` | CRC-32 (IEEE) |
| `xxh32` / `xxh64` / `xxh128` | xxHash |
| `blake2b` / `blake2s` | BLAKE2 |

## Examples

```shell
tweak md5 "Hello World"
tweak sha256 file.txt
echo "hello" | tweak blake2b
```
