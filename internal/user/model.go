package user

import (
    "strings"
    "time"
)

type User struct {
    ID int64 `json:id`
    Name  string `json:"name"`
    Email string `json:"email"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

// Validate checks that a user's name and email are valid.
func (u *User) Validate() error {
    u.Name = strings.TrimSpace(u.Name)
    u.Email = strings.TrimSpace(u.Email)

    if u.Name == "" {
        return ErrInvalidName
    }

    if !strings.Contains(u.Email, "@") {
        return ErrInvalidEmail
    }

    return nil
}

// ErrInvalidName is returned when a user's name is invalid.
var ErrInvalidName = ValidationError("invalid name")

// ErrInvalidEmail is returned when a user's email is invalid.
var ErrInvalidEmail = ValidationError("invalid email")

// ValidationError is a custom error type that represents a validation error.
type ValidationError string

// Error returns the string representation of a validation error.
func (e ValidationError) Error() string {
    return string(e)
}
