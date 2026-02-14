package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"log/slog"
	"net"

	"github.com/BurntSushi/toml"
)

type serverConfig struct {
	Server_port string `toml:"server_port"`
}

func main() {

	var cfg serverConfig

	if _, err := toml.DecodeFile("server.toml", &cfg); err != nil {
		slog.Error("failed to decode server toml", slog.Any("error", err))
		return
	}

	StartServer(cfg.Server_port)
}

func StartServer(port string) {

	// we listen on server port
	listener, err := net.Listen("tcp", ":"+port)

	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}

	defer listener.Close()

	fmt.Println("Server running on port :" + port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error Accepting:", err)
			continue
		}
		fmt.Println("server " + port + " has recieved a connection")
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	// always close the connection at the end
	defer func() {
		fmt.Println("Closing connection")
		err := conn.Close()
		if err != nil {
			fmt.Println("Failed closing connection:", err)
		} else {
			fmt.Println("Connection closed")
		}
	}()

	// len of data at the start of the buffer, 4 bytes
	Rxbuffer := make([]byte, 4)
	Txbuffer := make([]byte, 4)

	for i := 0; i < 10; i++ {

		// read the message
		_, err := conn.Read(Rxbuffer)
		if err != nil && err != io.EOF {
			fmt.Println("Error reading:", err)
			return
		}
		msg := binary.LittleEndian.Uint32(Rxbuffer)
		fmt.Printf("Message Content is: %d \n", msg)

		// write back
		msg = msg + 3
		binary.LittleEndian.PutUint32(Txbuffer, msg)
		_, err = conn.Write(Txbuffer)
		if err != nil {
			fmt.Println("Error writing:", err)
		}

	}

}
