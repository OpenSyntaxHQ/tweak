---
hide:
  - navigation
---

# tweak

![banner](assets/banner.png)

**Fast CLI + TUI for text processing: encoding, hashing, transforms, and more.**

```shell
# Interactive TUI
tweak

# Direct input
tweak md5 "Hello World"
tweak base64-encode "secret"

# Pipe
echo "Hello World" | tweak md5
cat file.yaml | tweak yaml-json

# Chain
tweak md5 "hello" | tweak base64-encode
```

[Get Started](install.md){ .md-button .md-button--primary }
[View on GitHub](https://github.com/OpenSyntaxHQ/tweak){ .md-button }

---

## Features

- **118 built-in operations** — encoding, hashing, crypto, JSON, text, generators, validation
- **Interactive TUI** — fuzzy-searchable, keyboard-driven, no flags needed
- **Pipe-friendly** — works with `cat`, `curl`, `echo`, and every Unix tool
- **Cross-platform** — Linux, macOS, Windows, FreeBSD (amd64 / arm64 / i386)
- **Single binary** — no runtime dependencies, installs in seconds
- **Apache 2.0** — free for personal and commercial use
