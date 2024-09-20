package events

import (
	"fmt"
	"time"

	"github.com/yourusername/orquestador-microservicios/internal/container"
)

func StartEventListener(config map[string]string) {
	fmt.Println("Listening for events to generate microservices...")
	for {
		// Simulate event detection
		fmt.Println("Event detected!")
		container.CreateContainer(config["docker_image"])
		time.Sleep(10 * time.Second) // Simulate event interval
	}
}
