package yaml

import (
	"bytes"
	"io"

	"github.com/bronze1man/yaml2json/y2jLib"
	"github.com/itchyny/json2yaml"
)

// not use `gopkg.in/yaml.v3` directly

// FromJSON converts JSON to YAML
func FromJSON(input io.Reader) ([]byte, error) {
	var output bytes.Buffer
	if err := json2yaml.Convert(&output, input); err != nil {
		return nil, err
	}

	return output.Bytes(), nil
}

// ToJSON converts YAML to JSON
func ToJSON(input io.Reader) ([]byte, error) {
	var output bytes.Buffer
	if err := y2jLib.TranslateStream(input, &output); err != nil {
		return nil, err
	}

	return output.Bytes(), nil
}
