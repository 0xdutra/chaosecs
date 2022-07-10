package chaos

import (
	"chaosecs/src/provider"
	"fmt"
)

func ChaosECS(clusterName string, services []provider.Services) {
	for _, svc := range services {

		if validate(svc.DesiredCount, svc.PendingCount, svc.ServiceTasks, svc.Status) {
			task := random(svc.ServiceTasks)

			if err := taskstop(clusterName, task); err == nil {
				fmt.Printf("Stopping task %s of the service %s\n", task, svc.ServiceName)
			}
		}
	}
}
