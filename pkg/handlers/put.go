package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-character-data/pkg/data"
)

func (characterHandler *CharactersHandler) UpdateCharacters(responseWriter http.ResponseWriter, request *http.Request) {
	character := request.Context().Value(KeyCharacter{}).(data.Character)
	characterHandler.logger.Println("Handle PUT character", character.ID)

	// Update character
	err := data.UpdateCharacter(&character)
	if err == data.ErrorCharacterNotFound {
		characterHandler.logger.Println("[ERROR} character not found", err)
		http.Error(responseWriter, "Character not found", http.StatusNotFound)
		return
	}

	// Returns status, no content required
	responseWriter.WriteHeader(http.StatusNoContent)
}
