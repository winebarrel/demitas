# demitas

## Usage

```sh
demitas \
  -c '{image: busybox, command: [echo, test]}' \
  -t 'networkMode: awsvpc' \
  -s 'networkConfiguration: awsvpcConfiguration: securityGroups: [sg-123]' \
  -- --dry-run
```
