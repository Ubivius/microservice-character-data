package database

import (
	"context"
	"time"

	"github.com/Ubivius/microservice-character-data/pkg/data"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
)

type MockCharacters struct {
}

func NewMockCharacters() CharacterDB {
	log.Info("Connecting to mock database")
	return &MockCharacters{}
}

func (mp *MockCharacters) Connect() error {
	return nil
}

func (mp *MockCharacters) PingDB() error {
	return nil
}

func (mp *MockCharacters) CloseDB() {
	log.Info("Mocked DB connection closed")
}

func (mp *MockCharacters) GetCharacters(ctx context.Context) data.Characters {
	_, span := otel.Tracer("character-data").Start(ctx, "getCharactersDatabase")
	defer span.End()
	return characterList
}

func (mp *MockCharacters) GetCharacterByID(ctx context.Context, id string) (*data.Character, error) {
	_, span := otel.Tracer("character-data").Start(ctx, "getCharactersByIDDatabase")
	defer span.End()
	index := findIndexByCharacterID(id)
	if index == -1 {
		return nil, data.ErrorCharacterNotFound
	}
	return characterList[index], nil
}

func (mp *MockCharacters) GetCharactersByUserID(ctx context.Context, userID string) (data.Characters, error) {
	_, span := otel.Tracer("character-data").Start(ctx, "getCharactersByUserIdDatabase")
	defer span.End()
	charactersList := findCharactersListByUserID(userID)
	if len(charactersList) == 0 {
		return nil, data.ErrorCharacterNotFound
	}
	return charactersList, nil
}

func (mp *MockCharacters) GetCharactersAliveByUserID(ctx context.Context, userID string) (data.Characters, error) {
	_, span := otel.Tracer("character-data").Start(ctx, "getCharactersAliveByUserIdDatabase")
	defer span.End()
	charactersList := findCharactersAliveListByUserID(userID)
	if len(charactersList) == 0 {
		return nil, data.ErrorCharacterNotFound
	}
	return charactersList, nil
}

func (mp *MockCharacters) UpdateCharacter(ctx context.Context, character *data.Character) error {
	_, span := otel.Tracer("character-data").Start(ctx, "updateCharactersByIdDatabase")
	defer span.End()
	index := findIndexByCharacterID(character.ID)
	if index == -1 {
		return data.ErrorCharacterNotFound
	}
	characterList[index] = character
	return nil
}

func (mp *MockCharacters) AddCharacter(ctx context.Context, character *data.Character) error {
	_, span := otel.Tracer("character-data").Start(ctx, "addCharacterDatabase")
	defer span.End()
	if !mp.validateUserExist(character.UserID) {
		return data.ErrorUserNotFound
	}

	character.ID = uuid.NewString()
	characterList = append(characterList, character)
	return nil
}

func (mp *MockCharacters) DeleteCharacter(ctx context.Context, id string) error {
	_, span := otel.Tracer("character-data").Start(ctx, "deleteCharacterByIdDatabase")
	defer span.End()
	index := findIndexByCharacterID(id)
	if index == -1 {
		return data.ErrorCharacterNotFound
	}

	characterList = append(characterList[:index], characterList[index+1:]...)

	return nil
}

// Returns an array of characters in the database
// Returns -1 when no character is found
func findCharactersListByUserID(userID string) data.Characters {
	var charactersList data.Characters
	for _, character := range characterList {
		if character.UserID == userID {
			charactersList = append(charactersList, character)
		}
	}
	return charactersList
}

// Returns an array of characters in the database
// Returns -1 when no character is found
func findCharactersAliveListByUserID(userID string) data.Characters {
	var charactersList data.Characters
	for _ , character := range characterList {
		if character.UserID == userID && character.Alive {
			charactersList = append(charactersList, character)
		}
	}
	return charactersList
}

// Returns the index of a character in the database
// Returns -1 when no character is found
func findIndexByCharacterID(id string) int {
	for index, character := range characterList {
		if character.ID == id {
			return index
		}
	}
	return -1
}

func (mp *MockCharacters) validateUserExist(userID string) bool {
	return true
}

////////////////////////////////////////////////////////////////////////////////
/////////////////////////// Mocked database ///////////////////////////////////
//////////////////////////////////////////////////////////////////////////////

var characterList = []*data.Character{
	{
		ID:        "a2181017-5c53-422b-b6bc-036b27c04fc8",
		UserID:    "a2181017-5c53-422b-b6bc-036b27c04fc8",
		Name:      "ArcticWalrus",
		Alive:     true,
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
	{
		ID:        "e2382ea2-b5fa-4506-aa9d-d338aa52af44",
		UserID:    "e2382ea2-b5fa-4506-aa9d-d338aa52af44",
		Name:      "WinterSword",
		Alive:     true,
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
	{
		ID:        "aaaae510-956e-11eb-a8b3-0242ac130003",
		UserID:    "e2382ea2-b5fa-4506-aa9d-d338aa52af44",
		Name:      "ExistingCharacterName",
		Alive:     false,
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
}
