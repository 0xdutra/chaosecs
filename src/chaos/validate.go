package chaos

func validate(desiredCount int64, pendingCount int64, servicesTasks []string, status string) bool {
	if pendingCount == 0 && len(servicesTasks) == int(desiredCount) && status == "ACTIVE" {
		return true
	}
	return false
}
