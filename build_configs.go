package demitas

type BuildOptions struct {
	EcspressoConfigSrc       string
	ServiceDefSrc            string
	TaskDefSrc               string
	ContainerDefSrc          string
	EcspressoConfigOverrides []byte
	ServiceDefOverrides      []byte
	TaskDefOverrides         []byte
	ContainerDefOverrides    []byte
}

func BuildContainerDefinition(path string, overrides []byte) (*ContainerDefinition, error) {
	containerDef, err := NewContainerDefinition(path)

	if err != nil {
		return nil, err
	}

	if len(overrides) > 0 {
		err = containerDef.Patch(overrides)

		if err != nil {
			return nil, err
		}
	}

	return containerDef, nil
}

func BuildTaskDefinition(path string, overrides []byte, containerDef *ContainerDefinition) (*TaskDefinition, error) {
	taskDef, err := NewTaskDefinition(path)

	if err != nil {
		return nil, err
	}

	err = taskDef.Patch(overrides, containerDef)

	if err != nil {
		return nil, err
	}

	return taskDef, nil
}

func BuildServiceDefinition(path string, overrides []byte) (*ServiceDefinition, error) {
	svrDef, err := NewServiceDefinition(path)

	if err != nil {
		return nil, err
	}

	if len(overrides) > 0 {
		err = svrDef.Patch(overrides)

		if err != nil {
			return nil, err
		}
	}

	return svrDef, nil
}

func BuildEcspressoConfig(path string, overrides []byte) (*EcspressoConfig, error) {
	ecsConf, err := NewEcspressoConfig(path)

	if err != nil {
		return nil, err
	}

	if len(overrides) > 0 {
		err = ecsConf.Patch(overrides)

		if err != nil {
			return nil, err
		}
	}

	return ecsConf, nil
}
