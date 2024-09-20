package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/yourusername/orquestador-microservicios/internal/api"
	"github.com/yourusername/orquestador-microservicios/pkg/events"
)

func main() {
	fmt.Println("Starting the Ephemeral Microservices Orchestrator...")

	// Load configuration (simplified)
	config := LoadConfig()

	// Start API server
	go api.StartServer()

	// Start event listener
	go events.StartEventListener(config)

	// Wait for a signal to stop the orchestrator
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	fmt.Println("Stopping orchestrator.")
}

func LoadConfig() map[string]string {
	// Simplified configuration loading
	return map[string]string{
		"docker_image": "my_microservice_image",
		"network":      "default",
	}
}
