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
	if house.ID == uuid.Nil {
		house.ID = uuid.New()
	}

	// Set timestamps
	now := time.Now()
	house.CreatedAt = now
	house.UpdatedAt = now

	// Check for duplicate name
	for _, existingHouse := range r.houses {
		if existingHouse.Name == house.Name {
			return &DuplicateHouseError{Message: "house name must be unique"}
		}
	}

	r.houses[house.ID.String()] = house
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

// ListHouses retrieves all houses
func (r *InMemoryHouseRepository) ListHouses(ctx context.Context) ([]*House, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	houses := make([]*House, 0, len(r.houses))
	for _, house := range r.houses {
		houses = append(houses, house)
	}
	return houses, nil
}

// UpdateHouse updates an existing house
func (r *InMemoryHouseRepository) UpdateHouse(ctx context.Context, house *House) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Set updated timestamp
	house.UpdatedAt = time.Now()

	// Check for duplicate name (excluding the current house)
	for id, existingHouse := range r.houses {
		if existingHouse.Name == house.Name && id != house.ID.String() {
			return &DuplicateHouseError{Message: "house name must be unique"}
		}
	}

	r.houses[house.ID.String()] = house
	return nil
}

// DeleteHouse deletes a house by its ID
func (r *InMemoryHouseRepository) DeleteHouse(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.houses, id)
	return nil
}

// GetHouseByName retrieves a house by its name
func (r *InMemoryHouseRepository) GetHouseByName(ctx context.Context, name string) (*House, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, house := range r.houses {
		if house.Name == name {
			return house, nil
		}
	}
	return nil, &HouseNotFoundError{Message: "house not found"}
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
