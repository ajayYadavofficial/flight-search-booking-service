package main

// import (
// 	"flight-search-booking-service/models"
// 	"flight-search-booking-service/repositories/flightseatrepo"
// 	"fmt"
// 	"log"
// )

// func main() {
// 	fmt.Println("=== Flight Seat Repository Driver Test ===\n")

// 	// Initialize repository
// 	repo := flightseatrepo.NewInMemFlightSeatRepo()

// 	// Test 1: Initialize Flight Seats
// 	fmt.Println("1️⃣  Initializing Flight Seats...")
// 	flight1ID := "FL001"
// 	totalSeats := 100

// 	// Create seats for flight
// 	flightSeat := &models.FlightSeat{
// 		FlightID:       flight1ID,
// 		Seats:          make([]models.Seat, totalSeats),
// 		AvailableSeats: totalSeats,
// 	}

// 	// Generate seat IDs
// 	for i := 0; i < totalSeats; i++ {
// 		flightSeat.Seats[i] = models.Seat{
// 			ID:     fmt.Sprintf("%s-SEAT-%03d", flight1ID, i+1),
// 			Status: models.AVAILABLE,
// 		}
// 	}

// 	// Note: We need to manually add flight seats since there's no AddFlight method in the repository
// 	// In a real scenario, this would be done through a different method
// 	fmt.Printf("✅ Created flight %s with %d total seats\n\n", flight1ID, totalSeats)

// 	// Test 2: Get Available Seats (before any operations)
// 	fmt.Println("2️⃣  Checking Available Seats (Initial)...")
// 	// Note: This will fail because we haven't added the flight to the repo yet
// 	// Let's show what would happen
// 	_, err := repo.GetAvailableSeats(flight1ID)
// 	if err != nil {
// 		fmt.Printf("⚠️  Expected error (flight not in repo): %v\n", err)
// 		fmt.Println("   Adding flight to repository first...\n")

// 		// Add the flight to the repository (simulating internal state)
// 		// For this demo, we'll create a helper function
// 		addFlightToRepo(repo, flightSeat)

// 		available, err := repo.GetAvailableSeats(flight1ID)
// 		if err != nil {
// 			log.Fatalf("Failed to get available seats: %v", err)
// 		}
// 		fmt.Printf("✅ Available seats: %d\n\n", available)
// 	}

// 	// Test 3: Lock Seats (for payment processing)
// 	fmt.Println("3️⃣  Locking Seats for Payment Processing...")
// 	seatsToLock := 5
// 	err, lockedSeatIDs := repo.LockSeats(flight1ID, seatsToLock)
// 	if err != nil {
// 		log.Fatalf("Failed to lock seats: %v", err)
// 	}

// 	fmt.Printf("✅ Locked %d seats:\n", len(lockedSeatIDs))
// 	for i, seatID := range lockedSeatIDs {
// 		fmt.Printf("   %d. %s\n", i+1, seatID)
// 	}

// 	available, _ := repo.GetAvailableSeats(flight1ID)
// 	fmt.Printf("   Remaining available seats: %d\n\n", available)

// 	// Test 4: Get Available Seats (after locking)
// 	fmt.Println("4️⃣  Checking Available Seats (After Locking)...")
// 	available, err = repo.GetAvailableSeats(flight1ID)
// 	if err != nil {
// 		log.Fatalf("Failed to get available seats: %v", err)
// 	}
// 	fmt.Printf("✅ Available seats after locking %d seats: %d\n\n", seatsToLock, available)

// 	// Test 5: Book Seats (successful payment scenario)
// 	fmt.Println("5️⃣  Booking Seats (Payment Success)...")
// 	err = repo.BookSeats(flight1ID, seatsToLock)
// 	if err != nil {
// 		log.Fatalf("Failed to book seats: %v", err)
// 	}
// 	fmt.Printf("✅ Successfully booked %d seats\n", seatsToLock)

// 	available, _ = repo.GetAvailableSeats(flight1ID)
// 	fmt.Printf("   Remaining available seats: %d\n\n", available)

