package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error dialing:", err.Error())
		return
	}
	defer conn.Close()

	go func() {
		for {
			message := make([]byte, 4096)
			_, err := conn.Read(message)
			if err != nil {
				fmt.Println("Error reading:", err.Error())
				return
			}
			fmt.Println(string(message))
		}
	}()

	input := make([]byte, 4096)
	for {
		_, err := os.Stdin.Read(input)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			return
		}
		conn.Write(input)
	}
}
