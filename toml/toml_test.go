package toml_test

import (
	"strings"
	"testing"

	"github.com/elvuel/formato/toml"
)

func TestFromJSON(t *testing.T) {
	inputJSON := `{"name": "John", "age": 30}`
	expectedTOML := "age = 30.0\nname = 'John'\n"

	input := strings.NewReader(inputJSON)
	output, err := toml.FromJSON(input)
	if err != nil {
		t.Errorf("FromJSON returned an error: %v", err)
	}

	if string(output) != expectedTOML {
		t.Errorf("FromJSON returned unexpected output:\nExpected: %s\nActual: %s", expectedTOML, string(output))
	}
}

func TestToJSON(t *testing.T) {
	inputTOML := "name = \"John\"\nage = 30\n"
	expectedJSON := `{"age":30,"name":"John"}`

	input := strings.NewReader(inputTOML)
	output, err := toml.ToJSON(input)
	if err != nil {
		t.Errorf("ToJSON returned an error: %v", err)
	}

	if string(output) != expectedJSON {
		t.Errorf("ToJSON returned unexpected output:\nExpected: %s\nActual: %s", expectedJSON, string(output))
	}
}
