package flightsearchrepo

import (
	"context"
	"flight-search-booking-service/models"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type flightSearchMongoRepo struct {
	mongoClient *mongo.Client
	dbName      string
	collection  string
}

func NewFlightSearchMongoRepo(mongoClient *mongo.Client, dbName, collection string) *flightSearchMongoRepo {
	return &flightSearchMongoRepo{
		mongoClient: mongoClient,
		dbName:      dbName,
		collection:  collection,
	}
}

// getCollection returns the MongoDB collection
func (repo *flightSearchMongoRepo) getCollection() *mongo.Collection {
	return repo.mongoClient.Database(repo.dbName).Collection(repo.collection)
}

// getNextFlightID generates the next auto-increment flight ID
func (repo *flightSearchMongoRepo) getNextFlightID(ctx context.Context) (int64, error) {
	countersCollection := repo.mongoClient.Database(repo.dbName).Collection("counters")

	// Upsert and increment the counter
	opts := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)
	filter := bson.M{"_id": "flightID"}
	update := bson.M{"$inc": bson.M{"sequence_value": 1}}

	var result struct {
		SequenceValue int64 `bson:"sequence_value"`
	}

	err := countersCollection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&result)
	if err != nil {
		return 0, fmt.Errorf("error generating flight ID: %v", err)
	}

	return result.SequenceValue, nil
}

// AddFlight adds a flight to MongoDB with auto-generated ID and returns the flight ID
func (repo *flightSearchMongoRepo) AddFlight(ctx context.Context, flight models.Flight) (flightId string, err error) {
	collection := repo.getCollection()

	// Generate auto-increment flight ID
	nextID, err := repo.getNextFlightID(ctx)
	if err != nil {
		return "", err
	}

	// Set the auto-generated ID with FL prefix
	flight.ID = fmt.Sprintf("FL%04d", nextID)

	// Check if flight already exists (just in case)
	filter := bson.M{"_id": flight.ID}
	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return "", fmt.Errorf("error checking if flight exists: %v", err)
	}

	if count > 0 {
		// Update existing flight
		update := bson.M{"$set": flight}
		_, err := collection.UpdateOne(ctx, filter, update)
		if err != nil {
			return "", fmt.Errorf("error updating flight: %v", err)
		}
		return flight.ID, nil
	}

	// Insert new flight with auto-generated ID
	_, err = collection.InsertOne(ctx, flight)
	if err != nil {
		return "", fmt.Errorf("error inserting flight: %v", err)
	}

	return flight.ID, nil
}

// GetFlights returns all flights (filtering logic moved to service layer)
func (repo *flightSearchMongoRepo) GetFlights(ctx context.Context, req models.FlightSearchRequest) ([]models.Flight, error) {
	return repo.GetAllFlights(ctx)
}

// GetAllFlights retrieves all flights from MongoDB
func (repo *flightSearchMongoRepo) GetAllFlights(ctx context.Context) ([]models.Flight, error) {
	collection := repo.getCollection()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("error fetching flights: %v", err)
	}
	defer cursor.Close(ctx)

	var flights []models.Flight
	err = cursor.All(ctx, &flights)
	if err != nil {
		return nil, fmt.Errorf("error decoding flights: %v", err)
	}

	return flights, nil
}

// RemoveFlight removes a flight from MongoDB by ID
func (repo *flightSearchMongoRepo) RemoveFlight(ctx context.Context, flightID string) error {
	if flightID == "" {
		return fmt.Errorf("flight ID cannot be empty")
	}

	collection := repo.getCollection()

	filter := bson.M{"_id": flightID}
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("error deleting flight: %v", err)
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("flight with ID %s not found", flightID)
	}

	return nil
}
