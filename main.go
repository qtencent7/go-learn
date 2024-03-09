package main

import (
	"fmt"
	"net"
)

// ChatRoom 用于存储所有客户端的信息
type ChatRoom struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

// Client 代表一个连接的客户端
type Client struct {
	socket net.Conn
	data   chan []byte
}

func (c *Client) Read() {
	defer c.socket.Close()
	for {
		message := make([]byte, 4096)
		_, err := c.socket.Read(message)
		if err != nil {
			return
		}
		c.data <- message
	}
}

func (c *Client) Write() {
	defer c.socket.Close()
	for message := range c.data {
		_, err := c.socket.Write(message)
		if err != nil {
			return
		}
	}
}

func (cr *ChatRoom) Start() {
	for {
		select {
		case client := <-cr.register:
			cr.clients[client] = true
			fmt.Println(" 新客户端已连接 ")
		case client := <-cr.unregister:
			if _, ok := cr.clients[client]; ok {
				delete(cr.clients, client)
				close(client.data)
				fmt.Println("客户端已关闭")
			}
		case message := <-cr.broadcast:
			for client := range cr.clients {
				select {
				case client.data <- message:
				default:
					close(client.data)
					delete(cr.clients, client)
				}
			}
		}
	}
}

func (cr *ChatRoom) Run() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer listener.Close()

	go cr.Start()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			continue
		}
		client := &Client{socket: conn, data: make(chan []byte)}
		cr.register <- client

		go client.Read()
		go client.Write()
	}
}

func main() {
	chatRoom := &ChatRoom{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}

	chatRoom.Run()
}
