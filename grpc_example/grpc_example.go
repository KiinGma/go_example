package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials"
	"log"
	"time"
)

func main() {
	// 创建TLS凭证
	creds, err := credentials.NewClientTLSFromFile("path/to/your/cert.pem", "")
	if err != nil {
		log.Fatalf("Failed to create TLS credentials: %v", err)
	}

	// 创建gRPC连接
	conn, err := grpc.Dial("your_server_address", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// 监听连接状态
	go func() {
		for {
			if conn.WaitForStateChange(context.Background(), conn.GetState()) {
				if conn.GetState() == connectivity.Connecting {
					log.Println("Connection lost. Reconnecting...")
					// 连接断开，尝试重新建立连接
					for {
						conn.ResetConnectBackoff()
						ok := conn.WaitForStateChange(context.Background(), connectivity.Connecting)
						if !ok {
							log.Printf("Failed to wait for state change: %v", err)
							time.Sleep(1 * time.Second)
							continue
						}
						change := conn.WaitForStateChange(context.Background(), connectivity.Ready)
						if !change {
							log.Printf("Failed to wait for state change: %v", err)
							time.Sleep(1 * time.Second)
							continue
						}
						break
					}
					log.Println("Connection reestablished.")
					// 重新建立连接后，恢复发送请求和接收响应

				}
			}
		}
	}()
}
