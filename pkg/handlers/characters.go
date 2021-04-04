package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// KeyCharacter is a key used for the Character object inside context
type KeyCharacter struct{}

// Charactershandler used for getting and updating characters
type CharactersHandler struct {
	logger *log.Logger
}

func NewCharactersHandler(logger *log.Logger) *CharactersHandler {
	return &CharactersHandler{logger}
}

// getCharacterId extracts the character ID from the URL
// The verification of this variable is handled by gorilla/mux
// We panic if it is not valid because that means gorilla is failing
func getCharacterId(request *http.Request) int {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	return id
}
