package main

import (
	"lb/lb"
	"log/slog"

	"github.com/BurntSushi/toml"
)

func main() {
	var cfg lb.LbConfig
	if _, err := toml.DecodeFile("/app/lb.toml", &cfg); err != nil {
		slog.Error("failed to decode lb toml", slog.Any("error", err))
		return
	}
	if err := lb.StartLoadBalancer(cfg); err != nil {
		slog.Error("failed to run lb", slog.Any("error", err))
		return
	}
}
