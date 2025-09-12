package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/docWarlock/inventory-mk3/internal/houses"
)

// housesRepository implements the houses.Repository interface
type housesRepository struct {
	db *DB
}

// NewHousesRepository creates a new houses repository
func NewHousesRepository(db *DB) houses.Repository {
	return &housesRepository{
		db: db,
	}
}

// CreateHouse creates a new house in the database
func (r *housesRepository) CreateHouse(ctx context.Context, house *houses.House) error {
	query := `
	INSERT INTO houses (id, name, total_area, unit, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?, ?)
	`

	_, err := r.db.ExecContext(ctx, query, house.ID, house.Name, house.TotalArea, house.Unit, house.CreatedAt, house.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

// GetHouseByID retrieves a house by its ID
func (r *housesRepository) GetHouseByID(ctx context.Context, id string) (*houses.House, error) {
	query := `
	SELECT id, name, total_area, unit, created_at, updated_at
	FROM houses
	WHERE id = ?
	`

	row := r.db.QueryRowContext(ctx, query, id)
	return r.scanHouse(row)
}

// ListHouses retrieves all houses with optional pagination
func (r *housesRepository) ListHouses(ctx context.Context, limit, offset int) ([]*houses.House, error) {
	var query string
	var args []interface{}

	if limit > 0 {
		query = `
		SELECT id, name, total_area, unit, created_at, updated_at
		FROM houses
		ORDER BY name
		LIMIT ? OFFSET ?
		`
		args = []interface{}{limit, offset}
	} else {
		query = `
		SELECT id, name, total_area, unit, created_at, updated_at
		FROM houses
		ORDER BY name
		`
		args = []interface{}{}
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var houses []*houses.House
	for rows.Next() {
		house, err := r.scanHouse(rows)
		if err != nil {
			return nil, err
		}
		houses = append(houses, house)
	}

	return houses, nil
}

// UpdateHouse updates an existing house
func (r *housesRepository) UpdateHouse(ctx context.Context, id string, req *houses.HouseUpdateRequest) error {
	// Build dynamic update query
	setClause := ""
	args := []interface{}{}

	if req.Name != nil {
		setClause += "name = ?, "
		args = append(args, *req.Name)
	}

	if req.TotalArea != nil {
		setClause += "total_area = ?, "
		args = append(args, *req.TotalArea)
	}

	if req.Unit != "" {
		setClause += "unit = ?, "
		args = append(args, req.Unit)
	}

	// Remove trailing comma and space
	if len(setClause) > 0 {
		setClause = setClause[:len(setClause)-2]
	} else {
		return nil // No fields to update
	}

	// Add ID to args for WHERE clause
	args = append(args, id)

	query := "UPDATE houses SET " + setClause + ", updated_at = ? WHERE id = ?"
	args = append(args, time.Now())

	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

// DeleteHouse deletes a house by its ID
func (r *housesRepository) DeleteHouse(ctx context.Context, id string) error {
	query := "DELETE FROM houses WHERE id = ?"
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

// HouseExists checks if a house with the given name exists (excluding the house with given ID)
func (r *housesRepository) HouseExists(ctx context.Context, name string, excludeID *string) (bool, error) {
	var query string
	var args []interface{}

	if excludeID != nil {
		query = "SELECT EXISTS(SELECT 1 FROM houses WHERE name = ? AND id != ?)"
		args = []interface{}{name, *excludeID}
	} else {
		query = "SELECT EXISTS(SELECT 1 FROM houses WHERE name = ?)"
		args = []interface{}{name}
	}

	var exists bool
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

// scanHouse scans a database row into a House struct
func (r *housesRepository) scanHouse(row interface{ Scan(...interface{}) error }) (*houses.House, error) {
	var house houses.House
	var totalArea sql.NullFloat64
	var unit sql.NullString

	err := row.Scan(&house.ID, &house.Name, &totalArea, &unit, &house.CreatedAt, &house.UpdatedAt)
	if err != nil {
		return nil, err
	}

	// Set total area if it exists
	if totalArea.Valid {
		house.TotalArea = &totalArea.Float64
	} else {
		house.TotalArea = nil // Explicitly set to nil if not valid
	}

	// Set unit if it exists
	if unit.Valid {
		house.Unit = unit.String
	} else {
		house.Unit = "" // Explicitly set to empty string if not valid
	}

	return &house, nil
}
