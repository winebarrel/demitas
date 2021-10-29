# demitas

## Usage

```
demitas - Wrapper for ecspresso that creates task definitions at run time.

  Flags:
       --version                      Displays the program version string.
    -h --help                         Displays help with available flag, subcommand, and positional value parameters.
       --ecspresso-cmd                ecspresso command path. (default: ecspresso)
       --ecspresso-config-src         ecspresso config source path. (default: ~/.demitas/ecspresso.yml)
       --service-def-src              ECS service definition source path. (default: ~/.demitas/ecs-service-def.jsonnet)
       --task-def-src                 ECS task definition source path. (default: ~/.demitas/ecs-task-def.jsonnet)
       --container-def-src            ECS container definition source path. (default: ~/.demitas/ecs-container-def.jsonnet)
       --ecspresso-config-overrides   JSON/YAML string that overrides ecspresso config source.
    -s --service-def-overrides        JSON/YAML string that overrides ECS service definition source.
    -t --task-def-overrides           JSON/YAML string that overrides ECS task definition source.
    -c --container-def-overrides      JSON/YAML string that overrides ECS container definition source.

  Trailing Arguments:
    Arguments after "--" is passed to "ecspresso run".
    e.g. demitas -c 'image: ...' -- --color --wait-until=running --debug
                                 ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

  Environment Variables:
    DEMITAS_CONFIGS_DIR   Configuration file base directory.  (default: ~/.demitas)
    DEMITAS_PROFILE       Configuration profile directory.  (default: "")
                          If "database" is set, configs file will be read from "$DEMITAS_CONFIGS_DIR/database/..."
```

## Example Configurations


### `~/.demitas/ecspresso.yml`

```yml
region: ap-northeast-1
cluster: my-cluster
service: my-service
```

### `~/.demitas/ecs-service-def.jsonnet`

```jsonnet
{
  launchType: 'FARGATE',
  networkConfiguration: {
    awsvpcConfiguration: {
      assignPublicIp: 'DISABLED',
      securityGroups: ['sg-xxx'],
      subnets: ['subnet-xxx'],
    },
  },
  enableExecuteCommand: true,
}
```

### `~/.demitas/ecs-task-def.jsonnet`

```jsonnet
{
  family: 'my-oneshot-task',
  cpu: '256',
  memory: '512',
  networkMode: 'awsvpc',
  taskRoleArn: 'arn:aws:iam::xxx:role/my-role',
  executionRoleArn: 'arn:aws:iam::xxx:role/my-exec-role',
  requiresCompatibilities: ['FARGATE'],
}
```

### `~/.demitas/ecs-container-def.jsonnet`

```jsonnet
{
  name: 'oneshot',
  cpu: 0,
  essential: true,
  logConfiguration: {
    logDriver: 'awslogs',
    options: {
      'awslogs-group': '/ecs/oneshot',
      'awslogs-region': 'ap-northeast-1',
      'awslogs-stream-prefix': 'ecs',
    },
  },
}
```

## Execution Example

```sh
$ demitas -c '{command: [echo, hello], image: "public.ecr.aws/runecast/busybox:1.33.1"}' -- --dry-run
2021/10/10 22:33:44 my-cluster/my-service Running task
2021/10/10 22:33:44 my-cluster/my-service task definition:
{
  "containerDefinitions": [
    {
      "command": [
        "echo",
        "hello"
      ],
      "cpu": 0,
      "essential": true,
      "image": "public.ecr.aws/runecast/busybox:1.33.1",
      "logConfiguration": {
        "logDriver": "awslogs",
        "options": {
          "awslogs-group": "/ecs/busybox",
          "awslogs-region": "ap-northeast-1",
          "awslogs-stream-prefix": "ecs"
        }
      },
      "name": "oneshot"
    }
  ],
  "cpu": "256",
  "executionRoleArn": "arn:aws:iam::xxx:role/my-role",
  "family": "my-oneshot-task",
  "memory": "512",
  "networkMode": "awsvpc",
  "requiresCompatibilities": [
    "FARGATE"
  ],
  "taskRoleArn": "arn:aws:iam::xxx:role/my-exec-role",
}
2021/10/10 22:33:44 my-cluster/my-service DRY RUN OK
```

## FAQ

**Q:** Will the created ECS task definitions be deleted?

**A:** No.
