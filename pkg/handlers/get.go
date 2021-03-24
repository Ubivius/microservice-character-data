package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Ubivius/microservice-character-data/pkg/data"
)

// GET /characters
// Returns the full list of characters
func (characterHandler *CharactersHandler) GetCharacters(responseWriter http.ResponseWriter, request *http.Request) {
	characterHandler.logger.Println("Handle GET characters")
	characterList := data.GetCharacters()
	err := json.NewEncoder(responseWriter).Encode(characterList)
	if err != nil {
		characterHandler.logger.Println("[ERROR] serializing character", err)
		http.Error(responseWriter, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// GET /characters/{id}
// Returns a single character from the database
func (characterHandler *CharactersHandler) GetCharacterByID(responseWriter http.ResponseWriter, request *http.Request) {
	id := getCharacterId(request)

	characterHandler.logger.Println("[DEBUG] getting id", id)

	character, err := data.GetCharacterByID(id)
	switch err {
	case nil:
		err = json.NewEncoder(responseWriter).Encode(character)
		if err != nil {
			characterHandler.logger.Println("[ERROR] serializing character", err)
		}
	case data.ErrorCharacterNotFound:
		characterHandler.logger.Println("[ERROR] fetching character", err)
		http.Error(responseWriter, "Character not found", http.StatusBadRequest)
		return
	default:
		characterHandler.logger.Println("[ERROR] fetching character", err)
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
}
