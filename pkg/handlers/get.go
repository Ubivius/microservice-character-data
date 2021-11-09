package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Ubivius/microservice-character-data/pkg/data"
	"go.opentelemetry.io/otel"
)

// GET /characters
// Returns the full list of characters
func (characterHandler *CharactersHandler) GetCharacters(responseWriter http.ResponseWriter, request *http.Request) {
	_, span := otel.Tracer("character-data").Start(request.Context(), "getAllCharacters")
	defer span.End()
	log.Info("GetCharacters request")
	characterList := characterHandler.db.GetCharacters(request.Context())
	err := json.NewEncoder(responseWriter).Encode(characterList)
	if err != nil {
		log.Error(err, "Error serializing character")
		http.Error(responseWriter, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// GET /characters/{id}
// Returns a single character from the database
func (characterHandler *CharactersHandler) GetCharacterByID(responseWriter http.ResponseWriter, request *http.Request) {
	_, span := otel.Tracer("character-data").Start(request.Context(), "getCharacterById")
	defer span.End()
	id := getCharacterID(request)

	log.Info("GetCharacterByID request", "id", id)

	character, err := characterHandler.db.GetCharacterByID(request.Context(), id)
	switch err {
	case nil:
		err = json.NewEncoder(responseWriter).Encode(character)
		if err != nil {
			log.Error(err, "Error serializing character")
		}
		return
	case data.ErrorCharacterNotFound:
		log.Error(err, "Character not found")
		http.Error(responseWriter, "Character not found", http.StatusNotFound)
		return
	default:
		log.Error(err, "Error getting characters")
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
}

// GET /characters/user/{user_id}
// Returns an array of characters from the database
func (characterHandler *CharactersHandler) GetCharactersByUserID(responseWriter http.ResponseWriter, request *http.Request) {
	_, span := otel.Tracer("character-data").Start(request.Context(), "getCharactersByUserId")
	defer span.End()
	user_id := getUserID(request)

	log.Info("GetCharactersByUserID request", "user_id", user_id)

	characters, err := characterHandler.db.GetCharactersByUserID(request.Context(), user_id)
	switch err {
	case nil:
		err = json.NewEncoder(responseWriter).Encode(characters)
		if err != nil {
			log.Error(err, "Error serializing characters")
		}
		return
	case data.ErrorCharacterNotFound:
		log.Error(err, "Characters not found")
		http.Error(responseWriter, "Characters not found", http.StatusNotFound)
		return
	default:
		log.Error(err, "Error getting characters")
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
}

// GET /characters/alive/user/{user_id}
// Returns an array of characters alive for a user from the database
func (characterHandler *CharactersHandler) GetCharactersAliveByUserID(responseWriter http.ResponseWriter, request *http.Request) {
	user_id := getUserID(request)

	log.Info("GetCharactersAliveByUserID request", "user_id", user_id)

	characters, err := characterHandler.db.GetCharactersAliveByUserID(request.Context(), user_id)
	switch err {
	case nil:
		err = json.NewEncoder(responseWriter).Encode(characters)
		if err != nil {
			log.Error(err, "Error serializing characters")
		}
		return
	case data.ErrorCharacterNotFound:
		log.Error(err, "Characters not found")
		http.Error(responseWriter, "Characters not found", http.StatusNotFound)
		return
	default:
		log.Error(err, "Error getting characters")
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
}
