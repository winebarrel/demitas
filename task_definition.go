package demitas

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	jsonpatch "github.com/evanphx/json-patch"
	"github.com/winebarrel/demitas/utils"
)

const (
	TaskDefinitionName = "ecs-task-def.json"
)

type TaskDefinition struct {
	Content []byte
}

func NewTaskDefinition(path string) (*TaskDefinition, error) {
	content, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, fmt.Errorf("Failed to load ECS task definition: %w: %s", err, path)
	}

	if filepath.Ext(path) == ".jsonnet" {
		content, err = utils.ParseJsonnet(filepath.Base(path), content)

		if err != nil {
			return nil, fmt.Errorf("Failed to parse ECS task definition jsonnet: %w: %s", err, path)
		}
	} else if !utils.IsJSON(content) {
		return nil, fmt.Errorf("ECS task definition is not JSON: %s", path)
	}

	taskDef := &TaskDefinition{
		Content: content,
	}

	return taskDef, nil
}

func (taskDef *TaskDefinition) Patch(overrides []byte, containerDef *ContainerDefinition) error {
	patchedContent := taskDef.Content
	var err error

	if len(overrides) > 0 {
		patchedContent, err = jsonpatch.MergePatch(patchedContent, overrides)

		if err != nil {
			return fmt.Errorf("Failed to patch ECS task definition: %w", err)
		}
	}

	containerDefinitions := fmt.Sprintf(`{"containerDefinitions":[%s]}`, string(containerDef.Content))
	patchedContent, err = jsonpatch.MergePatch(patchedContent, []byte(containerDefinitions))

	if err != nil {
		return fmt.Errorf("Failed to patch containerDefinitions: %w", err)
	}

	taskDef.Content = patchedContent

	return nil
}
