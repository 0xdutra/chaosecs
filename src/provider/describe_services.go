package provider

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecs"
)

type Services struct {
	DesiredCount int64
	PendingCount int64
	RunningCount int64
	Status       string
	ServiceName  string
	ServiceTasks []string
}

func DescribeECSServices(clusterName string, serviceArn string) Services {
	sess := NewSession()
	svc := ecs.New(sess)

	result, err := svc.DescribeServices(&ecs.DescribeServicesInput{
		Cluster: aws.String(clusterName),
		Services: []*string{
			aws.String(serviceArn),
		},
	})

	if err != nil {
		return Services{}
	}

	var services Services

	for _, service := range result.Services {
		services.DesiredCount = *service.DesiredCount
		services.PendingCount = *service.PendingCount
		services.RunningCount = *service.RunningCount
		services.ServiceName = *service.ServiceName
		services.Status = *service.Status
		services.ServiceTasks = ListECSTasks(clusterName, *service.ServiceName)
	}

	return services
}
