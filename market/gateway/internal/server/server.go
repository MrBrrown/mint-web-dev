package server

import (
	"context"
	"log"
	"marketapi/gateway/internal/config"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func Start(serInfo config.Server, handler http.Handler) error {
	srv := &http.Server{
		Addr:         serInfo.Address,
		Handler:      handler,
		ReadTimeout:  serInfo.ReadTimeout,
		WriteTimeout: serInfo.WriteTimeout,
	}

	errCh := make(chan error, 1)
	go func() {
		log.Printf("starting server on %s\n", serInfo.Address)
		errCh <- srv.ListenAndServe()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-quit:
		log.Printf("caught signal %s: shutting down...", sig)
	case err := <-errCh:
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), serInfo.ShutdownTimeout)
	defer cancel()
	return srv.Shutdown(ctx)
}
