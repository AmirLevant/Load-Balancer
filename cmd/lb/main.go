package main

import (
	"lb/lb"
	"log/slog"
)

func main() {
	var (
		lbPort      string = "8080"
		serverPorts        = []string{"9090", "9091", "9092"}
	)
	if err := lb.StartLoadBalancer(lbPort, serverPorts); err != nil {
		slog.Error("Failed to run lb", slog.Any("error", err))
	}
}
