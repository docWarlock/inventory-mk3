package rooms

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockRepository is a mock implementation of the Repository interface for testing
type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) CreateRoom(ctx context.Context, room *Room) error {
	args := m.Called(ctx, room)
	// Set ID on the room like the real repository would do
	if room.ID == "" {
		room.ID = "mock-room-id"
	}
	return args.Error(0)
}

func (m *MockRepository) GetRoomByID(ctx context.Context, id string) (*Room, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*Room), args.Error(1)
}

func (m *MockRepository) ListRoomsByHouseID(ctx context.Context, houseID string) ([]*Room, error) {
	args := m.Called(ctx, houseID)
	return args.Get(0).([]*Room), args.Error(1)
}

func (m *MockRepository) UpdateRoom(ctx context.Context, room *Room) error {
	args := m.Called(ctx, room)
	return args.Error(0)
}

func (m *MockRepository) DeleteRoom(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestCreateRoom(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(MockRepository)

	service := NewService(mockRepo)

	// Test valid request
	req := &RoomCreateRequest{
		Name:    "Living Room",
		HouseID: "house-123",
	}

	// We don't know the exact ID since it's generated dynamically, so we'll just check the other fields
	now := time.Now()
	expectedRoom := &Room{
		Name:        req.Name,
		HouseID:     req.HouseID,
		Description: "",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	mockRepo.On("CreateRoom", ctx, mock.AnythingOfType("*rooms.Room")).Return(nil)

	room, err := service.CreateRoom(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, room)
	assert.Equal(t, expectedRoom.Name, room.Name)
	assert.Equal(t, expectedRoom.HouseID, room.HouseID)
	assert.NotEmpty(t, room.ID) // ID should be generated

	mockRepo.AssertExpectations(t)
}

func TestGetRoom(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(MockRepository)

	service := NewService(mockRepo)

	now := time.Now()
	expectedRoom := &Room{
		ID:          "room-123",
		Name:        "Living Room",
		HouseID:     "house-123",
		Description: "",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	mockRepo.On("GetRoomByID", ctx, "room-123").Return(expectedRoom, nil)

	room, err := service.GetRoom(ctx, "room-123")
	assert.NoError(t, err)
	assert.NotNil(t, room)
	assert.Equal(t, expectedRoom.ID, room.ID)

	mockRepo.AssertExpectations(t)
}

func TestListRooms(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(MockRepository)

	service := NewService(mockRepo)

	now := time.Now()
	expectedRooms := []*Room{
		{
			ID:          "room-123",
			Name:        "Living Room",
			HouseID:     "house-123",
			Description: "",
			CreatedAt:   now,
			UpdatedAt:   now,
		},
	}

	mockRepo.On("ListRoomsByHouseID", ctx, "house-123").Return(expectedRooms, nil)

	rooms, err := service.ListRooms(ctx, "house-123")
	assert.NoError(t, err)
	assert.NotNil(t, rooms)
	assert.Len(t, rooms, 1)

	mockRepo.AssertExpectations(t)
}

func TestUpdateRoom(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(MockRepository)

	service := NewService(mockRepo)

	// Test updating room name
	req := &RoomUpdateRequest{
		Name: &[]string{"Updated Living Room"}[0],
	}

	now := time.Now()
	existingRoom := &Room{
		ID:          "room-123",
		Name:        "Living Room",
		HouseID:     "house-123",
		Description: "",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	mockRepo.On("GetRoomByID", ctx, "room-123").Return(existingRoom, nil)
	mockRepo.On("UpdateRoom", ctx, mock.AnythingOfType("*rooms.Room")).Return(nil)

	room, err := service.UpdateRoom(ctx, "room-123", req)
	assert.NoError(t, err)
	assert.NotNil(t, room)
	assert.Equal(t, *req.Name, room.Name)

	mockRepo.AssertExpectations(t)
}

func TestDeleteRoom(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(MockRepository)

	service := NewService(mockRepo)

	mockRepo.On("DeleteRoom", ctx, "room-123").Return(nil)

	err := service.DeleteRoom(ctx, "room-123")
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}
