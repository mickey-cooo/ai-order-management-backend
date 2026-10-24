package main

import "github.com/Phanupong-Phurahong/ai-order-management-backend/pkg/server"

func main() {
	server.Serve("orders", "8082")
}
