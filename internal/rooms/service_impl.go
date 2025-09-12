package rooms

import (
	"context"
	"time"
)

// serviceImpl implements the RoomService interface
type serviceImpl struct {
	repository Repository
}

// NewService creates a new room service instance
func NewService(repository Repository) *serviceImpl {
	return &serviceImpl{
		repository: repository,
	}
}

// CreateRoom creates a new room with validation
func (s *serviceImpl) CreateRoom(ctx context.Context, req *RoomCreateRequest) (*Room, error) {
	// Validate request
	if req.Name == "" {
		return nil, &ValidationError{Field: "name", Message: "Name is required"}
	}

	if req.HouseID == "" {
		return nil, &ValidationError{Field: "house_id", Message: "House ID is required"}
	}

	// Create room object - let the repository handle ID generation
	room := &Room{
		Name:        req.Name,
		HouseID:     req.HouseID,
		Description: req.Description,
		Dimensions:  req.Dimensions,
		Area:        req.Area,
		Unit:        req.Unit,
	}

	// Save to database
	err := s.repository.CreateRoom(ctx, room)
	if err != nil {
		return nil, err
	}

	return room, nil
}

// GetRoom retrieves a room by its ID
func (s *serviceImpl) GetRoom(ctx context.Context, id string) (*Room, error) {
	room, err := s.repository.GetRoomByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return room, nil
}

// ListRooms lists all rooms for a specific house
func (s *serviceImpl) ListRooms(ctx context.Context, houseID string) ([]*Room, error) {
	rooms, err := s.repository.ListRoomsByHouseID(ctx, houseID)
	if err != nil {
		return nil, err
	}

	return rooms, nil
}

// UpdateRoom updates an existing room with validation
func (s *serviceImpl) UpdateRoom(ctx context.Context, id string, req *RoomUpdateRequest) (*Room, error) {
	// Get the existing room
	room, err := s.repository.GetRoomByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Apply updates if provided
	if req.Name != nil {
		room.Name = *req.Name
	}

	if req.Description != nil {
		room.Description = *req.Description
	}

	if req.Dimensions != nil {
		room.Dimensions = req.Dimensions
	}

	if req.Area != nil {
		room.Area = *req.Area
	}

	if req.Unit != nil {
		room.Unit = *req.Unit
	}

	room.UpdatedAt = time.Now()

	// Update in database
	err = s.repository.UpdateRoom(ctx, room)
	if err != nil {
		return nil, err
	}

	return room, nil
}

// DeleteRoom deletes a room by its ID
func (s *serviceImpl) DeleteRoom(ctx context.Context, id string) error {
	err := s.repository.DeleteRoom(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

// ValidationError represents a validation error
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}
