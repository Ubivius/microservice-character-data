package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Ubivius/microservice-character-data/pkg/data"
	"github.com/Ubivius/microservice-character-data/pkg/database"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func newCharacterDB() database.CharacterDB {
	return database.NewMockCharacters()
}

func TestGetCharacters(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/characters", nil)
	response := httptest.NewRecorder()

	characterHandler := NewCharactersHandler(newCharacterDB())
	characterHandler.GetCharacters(response, request)

	if response.Code != 200 {
		t.Errorf("Expected status code 200 but got : %d", response.Code)
	}

	if !strings.Contains(response.Body.String(), "a2181017-5c53-422b-b6bc-036b27c04fc8") || !strings.Contains(response.Body.String(), "e2382ea2-b5fa-4506-aa9d-d338aa52af44") {
		t.Error("Missing elements from expected results")
	}
}

func TestGetExistingCharacterByID(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/characters/a2181017-5c53-422b-b6bc-036b27c04fc8", nil)
	response := httptest.NewRecorder()

	characterHandler := NewCharactersHandler(newCharacterDB())

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"id": "a2181017-5c53-422b-b6bc-036b27c04fc8",
	}
	request = mux.SetURLVars(request, vars)

	characterHandler.GetCharacterByID(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got : %d", http.StatusOK, response.Code)
	}
	if !strings.Contains(response.Body.String(), "a2181017-5c53-422b-b6bc-036b27c04fc8") {
		t.Error("Missing elements from expected results")
	}
}

func TestGetNonExistingCharacterByID(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/characters/4", nil)
	response := httptest.NewRecorder()

	characterHandler := NewCharactersHandler(newCharacterDB())

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"id": uuid.NewString(),
	}
	request = mux.SetURLVars(request, vars)

	characterHandler.GetCharacterByID(response, request)

	if response.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d but got : %d", http.StatusNotFound, response.Code)
	}
	if !strings.Contains(response.Body.String(), "Character not found") {
		t.Error("Expected response : Character not found")
	}
}

func TestGetExistingCharactersByUserID(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/characters/user/a2181017-5c53-422b-b6bc-036b27c04fc8", nil)
	response := httptest.NewRecorder()

	characterHandler := NewCharactersHandler(newCharacterDB())

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"user_id": "a2181017-5c53-422b-b6bc-036b27c04fc8",
	}
	request = mux.SetURLVars(request, vars)

	characterHandler.GetCharactersByUserID(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got : %d", http.StatusOK, response.Code)
	}
	if !strings.Contains(response.Body.String(), "a2181017-5c53-422b-b6bc-036b27c04fc8") {
		t.Error("Missing elements from expected results")
	}
}

func TestGetNonExistingCharactersByUserID(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/characters/user/4", nil)
	response := httptest.NewRecorder()

	characterHandler := NewCharactersHandler(newCharacterDB())

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"user_id": uuid.NewString(),
	}
	request = mux.SetURLVars(request, vars)

	characterHandler.GetCharactersByUserID(response, request)

	if response.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d but got : %d", http.StatusNotFound, response.Code)
	}
	if !strings.Contains(response.Body.String(), "Characters not found") {
		t.Error("Expected response : Characters not found")
	}
}

func TestGetOneExistingAliveCharactersByUserID(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/characters/alive/user/e2382ea2-b5fa-4506-aa9d-d338aa52af44", nil)
	response := httptest.NewRecorder()

	characterHandler := NewCharactersHandler(newCharacterDB())

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"user_id": "e2382ea2-b5fa-4506-aa9d-d338aa52af44",
	}
	request = mux.SetURLVars(request, vars)

	characterHandler.GetAliveCharactersByUserID(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got : %d", http.StatusOK, response.Code)
	}
	if !strings.Contains(response.Body.String(), "e2382ea2-b5fa-4506-aa9d-d338aa52af44") {
		t.Error("Missing elements from expected results")
	}
	if strings.Contains(response.Body.String(), "aaaae510-956e-11eb-a8b3-0242ac130003") {
		t.Error("Unexpected element present in the results")
	}
}

func TestGetTwoExistingAliveCharactersByUserID(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/characters/alive/user/40a2708f-21e4-439c-a019-3dc339d5717d", nil)
	response := httptest.NewRecorder()

	characterHandler := NewCharactersHandler(newCharacterDB())

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"user_id": "40a2708f-21e4-439c-a019-3dc339d5717d",
	}
	request = mux.SetURLVars(request, vars)

	characterHandler.GetAliveCharactersByUserID(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got : %d", http.StatusOK, response.Code)
	}
	if !strings.Contains(response.Body.String(), "2a0d39d5-5004-45eb-9aea-32f917956ba4") && 
	!strings.Contains(response.Body.String(), "91abbbc5-87d6-4ec0-8ccc-67b19b02f831") {
		t.Error("Missing elements from expected results")
	}
}

