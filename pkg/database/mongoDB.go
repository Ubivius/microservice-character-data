package database

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Ubivius/microservice-character-data/pkg/data"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"
)

// ErrorEnvVar : Environment variable error
var ErrorEnvVar = fmt.Errorf("missing environment variable")

type MongoCharacters struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoCharacters() CharacterDB {
	mp := &MongoCharacters{}
	err := mp.Connect()
	// If connect fails, kill the program
	if err != nil {
		log.Error(err, "MongoDB setup failed")
		os.Exit(1)
	}
	return mp
}

func (mp *MongoCharacters) Connect() error {
	uri := mongodbURI()

	// Setting client options
	opts := options.Client()
	clientOptions := opts.ApplyURI(uri)
	opts.Monitor = otelmongo.NewMonitor()

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil || client == nil {
		log.Error(err, "Failed to connect to database. Shutting down service")
		os.Exit(1)
	}

	// Ping DB
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Error(err, "Failed to ping database. Shutting down service")
		os.Exit(1)
	}

	log.Info("Connection to MongoDB established")

	collection := client.Database("ubivius").Collection("characters")

	// Assign client and collection to the MongoCharacters struct
	mp.collection = collection
	mp.client = client
	return nil
}

func (mp *MongoCharacters) PingDB() error {
	return mp.client.Ping(context.Background(), nil)
}

func (mp *MongoCharacters) CloseDB() {
	err := mp.client.Disconnect(context.Background())
	if err != nil {
		log.Error(err, "Error while disconnecting from database")
	}
}

func (mp *MongoCharacters) GetCharacters(ctx context.Context) data.Characters {
	// characters will hold the array of Characters
	var characters data.Characters

	// Find returns a cursor that must be iterated through
	cursor, err := mp.collection.Find(ctx, bson.D{})
	if err != nil {
		log.Error(err, "Error getting characters from database")
	}

	// Iterating through cursor
	for cursor.Next(ctx) {
		var result data.Character
		err := cursor.Decode(&result)
		if err != nil {
			log.Error(err, "Error decoding character from database")
		}
		characters = append(characters, &result)
	}

	if err := cursor.Err(); err != nil {
		log.Error(err, "Error in cursor after iteration")
	}

	// Close the cursor once finished
	cursor.Close(ctx)

	return characters
}

func (mp *MongoCharacters) GetCharacterByID(ctx context.Context, id string) (*data.Character, error) {
	// MongoDB search filter
	filter := bson.D{{Key: "_id", Value: id}}

	// Holds search result
	var result data.Character

	// Find a single matching item from the database
	err := mp.collection.FindOne(ctx, filter).Decode(&result)

	// Parse result into the returned character
	return &result, err
}

func (mp *MongoCharacters) GetCharactersByUserID(ctx context.Context, userID string) (data.Characters, error) {
	// MongoDB search filter
	filter := bson.D{{Key: "user_id", Value: userID}}

	// characters will hold the array of Characters
	var characters data.Characters

	// Find returns a cursor that must be iterated through
	cursor, err := mp.collection.Find(ctx, filter)
	if err != nil {
		log.Error(err, "Error getting characters by userID from database")
	}

	// Iterating through cursor
	for cursor.Next(ctx) {
		var result data.Character
		err := cursor.Decode(&result)
		if err != nil {
			log.Error(err, "Error decoding characters from database")
		}
		characters = append(characters, &result)
	}

	if err := cursor.Err(); err != nil {
		log.Error(err, "Error in cursor after iteration")
	}

	// Close the cursor once finished
	cursor.Close(ctx)

	return characters, err
}

func (mp *MongoCharacters) GetAliveCharactersByUserID(ctx context.Context, userID string) (data.Characters, error) {
	// MongoDB search filter
	filter := bson.D{{Key: "user_id", Value: userID}, {Key: "alive", Value: true}}

	// characters will hold the array of Characters
	var characters data.Characters

	// Find returns a cursor that must be iterated through
	cursor, err := mp.collection.Find(ctx, filter)
	if err != nil {
		log.Error(err, "Error getting characters by userID from database")
	}

	// Iterating through cursor
	for cursor.Next(ctx) {
		var result data.Character
		err := cursor.Decode(&result)
		if err != nil {
			log.Error(err, "Error decoding characters from database")
		}
		characters = append(characters, &result)
	}

	if err := cursor.Err(); err != nil {
		log.Error(err, "Error in cursor after iteration")
	}

	// Close the cursor once finished
	cursor.Close(ctx)

	return characters, err
}

func (mp *MongoCharacters) UpdateCharacter(ctx context.Context, character *data.Character) error {
	// Set updated timestamp in character
	character.UpdatedOn = time.Now().UTC().String()

	// MongoDB search filter
	filter := bson.D{{Key: "_id", Value: character.ID}}

	// Update sets the matched characters in the database to character
	update := bson.M{"$set": character}

	// Update a single item in the database with the values in update that match the filter
	updateResult, err := mp.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Error(err, "Error updating character.")
	}
	if updateResult.MatchedCount != 1 {
		log.Error(data.ErrorCharacterNotFound, "No matches found for update")
		return err
	}

	return err
}

func (mp *MongoCharacters) AddCharacter(ctx context.Context, character *data.Character) error {
	if !mp.validateUserExist(character.UserID) {
		return data.ErrorUserNotFound
	}

	character.ID = uuid.NewString()
	character.Alive = true;
	// Adding time information to new character
	character.CreatedOn = time.Now().UTC().String()
	character.UpdatedOn = time.Now().UTC().String()

	// Inserting the new character into the database
	insertResult, err := mp.collection.InsertOne(ctx, character)
	if err != nil {
		return err
	}

	log.Info("Inserting character", "Inserted ID", insertResult.InsertedID)
	return nil
}

func (mp *MongoCharacters) DeleteCharacter(ctx context.Context, id string) error {
	// MongoDB search filter
	filter := bson.D{{Key: "_id", Value: id}}

	// Delete a single item matching the filter
	result, err := mp.collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Error(err, "Error deleting character")
	}

	log.Info("Deleted documents in characters collection", "delete_count", result.DeletedCount)
	return nil
}

func (mp *MongoCharacters) validateUserExist(userID string) bool {
	getUserByIDPath := data.MicroserviceUserPath + "/users/" + userID
	resp, err := http.Get(getUserByIDPath)
	return err == nil && resp.StatusCode == 200
}

func deleteAllCharactersFromMongoDB() error {
	uri := mongodbURI()

	// Setting client options
	opts := options.Client()
	clientOptions := opts.ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil || client == nil {
		log.Error(err, "Failed to connect to database. Failing test")
		return err
	}
	collection := client.Database("ubivius").Collection("characters")
	_, err = collection.DeleteMany(context.Background(), bson.D{{}})
	return err
}

func mongodbURI() string {
	hostname := os.Getenv("DB_HOSTNAME")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")

	if hostname == "" || port == "" || username == "" || password == "" {
		log.Error(ErrorEnvVar, "Some environment variables are not available for the DB connection. DB_HOSTNAME, DB_PORT, DB_USERNAME, DB_PASSWORD")
		os.Exit(1)
	}

	return "mongodb://" + username + ":" + password + "@" + hostname + ":" + port + "/?authSource=admin"
}
