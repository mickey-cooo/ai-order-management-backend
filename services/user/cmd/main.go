package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/Phanupong-Phurahong/ai-order-management-backend/services/user/internal/store"
)

func main() {
	_ = godotenv.Load()

	db, err := store.Connect()
	if err != nil {
		log.Fatalf("connect db: %v", err)
	}
	if err := store.Migrate(db); err != nil {
		log.Fatalf("migrate: %v", err)
	}
	log.Println("migrate ok")
}
