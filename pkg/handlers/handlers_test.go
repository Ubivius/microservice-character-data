package handlers

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/Ubivius/microservice-character-data/pkg/data"
	"github.com/gorilla/mux"
)

func NewTestLogger() *log.Logger {
	return log.New(os.Stdout, "Tests", log.LstdFlags)
}

func TestGetCharacters(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/characters", nil)
	response := httptest.NewRecorder()

	characterHandler := NewCharactersHandler(NewTestLogger())
	characterHandler.GetCharacters(response, request)

	if response.Code != 200 {
		t.Errorf("Expected status code 200 but got : %d", response.Code)
	}
	if !strings.Contains(response.Body.String(), "\"id\":2") {
		t.Error("Missing elements from expected results")
	}
}

func TestGetExistingCharacterByID(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/characters/1", nil)
	response := httptest.NewRecorder()

	characterHandler := NewCharactersHandler(NewTestLogger())

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"id": "1",
	}
	request = mux.SetURLVars(request, vars)

	characterHandler.GetCharacterByID(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got : %d", http.StatusOK, response.Code)
	}
	if !strings.Contains(response.Body.String(), "\"id\":1") {
		t.Error("Missing elements from expected results")
	}
}

func TestGetNonExistingCharacterByID(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/characters/4", nil)
	response := httptest.NewRecorder()

	characterHandler := NewCharactersHandler(NewTestLogger())

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"id": "4",
	}
	request = mux.SetURLVars(request, vars)

	characterHandler.GetCharacterByID(response, request)

	if response.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got : %d", http.StatusBadRequest, response.Code)
	}
	if !strings.Contains(response.Body.String(), "Character not found") {
		t.Error("Expected response : Character not found")
	}
}

func TestDeleteNonExistantCharacter(t *testing.T) {
	request := httptest.NewRequest(http.MethodDelete, "/characters/4", nil)
	response := httptest.NewRecorder()

	characterHandler := NewCharactersHandler(NewTestLogger())

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"id": "4",
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
		UserID: 1,
	}

	request := httptest.NewRequest(http.MethodPost, "/characters", nil)
	response := httptest.NewRecorder()

	// Add the body to the context since we arent passing through middleware
	ctx := context.WithValue(request.Context(), KeyCharacter{}, body)
	request = request.WithContext(ctx)

	characterHandler := NewCharactersHandler(NewTestLogger())
	characterHandler.AddCharacter(response, request)

	if response.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d, but got %d", http.StatusNoContent, response.Code)
	}
}

func TestAddExistingCharacter(t *testing.T) {
	// Creating request body
	body := &data.Character{
		Name:   "ExistingCharacterName",
		UserID: 1,
	}

	request := httptest.NewRequest(http.MethodPost, "/characters", nil)
	response := httptest.NewRecorder()

	// Add the body to the context since we arent passing through middleware
	ctx := context.WithValue(request.Context(), KeyCharacter{}, body)
	request = request.WithContext(ctx)

	characterHandler := NewCharactersHandler(NewTestLogger())
	characterHandler.AddCharacter(response, request)

	if response.Code != http.StatusConflict {
		t.Errorf("Expected status code %d, but got %d", http.StatusConflict, response.Code)
	}
}

func TestUpdateCharacter(t *testing.T) {
	// Creating request body
	body := &data.Character{
		ID:     1,
		Name:   "newName",
		UserID: 1,
	}

	request := httptest.NewRequest(http.MethodPut, "/characters", nil)
	response := httptest.NewRecorder()

	// Add the body to the context since we arent passing through middleware
	ctx := context.WithValue(request.Context(), KeyCharacter{}, body)
	request = request.WithContext(ctx)

	characterHandler := NewCharactersHandler(NewTestLogger())
	characterHandler.UpdateCharacters(response, request)

	if response.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d, but got %d", http.StatusNoContent, response.Code)
	}
}

func TestDeleteExistingCharacter(t *testing.T) {
	request := httptest.NewRequest(http.MethodDelete, "/characters/1", nil)
	response := httptest.NewRecorder()

	characterHandler := NewCharactersHandler(NewTestLogger())

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"id": "1",
	}
	request = mux.SetURLVars(request, vars)

	characterHandler.Delete(response, request)
	if response.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d but got : %d", http.StatusNoContent, response.Code)
	}
}
