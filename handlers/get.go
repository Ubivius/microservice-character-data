package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-character-data/data"
)

// GET /characters
// Returns the full list of characters
func (characterHandler *CharactersHandler) GetCharacters(responseWriter http.ResponseWriter, request *http.Request) {
	characterHandler.logger.Println("Handle GET characters")
	characterList := data.GetCharacters()
	err := data.ToJSON(characterList, responseWriter)
	if err != nil {
		characterHandler.logger.Println("[ERROR] serializing character", err)
		http.Error(responseWriter, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// GET /characters/{id}
// Returns a single character from the database
func (characterHandler *CharactersHandler) GetCharacterById(responseWriter http.ResponseWriter, request *http.Request) {
	id := getCharacterId(request)

	characterHandler.logger.Println("[DEBUG] getting id", id)

	character, err := data.GetCharacterById(id)
	switch err {
	case nil:
	case data.ErrorCharacterNotFound:
		characterHandler.logger.Println("[ERROR] fetching character", err)
		http.Error(responseWriter, "Character not found", http.StatusBadRequest)
		return
	default:
		characterHandler.logger.Println("[ERROR] fetching character", err)
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	err = data.ToJSON(character, responseWriter)
	if err != nil {
		characterHandler.logger.Println("[ERROR] serializing character", err)
	}
}
