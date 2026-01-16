package main

import (
	"context"
	"flight-search-booking-service/repositories/flightbookingrepo"
	"flight-search-booking-service/repositories/flightsearchrepo"
	"flight-search-booking-service/repositories/flightseatrepo"
	"flight-search-booking-service/services/flightbookingsvc"
	"flight-search-booking-service/services/flightsearchsvc"
	"flight-search-booking-service/services/paymentsvc"
	"flight-search-booking-service/services/seatcontroller"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type APIServer struct {
	flightSearchService flightsearchsvc.FlightController
	bookingService      flightbookingsvc.Booking
}

func main() {
	// MongoDB configuration - read from environment variables with defaults
	mongoURI := getEnv("MONGO_URI", "mongodb://localhost:27017")
	dbName := getEnv("DB_NAME", "flight_booking_db")
	flightsCollection := "flights"
	flightSeatsCollection := "flight_seats"
	bookingsCollection := "bookings"
	port := getEnv("PORT", ":8080")

	// Create context with timeout for MongoDB connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	log.Println("Connecting to MongoDB...")
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer func() {
		if err = mongoClient.Disconnect(context.Background()); err != nil {
			log.Printf("Failed to disconnect from MongoDB: %v", err)
		}
	}()

	// Verify MongoDB connection
	err = mongoClient.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}
	log.Printf("✅ Connected to MongoDB at %s", mongoURI)

	// Initialize MongoDB repositories
	flightSearchRepo := flightsearchrepo.NewFlightSearchMongoRepo(mongoClient, dbName, flightsCollection)
	flightSeatRepo := flightseatrepo.NewFlightSeatMongoRepo(mongoClient, dbName, flightSeatsCollection)
	bookingRepo := flightbookingrepo.NewFlightBookingMongoRepo(mongoClient, dbName, bookingsCollection)

	// Initialize services
	seatControllerService := seatcontroller.NewSeatControllerService(flightSeatRepo)
	flightSearchService := flightsearchsvc.NewFlightSearchService(flightSearchRepo, seatControllerService)
	paymentService := paymentsvc.NewPaymentService()
	bookingService := flightbookingsvc.NewFlightBookingService(paymentService, seatControllerService, bookingRepo)

	// Create API server
	apiServer := &APIServer{
		flightSearchService: flightSearchService,
		bookingService:      bookingService,
	}

	// Setup routes
	router := mux.NewRouter()
	router.HandleFunc("/searchflight", apiServer.handleSearchFlight).Methods("GET")
	router.HandleFunc("/addflight", apiServer.handleAddFlight).Methods("POST")
	router.HandleFunc("/bookflight", apiServer.handleBookFlight).Methods("POST")
	router.HandleFunc("/getbooking", apiServer.handleGetBooking).Methods("GET")
	router.HandleFunc("/health", apiServer.handleHealth).Methods("GET")

	// Start HTTP server
	log.Printf("🚀 Server starting on port %s", port)
	log.Printf("   GET  /searchflight - Search for flights")
	log.Printf("   POST /addflight - Add a new flight")
	log.Printf("   POST /bookflight - Book a flight")
	log.Printf("   GET  /getbooking - Get booking details")
	log.Printf("   GET  /health - Health check")

	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// getEnv gets environment variable or returns default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
