package main

import (
	"awesomeProject3/Router"
	"fmt"
	"net"
)

func main() {
	fmt.Println("Listening to server")
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println("Error in starting server")
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error in Accepting connection")
			return
		}

		go Router.LoadHandler(conn)
	}
}
