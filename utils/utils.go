package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

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

func EvaluateJsonnet(filename string) ([]byte, error) {
	vm := jsonnet.MakeVM()
	js, err := vm.EvaluateFile(filename)

	if err != nil {
		return nil, err
	}

	return []byte(js), nil
}

func PrettyJSON(data []byte) string {
	var js json.RawMessage
	_ = json.Unmarshal(data, &js)
	js, _ = json.MarshalIndent(js, "", "  ")
	return string(js)
}

func ReadJSONorJsonnet(path string) ([]byte, error) {
	var content []byte

	_, err := os.Stat(path)

	if err != nil {
		return nil, err
	}

	if filepath.Ext(path) == ".jsonnet" {
		content, err = EvaluateJsonnet(path)

		if err != nil {
			return nil, err
		}
	} else {
		content, err = ioutil.ReadFile(path)

		if err != nil {
			return nil, err
		}

		if !IsJSON(content) {
			return nil, fmt.Errorf("definition is not JSON: %s", path)
		}
	}

	return content, nil
}
