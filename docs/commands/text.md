# Text Processing

| Command | Description |
|---------|-------------|
| `grep` | Filter lines by regex |
| `column` | Extract field by position |
| `pad-left` / `pad-right` | Pad string to width |
| `regex-match` | Extract regex matches |
| `regex-replace` | Replace by regex |
| `remove-spaces` | Remove all spaces |
| `remove-newlines` | Remove all newlines |
| `count-words` | Word count |
| `count-chars` | Character count |
| `sort-lines` | Sort alphabetically |
| `reverse-lines` | Reverse line order |
| `shuffle-lines` | Shuffle randomly |
| `unique-lines` | Remove duplicates |
| `count-lines` | Line count |
| `number-lines` | Prepend line numbers |
| `morse-encode` / `morse-decode` | Morse code |
| `caesar-encode` / `caesar-decode` | Caesar cipher |
| `markdown-html` | Markdown → HTML |
| `base-convert` | Arbitrary base conversion |
| `zeropad` | Zero-pad a number |

## Examples

```shell
cat file.txt | tweak grep --pattern "^error"
echo "one two three" | tweak column --field 2
echo "5" | tweak pad-left --width 4 --char 0
echo "hello" | tweak regex-replace --pattern "l+" --replace "L"
echo "SOS" | tweak morse-encode
tweak zeropad --n 6 "42"
```
