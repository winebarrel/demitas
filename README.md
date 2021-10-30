# demitas

Wrapper for [ecspresso](https://github.com/kayac/ecspresso) that creates task definitions at run time.

## Usage

```
demitas - Wrapper for ecspresso that creates task definitions at run time.

  Flags:
       --version         Displays the program version string.
    -h --help            Displays help with available flag, subcommand, and positional value parameters.
       --ecsp-cmd        ecspresso command path. (default: ecspresso)
    -d --conf-dir        Configuration file base directory. (default: ~/.demitas)
    -p --profile         Configuration profile directory.
    -E --ecsp-conf-src   ecspresso config source path. (default: ~/.demitas/ecspresso.yml)
    -S --svr-def-src     ECS service definition source path. (default: ~/.demitas/ecs-service-def.jsonnet)
    -T --task-def-src    ECS task definition source path. (default: ~/.demitas/ecs-task-def.jsonnet)
    -C --cont-def-src    ECS container definition source path. (default: ~/.demitas/ecs-container-def.jsonnet)
    -e --ecsp-conf-ovr   JSON/YAML string that overrides ecspresso config source.
    -s --svr-def-ovr     JSON/YAML string that overrides ECS service definition source.
    -t --task-def-ovr    JSON/YAML string that overrides ECS task definition source.
    -c --cont-def-ovr    JSON/YAML string that overrides ECS container definition source.
    -n --print-conf      Display configs only.

  Trailing Arguments:
    Arguments after "--" is passed to "ecspresso run".
    e.g. demitas -c 'image: ...' -- --color --wait-until=running --debug
                                 ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

  Environment Variables:
    DMTS_CONF_DIR (--conf-dir)   Configuration file base directory.  (default: ~/.demitas)
    DMTS_PROFILE  (--profile)    Configuration profile directory.
                                 If "database" is set, configs file will be read from "$DMTS_CONF_DIR/database/..."
```

## Installation

```sh
brew tap winebarrel/demitas
brew install demitas
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
  //launchType: 'FARGATE',
  networkConfiguration: {
    awsvpcConfiguration: {
      assignPublicIp: 'DISABLED',
      securityGroups: ['sg-xxx'],
      subnets: ['subnet-xxx'],
    },
  },
  enableExecuteCommand: true,
  capacityProviderStrategy: [
    {
      capacityProvider: 'FARGATE_SPOT',
      weight: 1,
    },
  ],
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
$ demitas \
  -e 'service: my-service2' \
  -s 'networkConfiguration: {awsvpcConfiguration: {securityGroups: [sg-zzz]}}' \
  -c '{image: "public.ecr.aws/runecast/busybox:1.33.1", command: [echo, hello]}' \
  -- --dry-run

2021/10/10 22:33:44 my-service2/my-cluster Running task
2021/10/10 22:33:44 my-service2/my-cluster task definition:
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
2021/10/10 22:33:44 my-service2/my-cluster DRY RUN OK
```

## FAQ

**Q:** Will the created ECS task definitions be deleted?

**A:** No.

##

**Q:** Will IAM roles be created automatically?

**A:** No.
