package main

import (
	"log"
	"marketapi/auth/internal/config"
	repo "marketapi/auth/internal/repositories"
	"marketapi/auth/internal/server"
	"marketapi/auth/internal/transport"
	"marketapi/auth/internal/usecase"
	"os"

	"market/common/storage"
)

func main() {
	cnfName := os.Getenv("SERVICE_CONFIG_PATH")
	cnf := config.New(cnfName)

	db, err := storage.New(cnf.DB)
	if err != nil {
		log.Fatalf("[ERROR] can't connect to database")
	}

	repo := repo.New(db.DB)
	uc := usecase.New(repo, cnf.JWTSecret)

	handler := transport.NewAuthHandler(uc)
	router := server.NewRouter(handler)

	if err := server.Run(cnf.BinAddr, router); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
