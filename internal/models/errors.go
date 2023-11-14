package models

import "errors"

var (
	ErrNoRecord = errors.New("models: no matching record found")

	ErrInvalidCredential = errors.New("models: invalid credentials")

	ErrDuplicateEmail = errors.New("models: duplicate email")
)


