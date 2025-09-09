package houses

import (
	"context"
)

// Repository defines the interface for house data operations
type Repository interface {
	// CreateHouse creates a new house in the database
	CreateHouse(ctx context.Context, house *House) error

	// GetHouseByID retrieves a house by its ID
	GetHouseByID(ctx context.Context, id string) (*House, error)

	// ListHouses retrieves all houses with optional pagination
	ListHouses(ctx context.Context, limit, offset int) ([]*House, error)

	// UpdateHouse updates an existing house
	UpdateHouse(ctx context.Context, id string, house *HouseUpdateRequest) error

	// DeleteHouse deletes a house by its ID
	DeleteHouse(ctx context.Context, id string) error

	// HouseExists checks if a house with the given name exists (excluding the house with given ID)
	HouseExists(ctx context.Context, name string, excludeID *string) (bool, error)
}
