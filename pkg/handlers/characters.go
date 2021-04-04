package handlers

import (
	"log"
	"net/http"

	"github.com/Ubivius/microservice-character-data/pkg/database"
	"github.com/gorilla/mux"
)

// KeyCharacter is a key used for the Character object inside context
type KeyCharacter struct{}

// Charactershandler used for getting and updating characters
type CharactersHandler struct {
	logger *log.Logger
	db     database.CharacterDB
}

func NewCharactersHandler(logger *log.Logger, db database.CharacterDB) *CharactersHandler {
	return &CharactersHandler{logger, db}
}

// getCharacterID extracts the character ID from the URL
// The verification of this variable is handled by gorilla/mux
func getCharacterID(request *http.Request) string {
	vars := mux.Vars(request)
	id := vars["id"]
	return id
}

// getUserID extracts the user ID from the URL
// The verification of this variable is handled by gorilla/mux
func getUserID(request *http.Request) string {
	vars := mux.Vars(request)
	id := vars["user_id"]
	return id
}
