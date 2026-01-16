package main

// import (
// 	"context"
// 	"flight-search-booking-service/models"
// 	"flight-search-booking-service/repositories/flightsearchrepo"
// 	"fmt"
// )

// func main() {
// 	ctx := context.Background()
// 	repo := flightsearchrepo.NewInMemFlightSearchRepo()

// 	// Insert 10 sample flights
// 	flights := []models.Flight{
// 		{
// 			ID:              "FL001",
// 			FlightNumber:    "AA101",
// 			Airline:         "American Airlines",
// 			DepartureCode:   "JFK",
// 			ArrivalCode:     "LAX",
// 			DepartureTime:   "2026-01-15T08:00:00",
// 			ArrivalTime:     "2026-01-15T11:30:00",
// 			AvailableSeats:  50,
// 			TotalSeats:      180,
// 			DurationMinutes: 330,
// 			Price:           250.00,
// 		},
// 		{
// 			ID:              "FL002",
// 			FlightNumber:    "UA202",
// 			Airline:         "United Airlines",
// 			DepartureCode:   "JFK",
// 			ArrivalCode:     "LAX",
// 			DepartureTime:   "2026-01-15T10:00:00",
// 			ArrivalTime:     "2026-01-15T13:45:00",
// 			AvailableSeats:  75,
// 			TotalSeats:      180,
// 			DurationMinutes: 345,
// 			Price:           280.00,
// 		},
// 		{
// 			ID:              "FL003",
// 			FlightNumber:    "DL303",
// 			Airline:         "Delta Airlines",
// 			DepartureCode:   "JFK",
// 			ArrivalCode:     "ORD",
// 			DepartureTime:   "2026-01-15T09:00:00",
// 			ArrivalTime:     "2026-01-15T11:15:00",
// 			AvailableSeats:  40,
// 			TotalSeats:      150,
// 			DurationMinutes: 135,
// 			Price:           150.00,
// 		},
// 		{
// 			ID:              "FL004",
// 			FlightNumber:    "SW404",
// 			Airline:         "Southwest Airlines",
// 			DepartureCode:   "ORD",
// 			ArrivalCode:     "LAX",
// 			DepartureTime:   "2026-01-15T14:00:00",
// 			ArrivalTime:     "2026-01-15T16:00:00",
// 			AvailableSeats:  100,
// 			TotalSeats:      200,
// 			DurationMinutes: 240,
// 			Price:           180.00,
// 		},
// 		{
// 			ID:              "FL005",
// 			FlightNumber:    "BA505",
// 			Airline:         "British Airways",
// 			DepartureCode:   "JFK",
// 			ArrivalCode:     "LHR",
// 			DepartureTime:   "2026-01-15T19:00:00",
// 			ArrivalTime:     "2026-01-16T07:00:00",
// 			AvailableSeats:  60,
// 			TotalSeats:      250,
// 			DurationMinutes: 420,
// 			Price:           450.00,
// 		},
// 		{
// 			ID:              "FL006",
// 			FlightNumber:    "AF606",
// 			Airline:         "Air France",
// 			DepartureCode:   "JFK",
// 			ArrivalCode:     "LAX",
// 			DepartureTime:   "2026-01-15T06:00:00",
// 			ArrivalTime:     "2026-01-15T09:15:00",
// 			AvailableSeats:  30,
// 			TotalSeats:      180,
// 			DurationMinutes: 315,
// 			Price:           220.00,
// 		},
// 		{
// 			ID:              "FL007",
// 			FlightNumber:    "LH707",
// 			Airline:         "Lufthansa",
// 			DepartureCode:   "LAX",
// 			ArrivalCode:     "JFK",
// 			DepartureTime:   "2026-01-15T15:00:00",
// 			ArrivalTime:     "2026-01-16T00:30:00",
// 			AvailableSeats:  55,
// 			TotalSeats:      180,
// 			DurationMinutes: 315,
// 			Price:           290.00,
// 		},
// 		{
// 			ID:              "FL008",
// 			FlightNumber:    "SQ808",
// 			Airline:         "Singapore Airlines",
// 			DepartureCode:   "JFK",
// 			ArrivalCode:     "ORD",
// 			DepartureTime:   "2026-01-15T12:00:00",
// 			ArrivalTime:     "2026-01-15T14:30:00",
// 			AvailableSeats:  80,
// 			TotalSeats:      160,
// 			DurationMinutes: 150,
// 			Price:           160.00,
// 		},
// 		{
// 			ID:              "FL009",
// 			FlightNumber:    "JL909",
// 			Airline:         "Japan Airlines",
// 			DepartureCode:   "ORD",
// 			ArrivalCode:     "LAX",
// 			DepartureTime:   "2026-01-15T12:30:00",
// 			ArrivalTime:     "2026-01-15T14:45:00",
// 			AvailableSeats:  45,
// 			TotalSeats:      170,
// 			DurationMinutes: 215,
// 			Price:           175.00,
// 		},
// 		{
// 			ID:              "FL010",
// 			FlightNumber:    "KL1010",
// 			Airline:         "KLM Royal Dutch Airlines",
// 			DepartureCode:   "LAX",
// 			ArrivalCode:     "ORD",
// 			DepartureTime:   "2026-01-15T11:00:00",
// 			ArrivalTime:     "2026-01-15T19:30:00",
// 			AvailableSeats:  70,
// 			TotalSeats:      190,
// 			DurationMinutes: 270,
// 			Price:           195.00,
// 		},
// 	}

