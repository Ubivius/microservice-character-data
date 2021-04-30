package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Ubivius/microservice-character-data/pkg/data"
)

// Json Character Validation
func (characterHandler *CharactersHandler) MiddlewareCharacterValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		character := &data.Character{}

		err := json.NewDecoder(request.Body).Decode(character)
		if err != nil {
			log.Error(err, "Error deserializing character")
			http.Error(responseWriter, "Error reading character", http.StatusBadRequest)
			return
		}

		// validate the character
		err = character.ValidateCharacter()
		if err != nil {
			log.Error(err, "Error validating character")
			http.Error(responseWriter, fmt.Sprintf("Error validating character: %s", err), http.StatusBadRequest)
			return
		}

		// Add the character to the context
		context := context.WithValue(request.Context(), KeyCharacter{}, character)
		newRequest := request.WithContext(context)

		// Call the next handler, which can be another middleware or the final handler
		next.ServeHTTP(responseWriter, newRequest)
	})
}
