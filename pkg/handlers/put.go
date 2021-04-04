package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-character-data/pkg/data"
)

func (characterHandler *CharactersHandler) UpdateCharacters(responseWriter http.ResponseWriter, request *http.Request) {
	character := request.Context().Value(KeyCharacter{}).(*data.Character)
	characterHandler.logger.Println("Handle PUT character", character.ID)

	// Update character
	err := characterHandler.db.UpdateCharacter(character)
	switch err {
	case nil:
		responseWriter.WriteHeader(http.StatusNoContent)
		return
	case data.ErrorCharacterNotFound:
		characterHandler.logger.Println("[ERROR} character not found", err)
		http.Error(responseWriter, "Character not found", http.StatusNotFound)
		return
	default:
		characterHandler.logger.Println("[ERROR] updating character", err)
		http.Error(responseWriter, "Error updating character", http.StatusInternalServerError)
		return
	}
}
