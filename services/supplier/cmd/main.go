package main

import "github.com/Phanupong-Phurahong/ai-order-management-backend/pkg/server"

func main() {
	server.Serve("supplier", "8085")
}
