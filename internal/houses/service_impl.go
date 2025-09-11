package houses

import (
	"context"
)

// HouseServiceImpl is a concrete implementation of the House service
type HouseServiceImpl struct {
	repo Repository
}

// NewHouseService creates a new house service
func NewHouseService(repo Repository) *HouseServiceImpl {
	return &HouseServiceImpl{
		repo: repo,
	}
}

// CreateHouse creates a new house with the provided details
func (s *HouseServiceImpl) CreateHouse(ctx context.Context, req *HouseCreateRequest) (*House, error) {
	// Validate request
	if req.Name == "" {
		return nil, &ValidationError{Message: "house name is required"}
	}

	// Create house object
	house := &House{
		Name:       req.Name,
		Address:    req.Address,
		Dimensions: req.Dimensions,
	}

	// Create the house
	err := s.repo.CreateHouse(ctx, house)
	if err != nil {
		return nil, err
	}

	return house, nil
}

// GetHouse retrieves a house by its ID
func (s *HouseServiceImpl) GetHouse(ctx context.Context, id string) (*House, error) {
	return s.repo.GetHouseByID(ctx, id)
}

// ListHouses retrieves all houses
func (s *HouseServiceImpl) ListHouses(ctx context.Context) ([]*House, error) {
	return s.repo.ListHouses(ctx)
}

// UpdateHouse updates an existing house
func (s *HouseServiceImpl) UpdateHouse(ctx context.Context, id string, req *HouseUpdateRequest) (*House, error) {
	// Get the existing house
	currentHouse, err := s.repo.GetHouseByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Update fields if provided
	if req.Name != nil {
		currentHouse.Name = *req.Name
	}
	if req.Address != nil {
		currentHouse.Address = *req.Address
	}
	if req.Dimensions != nil {
		currentHouse.Dimensions = req.Dimensions
	}

	// Update the house
	err = s.repo.UpdateHouse(ctx, currentHouse)
	if err != nil {
		return nil, err
	}

	return currentHouse, nil
}

// DeleteHouse deletes a house by its ID
func (s *HouseServiceImpl) DeleteHouse(ctx context.Context, id string) error {
	return s.repo.DeleteHouse(ctx, id)
}

// ValidationError is returned when a validation error occurs
type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}
