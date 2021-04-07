package database

import (
	"github.com/Ubivius/microservice-character-data/pkg/data"
)

// The interface that any kind of database must implement
type CharacterDB interface {
	GetCharacters() data.Characters
	GetCharacterByID(id string) (*data.Character, error)
	GetCharactersByUserID(userID string) (data.Characters, error)
	UpdateCharacter(character *data.Character) error
	AddCharacter(character *data.Character) error
	DeleteCharacter(id string) error
	validateUserExist(userID string) bool
	Connect() error
	PingDB() error
	CloseDB()
}
