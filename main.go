package main

import (
	"context"
<<<<<<< HEAD
	"fmt"
=======
>>>>>>> 0fc602b3644a10308454371779653eba100db7ab
	"log"
	"net/http"
	"os"
	"os/signal"
<<<<<<< HEAD
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
=======
	"syscall"
	"time"

	"github.com/docWarlock/inventory-mk3/internal/database"
	"github.com/docWarlock/inventory-mk3/internal/houses"
	"github.com/go-chi/chi/v5"
)

func main() {
	// Create database connection
	db, err := database.NewDB("houses.db")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Initialize schema
	if err = db.InitSchema(); err != nil {
		log.Fatal("Failed to initialize schema:", err)
	}

	// Create repository
	housesRepo := database.NewHousesRepository(db)

	// Create service
	housesService := houses.NewService(housesRepo)

	// Create handler
	housesHandler := houses.NewHandler(housesService)

	// Setup HTTP router
	router := setupRouter(housesHandler)
>>>>>>> 0fc602b3644a10308454371779653eba100db7ab

	// Start server
	server := &http.Server{
		Addr:    ":8080",
<<<<<<< HEAD
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
=======
		Handler: router,
	}

	// Graceful shutdown
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Println("Server starting on :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Server failed to start:", err)
		}
	}()

	<-done
	log.Println("Server stopping...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown error:", err)
	}

	log.Println("Server stopped")
}

func setupRouter(housesHandler houses.Handler) http.Handler {
	r := chi.NewRouter()

	// House endpoints
	r.Post("/houses", housesHandler.CreateHouse)
	r.Get("/houses/{id}", housesHandler.GetHouse)
	r.Get("/houses", housesHandler.ListHouses)
	r.Put("/houses/{id}", housesHandler.UpdateHouse)
	r.Delete("/houses/{id}", housesHandler.DeleteHouse)

	return r
>>>>>>> 0fc602b3644a10308454371779653eba100db7ab
}
