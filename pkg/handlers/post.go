package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-character-data/pkg/data"
)

// /POST /characters
// Creates a new character
func (characterHandler *CharactersHandler) AddCharacter(responseWriter http.ResponseWriter, request *http.Request) {
	characterHandler.logger.Println("Handle POST Character")
	character := request.Context().Value(KeyCharacter{}).(*data.Character)

	err := data.AddCharacter(character)

	if err != nil {
		characterHandler.logger.Println("[ERROR] adding character", err)
		http.Error(responseWriter, "Error adding character", http.StatusInternalServerError)
		return
	}

	responseWriter.WriteHeader(http.StatusNoContent)
}
