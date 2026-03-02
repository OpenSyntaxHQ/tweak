package processors

import "fmt"

type Processor interface {
	Name() string
	Alias() []string
	Transform(data []byte, opts ...Flag) (string, error)
	Flags() []Flag
}

type Generator interface {
	IsGenerator() bool
}

func IsGenerator(p Processor) bool {
	if g, ok := p.(Generator); ok {
		return g.IsGenerator()
	}
	return false
}

type FlagType string

func (f FlagType) String() string { return string(f) }
func (f FlagType) IsString() bool { return f == FlagString }

const (
	FlagInt    FlagType = "Int"
	FlagUint   FlagType = "Uint"
	FlagBool   FlagType = "Bool"
	FlagString FlagType = "String"
)

type Flag struct {
	Name      string
	Short     string
	Desc      string
	Type      FlagType
	Value     any
	Required  bool
	Sensitive bool
}

func (f Flag) HelpLabel() string {
	if f.Required {
		return fmt.Sprintf("--%s (required)", f.Name)
	}
	return fmt.Sprintf("--%s", f.Name)
}
