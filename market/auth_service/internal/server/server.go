package server

import (
	"fmt"
	"net/http"
)

func Run(bindAddr string, handler http.Handler) error {
	fmt.Printf("Starting HTTP server on %s\n", bindAddr)
	return http.ListenAndServe(bindAddr, handler)
}
