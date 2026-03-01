# Contributing to Tweak

Thank you for your interest in contributing!

## Development Setup

```bash
git clone https://github.com/OpenSyntaxHQ/tweak.git
cd tweak
go mod tidy
go generate ./...
go build ./...
go test ./tests/...
```

## Adding a New Processor

1. Create a file in `processors/` implementing the `Processor` interface:

```go
type MyProcessor struct{}

func (p MyProcessor) Name() string        { return "my-proc" }
func (p MyProcessor) Alias() []string     { return nil }
func (p MyProcessor) Title() string       { return fmt.Sprintf("My Processor (%s)", p.Name()) }
func (p MyProcessor) Description() string { return "Short description" }
func (p MyProcessor) FilterValue() string { return p.Title() }
func (p MyProcessor) Flags() []Flag       { return nil }
func (p MyProcessor) Transform(data []byte, f ...Flag) (string, error) {
    // implementation
}
```

2. Register it in `processors/all.go` (the list returned by `All()`).

3. Add a test in `tests/` using the table-driven pattern:

```go
func TestMyProcessor_Transform(t *testing.T) {
    tests := []struct{ name, in, want string }{
        {"basic", "input", "output"},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            assertTransform(t, processors.MyProcessor{}, tt.in, nil, tt.want, false)
        })
    }
}
```

## Commit Guidelines

- Use [Conventional Commits](https://www.conventionalcommits.org/)
- Sign off every commit: `git commit -s`
- Keep commits small and focused

## Running Tests

```bash
go test -v -race -count=1 ./tests/...
```

## License

By contributing, you agree your contributions will be licensed under the Apache 2.0 License.
