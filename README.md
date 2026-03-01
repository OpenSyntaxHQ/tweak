<p align="center">
  <img src="./docs/assets/banner.png" width="100%" alt="tweak banner">
</p>


# tweak

[Website](https://opensyntaxhq.github.io/tweak/) | [Install](#installation) | [Usage](#usage) | [Commands](#commands) | [Contributing](CONTRIBUTING.md)

[![CI](https://github.com/OpenSyntaxHQ/tweak/actions/workflows/ci.yml/badge.svg)](https://github.com/OpenSyntaxHQ/tweak/actions/workflows/ci.yml)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)
[![Go Version](https://img.shields.io/badge/go-1.23+-00ADD8.svg)](go.mod)
[![Release](https://img.shields.io/github/v/release/OpenSyntaxHQ/tweak)](https://github.com/OpenSyntaxHQ/tweak/releases)
[![Platform](https://img.shields.io/badge/platform-linux%20%7C%20macOS%20%7C%20windows%20%7C%20freebsd-lightgrey)](https://github.com/OpenSyntaxHQ/tweak/releases)

**`tweak`** is a fast, cross-platform CLI and interactive TUI for text processing. Encode, hash, transform, generate, and validate — 118 built-in operations, zero dependencies at runtime.

```shell
# Interactive TUI
tweak

# Direct input
tweak md5 "Hello World"
tweak base64-encode "secret"
tweak json-yaml file.json

# Pipe from anything
echo "Hello World" | tweak md5
cat file.txt | tweak base64-encode
curl https://api.example.com/data | tweak json

# Chain operations
tweak md5 "hello" | tweak base64-encode

# Write to file
tweak yaml-json config.yaml > config.json
```

---

## Installation

### Homebrew (macOS / Linux)

```shell
brew install OpenSyntaxHQ/tweak/tweak
```

### Go

```shell
go install github.com/OpenSyntaxHQ/tweak@latest
```

### Quick Install (curl)

```shell
curl -sfL https://raw.githubusercontent.com/OpenSyntaxHQ/tweak/main/install.sh | sh
```

### Binaries

**macOS** (Universal — arm64 + amd64)
```
tweak_Darwin_all.tar.gz
```

**Linux**
| Architecture | Download |
|---|---|
| amd64 | [tweak_Linux_x86_64.tar.gz](https://github.com/OpenSyntaxHQ/tweak/releases/latest/download/tweak_Linux_x86_64.tar.gz) |
| arm64 | [tweak_Linux_arm64.tar.gz](https://github.com/OpenSyntaxHQ/tweak/releases/latest/download/tweak_Linux_arm64.tar.gz) |
| i386 | [tweak_Linux_i386.tar.gz](https://github.com/OpenSyntaxHQ/tweak/releases/latest/download/tweak_Linux_i386.tar.gz) |

**Windows**
| Architecture | Download |
|---|---|
| amd64 | [tweak_Windows_x86_64.zip](https://github.com/OpenSyntaxHQ/tweak/releases/latest/download/tweak_Windows_x86_64.zip) |
| arm64 | [tweak_Windows_arm64.zip](https://github.com/OpenSyntaxHQ/tweak/releases/latest/download/tweak_Windows_arm64.zip) |
| i386 | [tweak_Windows_i386.zip](https://github.com/OpenSyntaxHQ/tweak/releases/latest/download/tweak_Windows_i386.zip) |

**FreeBSD**
| Architecture | Download |
|---|---|
| amd64 | [tweak_Freebsd_x86_64.tar.gz](https://github.com/OpenSyntaxHQ/tweak/releases/latest/download/tweak_Freebsd_x86_64.tar.gz) |
| arm64 | [tweak_Freebsd_arm64.tar.gz](https://github.com/OpenSyntaxHQ/tweak/releases/latest/download/tweak_Freebsd_arm64.tar.gz) |

**Linux Packages**
```shell
# Debian / Ubuntu
dpkg -i tweak_*.deb

# RHEL / Fedora / CentOS
rpm -i tweak_*.rpm

# Arch Linux
pacman -U tweak_*.pkg.tar.zst
```

---

## Usage

```shell
# Launch interactive TUI
tweak

# Get help on any command
tweak --help
tweak md5 --help

# Use with a file
tweak base64-encode image.jpg
tweak md5 file.txt

# Pipe from another command
curl https://jsonplaceholder.typicode.com/users | tweak json
cat file.yaml | tweak yaml-json

# Chain processors
echo "Hello World" | tweak base64-encode | tweak md5

# Write output to file
tweak yaml-json input.yaml > output.json
```

---

## Commands

### 🔐 Encoding & Decoding

| Command | Description |
|---------|-------------|
| `base32-encode` / `base32-decode` | Base32 encode / decode |
| `base58-encode` / `base58-decode` | Base58 encode / decode |
| `base62-encode` / `base62-decode` | Base62 encode / decode |
| `base64-encode` / `base64-decode` | Base64 encode / decode |
| `base64url-encode` / `base64url-decode` | Base64 URL-safe encode / decode |
| `ascii85-encode` / `ascii85-decode` | Ascii85 encode / decode |
| `crockford-encode` / `crockford-decode` | Crockford Base32 encode / decode |
| `hex-encode` / `hex-decode` | Hexadecimal encode / decode |
| `binary-encode` / `binary-decode` | Binary encode / decode |
| `html-encode` / `html-decode` | HTML escape / unescape |
| `url-encode` / `url-decode` | URL encode / decode |

### #️⃣ Hashing & Checksums

| Command | Description |
|---------|-------------|
| `md5` | MD5 checksum |
| `sha1` | SHA-1 checksum |
| `sha224` | SHA-224 checksum |
| `sha256` | SHA-256 checksum |
| `sha384` | SHA-384 checksum |
| `sha512` | SHA-512 checksum |
| `bcrypt` | bcrypt hash |
| `adler32` | Adler-32 checksum |
| `crc32` | CRC-32 checksum |
| `xxh32` / `xxh64` / `xxh128` | xxHash checksums |
| `blake2b` / `blake2s` | BLAKE2 checksums |

### 🔑 Cryptography & Security

| Command | Description |
|---------|-------------|
| `hmac-sha256` / `hmac-sha512` | HMAC-SHA256 / HMAC-SHA512 |
| `aes-encrypt` / `aes-decrypt` | AES-256-GCM encrypt / decrypt |
| `argon2` | Argon2id password hash |
| `jwt-encode` / `jwt-decode` | JWT encode / decode |
| `password-gen` | Secure password generator |
| `totp` | TOTP code generator (RFC 6238) |
| `checksum-verify` | Verify file checksums |

### 🔡 String Operations

| Command | Description |
|---------|-------------|
| `lower` / `upper` / `title` | Change case |
| `snake` / `kebab` / `camel` / `pascal` / `slug` | Convert case format |
| `reverse` | Reverse text |
| `trim` | Trim whitespace |
| `repeat` | Repeat text N times |
| `wrap` | Wrap text at width |
| `replace` | Find and replace |
| `escape-quotes` | Escape single and double quotes |
| `rot13` | ROT13 encode |
| `char-freq` | Character frequency analysis |
| `word-freq` | Word frequency analysis |

### 📋 Line Operations

| Command | Description |
|---------|-------------|
| `sort-lines` | Sort lines alphabetically |
| `reverse-lines` | Reverse line order |
| `shuffle-lines` | Shuffle lines randomly |
| `unique-lines` | Remove duplicate lines |
| `count-lines` | Count lines |
| `number-lines` | Prepend line numbers |

### 📊 JSON & Data Formats

| Command | Description |
|---------|-------------|
| `json` | Format / pretty-print JSON |
| `json-minify` | Minify JSON |
| `json-escape` / `json-unescape` | JSON escape / unescape |
| `json-yaml` / `yaml-json` | Convert between JSON and YAML |
| `json-toml` / `toml-json` | Convert between JSON and TOML |
| `json-xml` / `xml-json` | Convert between JSON and XML |
| `json-csv` / `csv-json` | Convert between JSON and CSV |
| `json-msgpack` / `msgpack-json` | Convert between JSON and MessagePack |
| `jsonl-json` | Convert JSONL/NDJSON to JSON array |

### ⚙️ Generators

| Command | Description |
|---------|-------------|
| `uuid` | UUID v4 generator |
| `ulid` | ULID generator |
| `nanoid` | NanoID generator |
| `lorem` | Lorem ipsum generator |
| `now` | Current timestamp |
| `epoch` | Unix timestamp ↔ human date |
| `password-gen` | Secure password generator |
| `totp` | TOTP code (RFC 6238) |

### 🎨 Developer Utilities

| Command | Description |
|---------|-------------|
| `rgb-hex` | RGB → hex color |
| `hsl-hex` | HSL → hex color |
| `hex-rgb` | Hex → RGB color |
| `base-convert` | Arbitrary base conversion |
| `zeropad` | Pad number with zeros |
| `qrcode` | Generate QR code in terminal |
| `morse-encode` / `morse-decode` | Morse code encode / decode |
| `caesar-encode` / `caesar-decode` | Caesar cipher |
| `markdown-html` | Markdown → HTML |

### 🔍 Text Processing

| Command | Description |
|---------|-------------|
| `grep` | Filter lines by regex |
| `column` | Extract a field by position |
| `pad-left` / `pad-right` | Pad string to width |
| `regex-match` | Extract regex matches |
| `regex-replace` | Replace by regex |
| `remove-spaces` | Remove all spaces |
| `remove-newlines` | Remove all newlines |
| `count-words` / `count-chars` | Word / character count |

### ✅ Validation & Detection

| Command | Description |
|---------|-------------|
| `validate-json` | Validate JSON |
| `validate-email` | Validate email address |
| `validate-url` | Validate URL |
| `detect` | Auto-detect and decode encoding |
| `extract-emails` | Extract all emails from text |
| `extract-urls` | Extract all URLs from text |
| `extract-ips` | Extract all IP addresses from text |

---

## Contributing

Contributions are welcome! See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

## License

Apache 2.0 — see [LICENSE](LICENSE).
