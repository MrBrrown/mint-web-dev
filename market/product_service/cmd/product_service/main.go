package main

import (
	"marketapi/products/internal/config"
	"os"
)

func main() {
	cnfName := os.Getenv("SERVICE_CONFIG_PATH")
	cnf := config.New(cnfName)
	defer print(cnf)
}
