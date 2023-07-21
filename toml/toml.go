package toml

import (
	"encoding/json"
	"io"

	gotoml "github.com/pelletier/go-toml/v2"
)

func FromJSON(input io.Reader) ([]byte, error) {
	var output interface{}

	err := json.NewDecoder(input).Decode(&output)
	if err != nil {
		return nil, err
	}

	return gotoml.Marshal(output)
}

func ToJSON(input io.Reader) ([]byte, error) {
	data, err := io.ReadAll(input)
	if err != nil {
		return nil, err
	}

	var output interface{}
	err = gotoml.Unmarshal(data, &output)
	if err != nil {
		return nil, err
	}

	return json.Marshal(&output)
}
