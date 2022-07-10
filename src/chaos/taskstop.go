package chaos

import (
	"chaosecs/src/provider"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecs"
)

func taskstop(clusterName string, taskArn string) error {
	sess := provider.NewSession()
	svc := ecs.New(sess)

	_, err := svc.StopTask(&ecs.StopTaskInput{
		Cluster: aws.String(clusterName),
		Task:    aws.String(taskArn),
		Reason:  aws.String("Task stopped by chaosecs"),
	})

	if err != nil {
		return err
	}

	return nil
}
