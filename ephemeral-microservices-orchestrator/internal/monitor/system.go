package monitor

import (
	"fmt"
	"time"
)

func StartMonitoring() {
	for {
		fmt.Println("Monitoring microservices...")
		time.Sleep(10 * time.Second)
	}
}
