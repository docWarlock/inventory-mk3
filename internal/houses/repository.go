package houses

import (
	"context"
)

// Repository defines the interface for house data operations
type Repository interface {
<<<<<<< HEAD
	// CreateHouse creates a new house
=======
	// CreateHouse creates a new house in the database
>>>>>>> 0fc602b3644a10308454371779653eba100db7ab
	CreateHouse(ctx context.Context, house *House) error

	// GetHouseByID retrieves a house by its ID
	GetHouseByID(ctx context.Context, id string) (*House, error)

<<<<<<< HEAD
	// ListHouses retrieves all houses
	ListHouses(ctx context.Context) ([]*House, error)

	// UpdateHouse updates an existing house
	UpdateHouse(ctx context.Context, house *House) error
=======
	// ListHouses retrieves all houses with optional pagination
	ListHouses(ctx context.Context, limit, offset int) ([]*House, error)

	// UpdateHouse updates an existing house
	UpdateHouse(ctx context.Context, id string, house *HouseUpdateRequest) error
>>>>>>> 0fc602b3644a10308454371779653eba100db7ab

	// DeleteHouse deletes a house by its ID
	DeleteHouse(ctx context.Context, id string) error

<<<<<<< HEAD
	// GetHouseByName retrieves a house by its name (used for uniqueness validation)
	GetHouseByName(ctx context.Context, name string) (*House, error)
=======
	// HouseExists checks if a house with the given name exists (excluding the house with given ID)
	HouseExists(ctx context.Context, name string, excludeID *string) (bool, error)
>>>>>>> 0fc602b3644a10308454371779653eba100db7ab
}
