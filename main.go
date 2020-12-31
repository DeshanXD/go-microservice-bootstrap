package main

import (
	"./homepage"
	"./server"
	"log"
	"net/http"
	"os"
)

var (
	CertFile    = os.Getenv("CERT_FILE")
	KeyFile     = os.Getenv("KEY_FILE")
	ServiceAddr = os.Getenv("SERVICE_ADDR")
)

func main() {
	// Making a logger
	logger := log.New(os.Stdout, "GM : ", log.LstdFlags|log.Lshortfile)

	h := homepage.NewHandlers(logger)

	mux := http.NewServeMux()

	h.SetupRoutes(mux)

	srv := server.New(mux, ServiceAddr)
	logger.Println("Server Starting!")

	err := srv.ListenAndServeTLS(CertFile, KeyFile)
	if err == nil {
		logger.Fatalf("Server failed to start %v", err)
	}
}
