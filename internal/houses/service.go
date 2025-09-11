package houses

import (
	"context"
<<<<<<< HEAD
=======
	"fmt"
	"time"
>>>>>>> 0fc602b3644a10308454371779653eba100db7ab
)

// Service defines the interface for house business logic
type Service interface {
<<<<<<< HEAD
	// CreateHouse creates a new house with the provided details
	CreateHouse(ctx context.Context, req *HouseCreateRequest) (*House, error)
=======
	// CreateHouse creates a new house
	CreateHouse(ctx context.Context, req HouseCreateRequest) (*House, error)
>>>>>>> 0fc602b3644a10308454371779653eba100db7ab

	// GetHouse retrieves a house by its ID
	GetHouse(ctx context.Context, id string) (*House, error)

	// ListHouses retrieves all houses
<<<<<<< HEAD
	ListHouses(ctx context.Context) ([]*House, error)

	// UpdateHouse updates an existing house
	UpdateHouse(ctx context.Context, id string, req *HouseUpdateRequest) (*House, error)
=======
	ListHouses(ctx context.Context, limit, offset int) ([]*House, error)

	// UpdateHouse updates an existing house
	UpdateHouse(ctx context.Context, id string, req HouseUpdateRequest) (*House, error)
>>>>>>> 0fc602b3644a10308454371779653eba100db7ab

	// DeleteHouse deletes a house by its ID
	DeleteHouse(ctx context.Context, id string) error
}
<<<<<<< HEAD
=======

// houseService implements the Service interface
type houseService struct {
	repo Repository
}

// NewService creates a new house service
func NewService(repo Repository) Service {
	return &houseService{
		repo: repo,
	}
}

// CreateHouse creates a new house with validation
func (s *houseService) CreateHouse(ctx context.Context, req HouseCreateRequest) (*House, error) {
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
		ID:        generateUUID(),
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

// generateUUID generates a UUID string (placeholder implementation)
func generateUUID() string {
	// Simple placeholder - in production this would use a proper UUID library
	// For now, we'll use a timestamp-based approach to ensure uniqueness
	return "house-" + time.Now().Format("20060102150405")
}

// GetHouse retrieves a house by its ID
func (s *houseService) GetHouse(ctx context.Context, id string) (*House, error) {
	return s.repo.GetHouseByID(ctx, id)
}

// ListHouses retrieves all houses with pagination
func (s *houseService) ListHouses(ctx context.Context, limit, offset int) ([]*House, error) {
	return s.repo.ListHouses(ctx, limit, offset)
}

// UpdateHouse updates an existing house
func (s *houseService) UpdateHouse(ctx context.Context, id string, req HouseUpdateRequest) (*House, error) {
	// Validate request - at least one field must be provided
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
			return nil, err
		}
		if exists {
			return nil, &DuplicateError{Message: "House with this name already exists"}
		}
	}

	// Update the house entity
	updatedHouse := &House{
		ID:        id,
		Name:      getOrDefaultString(req.Name, currentHouse.Name),
		TotalArea: getOrDefaultTotalArea(req.TotalArea, currentHouse.TotalArea),
		Unit:      getOrDefaultUnit(req.Unit, currentHouse.Unit),
		CreatedAt: currentHouse.CreatedAt,
		UpdatedAt: time.Now(),
	}

	// Save to repository
	err = s.repo.UpdateHouse(ctx, id, &HouseUpdateRequest{
		Name:      req.Name,
		TotalArea: req.TotalArea,
		Unit:      req.Unit,
	})
	if err != nil {
		return nil, err
	}

	return updatedHouse, nil
}

// DeleteHouse deletes a house by its ID
func (s *houseService) DeleteHouse(ctx context.Context, id string) error {
	return s.repo.DeleteHouse(ctx, id)
}

// Helper functions for validation and data handling

// getOrDefaultString returns the value if not nil, otherwise returns the default
func getOrDefaultString(value *string, defaultValue string) string {
	if value != nil {
		return *value
	}
	return defaultValue
}

// getOrDefaultUnit returns the value if not nil, otherwise returns the default
func getOrDefaultUnit(value string, defaultValue string) string {
	if value != "" {
		return value
	}
	return defaultValue
}

// getOrDefaultTotalArea returns the value if not nil, otherwise returns the default
func getOrDefaultTotalArea(value *float64, defaultValue *float64) *float64 {
	if value != nil {
		return value
	}
	return defaultValue
}

// ValidationError represents a validation error
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return "validation error on field " + e.Field + ": " + e.Message
}

// DuplicateError represents a duplicate entry error
type DuplicateError struct {
	Message string
}

func (e *DuplicateError) Error() string {
	return "duplicate error: " + e.Message
}
>>>>>>> 0fc602b3644a10308454371779653eba100db7ab
