package main

import (
	"log"
	"os"
	"time"

	"market/common/yamlconf"
	"marketapi/gateway/internal/config"
	"marketapi/gateway/internal/server"
	"marketapi/gateway/internal/transport"
)

func main() {
	cfgFile := os.Getenv("SERVICE_CONFIG_PATH")
	cfg := &config.Config{}
	err := yamlconf.Load(cfgFile, cfg)
	if err != nil {
		log.Fatal(err)
	}

	handler, err := transport.NewHandler(cfg.Proxies)
	if err != nil {
		log.Fatal(err)
	}

	corsCfg := server.CORSConfig{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"OPTIONS",
		},
		AllowedHeaders:   []string{"Content-Type", "application/json"},
		AllowCredentials: true,

		MaxAge: 12 * time.Hour,
	}
	requestTimeout := 10 * time.Second

	router := server.NewRouter(handler, requestTimeout, corsCfg)

	if err := server.Start(cfg.Server, router); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
