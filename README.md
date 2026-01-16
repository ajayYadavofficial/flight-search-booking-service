# Flight Search and Booking Service

A comprehensive flight search and booking system built with Go, featuring MongoDB integration and RESTful APIs.

## Features

- 🔍 **Flight Search**: Search for direct and connecting flights
- ✈️ **Flight Management**: Add and manage flights
- 💺 **Seat Management**: Lock, book, and unlock seats with real-time availability
- 💳 **Booking System**: Complete booking flow with payment processing
- 🗄️ **MongoDB Integration**: Persistent data storage with auto-increment IDs
- 🐳 **Docker Support**: Containerized deployment with Docker Compose
- 🎯 **Custom Error Handling**: Detailed error categorization (Repository, Payment, Seat errors)

## Architecture

```
├── handlers.go              # HTTP request handlers
├── main.go                  # Application entry point & server setup
├── models/                  # Data models
│   ├── models.go
│   └── status.go
├── services/                # Business logic layer
│   ├── flightbookingsvc/   # Booking service
│   ├── flightsearchsvc/    # Flight search service
│   ├── paymentsvc/         # Payment service
│   └── seatcontroller/     # Seat management service
├── repositories/            # Data access layer
│   ├── flightbookingrepo/  # Booking repository
│   ├── flightsearchrepo/   # Flight repository
│   └── flightseatrepo/     # Seat repository
└── cerrors/                 # Custom error definitions
```

## API Endpoints

### Health Check
```bash
GET /health
```

### Search Flights
```bash
GET /searchflight?departure_code=DEL&arrival_code=BOM&departure_date=2026-01-25&is_price_low_to_high=true&is_direct_flight=false&passengers=2
```

### Add Flight
```bash
POST /addflight
Content-Type: application/json

{
  "airline": "IndiGo",
  "departure_code": "DEL",
  "arrival_code": "BOM",
  "departure_date": "2026-01-25 14:30",
  "arrival_date": "2026-01-25 17:00",
  "total_seats": 100,
  "duration_minutes": 150,
  "price": 4500.50
}
```

### Book Flight
```bash
POST /bookflight
Content-Type: application/json

{
  "user_info": {
    "name": "John Doe",
    "email": "john@example.com",
    "phone": "+1234567890"
  },
  "flight_id": "FL0001",
  "seats": 2
}
```

### Get Booking Details
```bash
GET /getbooking?booking_id=BK0001
```

## Prerequisites

- Go 1.21+
- MongoDB 7.0+
- Docker & Docker Compose (optional)

## Installation & Setup

### Local Development

1. **Clone the repository**
```bash
git clone https://github.com/yourusername/flight-search-booking-service.git
cd flight-search-booking-service
```

2. **Install dependencies**
```bash
go mod download
```

3. **Start MongoDB**
```bash
# Using Docker
docker run -d -p 27017:27017 --name mongodb mongo:7.0

# Or use your local MongoDB installation
```

4. **Run the application**
```bash
go run .
```

The server will start on `http://localhost:8080`

### Docker Deployment

1. **Using Docker Compose (Recommended)**
```bash
docker-compose up --build
```

This starts both MongoDB and the application in containers.

2. **Build and run Docker image manually**
```bash
# Build the image
docker build -t flight-booking-service .

# Run with connection to host MongoDB
docker run -p 8080:8080 \
  -e MONGO_URI=mongodb://host.docker.internal:27017 \
  flight-booking-service
```

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `MONGO_URI` | MongoDB connection string | `mongodb://localhost:27017` |
| `DB_NAME` | Database name | `flight_booking_db` |
| `PORT` | Server port | `:8080` |

## Technologies Used

- **Language**: Go 1.25
- **Database**: MongoDB 7.0
- **Router**: Gorilla Mux
- **Containerization**: Docker & Docker Compose
- **Database Driver**: mongo-go-driver

## Project Highlights

- **Clean Architecture**: Separation of concerns with handlers, services, and repositories
- **Context Propagation**: Request context flows through all layers
- **Custom Error Handling**: Domain-specific errors with clear categorization
- **Auto-increment IDs**: MongoDB counter-based ID generation (FL0001, BK0001)
- **Seat State Management**: AVAILABLE → LOCKED → BOOKED flow
- **Connecting Flights**: Automatic detection of one-stop connecting flights
- **Real-time Availability**: Dynamic seat availability tracking

## Database Collections

- `flights`: Flight information
- `flight_seats`: Seat availability and status
- `bookings`: Booking records
- `counters`: Auto-increment sequence generators

