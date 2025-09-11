# Room Model Implementation Plan

**Feature Branch**: `002-room-model`  
**Created**: 2025-09-10  
**Status**: Draft  

## Overview
This plan outlines the implementation of the room model for the home inventory system, following the same patterns and architecture as the existing house model.

## Implementation Steps

### Phase 1: Model Definition
1. Create `internal/rooms/models.go` with:
   - Room struct definition
   - RoomDimensions struct 
   - RoomCreateRequest struct
   - RoomUpdateRequest struct

### Phase 2: Repository Interface
2. Create `internal/rooms/repository.go` with:
   - Repository interface defining all required operations

### Phase 3: Service Interface
3. Create `internal/rooms/service.go` with:
   - Service interface defining all required business logic operations

### Phase 4: Handler Interface
4. Create `internal/rooms/handler.go` with:
   - Handler interface defining all HTTP endpoint operations

### Phase 5: Implementation
5. Create `internal/rooms/repository_impl.go` with:
   - Repository implementation using database operations
   - All required CRUD operations

6. Create `internal/rooms/service_impl.go` with:
   - Service implementation with business logic
   - Validation and error handling
   - All required service methods

7. Create `internal/rooms/handler_impl.go` with:
   - HTTP handler implementation
   - Request parsing and response formatting
   - All required endpoint handlers

### Phase 6: Integration
8. Update `main.go` to register room routes:
   - Add room router registration
   - Integrate with existing routing structure

### Phase 7: Testing
9. Create or update tests in `internal/rooms/service_test.go`:
   - Unit tests for service methods
   - Test validation logic
   - Test error conditions

## Technical Details

### Data Model Relationships
- Rooms are associated with houses via house_id foreign key
- Room names must be unique within the context of a house
- All entities follow the same timestamp pattern (created_at, updated_at)

### Database Considerations
- Need to create room table in database schema
- Must implement proper foreign key constraints for house_id
- Indexes needed on house_id for efficient querying

### API Endpoints
- POST /houses/{house_id}/rooms - Create room
- GET /houses/{house_id}/rooms - List rooms for a house
- GET /rooms/{room_id} - Get specific room
- PUT /rooms/{room_id} - Update room
- DELETE /rooms/{room_id} - Delete room

## Dependencies
- Database connection (already exists)
- House model (already implemented)
- Standard Go libraries (context, time, etc.)

## Success Criteria
- All API endpoints function correctly
- Database constraints are properly enforced
- Service validation works as expected
- Unit tests pass for all components
- Integration with existing house model works seamlessly
