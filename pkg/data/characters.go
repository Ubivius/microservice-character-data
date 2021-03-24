package data

import (
	"fmt"
	"time"
)

// Character specific errors
var ErrorCharacterNotFound = fmt.Errorf("Character not found")
var ErrorCharacterNameAlreadyExist = fmt.Errorf("Character name is already used")

// Character defines the structure for an API character.
// Formatting done with json tags to the right. "-" : don't include when encoding to json
type Character struct {
	ID        int    `json:"id"`
	UserID    int    `json:"userid" validate:"required"`
	Name      string `json:"name" validate:"required,name"`
	CreatedOn string `json:"-"`
	UpdatedOn string `json:"-"`
	DeletedOn string `json:"-"`
}

// Characters is a collection of Character
type Characters []*Character

// All of these functions will become database calls in the future
// GETTING CHARACTERS

// Returns the list of characters
func GetCharacters() Characters {
	return characterList
}

// Returns a single character with the given id
func GetCharacterByID(id int) (*Character, error) {
	index := FindIndexByCharacterID(id)
	if index == -1 {
		return nil, ErrorCharacterNotFound
	}
	return characterList[index], nil
}

// UPDATING CHARACTERS

func UpdateCharacter(character *Character) error {
	index := FindIndexByCharacterID(character.ID)
	if index == -1 {
		return ErrorCharacterNotFound
	}
	characterList[index] = character
	return nil
}

// ADD A CHARACTER
func AddCharacter(character *Character) error {
	err := CharacterNameExist(character.Name)
	if err != nil {
		return err
	}

	// TODO: Verify that the user exist
	character.ID = getNextId()
	characterList = append(characterList, character)
	return nil
}

// DELETING A CHARACTER
func DeleteCharacter(id int) error {
	index := FindIndexByCharacterID(id)
	if index == -1 {
		return ErrorCharacterNotFound
	}

	characterList = append(characterList[:index], characterList[index+1:]...)

	return nil
}

// Returns the index of a character in the database
// Returns -1 when no character is found
func FindIndexByCharacterID(id int) int {
	for index, character := range characterList {
		if character.ID == id {
			return index
		}
	}
	return -1
}

func CharacterNameExist(name string) error {
	for _, character := range characterList {
		if character.Name == name {
			return ErrorCharacterNameAlreadyExist
		}
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////
/////////////////////////// Fake database ///////////////////////////////////
///// DB connection setup and docker file will be done in sprint 8 /////////
///////////////////////////////////////////////////////////////////////////

// Finds the maximum index of our fake database and adds 1
func getNextId() int {
	lastCharacter := characterList[len(characterList)-1]
	return lastCharacter.ID + 1
}

// characterList is a hard coded list of characters for this
// example data source. Should be replaced by database connection
var characterList = []*Character{
	{
		ID:        1,
		UserID:    1,
		Name:      "ArcticWalrus",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
	{
		ID:        2,
		UserID:    2,
		Name:      "WinterSword",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
	{
		ID:        3,
		UserID:    2,
		Name:      "ExistingCharacterName",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
}
