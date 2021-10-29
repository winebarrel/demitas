package main

import (
	"log"
	"os"
	"os/user"
	"path/filepath"

	"github.com/integrii/flaggy"
	"github.com/winebarrel/demitas"
	"github.com/winebarrel/demitas/utils"
)

type Options struct {
	demitas.RunOptions
	demitas.BuildOptions
}

var version string

const (
	Description               = "Wrapper for ecspresso that creates task definitions at run time."
	DefaultEcspressoCmd       = "ecspresso"
	DefaultConfigsDir         = ".demitas"
	DefaultEcspressoConfigSrc = "ecspresso.yml"
	DefaultServiceDefSrc      = "ecs-service-def.json"
	DefaultTaskDefSrc         = "ecs-task-def.json"
	DefaultContainerDefSrc    = "ecs-container-def.json"
)

func getEnv(k, defval string) string {
	v, ok := os.LookupEnv(k)

	if ok {
		return v
	} else {
		return defval
	}
}

func parseArgs() *Options {
	currUser, err := user.Current()

	if err != nil {
		log.Fatalf("Failed to get the current user: %s", err)
	}

	configs_dir := getEnv("DEMITAS_CONFIGS_DIR", DefaultConfigsDir)
	profile := getEnv("DEMITAS_PROFILE", "")

	opts := &Options{
		RunOptions: demitas.RunOptions{
			EcspressoPath: DefaultEcspressoCmd,
		},
		BuildOptions: demitas.BuildOptions{
			EcspressoConfigSrc: filepath.Join(currUser.HomeDir, configs_dir, profile, DefaultEcspressoConfigSrc),
			ServiceDefSrc:      filepath.Join(currUser.HomeDir, configs_dir, profile, DefaultServiceDefSrc),
			TaskDefSrc:         filepath.Join(currUser.HomeDir, configs_dir, profile, DefaultTaskDefSrc),
			ContainerDefSrc:    filepath.Join(currUser.HomeDir, configs_dir, profile, DefaultContainerDefSrc),
		},
	}

	var ecsConfOverridesStr string
	var svrDefOverridesStr string
	var taskDefOverridesStr string
	var containerDefOverridesStr string

	flaggy.DefaultParser.AdditionalHelpAppend = `
  Trailing Arguments:
    Arguments after "--" is passed to "ecspresso run".
    e.g. demitas -c 'image: ...' -- --color --wait-until=running --debug
                                 ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

  Environment Variables:
    DEMITAS_CONFIGS_DIR   Configuration file base directory.  (default: ` + filepath.Join(currUser.HomeDir, DefaultConfigsDir) + `)
    DEMITAS_PROFILE       Configuration profile directory.  (default: "")
                          If "database" is set, configs file will be read from "$DEMITAS_CONFIGS_DIR/database"
`

	flaggy.SetDescription(Description)
	flaggy.SetVersion(version)
	flaggy.String(&opts.EcspressoPath, "", "ecspresso-cmd", "ecspresso command path.")
	flaggy.String(&opts.EcspressoConfigSrc, "", "ecspresso-config-src", "ecspresso config source path.")
	flaggy.String(&opts.ServiceDefSrc, "", "service-def-src", "ECS service definition source path.")
	flaggy.String(&opts.TaskDefSrc, "", "task-def-src", "ECS task definition source path.")
	flaggy.String(&opts.ContainerDefSrc, "", "container-def-src", "ECS container definition source path.")
	flaggy.String(&ecsConfOverridesStr, "", "ecspresso-config-overrides", "JSON/YAML string that overrides ecspresso config source.")
	flaggy.String(&svrDefOverridesStr, "s", "service-def-overrides", "JSON/YAML string that overrides ECS service definition source.")
	flaggy.String(&taskDefOverridesStr, "t", "task-def-overrides", "JSON/YAML string that overrides ECS task definition source.")
	flaggy.String(&containerDefOverridesStr, "c", "container-def-overrides", "JSON/YAML string that overrides ECS container definition source.")
	flaggy.Parse()

	opts.EcspressoConfigOverrides = []byte(ecsConfOverridesStr)
	opts.ServiceDefOverrides = []byte(svrDefOverridesStr)
	opts.TaskDefOverrides = []byte(taskDefOverridesStr)
	opts.ContainerDefOverrides = []byte(containerDefOverridesStr)

	if len(opts.EcspressoConfigOverrides) > 0 && !utils.IsJSON(opts.EcspressoConfigOverrides) {
		js, err := utils.YAMLToJSON(opts.EcspressoConfigOverrides)

		if err != nil {
			log.Fatalf("'--ecspresso-config-overrides' value is not valid: %s", string(opts.EcspressoConfigOverrides))
		}

		opts.EcspressoConfigOverrides = js
	}

	if len(opts.ServiceDefOverrides) > 0 && !utils.IsJSON(opts.ServiceDefOverrides) {
		js, err := utils.YAMLToJSON(opts.ServiceDefOverrides)

		if err != nil {
			log.Fatalf("'--service-def-overrides' value is not valid: %s", string(opts.ServiceDefOverrides))
		}

		opts.ServiceDefOverrides = js
	}

	if len(opts.TaskDefOverrides) > 0 && !utils.IsJSON(opts.TaskDefOverrides) {
		js, err := utils.YAMLToJSON(opts.TaskDefOverrides)

		if err != nil {
			log.Fatalf("'--task-def-overrides' value is not valid: %s", string(opts.TaskDefOverrides))
		}

		opts.TaskDefOverrides = js
	}

	if len(opts.ContainerDefOverrides) > 0 && !utils.IsJSON(opts.ContainerDefOverrides) {
		js, err := utils.YAMLToJSON(opts.ContainerDefOverrides)

		if err != nil {
			log.Fatalf("'--container-def-overrides' value is not valid: %s", string(opts.ContainerDefOverrides))
		}

		opts.ContainerDefOverrides = js
	}

	opts.EcspressoOptions = make([]string, len(flaggy.TrailingArguments))
	copy(opts.EcspressoOptions, flaggy.TrailingArguments)

	return opts
}
