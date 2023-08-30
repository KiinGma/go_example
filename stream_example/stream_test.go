package stream_example

import (
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"
)

func TestStream(t *testing.T) {
	http.HandleFunc("/stream", streamHandler)
	http.ListenAndServe(":8080", nil)
}

func streamHandler(w http.ResponseWriter, r *http.Request) {
	// 设置HTTP响应头，表示数据以事件流的方式返回给前端
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	// 创建一个单向通道，将数据通过通道传递给前端
	dataChan := make(chan string)
	go generateData(dataChan)
	defer close(dataChan)
	// 不断循环读取通道中的数据，并将其以事件流的方式写入HTTP响应体中
	for {
		select {
		case data := <-dataChan:
			_, err := io.WriteString(w, "data: "+data+"\n\n")
			if err != nil {
				fmt.Println(err)
			}
			// 强制刷新HTTP响应
			w.(http.Flusher).Flush()
		case <-time.After(10 * time.Second):
			// 在10秒内没有接收到数据，则发送一条注释，避免浏览器超时
			_, err := io.WriteString(w, ": comment\n\n")
			if err != nil {
				fmt.Println(err)
			}
			w.(http.Flusher).Flush()
		}
	}
}
func generateData(dataChan chan string) {
	// 模拟生成数据的过程，将数据写入通道中
	for i := 0; i < 10; i++ {
		data := fmt.Sprintf("data %d", i+1)
		// 将数据写入通道中
		dataChan <- data
		time.Sleep(1 * time.Second)
	}
}
