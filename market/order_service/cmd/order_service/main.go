package main

import (
	"log"
	"market/common/storage"
	"marketapi/orders/internal/config"
	orderrepo "marketapi/orders/internal/repositories/order_repo"
	"marketapi/orders/internal/server"
	"marketapi/orders/internal/transport"
	orderusecase "marketapi/orders/internal/usecase"
	"os"
)

func main() {
	cnfPath := os.Getenv("SERVICE_CONFIG_PATH")
	if cnfPath == "" {
		log.Fatal("SERVICE_CONFIG_PATH not set")
	}

	cnf := config.New(cnfPath)

	db, err := storage.New(cnf.DB)
	if err != nil {
		log.Fatalf("[ERROR] Can't connect to database: %v", err)
	}

	orderRepo := orderrepo.New(db.DB)
	orderUC := orderusecase.New(orderRepo)

	handler := transport.New(orderUC, cnf.JWTSecret)
	router := server.NewRouter(handler)

	if err := server.Run(cnf.BinAddr, router); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
