package demitas

import (
	"fmt"

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
	content, err := utils.ReadJSONorJsonnet(path)

	if err != nil {
		return nil, fmt.Errorf("failed to load ECS task definition: %w: %s", err, path)
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
			return fmt.Errorf("failed to patch ECS task definition: %w", err)
		}
	}

	containerDefinitions := fmt.Sprintf(`{"containerDefinitions":[%s]}`, string(containerDef.Content))
	patchedContent, err = jsonpatch.MergePatch(patchedContent, []byte(containerDefinitions))
	fmt.Println(utils.PrettyJSON(patchedContent))

	if err != nil {
		return fmt.Errorf("failed to patch containerDefinitions: %w", err)
	}

	taskDef.Content = patchedContent

	return nil
}

func (taskDef *TaskDefinition) Print() {
	fmt.Printf("# %s\n%s\n", TaskDefinitionName, utils.PrettyJSON(taskDef.Content))
}
