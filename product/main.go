package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	configs "github.com/4lerman/e_com/common/config"
	"github.com/4lerman/e_com/common/db"
	"github.com/4lerman/e_com/product/routes"
	"github.com/4lerman/e_com/product/store"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	db, err := db.InitStorage()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Db connected successfully!")

	productStore := store.NewStore(db)
	productHandler := routes.NewHandler(productStore)
	
	router := mux.NewRouter()
	productRouter := router.PathPrefix("/products").Subrouter()
	productHandler.RegisterRoutes(productRouter)

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{configs.Envs.Base_Url},		
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	}).Handler(productRouter)

	port := configs.Envs.Products_Port
	server := &http.Server{
		Addr:   ":" + port,
		Handler: corsHandler,
	}

	go gracefulShutdown(server)

	log.Printf("Server(Products) is starting on port %s\n", port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server startup failed: %v\n", err)
	}

	log.Println("Server(Products) gracefully stopped")
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