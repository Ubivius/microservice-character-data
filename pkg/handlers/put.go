package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-character-data/pkg/data"
)

func (characterHandler *CharactersHandler) UpdateCharacters(responseWriter http.ResponseWriter, request *http.Request) {
	character := request.Context().Value(KeyCharacter{}).(*data.Character)
	log.Info("UpdateCharacters request", "id", character.ID)

	// Update character
	err := characterHandler.db.UpdateCharacter(character)
	switch err {
	case nil:
		responseWriter.WriteHeader(http.StatusNoContent)
		return
	case data.ErrorCharacterNotFound:
		log.Error(err, "Character not found")
		http.Error(responseWriter, "Character not found", http.StatusNotFound)
		return
	default:
		log.Error(err, "Error updating character")
		http.Error(responseWriter, "Error updating character", http.StatusInternalServerError)
		return
	}
}
