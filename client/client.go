package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	conn, err := net.Dial("tcp", ":8080")

	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}

	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Send a message:")

	message, _ := reader.ReadString('\n')
	fmt.Fprintf(conn, message)

}
