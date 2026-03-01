# Usage

## Interactive TUI

Run `tweak` with no arguments to open the interactive TUI. Use `/` to search, arrow keys to navigate, `Enter` to select.

```shell
tweak
```

## Direct Command

```shell
tweak <command> [input] [flags]

# Examples
tweak md5 "Hello World"
tweak base64-encode "hello"
tweak json-yaml config.json
```

## File Input

```shell
tweak base64-encode image.jpg
tweak md5 file.txt
tweak yaml-json config.yaml
```

## Pipe Input

```shell
echo "Hello World" | tweak md5
cat file.txt | tweak base64-encode
curl https://api.example.com | tweak json
```

## Chaining

```shell
# MD5, then base64-encode the result
tweak md5 "hello" | tweak base64-encode

# Multi-step pipeline
echo "Hello World" | tweak base64-encode | tweak md5
```

## Output to File

```shell
tweak yaml-json config.yaml > config.json
tweak json-minify input.json > output.json
```

## Help

```shell
tweak --help
tweak md5 --help
tweak base64-encode --help
```
