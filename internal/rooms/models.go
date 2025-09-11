package rooms

import (
	"time"

	"github.com/google/uuid"
)

// Room represents a room within a house in the inventory system
type Room struct {
	// ID is a unique identifier for the room (UUID)
	ID uuid.UUID `json:"id" db:"id"`

	// Name is the human-readable name of the room (required, unique within house)
	Name string `json:"name" db:"name"`

	// HouseID is the identifier of the house this room belongs to (required)
	HouseID uuid.UUID `json:"house_id" db:"house_id"`

	// Description is a description of the room (optional)
	Description *string `json:"description,omitempty" db:"description"`

	// Dimensions represents standard dimensions for the room (optional)
	// This could include Gridfinity container dimensions
	Dimensions *RoomDimensions `json:"dimensions,omitempty" db:"dimensions"`

	// CreatedAt timestamp when room was created
	CreatedAt time.Time `json:"created_at" db:"created_at"`

	// UpdatedAt timestamp when room was last updated
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// RoomDimensions represents the dimensions of a room for organizational purposes
type RoomDimensions struct {
	// Width in Gridfinity units
	Width int `json:"width" db:"width"`

	// Depth in Gridfinity units
	Depth int `json:"depth" db:"depth"`

	// Height in Gridfinity units (optional)
	Height *int `json:"height,omitempty" db:"height"`
}

// RoomCreateRequest represents the request body for creating a new room
type RoomCreateRequest struct {
	// Name is the human-readable name of the room (required, unique within house)
	Name string `json:"name"`

	// HouseID is the identifier of the house this room belongs to (required)
	HouseID uuid.UUID `json:"house_id"`

	// Description is a description of the room (optional)
	Description *string `json:"description,omitempty"`

	// Dimensions represents standard dimensions for the room (optional)
	Dimensions *RoomDimensions `json:"dimensions,omitempty"`
}

// RoomUpdateRequest represents the request body for updating an existing room
type RoomUpdateRequest struct {
	// Name is the human-readable name of the room (optional, unique within house if provided)
	Name *string `json:"name,omitempty"`

	// Description is a description of the room (optional)
	Description *string `json:"description,omitempty"`

	// Dimensions represents standard dimensions for the room (optional)
	Dimensions *RoomDimensions `json:"dimensions,omitempty"`
}
