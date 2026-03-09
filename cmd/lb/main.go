package main

import (
	"flag"
	"lb/lb"
	"log/slog"
	"os"

	"github.com/BurntSushi/toml"
)

func main() {
	pathPtr := flag.String("config", "/app/lb.toml", "the config path")
	flag.Parse()

	var cfg lb.LbConfig
	if _, err := toml.DecodeFile(*pathPtr, &cfg); err != nil {
		slog.Error("Failed decoding lb toml", slog.Any("error", err))
		os.Exit(1)
	}
	if err := lb.StartLoadBalancer(cfg); err != nil {
		slog.Error("Failed running lb", slog.Any("error", err))
		os.Exit(1)
	}
}
