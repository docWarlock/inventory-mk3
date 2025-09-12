package rooms

import (
	"net/http"
)

// Handler defines the interface for room HTTP handlers
type Handler interface {
	// CreateRoom handles POST /rooms
	CreateRoom(w http.ResponseWriter, r *http.Request)

	// GetRoom handles GET /rooms/{id}
	GetRoom(w http.ResponseWriter, r *http.Request)

	// ListRooms handles GET /houses/{house_id}/rooms
	ListRooms(w http.ResponseWriter, r *http.Request)

	// UpdateRoom handles PUT /rooms/{id}
	UpdateRoom(w http.ResponseWriter, r *http.Request)

	// DeleteRoom handles DELETE /rooms/{id}
	DeleteRoom(w http.ResponseWriter, r *http.Request)
}
