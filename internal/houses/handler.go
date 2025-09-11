package houses

import (
<<<<<<< HEAD
	"net/http"
)

// Handler defines the interface for house HTTP endpoints
=======
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// Handler defines the HTTP handler interface for house endpoints
>>>>>>> 0fc602b3644a10308454371779653eba100db7ab
type Handler interface {
	// CreateHouse handles POST /houses
	CreateHouse(w http.ResponseWriter, r *http.Request)

	// GetHouse handles GET /houses/{id}
	GetHouse(w http.ResponseWriter, r *http.Request)

	// ListHouses handles GET /houses
	ListHouses(w http.ResponseWriter, r *http.Request)

	// UpdateHouse handles PUT /houses/{id}
	UpdateHouse(w http.ResponseWriter, r *http.Request)

	// DeleteHouse handles DELETE /houses/{id}
	DeleteHouse(w http.ResponseWriter, r *http.Request)
}
<<<<<<< HEAD
=======

// houseHandler implements the Handler interface
type houseHandler struct {
	service Service
}

// NewHandler creates a new house handler
func NewHandler(service Service) Handler {
	return &houseHandler{
		service: service,
	}
}

// CreateHouse handles POST /houses
func (h *houseHandler) CreateHouse(w http.ResponseWriter, r *http.Request) {
	var req HouseCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	house, err := h.service.CreateHouse(r.Context(), req)
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(house)
}

// GetHouse handles GET /houses/{id}
func (h *houseHandler) GetHouse(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	house, err := h.service.GetHouse(r.Context(), id)
	if err != nil {
		switch err.(type) {
		case *ValidationError:
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		default:
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(house)
}

// ListHouses handles GET /houses
func (h *houseHandler) ListHouses(w http.ResponseWriter, r *http.Request) {
	// Get pagination parameters
	limit := 10
	if limitStr := r.URL.Query().Get("limit"); limitStr != "" {
		if parsedLimit, err := strconv.Atoi(limitStr); err == nil {
			limit = parsedLimit
		}
	}

	offset := 0
	if offsetStr := r.URL.Query().Get("offset"); offsetStr != "" {
		if parsedOffset, err := strconv.Atoi(offsetStr); err == nil {
			offset = parsedOffset
		}
	}

	houses, err := h.service.ListHouses(r.Context(), limit, offset)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(houses)
}

// UpdateHouse handles PUT /houses/{id}
func (h *houseHandler) UpdateHouse(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	var req HouseUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	house, err := h.service.UpdateHouse(r.Context(), id, req)
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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(house)
}

// DeleteHouse handles DELETE /houses/{id}
func (h *houseHandler) DeleteHouse(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	err := h.service.DeleteHouse(r.Context(), id)
	if err != nil {
		switch err.(type) {
		case *ValidationError:
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		default:
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}
>>>>>>> 0fc602b3644a10308454371779653eba100db7ab
