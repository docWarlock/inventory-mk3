package houses

import (
	"context"
	"sync"
	"time"

	"github.com/google/uuid"
)

// InMemoryHouseRepository is a simple in-memory implementation of the House repository
type InMemoryHouseRepository struct {
	houses map[string]*House
	mu     sync.RWMutex
}

// NewInMemoryHouseRepository creates a new in-memory house repository
func NewInMemoryHouseRepository() *InMemoryHouseRepository {
	return &InMemoryHouseRepository{
		houses: make(map[string]*House),
	}
}

// CreateHouse creates a new house
func (r *InMemoryHouseRepository) CreateHouse(ctx context.Context, house *House) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Set UUID if not already set
	if house.ID == "" {
		house.ID = uuid.New().String()
	}

	// Set timestamps
	now := time.Now()
	house.CreatedAt = now
	house.UpdatedAt = now

	r.houses[house.ID] = house
	return nil
}

// GetHouseByID retrieves a house by its ID
func (r *InMemoryHouseRepository) GetHouseByID(ctx context.Context, id string) (*House, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	house, exists := r.houses[id]
	if !exists {
		return nil, &HouseNotFoundError{Message: "house not found"}
	}
	return house, nil
}

// ListHouses retrieves all houses with optional pagination
func (r *InMemoryHouseRepository) ListHouses(ctx context.Context, limit, offset int) ([]*House, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	houses := make([]*House, 0, len(r.houses))
	for _, house := range r.houses {
		houses = append(houses, house)
	}
	return houses, nil
}

// UpdateHouse updates an existing house
func (r *InMemoryHouseRepository) UpdateHouse(ctx context.Context, id string, updateReq *HouseUpdateRequest) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	house, exists := r.houses[id]
	if !exists {
		return &HouseNotFoundError{Message: "house not found"}
	}

	// Update fields if provided
	if updateReq.Name != nil {
		house.Name = *updateReq.Name
	}
	if updateReq.TotalArea != nil {
		house.TotalArea = updateReq.TotalArea
	}
	if updateReq.Unit != "" {
		house.Unit = updateReq.Unit
	}

	// Set updated timestamp
	house.UpdatedAt = time.Now()

	r.houses[id] = house
	return nil
}

// DeleteHouse deletes a house by its ID
func (r *InMemoryHouseRepository) DeleteHouse(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.houses, id)
	return nil
}

// HouseExists checks if a house with the given name exists (excluding the house with given ID)
func (r *InMemoryHouseRepository) HouseExists(ctx context.Context, name string, excludeID *string) (bool, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for id, house := range r.houses {
		if house.Name == name {
			// If excludeID is provided and matches the current ID, skip this check
			if excludeID != nil && *excludeID == id {
				continue
			}
			return true, nil
		}
	}
	return false, nil
}

// HouseNotFoundError is returned when a house is not found
type HouseNotFoundError struct {
	Message string
}

func (e *HouseNotFoundError) Error() string {
	return e.Message
}

// DuplicateHouseError is returned when a duplicate house name is used
type DuplicateHouseError struct {
	Message string
}

func (e *DuplicateHouseError) Error() string {
	return e.Message
}
