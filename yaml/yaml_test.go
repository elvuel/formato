package yaml_test

import (
	"strings"
	"testing"

	"github.com/elvuel/formato/yaml"
)

func TestFromJSON(t *testing.T) {
	inputJSON := `{"name": "John", "age": 30}`
	expectedYAML := "name: John\nage: 30\n"

	input := strings.NewReader(inputJSON)
	output, err := yaml.FromJSON(input)
	if err != nil {
		t.Errorf("FromJSON returned an error: %v", err)
	}

	if string(output) != expectedYAML {
		t.Errorf("FromJSON returned unexpected output:\nExpected: %s\nActual: %s", expectedYAML, string(output))
	}
}

func TestToJSON(t *testing.T) {
	inputYAML := "name: John\nage: 30\n"
	expectedJSON := `{"age":30,"name":"John"}`

	input := strings.NewReader(inputYAML)
	output, err := yaml.ToJSON(input)
	if err != nil {
		t.Errorf("ToJSON returned an error: %v", err)
	}
	chomped := strings.TrimSuffix(string(output), "\n")

	if chomped != expectedJSON {
		t.Errorf("ToJSON returned unexpected output:\nExpected: %#v\nActual: %#v", expectedJSON, string(output))
	}
}
