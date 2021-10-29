# demitas

## Usage

```
demitas - TODO

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
```

```sh
demitas \
  -c '{image: busybox, command: [echo, test]}' \
  -t 'networkMode: awsvpc' \
  -s 'networkConfiguration: awsvpcConfiguration: securityGroups: [sg-123]' \
  -- --dry-run
```
