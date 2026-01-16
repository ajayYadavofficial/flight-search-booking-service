package flightbookingrepo

import (
	"context"
	"flight-search-booking-service/models"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	GetBookingDetails(ctx context.Context, bookingID string) (*models.Booking, error)
	BookFlight(ctx context.Context, bookingReq models.BookingRequest) (bookingId string, isSuccessful bool, err error)
}

type bookingMongoRepo struct {
	mongoClient *mongo.Client
	dbName      string
	collection  string
}

func NewFlightBookingMongoRepo(mongoClient *mongo.Client, dbName, collection string) *bookingMongoRepo {
	return &bookingMongoRepo{
		mongoClient: mongoClient,
		dbName:      dbName,
		collection:  collection,
	}
}

// getCollection returns the MongoDB collection
func (repo *bookingMongoRepo) getCollection() *mongo.Collection {
	return repo.mongoClient.Database(repo.dbName).Collection(repo.collection)
}

// getCountersCollection returns the counters collection for generating IDs
func (repo *bookingMongoRepo) getCountersCollection() *mongo.Collection {
	return repo.mongoClient.Database(repo.dbName).Collection("counters")
}

// getNextBookingID generates the next booking ID
func (repo *bookingMongoRepo) getNextBookingID(ctx context.Context) (string, error) {
	countersCollection := repo.getCountersCollection()

	opts := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)
	filter := bson.M{"_id": "bookingID"}
	update := bson.M{"$inc": bson.M{"sequence_value": 1}}

	var result bson.M
	err := countersCollection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&result)
	if err != nil {
		return "", fmt.Errorf("failed to generate booking ID: %v", err)
	}

	// Handle type conversion for sequence_value (could be int32, int64, or float64)
	var sequenceValue int64
	switch v := result["sequence_value"].(type) {
	case int32:
		sequenceValue = int64(v)
	case int64:
		sequenceValue = v
	case float64:
		sequenceValue = int64(v)
	default:
		return "", fmt.Errorf("unexpected type for sequence_value: %T", v)
	}

	bookingID := fmt.Sprintf("BK%04d", sequenceValue)
	return bookingID, nil
}

// BookFlight books a flight and returns the booking ID
func (repo *bookingMongoRepo) BookFlight(ctx context.Context, bookingReq models.BookingRequest) (bookingId string, isSuccessful bool, err error) {
	collection := repo.getCollection()

	// Generate booking ID
	bookingID, err := repo.getNextBookingID(ctx)
	if err != nil {
		return "", false, err
	}

	// Create booking object
	booking := models.Booking{
		ID:          bookingID,
		FlightID:    bookingReq.FlightID,
		SeatsBooked: bookingReq.Seats,
		UserInfo:    bookingReq.UserInfo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// Insert booking into MongoDB
	_, err = collection.InsertOne(ctx, booking)
	if err != nil {
		return "", false, fmt.Errorf("failed to insert booking: %v", err)
	}

	return bookingID, true, nil
}

// GetBookingDetails retrieves booking details by booking ID
func (repo *bookingMongoRepo) GetBookingDetails(ctx context.Context, bookingID string) (*models.Booking, error) {
	collection := repo.getCollection()

	filter := bson.M{"_id": bookingID}
	var booking models.Booking

	err := collection.FindOne(ctx, filter).Decode(&booking)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("booking not found: %s", bookingID)
		}
		return nil, fmt.Errorf("failed to retrieve booking: %v", err)
	}

	return &booking, nil
}
