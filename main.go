package main

import (
	"chaosecs/src/chaos"
	"chaosecs/src/provider"
	"flag"
)

var (
	clusterName *string
)

func init() {
	clusterName = flag.String("cluster", "", "The name of the ECS cluster")
}

func main() {
	flag.Parse()

	servicesArns, _ := provider.ListECSServices(*clusterName)
	var services []provider.Services

	for _, serviceArn := range servicesArns {
		services = append(services,
			provider.DescribeECSServices(*clusterName, serviceArn),
		)
	}

	chaos.ChaosECS(*clusterName, services)
}
