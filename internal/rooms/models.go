package rooms

import (
	"time"
)

// Room represents a room within a house
type Room struct {
	ID          string          `json:"id"`
	Name        string          `json:"name"`
	HouseID     string          `json:"house_id"`
	Description string          `json:"description,omitempty"`
	Dimensions  *RoomDimensions `json:"dimensions,omitempty"`
	Area        float64         `json:"area,omitempty"`
	Unit        string          `json:"unit,omitempty"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}

// RoomDimensions represents the physical dimensions of a room
type RoomDimensions struct {
	Length float64 `json:"length,omitempty"`
	Width  float64 `json:"width,omitempty"`
	Height float64 `json:"height,omitempty"`
}

// RoomCreateRequest represents the request body for creating a room
type RoomCreateRequest struct {
	Name        string          `json:"name"`
	HouseID     string          `json:"house_id"`
	Description string          `json:"description,omitempty"`
	Dimensions  *RoomDimensions `json:"dimensions,omitempty"`
	Area        float64         `json:"area,omitempty"`
	Unit        string          `json:"unit,omitempty"`
}

// RoomUpdateRequest represents the request body for updating a room
type RoomUpdateRequest struct {
	Name        *string         `json:"name,omitempty"`
	HouseID     *string         `json:"house_id,omitempty"`
	Description *string         `json:"description,omitempty"`
	Dimensions  *RoomDimensions `json:"dimensions,omitempty"`
	Area        *float64        `json:"area,omitempty"`
	Unit        *string         `json:"unit,omitempty"`
}
