package demitas

import (
	"fmt"
	"os"

	jsonpatch "github.com/evanphx/json-patch"
	"github.com/valyala/fastjson"
	"github.com/winebarrel/demitas/utils"
)

const (
	ContainerDefinitionName = "ecs-container-def.json"
)

type ContainerDefinition struct {
	Content []byte
}

func NewContainerDefinition(path string, taskDefPath string) (*ContainerDefinition, error) {
	var content []byte
	var err error

	if _, err = os.Stat(path); err != nil {
		content, err = readContainerDefFromTaskDef(taskDefPath)

		if err != nil {
			return nil, fmt.Errorf("failed to load ECS task definition (instead of ECS container definition): %w: %s", err, taskDefPath)
		}
	} else {
		content, err = utils.ReadJSONorJsonnet(path)

		if err != nil {
			return nil, fmt.Errorf("failed to load ECS container definition: %w: %s", err, path)
		}
	}

	containerDef := &ContainerDefinition{
		Content: content,
	}

	return containerDef, nil
}

func (containerDef *ContainerDefinition) Patch(overrides []byte) error {
	patchedContent0, err := jsonpatch.MergePatch(containerDef.Content, []byte(`{"logConfiguration":null}`))

	if err != nil {
		return fmt.Errorf("failed to patch ECS container definition: %w", err)
	}

	patchedContent, err := jsonpatch.MergePatch(patchedContent0, overrides)

	if err != nil {
		return fmt.Errorf("failed to patch ECS container definition: %w", err)
	}

	containerDef.Content = patchedContent

	return nil
}

func readContainerDefFromTaskDef(path string) ([]byte, error) {
	content, err := utils.ReadJSONorJsonnet(path)

	if err != nil {
		return nil, err
	}

	var p fastjson.Parser
	v, err := p.ParseBytes(content)

	if err != nil {
		return nil, err
	}

	containerDef := v.GetObject("containerDefinitions", "0")

	// NOTE: Ignore dependsOn
	containerDef.Del("dependsOn")

	if containerDef == nil {
		return nil, fmt.Errorf("'containerDefinitions.0' is not found in ECS task definition: %s", path)
	}

	return containerDef.MarshalTo(nil), nil
}
