package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"log/slog"
	"net"
	"os"

	"github.com/BurntSushi/toml"
)

type clientConfig struct {
	LbAddress string `toml:"lb_address"`
}

func main() {
	path := flag.String("config", "client.toml", "client config path")
	flag.Parse()

	var cfg clientConfig
	if _, err := toml.DecodeFile(*path, &cfg); err != nil {
		slog.Error("Failed decoding client toml", slog.Any("error", err))
		os.Exit(1)
	}

	err := startClient(6, cfg.LbAddress)
	if err != nil {
		slog.Error("Failed starting client", slog.Any("error", err))
		os.Exit(1)
	}
}

func startClient(data uint32, lbAddress string) error {
	lbConn, err := net.Dial("tcp", lbAddress)
	if err != nil {
		return fmt.Errorf("failed dialing lb: %w, %s", err, lbAddress)
	}
	defer func() {
		slog.Info("Closing connection")
		err := lbConn.Close()
		if err != nil {
			slog.Error("Failed closing connection", slog.Any("error", err))
		} else {
			slog.Info("Connection closed")
		}
	}()

	slog.Info("Connected", slog.Any("address", lbConn.RemoteAddr()))

	txBuffer := make([]byte, 4)
	rxBuffer := make([]byte, 4)

	// Write the number 10 times to the lb
	for range 10 {
		// Serialise the number we want to send into bytes,
		// and send it over the connection
		binary.LittleEndian.PutUint32(txBuffer, data)

		_, err := lbConn.Write(txBuffer)
		if err != nil {
			return fmt.Errorf("writing to lb: %w", err)
		}

		// Read the modified number
		_, err = lbConn.Read(rxBuffer)
		if err != nil {
			return fmt.Errorf("reading from lb: %w", err)
		}

		msg := binary.LittleEndian.Uint32(rxBuffer)
		slog.Info("Message received from server", slog.Uint64("msg", uint64(msg)))
	}
	return nil
}
