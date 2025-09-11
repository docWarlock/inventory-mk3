package houses

import (
	"context"
)

// Repository defines the interface for house data operations
type Repository interface {
	// CreateHouse creates a new house
	CreateHouse(ctx context.Context, house *House) error

	// GetHouseByID retrieves a house by its ID
	GetHouseByID(ctx context.Context, id string) (*House, error)

	// ListHouses retrieves all houses
	ListHouses(ctx context.Context) ([]*House, error)

	// UpdateHouse updates an existing house
	UpdateHouse(ctx context.Context, house *House) error

	// DeleteHouse deletes a house by its ID
	DeleteHouse(ctx context.Context, id string) error

	// GetHouseByName retrieves a house by its name (used for uniqueness validation)
	GetHouseByName(ctx context.Context, name string) (*House, error)
}
