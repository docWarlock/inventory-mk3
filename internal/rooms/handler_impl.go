package rooms

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// handlerImpl implements the RoomHandler interface
type handlerImpl struct {
	service Service
}

// NewHandler creates a new room handler instance
func NewHandler(service Service) *handlerImpl {
	return &handlerImpl{
		service: service,
	}
}

// CreateRoom handles POST /rooms
func (h *handlerImpl) CreateRoom(w http.ResponseWriter, r *http.Request) {
	var req RoomCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	room, err := h.service.CreateRoom(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(room)
}

// GetRoom handles GET /rooms/{id}
func (h *handlerImpl) GetRoom(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Room ID is required", http.StatusBadRequest)
		return
	}

	room, err := h.service.GetRoom(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(room)
}

// ListRooms handles GET /houses/{house_id}/rooms
func (h *handlerImpl) ListRooms(w http.ResponseWriter, r *http.Request) {
	houseID := chi.URLParam(r, "house_id")
	if houseID == "" {
		http.Error(w, "House ID is required", http.StatusBadRequest)
		return
	}

	rooms, err := h.service.ListRooms(r.Context(), houseID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rooms)
}

// UpdateRoom handles PUT /rooms/{id}
func (h *handlerImpl) UpdateRoom(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Room ID is required", http.StatusBadRequest)
		return
	}

	var req RoomUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	room, err := h.service.UpdateRoom(r.Context(), id, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(room)
}

// DeleteRoom handles DELETE /rooms/{id}
func (h *handlerImpl) DeleteRoom(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Room ID is required", http.StatusBadRequest)
		return
	}

	err := h.service.DeleteRoom(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
