package houses

import (
	"time"
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
}

// HouseUpdateRequest represents the request body for updating an existing house
type HouseUpdateRequest struct {
	// Name is the human-readable name for this house
	Name *string `json:"name,omitempty"`

	// Total area of the house
	TotalArea *float64 `json:"total_area,omitempty"`

	// Unit of measurement for total area (square feet, square meters, or pyeong)
	Unit string `json:"unit,omitempty"`
}
