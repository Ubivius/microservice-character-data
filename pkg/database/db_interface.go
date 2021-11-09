package database

import (
	"context"

	"github.com/Ubivius/microservice-character-data/pkg/data"
)

// The interface that any kind of database must implement
type CharacterDB interface {
	GetCharacters(ctx context.Context) data.Characters
	GetCharacterByID(ctx context.Context, id string) (*data.Character, error)
	GetCharactersByUserID(ctx context.Context, userID string) (data.Characters, error)
	GetCharactersAliveByUserID(ctx context.Context, userID string) (data.Characters, error)
	UpdateCharacter(ctx context.Context, character *data.Character) error
	AddCharacter(ctx context.Context, character *data.Character) error
	DeleteCharacter(ctx context.Context, id string) error
	validateUserExist(userID string) bool
	Connect() error
	PingDB() error
	CloseDB()
}
