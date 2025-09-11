# Room Model Tasks

**Feature Branch**: `002-room-model`  
**Created**: 2025-09-10  

## Task Breakdown

### Task 1: Create Room Models
**File**: `internal/rooms/models.go`
**Description**: Define the Room data structures following the same pattern as houses
**Subtasks**:
- [ ] Define Room struct with id, name, house_id, description, dimensions, created_at, updated_at
- [ ] Define RoomDimensions struct 
- [ ] Define RoomCreateRequest struct
- [ ] Define RoomUpdateRequest struct

### Task 2: Create Repository Interface
**File**: `internal/rooms/repository.go`
**Description**: Define the repository interface for room operations
**Subtasks**:
- [ ] Define Repository interface with CreateRoom, GetRoomByID, ListRoomsByHouseID, UpdateRoom, DeleteRoom methods

### Task 3: Create Service Interface
**File**: `internal/rooms/service.go`
**Description**: Define the service interface for room business logic
**Subtasks**:
- [ ] Define Service interface with CreateRoom, GetRoom, ListRooms, UpdateRoom, DeleteRoom methods

### Task 4: Create Handler Interface
**File**: `internal/rooms/handler.go`
**Description**: Define the handler interface for room HTTP endpoints
**Subtasks**:
- [ ] Define Handler interface with CreateRoom, GetRoom, ListRooms, UpdateRoom, DeleteRoom methods

### Task 5: Implement Repository
**File**: `internal/rooms/repository_impl.go`
**Description**: Implement the repository using database operations
**Subtasks**:
- [ ] Implement CreateRoom method
- [ ] Implement GetRoomByID method
- [ ] Implement ListRoomsByHouseID method
- [ ] Implement UpdateRoom method
- [ ] Implement DeleteRoom method

### Task 6: Implement Service
**File**: `internal/rooms/service_impl.go`
**Description**: Implement the service with business logic and validation
**Subtasks**:
- [ ] Implement CreateRoom method with validation
- [ ] Implement GetRoom method
- [ ] Implement ListRooms method
- [ ] Implement UpdateRoom method with validation
- [ ] Implement DeleteRoom method

### Task 7: Implement Handler
**File**: `internal/rooms/handler_impl.go`
**Description**: Implement HTTP handlers for room endpoints
**Subtasks**:
- [ ] Implement CreateRoom handler
- [ ] Implement GetRoom handler
- [ ] Implement ListRooms handler
- [ ] Implement UpdateRoom handler
- [ ] Implement DeleteRoom handler

### Task 8: Update Main Application
**File**: `main.go`
**Description**: Register room routes in the main application
**Subtasks**:
- [ ] Import room packages
- [ ] Create room router
- [ ] Register room routes with house router

### Task 9: Add Tests
**File**: `internal/rooms/service_test.go`
**Description**: Write unit tests for room service
**Subtasks**:
- [ ] Test CreateRoom validation
- [ ] Test GetRoom functionality
- [ ] Test ListRooms functionality
- [ ] Test UpdateRoom validation
- [ ] Test DeleteRoom functionality
- [ ] Test error conditions

## Dependencies
- Database connection (already exists)
- House model (already implemented)

## Success Criteria
- All tasks completed according to plan
- All tests pass
- API endpoints function correctly
- Room model integrates properly with house model
