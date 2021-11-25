package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-character-data/pkg/data"
	"go.opentelemetry.io/otel"
)

// /POST /characters
// Creates a new character
func (characterHandler *CharactersHandler) AddCharacter(responseWriter http.ResponseWriter, request *http.Request) {
	_, span := otel.Tracer("character-data").Start(request.Context(), "addCharacter")
	defer span.End()
	log.Info("AddCharacter request")
	character := request.Context().Value(KeyCharacter{}).(*data.Character)

	err := characterHandler.db.AddCharacter(request.Context(), character)
	switch err {
	case nil:
		responseWriter.WriteHeader(http.StatusNoContent)
		return
	case data.ErrorUserNotFound:
		log.Error(err, "UserID doesn't exist")
		http.Error(responseWriter, "UserID doesn't exist", http.StatusBadRequest)
		return
	default:
		log.Error(err, "Error adding character")
		http.Error(responseWriter, "Error adding character", http.StatusInternalServerError)
		return
	}
}
