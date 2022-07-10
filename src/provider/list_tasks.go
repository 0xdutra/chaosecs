package provider

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecs"
)

func ListECSTasks(clusterName string, serviceName string) []string {
	sess := NewSession()
	svc := ecs.New(sess)

	var nextToken string
	var taskArns []string

	for {
		result, err := svc.ListTasks(&ecs.ListTasksInput{
			Cluster:     aws.String(clusterName),
			ServiceName: aws.String(serviceName),
			NextToken:   &nextToken,
		})

		if err != nil {
			log.Fatal(err)
		}

		for _, taskArn := range result.TaskArns {
			taskArns = append(taskArns, *taskArn)
		}

		if result.NextToken == nil {
			break
		}

		nextToken = *result.NextToken
	}

	return taskArns
}
