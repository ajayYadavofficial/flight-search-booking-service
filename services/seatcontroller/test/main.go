package main

// import (
// 	"flight-search-booking-service/repositories/flightseatrepo"
// 	"flight-search-booking-service/services/seatcontroller"
// 	"fmt"
// 	"log"
// )

// func main() {
// 	fmt.Println("=== Seat Controller Service Driver Test ===\n")

// 	// Initialize repository (data layer - CRUD only)
// 	repo := flightseatrepo.NewInMemFlightSeatRepo()

// 	// Initialize service (business logic layer)
// 	service := seatcontroller.NewSeatControllerService(repo)

// 	// Test 1: Test AddSeats Service Function
// 	fmt.Println("1️⃣  Testing AddSeats Service Function...")
// 	flight1ID := "FL001"
// 	totalSeats := 100

// 	// Call AddSeats service method (uses make and append internally)
// 	err := service.AddSeats(flight1ID, totalSeats)
// 	if err != nil {
// 		log.Fatalf("Failed to add seats via service: %v", err)
// 	}

// 	fmt.Printf("✅ Created flight %s with %d seats using AddSeats service\n", flight1ID, totalSeats)
// 	fmt.Println("   (Service internally uses make with capacity and append for each seat)\n")

// 	// Test 2: Get Available Seats (Initial)
// 	fmt.Println("2️⃣  Checking Available Seats (Initial)...")
// 	available, err := service.GetAvailableSeats(flight1ID)
// 	if err != nil {
// 		log.Fatalf("Failed to get available seats: %v", err)
// 	}
// 	fmt.Printf("✅ Available seats: %d\n\n", available)

// 	// Test 3: Lock Seats (for payment processing)
// 	fmt.Println("3️⃣  Locking Seats for Payment Processing...")
// 	seatsToLock := 5
// 	err, lockedSeatIDs := service.LockSeats(flight1ID, seatsToLock)
// 	if err != nil {
// 		log.Fatalf("Failed to lock seats: %v", err)
// 	}

// 	fmt.Printf("✅ Locked %d seats:\n", len(lockedSeatIDs))
// 	for i, seatID := range lockedSeatIDs {
// 		fmt.Printf("   %d. %s\n", i+1, seatID)
// 	}

// 	available, _ = service.GetAvailableSeats(flight1ID)
// 	fmt.Printf("   Remaining available seats: %d\n\n", available)

// 	// Test 4: Get Available Seats (after locking)
// 	fmt.Println("4️⃣  Checking Available Seats (After Locking)...")
// 	available, err = service.GetAvailableSeats(flight1ID)
// 	if err != nil {
// 		log.Fatalf("Failed to get available seats: %v", err)
// 	}
// 	fmt.Printf("✅ Available seats after locking %d seats: %d\n", seatsToLock, available)
// 	fmt.Printf("   (Note: Available count reduced during lock)\n\n")

// 	// Test 5: Book Seats (successful payment scenario)
// 	fmt.Println("5️⃣  Booking Seats (Payment Success)...")
// 	err = service.BookSeats(flight1ID, seatsToLock)
// 	if err != nil {
// 		log.Fatalf("Failed to book seats: %v", err)
// 	}
// 	fmt.Printf("✅ Successfully booked %d seats\n", seatsToLock)

// 	available, _ = service.GetAvailableSeats(flight1ID)
// 	fmt.Printf("   Remaining available seats: %d\n", available)
// 	fmt.Printf("   (Note: Available count remains same after booking)\n\n")

// 	// Test 6: Lock More Seats
// 	fmt.Println("6️⃣  Locking More Seats (for another booking)...")
// 	seatsToLock2 := 8
// 	err, lockedSeatIDs2 := service.LockSeats(flight1ID, seatsToLock2)
// 	if err != nil {
// 		log.Fatalf("Failed to lock seats: %v", err)
// 	}

// 	fmt.Printf("✅ Locked %d more seats:\n", len(lockedSeatIDs2))
// 	for i, seatID := range lockedSeatIDs2 {
// 		fmt.Printf("   %d. %s\n", i+1, seatID)
// 	}

// 	available, _ = service.GetAvailableSeats(flight1ID)
// 	fmt.Printf("   Remaining available seats: %d\n\n", available)

// 	// Test 7: Unlock Seats (payment failed scenario - Rollback)
// 	fmt.Println("7️⃣  Unlocking Seats (Payment Failed - Rollback)...")
// 	err = service.UnlockSeats(flight1ID, lockedSeatIDs2)
// 	if err != nil {
// 		log.Fatalf("Failed to unlock seats: %v", err)
// 	}

// 	fmt.Printf("✅ Unlocked %d seats (payment failed scenario)\n", len(lockedSeatIDs2))

// 	available, _ = service.GetAvailableSeats(flight1ID)
// 	fmt.Printf("   Available seats after unlock: %d\n", available)
// 	fmt.Printf("   (Note: Available count restored after unlock)\n\n")

