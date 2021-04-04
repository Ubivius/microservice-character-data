package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-character-data/pkg/data"
)

// DELETE /characters/{id}
// Deletes a character with specified id from the database
func (characterHandler *CharactersHandler) Delete(responseWriter http.ResponseWriter, request *http.Request) {
	id := getCharacterID(request)
	characterHandler.logger.Println("Handle DELETE character", id)

	err := characterHandler.db.DeleteCharacter(id)
	switch err {
	case nil:
		responseWriter.WriteHeader(http.StatusNoContent)
		return
	case data.ErrorCharacterNotFound:
		characterHandler.logger.Println("[ERROR] deleting, id does not exist")
		http.Error(responseWriter, "Character not found", http.StatusNotFound)
		return
	default:
		characterHandler.logger.Println("[ERROR] deleting character", err)
		http.Error(responseWriter, "Error deleting character", http.StatusInternalServerError)
		return
	}	
}
