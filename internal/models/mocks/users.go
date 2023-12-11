package mocks

import "snippetbox.project/internal/models"

type UserModel struct{}

func (m *UserModel) Insert(name, email, password string) error {
	switch email {
	case "marcelo@example.com":
		return models.ErrDuplicateEmail
	default:
		return nil
	}
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
	if email == "galindo@example.com" && password == "pa$$word" {
		return 1, nil
	}

	return 0, models.ErrInvalidCredential
}

func (m *UserModel) Exists(id int) (bool, error) {
	switch id {
	case 1:
		return true, nil
	default:
		return false, nil
	}
}
