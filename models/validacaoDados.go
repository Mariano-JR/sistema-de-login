package models

import "gopkg.in/validator.v2"

func ValidaDados(user *User) error {
	if err := validator.Validate(user); err != nil {
		return err
	}
	return nil
}
