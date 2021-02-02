package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-character-data/data"
)

// /POST /characters
// Creates a new character
func (characterHandler *CharactersHandler) AddCharacter(responseWriter http.ResponseWriter, request *http.Request) {
	characterHandler.logger.Println("Handle POST Character")
	character := request.Context().Value(KeyCharacter{}).(*data.Character)

	data.AddCharacter(character)
}
