package templater

type MicroserviceTemplate struct {
	Image   string
	Command []string
}

func GetTemplate(taskType string) MicroserviceTemplate {
	switch taskType {
	case "worker":
		return MicroserviceTemplate{Image: "worker-image", Command: []string{"./run-worker"}}
	default:
		return MicroserviceTemplate{}
	}
}
