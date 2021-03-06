package database

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/Ubivius/microservice-character-data/pkg/data"
	"github.com/Ubivius/microservice-character-data/pkg/resources"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoCharacters struct {
	client           *mongo.Client
	collection       *mongo.Collection
	resourceManager  resources.ResourceManager
}

func NewMongoCharacters(r resources.ResourceManager) CharacterDB {
	mp := &MongoCharacters{resourceManager: r}
	err := mp.Connect()
	// If connect fails, kill the program
	if err != nil {
		log.Error(err, "MongoDB setup failed")
		os.Exit(1)
	}
	return mp
}

func (mp *MongoCharacters) Connect() error {
	// Getting mongodb secret
	password, err := mp.resourceManager.GetSecret("default", "mongodb", "mongodb-root-password")
	if err != nil {
		log.Error(err, "Failed to get mongodb secret")
		os.Exit(1)
	}
	
	// Setting client options
	clientOptions := options.Client().ApplyURI("mongodb://root:" + password + "@mongodb:27017/?authSource=admin")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil || client == nil {
		log.Error(err, "Failed to connect to database. Shutting down service")
		os.Exit(1)
	}

	// Ping DB
	err = client.Ping(context.TODO(), nil)
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
	return mp.client.Ping(context.TODO(), nil)
}

func (mp *MongoCharacters) CloseDB() {
	err := mp.client.Disconnect(context.TODO())
	if err != nil {
		log.Error(err, "Error while disconnecting from database")
	}
}

func (mp *MongoCharacters) GetCharacters() data.Characters {
	// characters will hold the array of Characters
	var characters data.Characters

	// Find returns a cursor that must be iterated through
	cursor, err := mp.collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Error(err, "Error getting characters from database")
	}

	// Iterating through cursor
	for cursor.Next(context.TODO()) {
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
	cursor.Close(context.TODO())

	return characters
}

func (mp *MongoCharacters) GetCharacterByID(id string) (*data.Character, error) {
	// MongoDB search filter
	filter := bson.D{{Key: "_id", Value: id}}

	// Holds search result
	var result data.Character

	// Find a single matching item from the database
	err := mp.collection.FindOne(context.TODO(), filter).Decode(&result)

	// Parse result into the returned character
	return &result, err
}

func (mp *MongoCharacters) GetCharactersByUserID(userID string) (data.Characters, error) {
	// MongoDB search filter
	filter := bson.D{{Key: "user_id", Value: userID}}

	// characters will hold the array of Messages
	var characters data.Characters

	// Find returns a cursor that must be iterated through
	cursor, err := mp.collection.Find(context.TODO(), filter)
	if err != nil {
		log.Error(err, "Error getting characters by userID from database")
	}

	// Iterating through cursor
	for cursor.Next(context.TODO()) {
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
	cursor.Close(context.TODO())

	return characters, err
}

func (mp *MongoCharacters) UpdateCharacter(character *data.Character) error {
	// Set updated timestamp in character
	character.UpdatedOn = time.Now().UTC().String()

	// MongoDB search filter
	filter := bson.D{{Key: "_id", Value: character.ID}}

	// Update sets the matched characters in the database to character
	update := bson.M{"$set": character}

	// Update a single item in the database with the values in update that match the filter
	_, err := mp.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Error(err, "Error updating character.")
	}

	return err
}

func (mp *MongoCharacters) AddCharacter(character *data.Character) error {
	if !mp.validateUserExist(character.UserID){
		return data.ErrorUserNotFound
	}

	character.ID = uuid.NewString()
	// Adding time information to new character
	character.CreatedOn = time.Now().UTC().String()
	character.UpdatedOn = time.Now().UTC().String()

	// Inserting the new character into the database
	insertResult, err := mp.collection.InsertOne(context.TODO(), character)
	if err != nil {
		return err
	}

	log.Info("Inserting character", "Inserted ID", insertResult.InsertedID)
	return nil
}

func (mp *MongoCharacters) DeleteCharacter(id string) error {
	// MongoDB search filter
	filter := bson.D{{Key: "_id", Value: id}}

	// Delete a single item matching the filter
	result, err := mp.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Error(err, "Error deleting character")
	}

	log.Info("Deleted documents in achievements collection", "delete_count", result.DeletedCount)
	return nil
}

func (mp *MongoCharacters) validateUserExist(userID string) bool {
	getUserByIDPath := data.MicroserviceUserPath + "/users/" + userID
	resp, err := http.Get(getUserByIDPath)
	return err == nil && resp.StatusCode == 200
}
