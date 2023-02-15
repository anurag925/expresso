package utils

import (
	"os"

	"gopkg.in/yaml.v2"
)

func LoadYaml(path string, data *map[any]any) error {
	file, err := os.ReadFile(path)

	if err != nil {
		return err
	}
	err2 := yaml.Unmarshal(file, &data)
	if err2 != nil {
		return err
	}
	return nil
}
