package rooms

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// ServiceImpl implements the RoomService interface
type ServiceImpl struct {
	repository Repository
}

// NewService creates a new room service instance
func NewService(repository Repository) *ServiceImpl {
	return &ServiceImpl{
		repository: repository,
	}
}

// CreateRoom creates a new room with the provided details
func (s *ServiceImpl) CreateRoom(ctx context.Context, req *RoomCreateRequest) (*Room, error) {
	// Generate a new UUID for the room
	id := uuid.New()

	// Create the room object
	room := &Room{
		ID:          id,
		Name:        req.Name,
		HouseID:     req.HouseID,
		Description: req.Description,
		Dimensions:  req.Dimensions,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// Save to database
	err := s.repository.CreateRoom(ctx, room)
	if err != nil {
		return nil, err
	}

	return room, nil
}

// GetRoom retrieves a room by its ID
func (s *ServiceImpl) GetRoom(ctx context.Context, id string) (*Room, error) {
	room, err := s.repository.GetRoomByID(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Return nil if not found
		}
		return nil, err
	}
	return room, nil
}

// ListRooms retrieves all rooms for a specific house
func (s *ServiceImpl) ListRooms(ctx context.Context, houseID string) ([]*Room, error) {
	rooms, err := s.repository.ListRoomsByHouseID(ctx, houseID)
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

// UpdateRoom updates an existing room
func (s *ServiceImpl) UpdateRoom(ctx context.Context, id string, req *RoomUpdateRequest) (*Room, error) {
	// First get the existing room
	room, err := s.repository.GetRoomByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if room == nil {
		return nil, nil // Room not found
	}

	// Update fields if provided
	if req.Name != nil {
		room.Name = *req.Name
	}
	if req.Description != nil {
		room.Description = req.Description
	}
	if req.Dimensions != nil {
		room.Dimensions = req.Dimensions
	}
	room.UpdatedAt = time.Now()

	// Save updated room to database
	err = s.repository.UpdateRoom(ctx, room)
	if err != nil {
		return nil, err
	}

	return room, nil
}

// DeleteRoom deletes a room by its ID
func (s *ServiceImpl) DeleteRoom(ctx context.Context, id string) error {
	return s.repository.DeleteRoom(ctx, id)
}
