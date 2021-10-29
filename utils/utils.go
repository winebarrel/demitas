package utils

import (
	"encoding/json"

	"github.com/goccy/go-yaml"
	"github.com/google/go-jsonnet"
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

func JSONToYAML(data []byte) ([]byte, error) {
	var v interface{}

	if err := yaml.UnmarshalWithOptions(data, &v, yaml.UseOrderedMap()); err != nil {
		return nil, err
	}

	ym, err := yaml.Marshal(v)

	if err != nil {
		return nil, err
	}
	return ym, nil
}

func ParseJsonnet(filename string, data []byte) ([]byte, error) {
	vm := jsonnet.MakeVM()
	js, err := vm.EvaluateAnonymousSnippet(filename, string(data))

	if err != nil {
		return nil, err
	}

	return []byte(js), nil
}
