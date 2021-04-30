package data

import (
	"regexp"

	"github.com/go-playground/validator"
)

func (character *Character) ValidateCharacter() error {
	validate := validator.New()
	err := validate.RegisterValidation("name", validateName)
	if err != nil {
		panic(err)
	}
	return validate.Struct(character)
}

func validateName(fieldLevel validator.FieldLevel) bool {
	// TODO: Chose what kind of name are accepted ( Special characters? Spaces? Numbers? Etc.)
	re := regexp.MustCompile(`[a-z0-9-]+`)
	matches := re.FindAllString(fieldLevel.Field().String(), -1)

	return len(matches) == 1
}
