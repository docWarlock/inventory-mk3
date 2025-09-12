package houses

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
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

// CreateHouse creates a new house with validation
func (s *HouseServiceImpl) CreateHouse(ctx context.Context, req *HouseCreateRequest) (*House, error) {
	// Validate request
	if req.Name == "" {
		return nil, &ValidationError{Field: "name", Message: "Name is required"}
	}

	// Validate total area if provided
	if req.TotalArea != nil && *req.TotalArea <= 0 {
		return nil, &ValidationError{Field: "total_area", Message: "Total area must be positive if provided"}
	}

	// Check if house with this name already exists
	exists, err := s.repo.HouseExists(ctx, req.Name, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to check house existence: %w", err)
	}
	if exists {
		return nil, &DuplicateError{Message: "House with this name already exists"}
	}

	// Create the house entity
	house := &House{
		ID:        uuid.New().String(),
		Name:      req.Name,
		TotalArea: req.TotalArea,
		Unit:      req.Unit,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Save to repository
	err = s.repo.CreateHouse(ctx, house)
	if err != nil {
		return nil, fmt.Errorf("failed to create house: %w", err)
	}

	return house, nil
}

// GetHouse retrieves a house by its ID
func (s *HouseServiceImpl) GetHouse(ctx context.Context, id string) (*House, error) {
	return s.repo.GetHouseByID(ctx, id)
}

// ListHouses retrieves all houses
func (s *HouseServiceImpl) ListHouses(ctx context.Context) ([]*House, error) {
	return s.repo.ListHouses(ctx, 0, 0)
}

// UpdateHouse updates an existing house
func (s *HouseServiceImpl) UpdateHouse(ctx context.Context, id string, req *HouseUpdateRequest) (*House, error) {
	// Validate request - at least one field must be provided for update
	if req.Name == nil && req.TotalArea == nil && req.Unit == "" {
		return nil, &ValidationError{Field: "name or total_area or unit", Message: "At least one field must be provided for update"}
	}

	// Validate total area if provided
	if req.TotalArea != nil && *req.TotalArea <= 0 {
		return nil, &ValidationError{Field: "total_area", Message: "Total area must be positive if provided"}
	}

	// Check if house exists
	currentHouse, err := s.repo.GetHouseByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// If name is being updated, check for duplicates (excluding current house)
	if req.Name != nil && *req.Name != currentHouse.Name {
		exists, err := s.repo.HouseExists(ctx, *req.Name, &id)
		if err != nil {
			return nil, fmt.Errorf("failed to check house existence: %w", err)
		}
		if exists {
			return nil, &DuplicateError{Message: "House with this name already exists"}
		}
	}

	// Update the house entity
	if req.Name != nil {
		currentHouse.Name = *req.Name
	}
	if req.TotalArea != nil {
		currentHouse.TotalArea = req.TotalArea
	}
	if req.Unit != "" {
		currentHouse.Unit = req.Unit
	}
	currentHouse.UpdatedAt = time.Now()

	// Save to repository
	err = s.repo.UpdateHouse(ctx, id, &HouseUpdateRequest{
		Name:      req.Name,
		TotalArea: req.TotalArea,
		Unit:      req.Unit,
	})
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
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return "validation error on field '" + e.Field + "': " + e.Message
}

// DuplicateError is returned when a duplicate house is found
type DuplicateError struct {
	Message string
}

func (e *DuplicateError) Error() string {
	return "duplicate error: " + e.Message
}