## High-Level Design (HLD)

### System Architecture

```
┌─────────────┐
│   Client    │
└──────┬──────┘
       │ HTTP Requests
       ▼
┌─────────────────────────────────────┐
│        API Layer (Handlers)         │
│  - Request Validation               │
│  - Response Formatting              │
└──────┬──────────────────────────────┘
       │
       ▼
┌─────────────────────────────────────┐
│       Service Layer                 │
│  - Business Logic                   │
│  - Orchestration                    │
│  ┌────────────────────────────────┐ │
│  │ Flight Search Service          │ │
│  │ Flight Booking Service         │ │
│  │ Seat Controller Service        │ │
│  │ Payment Service                │ │
│  └────────────────────────────────┘ │
└──────┬──────────────────────────────┘
       │
       ▼
┌─────────────────────────────────────┐
│     Repository Layer                │
│  - Data Access                      │
│  - CRUD Operations                  │
│  ┌────────────────────────────────┐ │
│  │ Flight Repository              │ │
│  │ Flight Seat Repository         │ │
│  │ Booking Repository             │ │
│  └────────────────────────────────┘ │
└──────┬──────────────────────────────┘
       │
       ▼
┌─────────────────────────────────────┐
│         MongoDB Database            │
│  - flights                          │
│  - flight_seats                     │
│  - bookings                         │
│  - counters                         │
└─────────────────────────────────────┘
```

### Key Design Decisions

1. **Layered Architecture**: Separation of HTTP handlers, business logic, and data access
2. **Repository Pattern**: Abstracts database operations with interface-based design
3. **Context Propagation**: Request context flows through all layers for tracing and cancellation
4. **Auto-increment IDs**: Counter-based ID generation using MongoDB atomic operations
5. **Seat Locking Mechanism**: Two-phase commit pattern (Lock → Payment → Book or Unlock)
6. **Error Categorization**: Custom errors clearly identify layer of failure (Repository, Payment, Seat, etc.)

## Low-Level Design (LLD)

### Component Details

#### 1. Flight Search Service
**Responsibility**: Search and filter flights based on criteria

**Key Methods**:
- `SearchFlights(ctx, req)`: Returns filtered and sorted flights
- `AddFlight(ctx, flightAddReq)`: Creates new flight with seats initialization
- `getDirectFlights()`: Filters flights matching exact departure/arrival
- `getConnectingFlights()`: Creates composite flights via intermediate stops

**Data Flow**:
```
Request → Parse Date → Get All Flights → Filter Direct → Filter Connecting 
→ Sort by Price → Populate Seat Availability → Return
```

**Error Handling**:
- `FLIGHT_SEARCH_REPO_ERROR`: Repository fetch failure
- `INVALID_DEPARTURE_DATE`: Date parsing error
- `FLIGHT_ADD_REPO_ERROR`: Flight creation failure
- `FLIGHT_SEAT_INIT_ERROR`: Seat initialization failure

#### 2. Seat Controller Service
**Responsibility**: Manage seat lifecycle and availability

**Key Methods**:
- `AddSeats(ctx, flightID, seats)`: Initialize seats for new flight
- `LockSeats(ctx, flightID, seats)`: Reserve seats (AVAILABLE → LOCKED)
- `UnlockSeats(ctx, flightID, seatIds)`: Release locked seats (LOCKED → AVAILABLE)
- `BookSeats(ctx, flightID, seats)`: Confirm booking (LOCKED → BOOKED)
- `GetAvailableSeats(ctx, flightID)`: Get available seat count

**State Transitions**:
```
AVAILABLE → (Lock) → LOCKED → (Book) → BOOKED
                         ↓
                    (Unlock)
                         ↓
                     AVAILABLE
```

**Error Handling**:
- `INSUFFICIENT_SEATS`: Not enough available seats
- `SEAT_REPO_NOT_FOUND`: Flight seats not found
- `SEAT_LOCK_ERROR`: Failed to lock seats
- `SEAT_BOOK_ERROR`: Failed to confirm booking

#### 3. Booking Service
**Responsibility**: Orchestrate complete booking flow

**Booking Flow**:
```
1. Check Seat Availability
   ↓
2. Lock Seats
   ↓
3. Process Payment
   ↓ (success)
4. Create Booking Record
   ↓
5. Confirm Seat Booking
   ↓
Return Booking ID

(Payment Failure → Unlock Seats)
(Booking Failure → Unlock Seats)
```

