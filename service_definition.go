package demitas

import (
	"fmt"

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
	content, err := utils.ReadJSONorJsonnet(path)

	if err != nil {
		return nil, fmt.Errorf("failed to load ECS service definition: %w: %s", err, path)
	}

	svrDef := &ServiceDefinition{
		Content: content,
	}

	return svrDef, nil
}

func (svrDef *ServiceDefinition) Patch(overrides []byte) error {
	patchedContent, err := jsonpatch.MergePatch(svrDef.Content, overrides)

	if err != nil {
		return fmt.Errorf("failed to patch ECS service definition: %w", err)
	}

	svrDef.Content = patchedContent

	return nil
}

func (svrDef *ServiceDefinition) Print() {
	fmt.Printf("# %s\n%s\n", ServiceDefinitionName, utils.PrettyJSON(svrDef.Content))
}
