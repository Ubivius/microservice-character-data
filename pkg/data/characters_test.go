package data

import "testing"

func TestChecksValidation(t *testing.T) {
	character := &Character{
		Name:   "Malcolm",
		UserID: "a2181017-5c53-422b-b6bc-036b27c04fc8",
	}

	err := character.ValidateCharacter()

	if err != nil {
		t.Fatal(err)
	}
}
