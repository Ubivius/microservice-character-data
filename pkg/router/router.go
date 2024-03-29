package router

import (
	"net/http"

	"github.com/Ubivius/microservice-character-data/pkg/handlers"
	"github.com/Ubivius/pkg-telemetry/metrics"
	tokenValidation "github.com/Ubivius/shared-authentication/pkg/auth"
	"github.com/gorilla/mux"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
)

// Mux route handling with gorilla/mux
func New(charactersHandler *handlers.CharactersHandler) *mux.Router {
	log.Info("Starting router")
	router := mux.NewRouter()
	router.Use(otelmux.Middleware("character-data"))
	router.Use(metrics.RequestCountMiddleware)

	// Get Router
	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.Use(tokenValidation.Middleware)
	getRouter.HandleFunc("/characters", charactersHandler.GetCharacters)
	getRouter.HandleFunc("/characters/{id:[0-9a-z-]+}", charactersHandler.GetCharacterByID)
	getRouter.HandleFunc("/characters/user/{user_id:[0-9a-z-]+}", charactersHandler.GetCharactersByUserID)
	getRouter.HandleFunc("/characters/alive/user/{user_id:[0-9a-z-]+}", charactersHandler.GetAliveCharactersByUserID)

	//Health Check
	healthRouter := router.Methods(http.MethodGet).Subrouter()
	healthRouter.HandleFunc("/health/live", charactersHandler.LivenessCheck)
	healthRouter.HandleFunc("/health/ready", charactersHandler.ReadinessCheck)

	// Put router
	putRouter := router.Methods(http.MethodPut).Subrouter()
	putRouter.Use(tokenValidation.Middleware)
	putRouter.HandleFunc("/characters", charactersHandler.UpdateCharacters)
	putRouter.Use(charactersHandler.MiddlewareCharacterValidation)

	// Post router
	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.Use(tokenValidation.Middleware)
	postRouter.HandleFunc("/characters", charactersHandler.AddCharacter)
	postRouter.Use(charactersHandler.MiddlewareCharacterValidation)

	// Delete router
	deleteRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.Use(tokenValidation.Middleware)
	deleteRouter.HandleFunc("/characters/{id:[0-9a-z-]+}", charactersHandler.Delete)

	return router
}
