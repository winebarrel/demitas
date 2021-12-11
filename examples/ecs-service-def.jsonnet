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
  //capacityProviderStrategy: [
  //  {
  //    capacityProvider: 'FARGATE_SPOT',
  //    weight: 1,
  //  },
  //],
}
