package rooms

import (
	"context"
)

// Repository defines the interface for room persistence operations
type Repository interface {
	// CreateRoom creates a new room in the database
	CreateRoom(ctx context.Context, room *Room) error

	// GetRoomByID retrieves a room by its ID
	GetRoomByID(ctx context.Context, id string) (*Room, error)

	// ListRoomsByHouseID retrieves all rooms for a specific house
	ListRoomsByHouseID(ctx context.Context, houseID string) ([]*Room, error)

	// UpdateRoom updates an existing room
	UpdateRoom(ctx context.Context, room *Room) error

	// DeleteRoom deletes a room by its ID
	DeleteRoom(ctx context.Context, id string) error
}
