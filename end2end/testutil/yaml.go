package testutil

import (
	"gopkg.in/yaml.v3"
	"os"
)

type TestData struct {
	Name    string
	Data    map[string]any
	Expect  map[string]any
	Control map[string]any
}

func ReadYamlFile(dir string) ([]TestData, error) {
	var result []TestData
	data, err := os.ReadFile(dir)
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}
