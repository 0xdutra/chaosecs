package chaos

import (
	"math/rand"
	"time"
)

func random(servicesTasks []string) string {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(servicesTasks))

	return servicesTasks[index]
}