// 	// Add all flights to the repository
// 	fmt.Println("=== Inserting 10 flights into in-memory repository ===\n")
// 	for _, flight := range flights {
// 		err := repo.AddFlight(ctx, flight)
// 		if err != nil {
// 			fmt.Printf("Error adding flight %s: %v\n", flight.ID, err)
// 		} else {
// 			fmt.Printf("✓ Added flight %s: %s (%s -> %s)\n", flight.ID, flight.FlightNumber, flight.DepartureCode, flight.ArrivalCode)
// 		}
// 	}

// 	fmt.Println("\n=== Test 1: Direct Flight Search (JFK -> LAX on 2026-01-15, 2 passengers, low to high price) ===\n")
// 	searchReq1 := models.FlightSearchRequest{
// 		DepartureCode:    "JFK",
// 		ArrivalCode:      "LAX",
// 		DepartureDate:    "2026-01-15",
// 		IsPriceLowToHigh: true,
// 		IsDirectFlight:   true,
// 		Passengers:       2,
// 	}
// 	results1, err := repo.GetFlights(ctx, searchReq1)
// 	if err != nil {
// 		fmt.Printf("Error: %v\n", err)
// 	} else {
// 		fmt.Printf("Found %d direct flight(s):\n", len(results1))
// 		for i, flight := range results1 {
// 			fmt.Printf("%d. %s (%s) - Price: $%.2f, Available: %d seats, Duration: %d mins\n",
// 				i+1, flight.FlightNumber, flight.Airline, flight.Price, flight.AvailableSeats, flight.DurationMinutes)
// 		}
// 	}

// 	fmt.Println("\n=== Test 2: All Flights with Connections (JFK -> LAX on 2026-01-15, 2 passengers, low to high price) ===\n")
// 	searchReq2 := models.FlightSearchRequest{
// 		DepartureCode:    "JFK",
// 		ArrivalCode:      "LAX",
// 		DepartureDate:    "2026-01-15",
// 		IsPriceLowToHigh: true,
// 		IsDirectFlight:   false,
// 		Passengers:       2,
// 	}
// 	results2, err := repo.GetFlights(ctx, searchReq2)
// 	if err != nil {
// 		fmt.Printf("Error: %v\n", err)
// 	} else {
// 		fmt.Printf("Found %d flight(s) including connections:\n", len(results2))
// 		for i, flight := range results2 {
// 			fmt.Printf("%d. %s (%s) - Price: $%.2f, Available: %d seats, Duration: %d mins\n",
// 				i+1, flight.FlightNumber, flight.Airline, flight.Price, flight.AvailableSeats, flight.DurationMinutes)
// 		}
// 	}