**Error Handling**:
- `SEAT_AVAILABILITY_ERROR`: Failed to check availability
- `SEAT_RESERVATION_ERROR`: Failed to lock seats
- `PAYMENT_FAILURE`: Payment processing failed
- `BOOKING_REPO_ERROR`: Failed to create booking
- `BOOKING_NOT_FOUND`: Booking not found in database

#### 4. Repository Layer

**Flight Repository**:
- Auto-generates Flight IDs (FL0001, FL0002...)
- Stores flight information with time-based fields
- Supports GetAllFlights for search operations

**Flight Seat Repository**:
- Manages seat status (AVAILABLE/LOCKED/BOOKED)
- Tracks available seat count
- Supports atomic updates for seat state changes

**Booking Repository**:
- Auto-generates Booking IDs (BK0001, BK0002...)
- Stores user info, flight details, seat IDs
- Maintains booking timestamps

### Database Schema

#### Flights Collection
```json
{
  "_id": "FL0001",
  "airline": "IndiGo",
  "departure_code": "DEL",
  "arrival_code": "BOM",
  "departure_time": ISODate("2026-01-25T14:30:00Z"),
  "arrival_time": ISODate("2026-01-25T17:00:00Z"),
  "total_seats": 100,
  "duration_minutes": 150,
  "price": 4500.50
}
```

#### Flight_Seats Collection
```json
{
  "_id": "FL0001",
  "seats": [
    {"_id": "1", "status": "AVAILABLE"},
    {"_id": "2", "status": "LOCKED"},
    {"_id": "3", "status": "BOOKED"}
  ],
  "available_seats": 98,
  "total_seats": 100
}
```

#### Bookings Collection
```json
{
  "_id": "BK0001",
  "flight_id": "FL0001",
  "seat_ids": ["1", "2"],
  "user_info": {
    "name": "John Doe",
    "email": "john@example.com",
    "phone": "+1234567890"
  },
  "created_at": ISODate("2026-01-16T03:00:00Z"),
  "updated_at": ISODate("2026-01-16T03:00:00Z")
}
```

#### Counters Collection
```json
{
  "_id": "flightID",
  "sequence_value": 10
},
{
  "_id": "bookingID",
  "sequence_value": 5
}
```

## API Flows

### Flight Search Flow

```
┌────────┐
│ Client │
└───┬────┘
    │ GET /searchflight?departure_code=DEL&arrival_code=BOM&departure_date=2026-01-25&passengers=2
    ▼
┌────────────────────┐
│  API Handler       │
│  - Parse params    │
│  - Validate inputs │
└───┬────────────────┘
    │
    ▼
┌─────────────────────────────┐
│  Flight Search Service      │
│  1. Get all flights (Repo)  │
│  2. Filter direct flights   │
│  3. Filter connecting       │
│  4. Sort by price           │
│  5. Get seat availability   │
└───┬─────────────────────────┘
    │
    ▼
┌────────────────────────────┐
│  Flight Repository         │
│  - Query MongoDB flights   │
│  - Return all flights      │
└───┬────────────────────────┘
    │
    ▼
┌────────────────────────────┐
│  Seat Controller           │
│  - Get available seats     │
│    for each flight         │
└───┬────────────────────────┘
    │
    ▼
┌────────────────────────────┐
│  Response                  │
│  {                         │
│    "success": true,        │
│    "count": 3,             │
│    "flights": [...]        │
│  }                         │
└────────────────────────────┘
```

### Flight Booking Flow

