{
  family: 'my-oneshot-task',
  cpu: '256',
  memory: '512',
  networkMode: 'awsvpc',
  taskRoleArn: 'arn:aws:iam::xxx:role/my-role',
  executionRoleArn: 'arn:aws:iam::xxx:role/my-exec-role',
  requiresCompatibilities: ['FARGATE'],
}