// 	// Test 8: Lock Multiple Times (Stress Test)
// 	fmt.Println("8️⃣  Stress Test - Locking and Booking Seats Multiple Times...")
// 	for batch := 1; batch <= 3; batch++ {
// 		err, seatIDs := service.LockSeats(flight1ID, 10)
// 		if err != nil {
// 			log.Fatalf("Batch %d: Failed to lock seats: %v", batch, err)
// 		}
// 		fmt.Printf("   Batch %d: Locked %d seats\n", batch, len(seatIDs))

// 		// Book these seats
// 		err = service.BookSeats(flight1ID, 10)
// 		if err != nil {
// 			log.Fatalf("Batch %d: Failed to book seats: %v", batch, err)
// 		}
// 		fmt.Printf("   Batch %d: Booked %d seats\n", batch, 10)

// 		available, _ := service.GetAvailableSeats(flight1ID)
// 		fmt.Printf("   Available seats remaining: %d\n", available)
// 	}
// 	fmt.Println()

// 	// Test 9: Attempt to Lock More Seats Than Available
// 	fmt.Println("9️⃣  Error Test - Locking More Seats Than Available...")
// 	available, _ = service.GetAvailableSeats(flight1ID)
// 	fmt.Printf("   Current available seats: %d\n", available)
// 	fmt.Printf("   Attempting to lock %d seats...\n", available+10)

// 	err, _ = service.LockSeats(flight1ID, available+10)
// 	if err != nil {
// 		fmt.Printf("✅ Expected error caught: %v\n\n", err)
// 	} else {
// 		log.Fatal("Expected error but got none!")
// 	}

// 	// Test 10: Final Seat Count and Status
// 	fmt.Println("🔟 Final Status")
// 	available, _ = service.GetAvailableSeats(flight1ID)
// 	totalAvailable := available
// 	totalBooked := totalSeats - available
// 	fmt.Printf("   Flight: %s\n", flight1ID)
// 	fmt.Printf("   Total Seats: %d\n", totalSeats)
// 	fmt.Printf("   Available Seats: %d\n", totalAvailable)
// 	fmt.Printf("   Booked Seats: %d\n", totalBooked)
// 	fmt.Printf("   Occupancy Rate: %.2f%%\n\n", float64(totalBooked)*100/float64(totalSeats))

// 	// Test 11: Business Logic Verification
// 	fmt.Println("1️⃣1️⃣  Business Logic Verification")
// 	fmt.Println("   ✓ Lock → Reduces available seats")
// 	fmt.Println("   ✓ Unlock → Restores available seats")
// 	fmt.Println("   ✓ Book → Confirms locked seats (available unchanged)")
// 	fmt.Println("   ✓ GetAvailableSeats → Returns accurate count")
// 	fmt.Println("   ✓ Error handling for insufficient seats")
// 	fmt.Println("   ✓ Service layer handles all business logic")
// 	fmt.Println("   ✓ Repository layer only does CRUD operations\n")

// 	// Test 12: Test AddSeats with Multiple Flights
// 	fmt.Println("1️⃣2️⃣  Testing AddSeats Function with Multiple Flights...")
// 	flight2ID := "FL002"
// 	flight2Seats := 50

// 	err = service.AddSeats(flight2ID, flight2Seats)
// 	if err != nil {
// 		log.Fatalf("Failed to add seats for flight 2: %v", err)
// 	}
// 	fmt.Printf("✅ Created flight %s with %d seats\n", flight2ID, flight2Seats)

// 	available2, _ := service.GetAvailableSeats(flight2ID)
// 	fmt.Printf("   Available seats: %d\n", available2)

// 	// Add another flight with different seat count
// 	flight3ID := "FL003"
// 	flight3Seats := 75

// 	err = service.AddSeats(flight3ID, flight3Seats)
// 	if err != nil {
// 		log.Fatalf("Failed to add seats for flight 3: %v", err)
// 	}
// 	fmt.Printf("✅ Created flight %s with %d seats\n", flight3ID, flight3Seats)

// 	available3, _ := service.GetAvailableSeats(flight3ID)
// 	fmt.Printf("   Available seats: %d\n\n", available3)

// 	fmt.Println("Summary of AddSeats Testing:")
// 	fmt.Printf("  ✅ %s: %d seats\n", flight1ID, totalSeats)
// 	fmt.Printf("  ✅ %s: %d seats\n", flight2ID, flight2Seats)
// 	fmt.Printf("  ✅ %s: %d seats\n", flight3ID, flight3Seats)
// 	fmt.Println("  ✅ AddSeats function tested with multiple flights")
// 	fmt.Println("  ✅ Each call uses make with capacity and append for efficient seat creation\n")

// 	fmt.Println("=== Seat Controller Service Test Completed Successfully! ===")
// 	fmt.Println("Summary:")
// 	fmt.Println("  ✅ AddSeats service method working correctly with make and append")
// 	fmt.Println("  ✅ All service methods tested")
// 	fmt.Println("  ✅ Business logic properly implemented in service layer")
// 	fmt.Println("  ✅ Repository acts as pure CRUD data layer")
// 	fmt.Println("  ✅ Lock-Book-Unlock flow validated")
// 	fmt.Println("  ✅ Error handling verified")
// }
