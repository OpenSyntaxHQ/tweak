package processors

import (
	"fmt"

	"github.com/google/uuid"
)

type UUID struct{}

func (p UUID) Name() string    { return "uuid" }
func (p UUID) Alias() []string { return []string{"uuid4", "gen-uuid"} }

func (p UUID) Transform(_ []byte, _ ...Flag) (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

func (p UUID) IsGenerator() bool    { return true }
func (p UUID) Flags() []Flag       { return nil }
func (p UUID) Title() string       { return fmt.Sprintf("UUID v4 Generator (%s)", p.Name()) }
func (p UUID) Description() string { return "Generate a UUID v4" }
func (p UUID) FilterValue() string { return p.Title() }
