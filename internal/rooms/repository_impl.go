package rooms

import (
	"context"
	"sync"
	"time"

	"github.com/google/uuid"
)

// inMemoryRoomRepository implements the Repository interface with in-memory storage
type inMemoryRoomRepository struct {
	rooms map[string]*Room
	mutex sync.RWMutex
}

// NewInMemoryRoomRepository creates a new in-memory room repository
func NewInMemoryRoomRepository() Repository {
	return &inMemoryRoomRepository{
		rooms: make(map[string]*Room),
	}
}

// CreateRoom creates a new room in memory
func (r *inMemoryRoomRepository) CreateRoom(ctx context.Context, room *Room) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// Generate ID if not provided
	if room.ID == "" {
		room.ID = uuid.New().String()
	}

	// Set timestamps
	now := time.Now()
	room.CreatedAt = now
	room.UpdatedAt = now

	r.rooms[room.ID] = room
	return nil
}

// GetRoomByID retrieves a room by its ID
func (r *inMemoryRoomRepository) GetRoomByID(ctx context.Context, id string) (*Room, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	room, exists := r.rooms[id]
	if !exists {
		return nil, nil // Return nil instead of error for "not found"
	}

	return room, nil
}

// ListRoomsByHouseID retrieves all rooms for a specific house
func (r *inMemoryRoomRepository) ListRoomsByHouseID(ctx context.Context, houseID string) ([]*Room, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	var rooms []*Room
	for _, room := range r.rooms {
		if room.HouseID == houseID {
			rooms = append(rooms, room)
		}
	}

	return rooms, nil
}

// UpdateRoom updates an existing room
func (r *inMemoryRoomRepository) UpdateRoom(ctx context.Context, room *Room) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// Set updated timestamp
	room.UpdatedAt = time.Now()

	r.rooms[room.ID] = room
	return nil
}

// DeleteRoom deletes a room by its ID
func (r *inMemoryRoomRepository) DeleteRoom(ctx context.Context, id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	delete(r.rooms, id)
	return nil
}
