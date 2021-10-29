package demitas

import (
	"fmt"
	"io/ioutil"

	jsonpatch "github.com/evanphx/json-patch"
	"github.com/winebarrel/demitas/utils"
)

const (
	EcspressoConfigName = "ecspresso.yml"
)

type EcspressoConfig struct {
	Content []byte
}

func NewEcspressoConfig(path string) (*EcspressoConfig, error) {
	content, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, fmt.Errorf("Failed to load ecspresso config: %w: %s", err, path)
	}

	js, err := utils.YAMLToJSON(content)

	if err != nil {
		return nil, fmt.Errorf("Failed to parse ecspresso config: %w: %s", err, path)
	}

	ecsConf := &EcspressoConfig{
		Content: js,
	}

	return ecsConf, nil
}

func (ecsConf *EcspressoConfig) Patch(overrides []byte) error {
	patchedContent, err := jsonpatch.MergePatch(ecsConf.Content, overrides)

	if err != nil {
		return fmt.Errorf("Failed to patch ecspresso config: %w", err)
	}

	ecsConf.Content = patchedContent

	return nil
}
