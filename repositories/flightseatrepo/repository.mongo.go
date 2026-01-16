package flightseatrepo

import (
	"context"
	"flight-search-booking-service/models"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type flightSeatMongoRepo struct {
	mongoClient *mongo.Client
	dbName      string
	collection  string
}

func NewFlightSeatMongoRepo(mongoClient *mongo.Client, dbName, collection string) *flightSeatMongoRepo {
	return &flightSeatMongoRepo{
		mongoClient: mongoClient,
		dbName:      dbName,
		collection:  collection,
	}
}

// getCollection returns the MongoDB collection for flight seats
func (r *flightSeatMongoRepo) getCollection() *mongo.Collection {
	return r.mongoClient.Database(r.dbName).Collection(r.collection)
}

// GetFlightSeats retrieves flight seat data by flight ID (CRUD - READ)
func (r *flightSeatMongoRepo) GetFlightSeats(ctx context.Context, flightID string) (*models.FlightSeat, error) {
	if flightID == "" {
		return nil, fmt.Errorf("flight ID cannot be empty")
	}

	collection := r.getCollection()
	filter := bson.M{"_id": flightID}

	var flightSeat models.FlightSeat
	err := collection.FindOne(ctx, filter).Decode(&flightSeat)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("flight with ID %s not found", flightID)
		}
		return nil, fmt.Errorf("error fetching flight seats: %v", err)
	}

	return &flightSeat, nil
}

// UpdateFlightSeats updates flight seat data (CRUD - UPDATE)
func (r *flightSeatMongoRepo) UpdateFlightSeats(ctx context.Context, flightID string, flightSeat *models.FlightSeat) error {
	if flightID == "" {
		return fmt.Errorf("flight ID cannot be empty")
	}

	collection := r.getCollection()
	filter := bson.M{"_id": flightID}
	update := bson.M{"$set": flightSeat}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("error updating flight seats: %v", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("flight with ID %s not found", flightID)
	}

	return nil
}

// AddFlightSeats adds a flight with seats to the repository (CRUD - CREATE)
func (r *flightSeatMongoRepo) AddFlightSeats(ctx context.Context, flightSeat *models.FlightSeat) error {
	if flightSeat.FlightID == "" {
		return fmt.Errorf("flight ID cannot be empty")
	}

	collection := r.getCollection()

	// Check if flight seats already exist
	filter := bson.M{"_id": flightSeat.FlightID}
	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return fmt.Errorf("error checking if flight seats exist: %v", err)
	}

	if count > 0 {
		// Update existing flight seats
		update := bson.M{"$set": flightSeat}
		_, err := collection.UpdateOne(ctx, filter, update)
		if err != nil {
			return fmt.Errorf("error updating flight seats: %v", err)
		}
		return nil
	}

	// Insert new flight seats
	_, err = collection.InsertOne(ctx, flightSeat)
	if err != nil {
		return fmt.Errorf("error inserting flight seats: %v", err)
	}

	return nil
}

// RemoveFlightSeats removes a flight from the repository (CRUD - DELETE)
func (r *flightSeatMongoRepo) RemoveFlightSeats(ctx context.Context, flightID string) error {
	if flightID == "" {
		return fmt.Errorf("flight ID cannot be empty")
	}

	collection := r.getCollection()
	filter := bson.M{"_id": flightID}

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("error deleting flight seats: %v", err)
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("flight with ID %s not found", flightID)
	}

	return nil
}
