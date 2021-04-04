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
	characterList := characterHandler.db.GetCharacters()
	err := json.NewEncoder(responseWriter).Encode(characterList)
	if err != nil {
		characterHandler.logger.Println("[ERROR] serializing character", err)
		http.Error(responseWriter, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// GET /characters/{id}
// Returns a single character from the database
func (characterHandler *CharactersHandler) GetCharacterByID(responseWriter http.ResponseWriter, request *http.Request) {
	id := getCharacterID(request)

	characterHandler.logger.Println("[DEBUG] getting id", id)

	character, err := characterHandler.db.GetCharacterByID(id)
	switch err {
	case nil:
		err = json.NewEncoder(responseWriter).Encode(character)
		if err != nil {
			characterHandler.logger.Println("[ERROR] serializing character", err)
		}
		return
	case data.ErrorCharacterNotFound:
		characterHandler.logger.Println("[ERROR] fetching character", err)
		http.Error(responseWriter, "Character not found", http.StatusNotFound)
		return
	default:
		characterHandler.logger.Println("[ERROR] fetching character", err)
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
}

// GET /characters/user/{user_id}
// Returns an array of characters from the database
func (characterHandler *CharactersHandler) GetCharactersByUserID(responseWriter http.ResponseWriter, request *http.Request) {
	user_id := getUserID(request)

	characterHandler.logger.Println("[DEBUG] getting id", user_id)

	characters, err := characterHandler.db.GetCharactersByUserID(user_id)
	switch err {
	case nil:
		err = json.NewEncoder(responseWriter).Encode(characters)
		if err != nil {
			characterHandler.logger.Println("[ERROR] serializing characters", err)
		}
		return
	case data.ErrorCharacterNotFound:
		characterHandler.logger.Println("[ERROR] fetching characters", err)
		http.Error(responseWriter, "Characters not found", http.StatusNotFound)
		return
	default:
		characterHandler.logger.Println("[ERROR] fetching characters", err)
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
}