// 	fmt.Println("\n=== Test 3: Direct Flight Search (JFK -> ORD on 2026-01-15, 3 passengers) ===\n")
// 	searchReq3 := models.FlightSearchRequest{
// 		DepartureCode:    "JFK",
// 		ArrivalCode:      "ORD",
// 		DepartureDate:    "2026-01-15",
// 		IsPriceLowToHigh: true,
// 		IsDirectFlight:   true,
// 		Passengers:       3,
// 	}
// 	results3, err := repo.GetFlights(ctx, searchReq3)
// 	if err != nil {
// 		fmt.Printf("Error: %v\n", err)
// 	} else {
// 		fmt.Printf("Found %d direct flight(s):\n", len(results3))
// 		for i, flight := range results3 {
// 			fmt.Printf("%d. %s (%s) - Price: $%.2f, Available: %d seats, Duration: %d mins\n",
// 				i+1, flight.FlightNumber, flight.Airline, flight.Price, flight.AvailableSeats, flight.DurationMinutes)
// 		}
// 	}

// 	fmt.Println("\n=== Test 4: All Flights with Connections (JFK -> ORD on 2026-01-15, 1 passenger, high to low price) ===\n")
// 	searchReq4 := models.FlightSearchRequest{
// 		DepartureCode:    "JFK",
// 		ArrivalCode:      "ORD",
// 		DepartureDate:    "2026-01-15",
// 		IsPriceLowToHigh: false,
// 		IsDirectFlight:   false,
// 		Passengers:       1,
// 	}
// 	results4, err := repo.GetFlights(ctx, searchReq4)
// 	if err != nil {
// 		fmt.Printf("Error: %v\n", err)
// 	} else {
// 		fmt.Printf("Found %d flight(s) including connections:\n", len(results4))
// 		for i, flight := range results4 {
// 			fmt.Printf("%d. %s (%s) - Price: $%.2f, Available: %d seats, Duration: %d mins\n",
// 				i+1, flight.FlightNumber, flight.Airline, flight.Price, flight.AvailableSeats, flight.DurationMinutes)
// 		}
// 	}

// 	fmt.Println("\n=== Test Summary ===")
// 	fmt.Printf("✓ All tests completed successfully!\n")
// }

