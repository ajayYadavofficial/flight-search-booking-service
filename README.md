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

## License

MIT

## Author

Ajay Yadav
