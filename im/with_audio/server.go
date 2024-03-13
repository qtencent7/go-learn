package with_audio

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool) // 所有连接的客户端
var broadcast = make(chan []byte)            // 广播通道

var upgrader = websocket.Upgrader{} // 使用默认的选项

func main() {
	// 配置WebSocket路由
	http.HandleFunc("/ws", handleConnections)

	// 启动广播协程
	go handleBroadcasts()

	// 开始监听HTTP服务
	log.Println("http server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	// 升级HTTP连接到WebSocket连接
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	// 注册新的客户端
	clients[ws] = true

	for {
		// 读取新的消息
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
		// 将消息发送到广播通道
		broadcast <- msg
	}
}

func handleBroadcasts() {
	for {
		// 从广播通道中获取消息
		msg := <-broadcast
		// 发送给所有客户端
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
