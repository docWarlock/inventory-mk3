package houses

import (
	"context"
)

// Service defines the interface for house business logic
type Service interface {
	// CreateHouse creates a new house with the provided details
	CreateHouse(ctx context.Context, req *HouseCreateRequest) (*House, error)

	// GetHouse retrieves a house by its ID
	GetHouse(ctx context.Context, id string) (*House, error)

	// ListHouses retrieves all houses
	ListHouses(ctx context.Context) ([]*House, error)

	// UpdateHouse updates an existing house
	UpdateHouse(ctx context.Context, id string, req *HouseUpdateRequest) (*House, error)

	// DeleteHouse deletes a house by its ID
	DeleteHouse(ctx context.Context, id string) error
}
