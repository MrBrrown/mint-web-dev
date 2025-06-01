package main

import (
	"log"
	"market/common/storage"
	"marketapi/products/internal/config"
	productrepo "marketapi/products/internal/repositories/product_repo"
	"marketapi/products/internal/server"
	"marketapi/products/internal/transport"
	productusecase "marketapi/products/internal/usecase"
	"os"
	"time"
)

func main() {
	cnfName := os.Getenv("SERVICE_CONFIG_PATH")
	cnf := config.New(cnfName)

	db, err := storage.New(cnf.DB)
	if err != nil {
		log.Print("[ERROR] can't connect to database")
	}

	productRepo := productrepo.New(db.DB)
	cacheRepo := productrepo.NewCacheRepo(productRepo, 10*time.Minute)
	productUC := productusecase.New(cacheRepo)

	handler := transport.New(productUC, cnf.JWTSecret)
	router := server.NewRouter(handler)

	if err := server.Run(cnf.BinAddr, router); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