// 	// Test 6: Lock More Seats
// 	fmt.Println("6️⃣  Locking More Seats (for another booking)...")
// 	seatsToLock2 := 8
// 	err, lockedSeatIDs2 := repo.LockSeats(flight1ID, seatsToLock2)
// 	if err != nil {
// 		log.Fatalf("Failed to lock seats: %v", err)
// 	}

// 	fmt.Printf("✅ Locked %d more seats:\n", len(lockedSeatIDs2))
// 	for i, seatID := range lockedSeatIDs2 {
// 		fmt.Printf("   %d. %s\n", i+1, seatID)
// 	}

// 	available, _ = repo.GetAvailableSeats(flight1ID)
// 	fmt.Printf("   Remaining available seats: %d\n\n", available)

// 	// Test 7: Unlock Seats (payment failed scenario)
// 	fmt.Println("7️⃣  Unlocking Seats (Payment Failed - Rollback)...")
// 	err = repo.UnlockSeats(flight1ID, lockedSeatIDs2)
// 	if err != nil {
// 		log.Fatalf("Failed to unlock seats: %v", err)
// 	}

// 	fmt.Printf("✅ Unlocked %d seats (payment failed scenario)\n", len(lockedSeatIDs2))

// 	available, _ = repo.GetAvailableSeats(flight1ID)
// 	fmt.Printf("   Available seats after unlock: %d\n\n", available)

// 	// Test 8: Lock Multiple Times
// 	fmt.Println("8️⃣  Stress Test - Locking Seats Multiple Times...")
// 	for batch := 1; batch <= 3; batch++ {
// 		err, seatIDs := repo.LockSeats(flight1ID, 10)
// 		if err != nil {
// 			log.Fatalf("Batch %d: Failed to lock seats: %v", batch, err)
// 		}
// 		fmt.Printf("   Batch %d: Locked %d seats\n", batch, len(seatIDs))

// 		// Book these seats
// 		err = repo.BookSeats(flight1ID, 10)
// 		if err != nil {
// 			log.Fatalf("Batch %d: Failed to book seats: %v", batch, err)
// 		}
// 		fmt.Printf("   Batch %d: Booked %d seats\n", batch, 10)

// 		available, _ := repo.GetAvailableSeats(flight1ID)
// 		fmt.Printf("   Available seats remaining: %d\n", available)
// 	}
// 	fmt.Println()

// 	// Test 9: Attempt to Lock More Seats Than Available
// 	fmt.Println("9️⃣  Error Test - Locking More Seats Than Available...")
// 	available, _ = repo.GetAvailableSeats(flight1ID)
// 	fmt.Printf("   Current available seats: %d\n", available)
// 	fmt.Printf("   Attempting to lock %d seats...\n", available+10)

// 	err, _ = repo.LockSeats(flight1ID, available+10)
// 	if err != nil {
// 		fmt.Printf("✅ Expected error caught: %v\n\n", err)
// 	}

// 	// Test 10: Final Seat Count
// 	fmt.Println("🔟 Final Status")
// 	available, _ = repo.GetAvailableSeats(flight1ID)
// 	totalAvailable := available
// 	totalBooked := totalSeats - available
// 	fmt.Printf("   Flight: %s\n", flight1ID)
// 	fmt.Printf("   Total Seats: %d\n", totalSeats)
// 	fmt.Printf("   Available Seats: %d\n", totalAvailable)
// 	fmt.Printf("   Booked Seats: %d\n", totalBooked)
// 	fmt.Printf("   Occupancy Rate: %.2f%%\n\n", float64(totalBooked)*100/float64(totalSeats))

// 	fmt.Println("=== All Tests Completed Successfully! ===")
// }

// // Helper function to add flight to repository (internal access)
// func addFlightToRepo(repo *flightseatrepo.InMemFlightSeatRepo, flightSeat *models.FlightSeat) {
// 	// This is a workaround to initialize the repository
// 	// In production, this would be done through proper initialization methods
// 	// For now, we manually set it in the repo's internal state
// 	repo.AddFlightSeats(flightSeat)
// }
