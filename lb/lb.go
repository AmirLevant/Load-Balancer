package lb

import (
	"fmt"
	"io"
	"log/slog"
	"net"

	"golang.org/x/sync/errgroup"
)

type LbConfig struct {
	LbPort  string   `toml:"lb_port"`
	Servers []string `toml:"servers"`
}

func StartLoadBalancer(cfg LbConfig) error {
	// Sets up a socket for lb to listen on,
	// for incoming connections
	address := ":" + cfg.LbPort
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	defer listener.Close()
	slog.Info("Running lb")

	slog.Info("Listening", slog.String("address", address))

	// Counts the number of total connections,
	// used for round robin load balancing
	robin := 0
	for {
		conn, err := listener.Accept()
		if err != nil {
			// TODO What happens if the TCP socket is closed for good?
			// handle different errs, which ones do we break with?
			slog.Error("failed accepting listener", slog.Any("error", err))
			continue
		}
		slog.Info("Accepted connection", slog.Any("address", conn.RemoteAddr()))
		go handleConnection(conn, cfg.Servers[robin%len(cfg.Servers)])
		robin += 1
	}
}

func handleConnection(clientConn net.Conn, serverAddress string) {
	// Ensure the client connection is closed even if
	// the function exits early, e.g. if we fail to
	// connect to the server
	defer clientConn.Close()

	// Connect to the server
	serverConn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		slog.Error("Failed to connect to server", slog.Any("error", err))
		return
	}
	defer serverConn.Close()

	logger := slog.With(
		slog.Any("client_address", clientConn.RemoteAddr()),
		slog.Any("server_address", serverConn.RemoteAddr()))

	var g errgroup.Group
	g.Go(func() error {
		_, err := io.Copy(clientConn, serverConn)
		if err != nil {
			return fmt.Errorf("proxy client to server: %w", err)
		}
		return nil
	})
	g.Go(func() error {
		_, err := io.Copy(serverConn, clientConn)
		if err != nil {
			return fmt.Errorf("proxy server to client: %w", err)
		}
		return nil
	})
	if err := g.Wait(); err != nil {
		logger.Error("Failed proxying", slog.Any("error", err))
	}
}
