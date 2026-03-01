package main

import (
	"github.com/OpenSyntaxHQ/tweak/cmd"
)

var version = "dev"

//go:generate go run cmd/generate.go
func main() {
	cmd.Version = version
	cmd.Execute()
}
