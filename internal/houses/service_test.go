package houses

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHouseService_CreateHouse(t *testing.T) {
	// Create an in-memory repository
	repo := NewInMemoryHouseRepository()
	service := NewHouseService(repo)

	ctx := context.Background()

	// Test creating a house
	req := &HouseCreateRequest{
		Name: "Test House",
	}

	house, err := service.CreateHouse(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, house)
	assert.Equal(t, "Test House", house.Name)
	assert.NotEmpty(t, house.ID)
	assert.False(t, house.CreatedAt.IsZero())
	assert.False(t, house.UpdatedAt.IsZero())
}

func TestHouseService_CreateHouse_UniqueName(t *testing.T) {
	// Create an in-memory repository
	repo := NewInMemoryHouseRepository()
	service := NewHouseService(repo)

	ctx := context.Background()

	// Create first house
	req1 := &HouseCreateRequest{
		Name: "Test House",
	}

	_, err1 := service.CreateHouse(ctx, req1)
	assert.NoError(t, err1)

	// Try to create another house with the same name
	req2 := &HouseCreateRequest{
		Name: "Test House",
	}

	_, err2 := service.CreateHouse(ctx, req2)
	assert.Error(t, err2)
	assert.Contains(t, err2.Error(), "House with this name already exists")
}

func TestHouseService_GetHouse(t *testing.T) {
	// Create an in-memory repository
	repo := NewInMemoryHouseRepository()
	service := NewHouseService(repo)

	ctx := context.Background()

	// Create a house
	req := &HouseCreateRequest{
		Name: "Test House",
	}

	house, err := service.CreateHouse(ctx, req)
	assert.NoError(t, err)

	// Retrieve the house
	retrievedHouse, err := service.GetHouse(ctx, house.ID)
	assert.NoError(t, err)
	assert.Equal(t, house.ID, retrievedHouse.ID)
	assert.Equal(t, "Test House", retrievedHouse.Name)
}

func TestHouseService_ListHouses(t *testing.T) {
	// Create an in-memory repository
	repo := NewInMemoryHouseRepository()
	service := NewHouseService(repo)

	ctx := context.Background()

	// Create multiple houses
	req1 := &HouseCreateRequest{
		Name: "Test House 1",
	}

	req2 := &HouseCreateRequest{
		Name: "Test House 2",
	}

	_, err1 := service.CreateHouse(ctx, req1)
	assert.NoError(t, err1)

	_, err2 := service.CreateHouse(ctx, req2)
	assert.NoError(t, err2)

	// List all houses
	houses, err := service.ListHouses(ctx)
	assert.NoError(t, err)
	assert.Len(t, houses, 2)
}
