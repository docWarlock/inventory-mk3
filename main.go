package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"inventory-mk3/internal/houses"
	"inventory-mk3/internal/rooms"
)

func main() {
	// Create repository and service
	houseRepo := houses.NewInMemoryHouseRepository()
	houseService := houses.NewHouseService(houseRepo)
	houseHandler := houses.NewHouseHandler(houseService)

	// Create room repository and service
	roomRepo := rooms.NewRepository(nil) // Will be updated with actual DB connection
	roomService := rooms.NewService(roomRepo)
	roomHandler := rooms.NewHandler(roomService)

	// Create a simple HTTP server
	mux := http.NewServeMux()

	// Define routes
	mux.HandleFunc("POST /houses", houseHandler.CreateHouse)
	mux.HandleFunc("GET /houses/{id}", houseHandler.GetHouse)
	mux.HandleFunc("GET /houses", houseHandler.ListHouses)
	mux.HandleFunc("PUT /houses/{id}", houseHandler.UpdateHouse)
	mux.HandleFunc("DELETE /houses/{id}", houseHandler.DeleteHouse)

	// Room routes
	mux.HandleFunc("POST /houses/{house_id}/rooms", roomHandler.CreateRoom)
	mux.HandleFunc("GET /rooms/{id}", roomHandler.GetRoom)
	mux.HandleFunc("GET /houses/{house_id}/rooms", roomHandler.ListRooms)
	mux.HandleFunc("PUT /rooms/{id}", roomHandler.UpdateRoom)
	mux.HandleFunc("DELETE /rooms/{id}", roomHandler.DeleteRoom)

	// Start server
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// Start server in a goroutine
	go func() {
		fmt.Println("Starting server on :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	fmt.Println("Shutting down server...")

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Shutdown server gracefully
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

	fmt.Println("Server stopped")
}
