package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
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

func setupRouter(housesHandler houses.Handler) http.Handler {
	r := chi.NewRouter()

	// House endpoints
	r.Post("/houses", housesHandler.CreateHouse)
	r.Get("/houses/{id}", housesHandler.GetHouse)
	r.Get("/houses", housesHandler.ListHouses)
	r.Put("/houses/{id}", housesHandler.UpdateHouse)
	r.Delete("/houses/{id}", housesHandler.DeleteHouse)

	return r
}
