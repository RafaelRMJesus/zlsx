package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Id   string `json:"id" validate:"nonzero,regexp=^[0-9]*$"`
	Name string `json:"name" validate:"nonzero,regexp=^[a-zA-Z]*$"`
}

func ValidateStudent(student *Student) error {
	if err := validator.Validate(student); err != nil {
		return err
	}
	return nil
}
