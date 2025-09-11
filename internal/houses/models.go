package houses

import (
	"time"

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
}

// HouseUpdateRequest represents the request body for updating an existing house
type HouseUpdateRequest struct {
	// Name is the human-readable name of the house (optional, unique if provided)
	Name *string `json:"name,omitempty"`

	// Address is the physical address of the house (optional)
	Address *string `json:"address,omitempty"`

	// Dimensions represents standard dimensions for the house (optional)
	Dimensions *HouseDimensions `json:"dimensions,omitempty"`
}
