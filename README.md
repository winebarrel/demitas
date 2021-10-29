# demitas

## Usage

```
demitas - Wrapper for ecspresso that creates task definitions at run time.

  Flags:
       --version                      Displays the program version string.
    -h --help                         Displays help with available flag, subcommand, and positional value parameters.
       --ecspresso-cmd                ecspresso command path. (default: ecspresso)
       --ecspresso-config-src         ecspresso config source path. (default: /Users/sugawara/.demitas/ecspresso.yml)
       --service-def-src              ECS service definition source path. (default: /Users/sugawara/.demitas/ecs-service-def.json)
       --task-def-src                 ECS task definition source path. (default: /Users/sugawara/.demitas/ecs-task-def.json)
       --container-def-src            ECS container definition source path. (default: /Users/sugawara/.demitas/ecs-container-def.json)
       --ecspresso-config-overrides   JSON/YAML string that overrides ecspresso config source.
    -s --service-def-overrides        JSON/YAML string that overrides ECS service definition source.
    -t --task-def-overrides           JSON/YAML string that overrides ECS task definition source.
    -c --container-def-overrides      JSON/YAML string that overrides ECS container definition source.

  Trailing Arguments:
    Arguments after "--" is passed to "ecspresso run".
    e.g. demitas -c 'image: ...' -- --color --wait-until=running --debug
                                 ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

  Environment Variables:
    DEMITAS_CONFIGS_DIR   Configuration file base directory.  (default: /Users/sugawara/.demitas)
    DEMITAS_PROFILE       Configuration profile directory.  (default: "")
                          If "database" is set, configs file will be read from "$DEMITAS_CONFIGS_DIR/database"
```
