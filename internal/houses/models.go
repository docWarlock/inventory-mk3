package houses

import (
	"time"
<<<<<<< HEAD

	"github.com/google/uuid"
)

// House represents a residential location in the inventory system
type House struct {
	// ID is a unique identifier for the house (UUID)
	ID uuid.UUID `json:"id" db:"id"`

	// Name is the human-readable name of the house (required, unique)
	Name string `json:"name" db:"name"`

	// Address is the physical address of the house (optional)
	Address string `json:"address,omitempty" db:"address"`

	// Dimensions represents standard dimensions for the house (optional)
	// This could include Gridfinity container dimensions
	Dimensions *HouseDimensions `json:"dimensions,omitempty" db:"dimensions"`

	// CreatedAt timestamp when house was created
	CreatedAt time.Time `json:"created_at" db:"created_at"`

	// UpdatedAt timestamp when house was last updated
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// HouseDimensions represents the dimensions of a house for organizational purposes
type HouseDimensions struct {
	// Width in Gridfinity units
	Width int `json:"width" db:"width"`

	// Depth in Gridfinity units
	Depth int `json:"depth" db:"depth"`

	// Height in Gridfinity units (optional)
	Height *int `json:"height,omitempty" db:"height"`
}

// HouseCreateRequest represents the request body for creating a new house
type HouseCreateRequest struct {
	// Name is the human-readable name of the house (required, unique)
	Name string `json:"name"`

	// Address is the physical address of the house (optional)
	Address string `json:"address,omitempty"`

	// Dimensions represents standard dimensions for the house (optional)
	Dimensions *HouseDimensions `json:"dimensions,omitempty"`
=======
)

// House represents a physical house or location in the inventory system
type House struct {
	// ID is the unique identifier for this house (UUID)
	ID string `json:"id" db:"id"`

	// Name is the human-readable name for this house (e.g., "Main House", "Summer Cottage")
	Name string `json:"name" db:"name"`

	// Total area of the house
	TotalArea *float64 `json:"total_area,omitempty" db:"total_area"`

	// Unit of measurement for total area (square feet, square meters, or pyeong)
	Unit string `json:"unit,omitempty" db:"unit"`

	// CreatedAt timestamp when this house was created
	CreatedAt time.Time `json:"created_at" db:"created_at"`

	// UpdatedAt timestamp when this house was last updated
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// HouseCreateRequest represents the request body for creating a new house
type HouseCreateRequest struct {
	// Name is the human-readable name for this house
	Name string `json:"name"`

	// Total area of the house
	TotalArea *float64 `json:"total_area,omitempty"`

	// Unit of measurement for total area (square feet, square meters, or pyeong)
	Unit string `json:"unit,omitempty"`
>>>>>>> 0fc602b3644a10308454371779653eba100db7ab
}

// HouseUpdateRequest represents the request body for updating an existing house
type HouseUpdateRequest struct {
<<<<<<< HEAD
	// Name is the human-readable name of the house (optional, unique if provided)
	Name *string `json:"name,omitempty"`

	// Address is the physical address of the house (optional)
	Address *string `json:"address,omitempty"`

	// Dimensions represents standard dimensions for the house (optional)
	Dimensions *HouseDimensions `json:"dimensions,omitempty"`
=======
	// Name is the human-readable name for this house
	Name *string `json:"name,omitempty"`

	// Total area of the house
	TotalArea *float64 `json:"total_area,omitempty"`

	// Unit of measurement for total area (square feet, square meters, or pyeong)
	Unit string `json:"unit,omitempty"`
>>>>>>> 0fc602b3644a10308454371779653eba100db7ab
}
