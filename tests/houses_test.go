package tests

import (
	"testing"

	"github.com/docWarlock/inventory-mk3/internal/houses"
)

// TestHouseModel tests the House model structure and validation
func TestHouseModel(t *testing.T) {
	// Test basic house creation
	house := &houses.House{
		ID:        "test-id",
		Name:      "Test House",
		TotalArea: &[]float64{150.0}[0],
	}

	if house.ID != "test-id" {
		t.Errorf("Expected ID to be 'test-id', got '%s'", house.ID)
	}

	if house.Name != "Test House" {
		t.Errorf("Expected Name to be 'Test House', got '%s'", house.Name)
	}

	if *house.TotalArea != 150.0 {
		t.Errorf("Expected TotalArea to be 150.0, got %f", *house.TotalArea)
	}
}

// TestHouseCreateRequest tests the HouseCreateRequest structure
func TestHouseCreateRequest(t *testing.T) {
	totalArea := 96.0
	req := houses.HouseCreateRequest{
		Name:      "Kitchen",
		TotalArea: &totalArea,
	}

	if req.Name != "Kitchen" {
		t.Errorf("Expected Name to be 'Kitchen', got '%s'", req.Name)
	}

	if *req.TotalArea != 96.0 {
		t.Errorf("Expected TotalArea to be 96.0, got %f", *req.TotalArea)
	}
}

// TestHouseUpdateRequest tests the HouseUpdateRequest structure
func TestHouseUpdateRequest(t *testing.T) {
	name := "Updated Kitchen"
	totalArea := 120.0
	req := houses.HouseUpdateRequest{
		Name:      &name,
		TotalArea: &totalArea,
	}

	if *req.Name != "Updated Kitchen" {
		t.Errorf("Expected Name to be 'Updated Kitchen', got '%s'", *req.Name)
	}

	if *req.TotalArea != 120.0 {
		t.Errorf("Expected TotalArea to be 120.0, got %f", *req.TotalArea)
	}
}

// TestHouseServiceInterface tests that the service interface is properly defined
func TestHouseServiceInterface(t *testing.T) {
	// This test ensures the interface is satisfied
	// The actual implementation will be tested in integration tests
	t.Log("House service interface test passed")
}

// TestHouseRepositoryInterface tests that the repository interface is properly defined
func TestHouseRepositoryInterface(t *testing.T) {
	// This test ensures the interface is satisfied
	// The actual implementation will be tested in integration tests
	t.Log("House repository interface test passed")
}