func TestGetNonExistingAliveCharactersByUserID(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/characters/alive/user/a2181017-5c53-422b-b6bc-036b27c04fc8", nil)
	response := httptest.NewRecorder()

	characterHandler := NewCharactersHandler(newCharacterDB())

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"user_id": "a2181017-5c53-422b-b6bc-036b27c04fc8",
	}
	request = mux.SetURLVars(request, vars)

	characterHandler.GetAliveCharactersByUserID(response, request)

	if response.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d but got : %d", http.StatusNotFound, response.Code)
	}
	if !strings.Contains(response.Body.String(), "Characters not found") {
		t.Error("Expected response : Characters not found")
	}
}

func TestDeleteNonExistantCharacter(t *testing.T) {
	request := httptest.NewRequest(http.MethodDelete, "/characters/4", nil)
	response := httptest.NewRecorder()

	characterHandler := NewCharactersHandler(newCharacterDB())

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"id": uuid.NewString(),
	}
	request = mux.SetURLVars(request, vars)

	characterHandler.Delete(response, request)
	if response.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d but got : %d", http.StatusNotFound, response.Code)
	}
	if !strings.Contains(response.Body.String(), "Character not found") {
		t.Error("Expected response : Character not found")
	}
}

func TestAddCharacter(t *testing.T) {
	// Creating request body
	body := &data.Character{
		Name:   "addName",
		UserID: "a2181017-5c53-422b-b6bc-036b27c04fc8",
	}

	request := httptest.NewRequest(http.MethodPost, "/characters", nil)
	response := httptest.NewRecorder()

	// Add the body to the context since we arent passing through middleware
	ctx := context.WithValue(request.Context(), KeyCharacter{}, body)
	request = request.WithContext(ctx)

	characterHandler := NewCharactersHandler(newCharacterDB())
	characterHandler.AddCharacter(response, request)

	if response.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d, but got %d", http.StatusNoContent, response.Code)
	}
}

func TestUpdateCharacter(t *testing.T) {
	// Creating request body
	body := &data.Character{
		ID:     "a2181017-5c53-422b-b6bc-036b27c04fc8",
		Name:   "newName",
		UserID: "a2181017-5c53-422b-b6bc-036b27c04fc8",
	}

	request := httptest.NewRequest(http.MethodPut, "/characters", nil)
	response := httptest.NewRecorder()

	// Add the body to the context since we arent passing through middleware
	ctx := context.WithValue(request.Context(), KeyCharacter{}, body)
	request = request.WithContext(ctx)

	characterHandler := NewCharactersHandler(newCharacterDB())
	characterHandler.UpdateCharacters(response, request)

	if response.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d, but got %d", http.StatusNoContent, response.Code)
	}
}

func TestUpdateNonExistingCharacter(t *testing.T) {
	// Creating request body
	body := &data.Character{
		ID:     "ba7fa838-9576-11eb-a8b3-0242ac130003",
		Name:   "newName",
		UserID: "a2181017-5c53-422b-b6bc-036b27c04fc8",
	}

	request := httptest.NewRequest(http.MethodPut, "/characters", nil)
	response := httptest.NewRecorder()

	// Add the body to the context since we arent passing through middleware
	ctx := context.WithValue(request.Context(), KeyCharacter{}, body)
	request = request.WithContext(ctx)

	characterHandler := NewCharactersHandler(newCharacterDB())
	characterHandler.UpdateCharacters(response, request)

	if response.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d but got : %d", http.StatusNotFound, response.Code)
	}
	if !strings.Contains(response.Body.String(), "Character not found") {
		t.Error("Expected response : Character not found")
	}
}

func TestDeleteExistingCharacter(t *testing.T) {
	request := httptest.NewRequest(http.MethodDelete, "/characters/a2181017-5c53-422b-b6bc-036b27c04fc8", nil)
	response := httptest.NewRecorder()

	characterHandler := NewCharactersHandler(newCharacterDB())

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"id": "a2181017-5c53-422b-b6bc-036b27c04fc8",
	}
	request = mux.SetURLVars(request, vars)

	characterHandler.Delete(response, request)
	if response.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d but got : %d", http.StatusNoContent, response.Code)
	}
}
