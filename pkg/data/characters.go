package data

import (
	"fmt"
)

// Character specific errors
var ErrorCharacterNotFound = fmt.Errorf("Character not found")

// Character defines the structure for an API character.
type Character struct {
	ID        string `json:"id" bson:"_id"`
	UserID    string `json:"user_id" bson:"user_id" validate:"required"`
	Name      string `json:"name" validate:"required,name"`
	CreatedOn string `json:"-"`
	UpdatedOn string `json:"-"`
}

// Characters is a collection of Character
type Characters []*Character
