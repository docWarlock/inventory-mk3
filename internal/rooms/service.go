package rooms

import (
	"context"
)

// Service defines the interface for room business logic
type Service interface {
	// CreateRoom creates a new room with the provided details
	CreateRoom(ctx context.Context, req *RoomCreateRequest) (*Room, error)

	// GetRoom retrieves a room by its ID
	GetRoom(ctx context.Context, id string) (*Room, error)

	// ListRooms retrieves all rooms for a specific house
	ListRooms(ctx context.Context, houseID string) ([]*Room, error)

	// UpdateRoom updates an existing room
	UpdateRoom(ctx context.Context, id string, req *RoomUpdateRequest) (*Room, error)

	// DeleteRoom deletes a room by its ID
	DeleteRoom(ctx context.Context, id string) error
}
