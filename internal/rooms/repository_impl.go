package rooms

import (
	"context"
	"database/sql"
)

// RepositoryImpl implements the RoomRepository interface
type RepositoryImpl struct {
	db *sql.DB
}

// NewRepository creates a new room repository instance
func NewRepository(db *sql.DB) *RepositoryImpl {
	return &RepositoryImpl{
		db: db,
	}
}

// CreateRoom creates a new room
func (r *RepositoryImpl) CreateRoom(ctx context.Context, room *Room) error {
	query := `
		INSERT INTO rooms (id, name, house_id, description, dimensions, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := r.db.ExecContext(ctx, query,
		room.ID,
		room.Name,
		room.HouseID,
		room.Description,
		room.Dimensions,
		room.CreatedAt,
		room.UpdatedAt,
	)
	return err
}

// GetRoomByID retrieves a room by its ID
func (r *RepositoryImpl) GetRoomByID(ctx context.Context, id string) (*Room, error) {
	query := `
		SELECT id, name, house_id, description, dimensions, created_at, updated_at
		FROM rooms
		WHERE id = $1
	`

	var room Room
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&room.ID,
		&room.Name,
		&room.HouseID,
		&room.Description,
		&room.Dimensions,
		&room.CreatedAt,
		&room.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &room, nil
}

// ListRoomsByHouseID retrieves all rooms for a specific house
func (r *RepositoryImpl) ListRoomsByHouseID(ctx context.Context, houseID string) ([]*Room, error) {
	query := `
		SELECT id, name, house_id, description, dimensions, created_at, updated_at
		FROM rooms
		WHERE house_id = $1
		ORDER BY name
	`

	rows, err := r.db.QueryContext(ctx, query, houseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rooms []*Room
	for rows.Next() {
		var room Room
		err := rows.Scan(
			&room.ID,
			&room.Name,
			&room.HouseID,
			&room.Description,
			&room.Dimensions,
			&room.CreatedAt,
			&room.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, &room)
	}

	return rooms, nil
}

// UpdateRoom updates an existing room
func (r *RepositoryImpl) UpdateRoom(ctx context.Context, room *Room) error {
	query := `
		UPDATE rooms
		SET name = $1, description = $2, dimensions = $3, updated_at = $4
		WHERE id = $5
	`

	_, err := r.db.ExecContext(ctx, query,
		room.Name,
		room.Description,
		room.Dimensions,
		room.UpdatedAt,
		room.ID,
	)
	return err
}

// DeleteRoom deletes a room by its ID
func (r *RepositoryImpl) DeleteRoom(ctx context.Context, id string) error {
	query := `
		DELETE FROM rooms
		WHERE id = $1
	`

	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
