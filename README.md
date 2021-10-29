# demitas

## Usage

```sh
demitas \
  -c '{image: busybo, command: [echo, test]}' \
  -t 'networkMode: awsvpc' \
  -s 'networkConfiguration: awsvpcConfiguration: securityGroups: [sg-123]' \
  -- --dry-run
```
