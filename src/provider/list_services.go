package provider

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecs"
)

func ListECSServices(clusterName string) ([]string, error) {
	sess := NewSession()
	svc := ecs.New(sess)

	var nextToken string
	var servicesArns []string

	for {
		result, err := svc.ListServices(&ecs.ListServicesInput{
			Cluster:   aws.String(clusterName),
			NextToken: &nextToken,
		})

		if err != nil {
			return nil, err
		}

		for _, value := range result.ServiceArns {
			servicesArns = append(servicesArns, *value)
		}

		if result.NextToken == nil {
			break
		}

		nextToken = *result.NextToken
	}

	return servicesArns, nil
}
