package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Ubivius/microservice-character-data/data"
)

// Errors should be templated in the future.
// A good starting reference can be found here : https://github.com/nicholasjackson/building-microservices-youtube/blob/episode_7/character-api/handlers/middleware.go
// We want our validation errors to have a standard format

// Json Character Validation
func (characterHandler *CharactersHandler) MiddlewareCharacterValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		character := &data.Character{}

		err := data.FromJSON(character, request.Body)
		if err != nil {
			characterHandler.logger.Println("[ERROR] deserializing character", err)
			http.Error(responseWriter, "Error reading character", http.StatusBadRequest)
			return
		}

		// validate the character
		err = character.ValidateCharacter()
		if err != nil {
			characterHandler.logger.Println("[ERROR] validating character", err)
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
