package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	r := gin.Default()

	r.GET("/ws", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			http.NotFound(c.Writer, c.Request)
			return
		}

		for {
			// 读取消息
			_, message, err := conn.ReadMessage()
			if err != nil {
				break
			}

			// 发送消息
			err = conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				break
			}
		}
	})

	r.Run(":8000")
}
