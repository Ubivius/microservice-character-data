package router

import (
	"log"
	"net/http"

	"github.com/Ubivius/microservice-character-data/pkg/handlers"
	"github.com/gorilla/mux"
)

// Mux route handling with gorilla/mux
func New(charactersHandler *handlers.CharactersHandler, logger *log.Logger) *mux.Router {
	router := mux.NewRouter()

	// Get Router
	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/characters", charactersHandler.GetCharacters)
	getRouter.HandleFunc("/characters/{id:[0-9a-z-]+}", charactersHandler.GetCharacterByID)
	getRouter.HandleFunc("/characters/user/{user_id:[0-9a-z-]+}", charactersHandler.GetCharactersByUserID)

	// Put router
	putRouter := router.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/characters", charactersHandler.UpdateCharacters)
	putRouter.Use(charactersHandler.MiddlewareCharacterValidation)

	// Post router
	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/characters", charactersHandler.AddCharacter)
	postRouter.Use(charactersHandler.MiddlewareCharacterValidation)

	// Delete router
	deleteRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/characters/{id:[0-9a-z-]+}", charactersHandler.Delete)

	return router
}