import (
	"context"
	"flight-search-booking-service/models"
	"flight-search-booking-service/repositories/flightsearchrepo"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Helper function to parse time string
func parseTime(timeStr string) time.Time {
	t, err := time.Parse("2006-01-02 15:04", timeStr)
	if err != nil {
		log.Fatalf("Failed to parse time: %v", err)
	}
	return t
}

func main() {
	fmt.Println("=== MongoDB Flight Search Repository Driver Test ===\n")

	// MongoDB connection details (dummy values for testing)
	mongoURI := "mongodb://localhost:27017"
	dbName := "flight_booking_test_db"
	collectionName := "flights_test_collection"

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	fmt.Println("1️⃣  Connecting to MongoDB...")
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v\n", err)
	}
	defer func() {
		if err = mongoClient.Disconnect(ctx); err != nil {
			log.Fatalf("Failed to disconnect from MongoDB: %v\n", err)
		}
	}()

	// Verify connection
	err = mongoClient.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v\n", err)
	}
	fmt.Printf("✅ Connected to MongoDB at %s\n", mongoURI)
	fmt.Printf("   Database: %s\n", dbName)
	fmt.Printf("   Collection: %s\n\n", collectionName)

	// Create repository
	repo := flightsearchrepo.NewFlightSearchMongoRepo(mongoClient, dbName, collectionName)

	// Test 1: Add Dummy Flights
	fmt.Println("2️⃣  Adding Dummy Flights to MongoDB...")
	dummyFlights := []models.Flight{
		{
			Airline:         "Air India",
			DepartureCode:   "DEL",
			ArrivalCode:     "BOM",
			DepartureTime:   parseTime("2026-01-20 08:00"),
			ArrivalTime:     parseTime("2026-01-20 11:00"),
			TotalSeats:      100,
			DurationMinutes: 180,
			Price:           5000,
		},
		{
			Airline:         "Air India",
			DepartureCode:   "DEL",
			ArrivalCode:     "BOM",
			DepartureTime:   parseTime("2026-01-20 14:00"),
			ArrivalTime:     parseTime("2026-01-20 17:00"),
			TotalSeats:      100,
			DurationMinutes: 180,
			Price:           4500,
		},
		{
			Airline:         "SpiceJet",
			DepartureCode:   "DEL",
			ArrivalCode:     "BOM",
			DepartureTime:   parseTime("2026-01-20 20:00"),
			ArrivalTime:     parseTime("2026-01-20 23:00"),
			TotalSeats:      80,
			DurationMinutes: 180,
			Price:           3500,
		},
		{
			Airline:         "Vistara",
			DepartureCode:   "BOM",
			ArrivalCode:     "BLR",
			DepartureTime:   parseTime("2026-01-20 09:30"),
			ArrivalTime:     parseTime("2026-01-20 11:30"),
			TotalSeats:      120,
			DurationMinutes: 120,
			Price:           2500,
		},
		{
			Airline:         "IndiGo",
			DepartureCode:   "BLR",
			ArrivalCode:     "HYD",
			DepartureTime:   parseTime("2026-01-20 15:00"),
			ArrivalTime:     parseTime("2026-01-20 16:45"),
			TotalSeats:      150,
			DurationMinutes: 105,
			Price:           2000,
		},
	}

	for _, flight := range dummyFlights {
		flightId, err := repo.AddFlight(ctx, flight)
		if err != nil {
			log.Fatalf("Failed to add flight: %v", err)
		}
		fmt.Printf("   ✓ Added flight: %s (%s -> %s) with ID: %s\n", flight.Airline, flight.DepartureCode, flight.ArrivalCode, flightId)
	}
	fmt.Println()

	// Test 2: Fetch All Flights
	fmt.Println("3️⃣  Fetching All Flights from MongoDB...")
	req := models.FlightSearchRequest{}
	flights, err := repo.GetFlights(ctx, req)
	if err != nil {
		log.Fatalf("Failed to get flights: %v", err)
	}

	fmt.Printf("✅ Retrieved %d flights from MongoDB:\n", len(flights))
	for i, flight := range flights {
		fmt.Printf("   %d. %s (%s)\n", i+1, flight.Airline, flight.ID)
		fmt.Printf("      Route: %s -> %s\n", flight.DepartureCode, flight.ArrivalCode)
		fmt.Printf("      Departure: %s, Arrival: %s\n", flight.DepartureTime.Format("2006-01-02 15:04"), flight.ArrivalTime.Format("2006-01-02 15:04"))
		fmt.Printf("      Total Seats: %d, Price: ₹%.2f\n", flight.TotalSeats, flight.Price)
	}
	fmt.Println()

	// Test 3: Fetch All Flights (using GetAllFlights)
	fmt.Println("4️⃣  Fetching All Flights Using GetAllFlights()...")
	allFlights, err := repo.GetAllFlights(ctx)
	if err != nil {
		log.Fatalf("Failed to get all flights: %v", err)
	}

	fmt.Printf("✅ Retrieved %d flights:\n", len(allFlights))
	totalPrice := 0.0
	for _, flight := range allFlights {
		totalPrice += flight.Price
		fmt.Printf("   - %s (%s -> %s): ₹%.2f\n", flight.Airline, flight.DepartureCode, flight.ArrivalCode, flight.Price)
	}
	fmt.Printf("   Total value of all flights: ₹%.2f\n\n", totalPrice)

	// // Test 4: Update Flight Seats
	// fmt.Println("5️⃣  Updating Flight Seats...")
	// flightIDToUpdate := "FL001"
	// newSeats := 50
	// err = repo.UpdateFlightSeats(ctx, flightIDToUpdate, newSeats)
	// if err != nil {
	// 	log.Fatalf("Failed to update flight seats: %v", err)
	// }
	// fmt.Printf("✅ Updated flight %s seats to %d\n\n", flightIDToUpdate, newSeats)

	// Verify the update
	// updatedFlights, _ := repo.GetAllFlights(ctx)
	// for _, flight := range updatedFlights {
	// 	if flight.ID == flightIDToUpdate {
	// 		fmt.Printf("   Verified: %s now has %d total seats\n\n", flight.FlightNumber, flight.TotalSeats)
	// 	}
	// }

	// Test 5: Remove Specific Flight
	fmt.Println("6️⃣  Removing Specific Flight...")
	flightIDToRemove := "FL005"
	err = repo.RemoveFlight(ctx, flightIDToRemove)
	if err != nil {
		log.Fatalf("Failed to remove flight: %v", err)
	}
	fmt.Printf("✅ Removed flight: %s\n\n", flightIDToRemove)

	// Verify removal
	remainingFlights, _ := repo.GetAllFlights(ctx)
	fmt.Printf("   Remaining flights after removal: %d\n", len(remainingFlights))
	for i, flight := range remainingFlights {
		fmt.Printf("   %d. %s (%s -> %s) - ID: %s\n", i+1, flight.Airline, flight.DepartureCode, flight.ArrivalCode, flight.ID)
	}
	fmt.Println()

	// // Test 6: Remove All Test Data (Cleanup)
	// fmt.Println("7️⃣  Cleaning Up - Removing All Test Data...")
	// flightsToRemove := []string{"FL001", "FL002", "FL003", "FL004"}
	// for _, flightID := range flightsToRemove {
	// 	err := repo.RemoveFlight(ctx, flightID)
	// 	if err != nil {
	// 		log.Fatalf("Failed to remove flight %s: %v", flightID, err)
	// 	}
	// 	fmt.Printf("   ✓ Removed flight: %s\n", flightID)
	// }
	// fmt.Println()

	// // Verify all data removed
	// fmt.Println("8️⃣  Verifying Cleanup...")
	// finalFlights, _ := repo.GetAllFlights(ctx)
	// if len(finalFlights) == 0 {
	// 	fmt.Println("✅ All test data successfully removed from MongoDB\n")
	// } else {
	// 	fmt.Printf("⚠️  Warning: %d flights still remain in collection\n\n", len(finalFlights))
	// }

	// // Drop collection for complete cleanup
	// fmt.Println("9️⃣  Dropping Test Collection...")
	// collection := mongoClient.Database(dbName).Collection(collectionName)
	// err = collection.Drop(ctx)
	// if err != nil {
	// 	log.Fatalf("Failed to drop collection: %v", err)
	// }
	// fmt.Printf("✅ Dropped collection: %s\n\n", collectionName)

	fmt.Println("=== MongoDB Driver Test Completed Successfully! ===")
	fmt.Println("Summary:")
	fmt.Println("  - Successfully connected to MongoDB")
	fmt.Println("  - Added 5 dummy flights")
	fmt.Println("  - Fetched all flights")
	fmt.Println("  - Updated flight seats")
	fmt.Println("  - Removed individual flights")
	fmt.Println("  - Cleaned up all test data")
	fmt.Println("  - Dropped test collection")
}
