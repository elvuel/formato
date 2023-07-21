package formato

import (
	"io"

	"github.com/elvuel/formato/toml"
	"github.com/elvuel/formato/yaml"
)

// Y2J converts YAML to JSON
func Y2J(input io.Reader) ([]byte, error) {
	return yaml.ToJSON(input)
}

// J2Y converts JSON to YAML
func J2Y(input io.Reader) ([]byte, error) {
	return yaml.FromJSON(input)
}

// T2J converts TOML to JSON
func T2J(input io.Reader) ([]byte, error) {
	return toml.ToJSON(input)
}

// J2T converts JSON to TOML
func J2T(input io.Reader) ([]byte, error) {
	return toml.FromJSON(input)
}
