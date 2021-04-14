package database

import (
	"testing"

	"github.com/Ubivius/microservice-character-data/pkg/data"
	"github.com/Ubivius/microservice-character-data/pkg/resources"
	"github.com/google/uuid"
)

func newResourceManager() resources.ResourceManager {
	return resources.NewMockResources()
}

func TestMongoDBConnectionAndShutdownIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}

	mp := NewMongoCharacters(newResourceManager())
	if mp == nil {
		t.Fail()
	}
	mp.CloseDB()
}

func TestMongoDBAddCharacterIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}

	character := &data.Character{
		Name:        "testName",
		UserID:      "a2181017-5c53-422b-b6bc-036b27c04fc8",
	}

	mp := NewMongoCharacters(newResourceManager())
	err := mp.AddCharacter(character)
	if err != nil {
		t.Errorf("Failed to add character to database")
	}
	mp.CloseDB()
}

func TestMongoDBUpdateCharacterIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}

	character := &data.Character{
		ID:          uuid.NewString(),
		UserID:      "a2181017-5c53-422b-b6bc-036b27c04fc8",
	}

	mp := NewMongoCharacters(newResourceManager())
	err := mp.UpdateCharacter(character)
	if err != nil {
		t.Fail()
	}
	mp.CloseDB()
}

func TestMongoDBGetCharactersIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}

	mp := NewMongoCharacters(newResourceManager())
	characters := mp.GetCharacters()
	if characters == nil {
		t.Fail()
	}

	mp.CloseDB()
}

func TestMongoDBGetCharacterByIDIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}

	mp := NewMongoCharacters(newResourceManager())
	_, err := mp.GetCharacterByID("a2181017-5c53-422b-b6bc-036b27c04fc8")
	if err != nil {
		t.Fail()
	}

	mp.CloseDB()
}

func TestMongoDBGetCharactersByUserIDIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}

	mp := NewMongoCharacters(newResourceManager())
	_, err := mp.GetCharactersByUserID("a2181017-5c53-422b-b6bc-036b27c04fc8")
	if err != nil {
		t.Fail()
	}

	mp.CloseDB()
}
