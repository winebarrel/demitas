package demitas

import (
	"fmt"
	"io/ioutil"
	"os"

	jsonpatch "github.com/evanphx/json-patch"
	"github.com/winebarrel/demitas/utils"
)

type RunOptions struct {
	EcspressoPath    string
	EcspressoOptions []string
	PrintConfig      bool
}

type Runner struct {
	*RunOptions
}

func RunTask(opts *RunOptions, ecsConf *EcspressoConfig, svrDef *ServiceDefinition, taskDef *TaskDefinition) (err error) {
	runInTempDir(func() {
		err = writeTemporaryConfigs(ecsConf, svrDef, taskDef)

		if err != nil {
			return
		}

		cmdWithArgs := []string{opts.EcspressoPath, "run"}

		if len(opts.EcspressoOptions) > 0 {
			cmdWithArgs = append(cmdWithArgs, opts.EcspressoOptions...)
		}

		err = runCommand(cmdWithArgs)

		if err != nil {
			return
		}
	})

	return
}

func runInTempDir(callback func()) {
	pwd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	tmp, err := ioutil.TempDir("", "demitas")

	if err != nil {
		panic(err)
	}

	defer func() {
		_ = os.Chdir(pwd)
		os.RemoveAll(tmp)
	}()

	err = os.Chdir(tmp)

	if err != nil {
		panic(err)
	}

	callback()
}

func writeTemporaryConfigs(ecsConf *EcspressoConfig, svrDef *ServiceDefinition, taskDef *TaskDefinition) error {
	err := ioutil.WriteFile(TaskDefinitionName, taskDef.Content, os.FileMode(0o644))

	if err != nil {
		return fmt.Errorf("failed to write ECS task definition: %w", err)
	}

	err = ioutil.WriteFile(ServiceDefinitionName, svrDef.Content, os.FileMode(0o644))

	if err != nil {
		return fmt.Errorf("failed to write ECS service definition: %w", err)
	}

	ecsConfOverrides := fmt.Sprintf(`{"service_definition":"%s","task_definition":"%s"}`, ServiceDefinitionName, TaskDefinitionName)
	ecsConfJson, err := jsonpatch.MergePatch(ecsConf.Content, []byte(ecsConfOverrides))

	if err != nil {
		return fmt.Errorf("failed to update ecspresso config: %w", err)
	}

	ecsConfYaml, err := utils.JSONToYAML(ecsConfJson)

	if err != nil {
		return fmt.Errorf("failed to convert ecspresso config: %w", err)
	}

	err = ioutil.WriteFile(EcspressoConfigName, ecsConfYaml, os.FileMode(0o644))

	if err != nil {
		return fmt.Errorf("failed to write ecspresso config: %w", err)
	}

	return nil
}
