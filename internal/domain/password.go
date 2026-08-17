package domain

import "github.com/google/uuid"

// PasswordEntry arepresents a credential stored in the vault.
type PasswordEntry struct {
	ID       uuid.UUID
	Name     string
	Username string
	Password string
	Tags     []string
	Favorite bool
}
