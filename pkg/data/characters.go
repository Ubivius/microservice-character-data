package data

import (
	"fmt"
)

// Character specific errors
var ErrorCharacterNotFound = fmt.Errorf("character not found")

// ErrorUserNotFound : User specific errors
var ErrorUserNotFound = fmt.Errorf("user not found")

// Character defines the structure for an API character.
type Character struct {
	ID             string `json:"id" bson:"_id"`
	UserID         string `json:"user_id" bson:"user_id" validate:"required"`
	Name           string `json:"name" validate:"required,name"`
	Alive          bool   `json:"alive"`
	GamesWon       int    `json:"games_won" bson:"games_won"`
	GamesPlayed    int    `json:"games_played" bson:"games_played"`
	EnemiesKilled  int    `json:"enemies_killed" bson:"enemies_killed"`
	CreatedOn      string `json:"created_on" bson:"created_on"`
	UpdatedOn      string `json:"updated_on" bson:"updated_on"`
}

// Characters is a collection of Character
type Characters []*Character

const MicroserviceUserPath = "http://microservice-user:9090"
