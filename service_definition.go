package demitas

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	jsonpatch "github.com/evanphx/json-patch"
	"github.com/winebarrel/demitas/utils"
)

const (
	ServiceDefinitionName = "ecs-service-def.json"
)

type ServiceDefinition struct {
	Content []byte
}

func NewServiceDefinition(path string) (*ServiceDefinition, error) {
	content, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, fmt.Errorf("Failed to load ECS service definition: %w: %s", err, path)
	}

	if filepath.Ext(path) == ".jsonnet" {
		content, err = utils.ParseJsonnet(filepath.Base(path), content)

		if err != nil {
			return nil, fmt.Errorf("Failed to parse ECS service definition jsonnet: %w: %s", err, path)
		}
	} else if !utils.IsJSON(content) {
		return nil, fmt.Errorf("ECS service definition is not JSON: %s", path)
	}

	svrDef := &ServiceDefinition{
		Content: content,
	}

	return svrDef, nil
}

func (svrDef *ServiceDefinition) Patch(overrides []byte) error {
	patchedContent, err := jsonpatch.MergePatch(svrDef.Content, overrides)

	if err != nil {
		return fmt.Errorf("Failed to patch ECS service definition: %w", err)
	}

	svrDef.Content = patchedContent

	return nil
}

func (svrDef *ServiceDefinition) Print() {
	fmt.Printf("# %s\n%s\n", ServiceDefinitionName, utils.PrettyJSON(svrDef.Content))
}
