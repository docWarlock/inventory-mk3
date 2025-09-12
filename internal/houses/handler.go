package houses

import "net/http"

// Handler defines the HTTP handler interface for house endpoints
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
