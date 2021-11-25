package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-character-data/pkg/data"
	"go.opentelemetry.io/otel"
)

// DELETE /characters/{id}
// Deletes a character with specified id from the database
func (characterHandler *CharactersHandler) Delete(responseWriter http.ResponseWriter, request *http.Request) {
	_, span := otel.Tracer("character-data").Start(request.Context(), "deleteCharacterById")
	defer span.End()
	id := getCharacterID(request)
	log.Info("Delete character by ID request", "id", id)

	err := characterHandler.db.DeleteCharacter(request.Context(), id)
	switch err {
	case nil:
		responseWriter.WriteHeader(http.StatusNoContent)
		return
	case data.ErrorCharacterNotFound:
		log.Error(err, "Error deleting character, id does not exist")
		http.Error(responseWriter, "Character not found", http.StatusNotFound)
		return
	default:
		log.Error(err, "Error deleting character")
		http.Error(responseWriter, "Error deleting character", http.StatusInternalServerError)
		return
	}
}
