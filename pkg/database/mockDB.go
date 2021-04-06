package database

import (
	"time"

	"github.com/Ubivius/microservice-character-data/pkg/data"
	"github.com/google/uuid"
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

func (mp *MockCharacters) CloseDB() {
	log.Info("Mocked DB connection closed")
}

func (mp *MockCharacters) GetCharacters() data.Characters {
	return characterList
}

func (mp *MockCharacters) GetCharacterByID(id string) (*data.Character, error) {
	index := findIndexByCharacterID(id)
	if index == -1 {
		return nil, data.ErrorCharacterNotFound
	}
	return characterList[index], nil
}

func (mp *MockCharacters) GetCharactersByUserID(userID string) (data.Characters, error) {
	charactersList := findCharactersListByUserID(userID)
	if len(charactersList) == 0 {
		return nil, data.ErrorCharacterNotFound
	}
	return charactersList, nil
}

func (mp *MockCharacters) UpdateCharacter(character *data.Character) error {
	index := findIndexByCharacterID(character.ID)
	if index == -1 {
		return data.ErrorCharacterNotFound
	}
	characterList[index] = character
	return nil
}

func (mp *MockCharacters) AddCharacter(character *data.Character) error {
	// TODO: Verify that the user exist
	character.ID = uuid.NewString()
	characterList = append(characterList, character)
	return nil
}

func (mp *MockCharacters) DeleteCharacter(id string) error {
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
	for _ , character := range characterList {
		if character.UserID == userID {
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

////////////////////////////////////////////////////////////////////////////////
/////////////////////////// Mocked database ///////////////////////////////////
//////////////////////////////////////////////////////////////////////////////

var characterList = []*data.Character{
	{
		ID:        "a2181017-5c53-422b-b6bc-036b27c04fc8",
		UserID:    "a2181017-5c53-422b-b6bc-036b27c04fc8",
		Name:      "ArcticWalrus",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
	{
		ID:        "e2382ea2-b5fa-4506-aa9d-d338aa52af44",
		UserID:    "e2382ea2-b5fa-4506-aa9d-d338aa52af44",
		Name:      "WinterSword",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
	{
		ID:        "aaaae510-956e-11eb-a8b3-0242ac130003",
		UserID:    "e2382ea2-b5fa-4506-aa9d-d338aa52af44",
		Name:      "ExistingCharacterName",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
}
