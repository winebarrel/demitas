package utils

import (
	"encoding/json"

	"github.com/goccy/go-yaml"
)

func IsJSON(data []byte) bool {
	var js json.RawMessage
	return json.Unmarshal(data, &js) == nil
}

func YAMLToJSON(data []byte) ([]byte, error) {
	var v interface{}

	if err := yaml.UnmarshalWithOptions(data, &v, yaml.UseOrderedMap()); err != nil {
		return nil, err
	}

	js, err := yaml.MarshalWithOptions(v, yaml.JSON())

	if err != nil {
		return nil, err
	}

	return js, nil
}

func JSONToYAML(bytes []byte) ([]byte, error) {
	var v interface{}

	if err := yaml.UnmarshalWithOptions(bytes, &v, yaml.UseOrderedMap()); err != nil {
		return nil, err
	}

	ym, err := yaml.Marshal(v)

	if err != nil {
		return nil, err
	}
	return ym, nil
}
