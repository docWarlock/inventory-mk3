package rooms

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

// HandlerImpl implements the RoomHandler interface
type HandlerImpl struct {
	service Service
}

// NewHandler creates a new room handler instance
func NewHandler(service Service) *HandlerImpl {
	return &HandlerImpl{
		service: service,
	}
}

// CreateRoom handles POST /houses/{house_id}/rooms
func (h *HandlerImpl) CreateRoom(w http.ResponseWriter, r *http.Request) {
	// Parse house ID from URL path
	// Note: This requires proper routing setup to extract house_id from the path
	// For now, we'll assume it's available in context or a different way

	// Parse request body
	var req RoomCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate house ID is provided
	if req.HouseID == uuid.Nil {
		http.Error(w, "House ID is required", http.StatusBadRequest)
		return
	}

	// Create the room
	ctx := context.Background()
	room, err := h.service.CreateRoom(ctx, &req)
	if err != nil {
		http.Error(w, "Failed to create room", http.StatusInternalServerError)
		return
	}

	// Return created room
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(room)
}

// GetRoom handles GET /rooms/{id}
func (h *HandlerImpl) GetRoom(w http.ResponseWriter, r *http.Request) {
	// Parse room ID from URL path
	// Note: This requires proper routing setup to extract id from the path

	// For now, we'll assume the ID is passed in a way that can be extracted
	// This would typically be done by the router implementation

	// Get room from service
	ctx := context.Background()
	room, err := h.service.GetRoom(ctx, "room_id_from_path")
	if err != nil {
		http.Error(w, "Failed to get room", http.StatusInternalServerError)
		return
	}

	if room == nil {
		http.Error(w, "Room not found", http.StatusNotFound)
		return
	}

	// Return room
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(room)
}

// ListRooms handles GET /houses/{house_id}/rooms
func (h *HandlerImpl) ListRooms(w http.ResponseWriter, r *http.Request) {
	// Parse house ID from URL path
	// Note: This requires proper routing setup to extract house_id from the path

	// Get rooms from service
	ctx := context.Background()
	rooms, err := h.service.ListRooms(ctx, "house_id_from_path")
	if err != nil {
		http.Error(w, "Failed to list rooms", http.StatusInternalServerError)
		return
	}

	// Return rooms
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rooms)
}

// UpdateRoom handles PUT /rooms/{id}
func (h *HandlerImpl) UpdateRoom(w http.ResponseWriter, r *http.Request) {
	// Parse room ID from URL path
	// Note: This requires proper routing setup to extract id from the path

	// Parse request body
	var req RoomUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Update room
	ctx := context.Background()
	room, err := h.service.UpdateRoom(ctx, "room_id_from_path", &req)
	if err != nil {
		http.Error(w, "Failed to update room", http.StatusInternalServerError)
		return
	}

	if room == nil {
		http.Error(w, "Room not found", http.StatusNotFound)
		return
	}

	// Return updated room
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(room)
}

// DeleteRoom handles DELETE /rooms/{id}
func (h *HandlerImpl) DeleteRoom(w http.ResponseWriter, r *http.Request) {
	// Parse room ID from URL path
	// Note: This requires proper routing setup to extract id from the path

	// Delete room
	ctx := context.Background()
	err := h.service.DeleteRoom(ctx, "room_id_from_path")
	if err != nil {
		http.Error(w, "Failed to delete room", http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
}
