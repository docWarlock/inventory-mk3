package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/docWarlock/inventory-mk3/internal/houses"
	"github.com/docWarlock/inventory-mk3/internal/rooms"
	"github.com/go-chi/chi/v5"
)

func main() {
	// Create repository and service
	houseRepo := houses.NewInMemoryHouseRepository()
	houseService := houses.NewHouseService(houseRepo)
	houseHandler := houses.NewHouseHandler(houseService)

	// Create room repository and service
	roomRepo := rooms.NewInMemoryRoomRepository()
	roomService := rooms.NewService(roomRepo)    // Fixed function name
	roomHandler := rooms.NewHandler(roomService) // Fixed function name

	// Setup HTTP router using Chi
	router := setupRouter(houseHandler, roomHandler)

	// Start server
	server := &http.Server{
		Addr:    ":8080",
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

func setupRouter(houseHandler houses.Handler, roomHandler rooms.Handler) http.Handler {
	router := chi.NewRouter()

	// House endpoints
	router.Post("/houses", houseHandler.CreateHouse)
	router.Get("/houses/{id}", houseHandler.GetHouse)
	router.Get("/houses", houseHandler.ListHouses)
	router.Put("/houses/{id}", houseHandler.UpdateHouse)
	router.Delete("/houses/{id}", houseHandler.DeleteHouse)

	// Room endpoints
	router.Post("/rooms", roomHandler.CreateRoom)
	router.Get("/rooms/{id}", roomHandler.GetRoom)
	router.Get("/houses/{house_id}/rooms", roomHandler.ListRooms)
	router.Put("/rooms/{id}", roomHandler.UpdateRoom)
	router.Delete("/rooms/{id}", roomHandler.DeleteRoom)

	return router
}
