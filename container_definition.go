package demitas

import (
	"fmt"
	"io/ioutil"

	jsonpatch "github.com/evanphx/json-patch"
	"github.com/winebarrel/demitas/utils"
)

const (
	ContainerDefinitionName = "ecs-container-def.json"
)

type ContainerDefinition struct {
	Content []byte
}

func NewContainerDefinition(path string) (*ContainerDefinition, error) {
	content, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, fmt.Errorf("Failed to load ECS container definition: %w: %s", err, path)
	}

	if !utils.IsJSON(content) {
		return nil, fmt.Errorf("ECS container definition is not JSON: %s", path)
	}

	containerDef := &ContainerDefinition{
		Content: content,
	}

	return containerDef, nil
}

func (containerDef *ContainerDefinition) Patch(overrides []byte) error {
	patchedContent, err := jsonpatch.MergePatch(containerDef.Content, overrides)

	if err != nil {
		return fmt.Errorf("Failed to patch ECS container definition: %w", err)
	}

	containerDef.Content = patchedContent

	return nil
}