```
┌────────┐
│ Client │
└───┬────┘
    │ POST /bookflight
    │ Body: { flight_id, seats, user_info }
    ▼
┌────────────────────┐
│  API Handler       │
│  - Parse JSON      │
│  - Validate inputs │
└───┬────────────────┘
    │
    ▼
┌─────────────────────────────┐
│  Booking Service            │
│  Step 1: Check Availability │
└───┬─────────────────────────┘
    │
    ▼
┌────────────────────────────┐
│  Seat Controller           │
│  GetAvailableSeats()       │
│  Returns: 50               │
└───┬────────────────────────┘
    │ Available ≥ Requested?
    ▼ YES
┌─────────────────────────────┐
│  Booking Service            │
│  Step 2: Lock Seats         │
└───┬─────────────────────────┘
    │
    ▼
┌────────────────────────────┐
│  Seat Controller           │
│  LockSeats(flightID, 2)    │
│  - Mark seats as LOCKED    │
│  - Decrement available     │
│  Returns: ["1", "2"]       │
└───┬────────────────────────┘
    │
    ▼
┌─────────────────────────────┐
│  Booking Service            │
│  Step 3: Process Payment    │
└───┬─────────────────────────┘
    │
    ▼
┌────────────────────────────┐
│  Payment Service           │
│  ProcessPayment()          │
│  Returns: true/false       │
└───┬────────────────────────┘
    │
    ├─── Payment Failed ────┐
    │                        │
    │ Payment Success        │ UnlockSeats()
    ▼                        │
┌─────────────────────────┐  │
│  Booking Service        │  │
│  Step 4: Create Booking │  │
└───┬─────────────────────┘  │
    │                        │
    ▼                        │
┌────────────────────────┐   │
│  Booking Repository    │   │
│  - Generate BK####     │   │
│  - Save booking        │   │
│  Returns: "BK0001"     │   │
└───┬────────────────────┘   │
    │                        │
    ▼                        │
┌─────────────────────────┐  │
│  Booking Service        │  │
│  Step 5: Confirm Seats  │  │
└───┬─────────────────────┘  │
    │                        │
    ▼                        │
┌────────────────────────┐   │
│  Seat Controller       │   │
│  BookSeats()           │   │
│  LOCKED → BOOKED       │   │
└───┬────────────────────┘   │
    │                        │
    ▼                        ▼
┌────────────────────────────┐
│  Response                  │
│  Success:                  │
│  {                         │
│    "success": true,        │
│    "booking_id": "BK0001"  │
│  }                         │
│                            │
│  Failure:                  │
│  {                         │
│    "success": false,       │
│    "error": "..."          │
│  }                         │
└────────────────────────────┘
```

### Key Flow Characteristics

**Flight Search**:
- Read-heavy operation
- No state changes
- Supports filtering, sorting, and pagination-ready
- Real-time seat availability lookup

**Flight Booking**:
- Write-heavy operation with multiple state transitions
- Transaction-like behavior with rollback on failure
- Idempotent seat locking mechanism
- Automatic cleanup on payment failure

## API Testing with cURL

### 1. Health Check
```bash
curl -X GET "http://localhost:8080/health"
```

**Response:**
```json
{
  "status": "healthy",
  "time": "2026-01-16T08:00:00Z"
}
```

### 2. Search Flights
```bash
curl -X GET "http://localhost:8080/searchflight?departure_code=DEL&arrival_code=BOM&departure_date=2026-01-25&is_price_low_to_high=true&is_direct_flight=false&passengers=2"
```

**Response:**
```json
{
  "success": true,
  "count": 3,
  "flights": [
    {
      "id": "FL0001",
      "airline": "IndiGo",
      "departure_code": "DEL",
      "arrival_code": "BOM",
      "departure_time": "2026-01-25T14:30:00Z",
      "arrival_time": "2026-01-25T17:00:00Z",
      "available_seats": 98,
      "total_seats": 100,
      "duration_minutes": 150,
      "price": 4500.50
    }
  ]
}
```

### 3. Add Flight
```bash
curl -X POST "http://localhost:8080/addflight" \
  -H "Content-Type: application/json" \
  -d '{
    "airline": "IndiGo",
    "departure_code": "DEL",
    "arrival_code": "BOM",
    "departure_date": "2026-01-25 14:30",
    "arrival_date": "2026-01-25 17:00",
    "total_seats": 100,
    "duration_minutes": 150,
    "price": 4500.50
  }'
```

**Response:**
```json
{
  "success": true,
  "flight_id": "FL0001",
  "message": "Flight added successfully"
}
```

### 4. Book Flight
```bash
curl -X POST "http://localhost:8080/bookflight" \
  -H "Content-Type: application/json" \
  -d '{
    "user_info": {
      "name": "John Doe",
      "email": "john@example.com",
      "phone": "+1234567890"
    },
    "flight_id": "FL0001",
    "seats": 2
  }'
```

**Response:**
```json
{
  "success": true,
  "booking_id": "BK0001",
  "message": "Flight booked successfully"
}
```

### 5. Get Booking Details
```bash
curl -X GET "http://localhost:8080/getbooking?booking_id=BK0001"
```

**Response:**
```json
{
  "success": true,
  "booking": {
    "id": "BK0001",
    "flight_id": "FL0001",
    "seat_ids": ["1", "2"],
    "user_info": {
      "name": "John Doe",
      "email": "john@example.com",
      "phone": "+1234567890"
    },
    "created_at": "2026-01-16T03:00:00Z",
    "updated_at": "2026-01-16T03:00:00Z"
  }
}
```

## License

MIT

## Author

Ajay Yadav
