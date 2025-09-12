# Rooms Package

This package implements the room model, service, repository, and handler for the inventory management system.

## Components

### Models
- `Room`: Represents a room in a house with fields like ID, name, house ID, description, dimensions, and timestamps
- `RoomCreateRequest`: Request structure for creating new rooms
- `RoomUpdateRequest`: Request structure for updating existing rooms

### Repository
- `Repository`: Interface defining the contract for room data operations
- `repositoryImpl`: Concrete implementation of the repository using database operations

### Service
- `Service`: Interface defining the contract for room business logic
- `serviceImpl`: Concrete implementation of the service with room management functionality

### Handler
- `Handler`: Interface defining the HTTP handler contract
- `handlerImpl`: Concrete implementation that handles HTTP requests for rooms

## Features

- Create, read, update, and delete rooms
- Associate rooms with houses
- Automatic timestamp management (CreatedAt, UpdatedAt)
- Dynamic ID generation for new rooms
- Proper error handling and validation

## Usage

The room package integrates with the house package, allowing rooms to be managed within the context of houses. Each room is associated with a house via the HouseID field.

## Testing

All components are fully tested with unit tests covering:
- Room creation with proper ID generation
- Room retrieval by ID
- Listing rooms by house ID
- Room updates with timestamp management
- Room deletion

Tests use mock repositories to ensure isolation and reliable testing of business logic.
