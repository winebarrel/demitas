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
