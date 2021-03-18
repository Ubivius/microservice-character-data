package data

import "testing"

func TestChecksValidation(t *testing.T) {
	character := &Character{
		Name:   "Malcolm",
		UserID: 1,
	}

	err := character.ValidateCharacter()

	if err != nil {
		t.Fatal(err)
	}
}
