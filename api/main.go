package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/4lerman/e_com/api/routes"
	"github.com/gorilla/mux"
	configs "github.com/4lerman/e_com/common/config"

	_ "github.com/4lerman/e_com/docs"
)


// @title E-commerce Service
// @version 1.0
// @description This is a API server for E-commerce service.
// @host e-comm-hl.onrender.com
// @BasePath /api/v1
func main() {
	router := mux.NewRouter()
	routes.Routes(router)

	port := configs.Envs.API_Port
	server := &http.Server{
		Addr:   ":" + port,
		Handler: router,
	}

	go gracefulShutdown(server)

	log.Printf("Server(API) is starting on port %s\n", port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server startup failed: %v\n", err)
	}

	log.Println("Server(API) gracefully stopped")
}

func gracefulShutdown(server *http.Server) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	<-signals

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Graceful shutdown failed: %v\n", err)
	}
}
