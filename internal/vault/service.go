package vault

import (
	"github.com/google/uuid"
	"github.com/th3oth3rjak3/password-vale/internal/domain"
)

type Service interface {
	ListEntries() ([]domain.PasswordEntry, error)
	CreateEntry(entry domain.PasswordEntry) error
	UpdateEntry(entry domain.PasswordEntry) error
	DeleteEntry(id uuid.UUID) error
}
