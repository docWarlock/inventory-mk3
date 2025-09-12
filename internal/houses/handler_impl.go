package houses

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// HouseHandler is a concrete implementation of the House HTTP handler
type HouseHandler struct {
	service Service
}

// NewHouseHandler creates a new house handler
func NewHouseHandler(service Service) *HouseHandler {
	return &HouseHandler{
		service: service,
	}
}

// CreateHouse handles POST /houses
func (h *HouseHandler) CreateHouse(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Parse request body
	var req HouseCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Create the house
	house, err := h.service.CreateHouse(ctx, &req)
	if err != nil {
		switch err.(type) {
		case *ValidationError:
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		case *DuplicateError:
			http.Error(w, err.Error(), http.StatusConflict)
			return
		default:
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	}

	// Return created house
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(house)
}

// GetHouse handles GET /houses/{id}
func (h *HouseHandler) GetHouse(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Extract ID from URL using Chi router
	id := chi.URLParam(r, "id")

	// Validate UUID
	if _, err := uuid.Parse(id); err != nil {
		http.Error(w, "Invalid house ID", http.StatusBadRequest)
		return
	}

	// Get the house
	house, err := h.service.GetHouse(ctx, id)
	if err != nil {
		switch err.(type) {
		case *HouseNotFoundError:
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		default:
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	}

	// Return house
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(house)
}

// ListHouses handles GET /houses
func (h *HouseHandler) ListHouses(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// List houses
	houses, err := h.service.ListHouses(ctx)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Return houses
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(houses)
}

// UpdateHouse handles PUT /houses/{id}
func (h *HouseHandler) UpdateHouse(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Extract ID from URL using Chi router
	id := chi.URLParam(r, "id")

	// Validate UUID
	if _, err := uuid.Parse(id); err != nil {
		http.Error(w, "Invalid house ID", http.StatusBadRequest)
		return
	}

	// Parse request body
	var req HouseUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Update the house
	house, err := h.service.UpdateHouse(ctx, id, &req)
	if err != nil {
		switch err.(type) {
		case *ValidationError:
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		case *HouseNotFoundError:
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		case *DuplicateError:
			http.Error(w, err.Error(), http.StatusConflict)
			return
		default:
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	}

	// Return updated house
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(house)
}

// DeleteHouse handles DELETE /houses/{id}
func (h *HouseHandler) DeleteHouse(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Extract ID from URL using Chi router
	id := chi.URLParam(r, "id")

	// Validate UUID
	if _, err := uuid.Parse(id); err != nil {
		http.Error(w, "Invalid house ID", http.StatusBadRequest)
		return
	}

	// Delete the house
	err := h.service.DeleteHouse(ctx, id)
	if err != nil {
		switch err.(type) {
		case *HouseNotFoundError:
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		default:
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	}

	// Return success
	w.WriteHeader(http.StatusOK)
}
