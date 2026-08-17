package vault

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/th3oth3rjak3/password-vale/internal/domain"
)

var _ Service = (*FakeService)(nil)

// FakeService provides an in-memory implementation of Service for development
// and testing.
type FakeService struct {
	entries []domain.PasswordEntry
}

// NewFakeService creates a FakeService populated with sample password entries.
func NewFakeService() Service {
	entries := []domain.PasswordEntry{
		{
			ID:       uuid.MustParse("0198a8c7-7c8a-7c5c-8c9a-123456789001"),
			Name:     "GitHub",
			Username: "jake@example.com",
			Password: "github-password",
			Tags:     []string{"Development", "Work"},
			Favorite: true,
		},
		{
			ID:       uuid.MustParse("0198a8c7-7c8a-7c5c-8c9a-123456789002"),
			Name:     "Amazon",
			Username: "jake@example.com",
			Password: "amazon-password",
			Tags:     []string{"Shopping"},
		},
		{
			ID:       uuid.MustParse("0198a8c7-7c8a-7c5c-8c9a-123456789003"),
			Name:     "Google",
			Username: "jake@example.com",
			Password: "google-password",
			Tags:     []string{"Personal", "Work"},
		},
		{
			ID:       uuid.MustParse("0198a8c7-7c8a-7c5c-8c9a-123456789004"),
			Name:     "Steam",
			Username: "jake@example.com",
			Password: "steam-password",
			Tags: []string{
				"Development",
				"Work",
				"Personal",
				"Important",
				"Programming",
				"GitHub",
				"Open Source",
				"Security",
				"Software",
				"Linux",
				"Go",
				"Testing",
				"Favorite",
				"Projects",
			},
		},
		{
			ID:       uuid.MustParse("0198a8c7-7c8a-7c5c-8c9a-123456789005"),
			Name:     "Discord",
			Username: "jake@example.com",
			Password: "discord-password",
			Tags:     []string{"Gaming", "Entertainment"},
		},
	}

	for i := 4; i <= 100; i++ {
		entries = append(entries, domain.PasswordEntry{
			ID:       uuid.Must(uuid.NewV7()),
			Name:     fmt.Sprintf("Test Password %d", i),
			Username: fmt.Sprintf("user%d@example.com", i),
			Password: fmt.Sprintf("password-%d", i),
			Tags:     fakeTags(i),
			Favorite: i%10 == 0,
		})
	}

	return &FakeService{
		entries: entries,
	}
}

func fakeTags(i int) []string {
	var tags []string

	if i%2 == 0 {
		tags = append(tags, "Development")
	}

	if i%3 == 0 {
		tags = append(tags, "Work")
	}

	if i%5 == 0 {
		tags = append(tags, "Gaming")
	}

	if i%7 == 0 {
		tags = append(tags, "Shopping")
	}

	if len(tags) == 0 {
		tags = append(tags, "Miscellaneous")
	}

	return tags
}

func (s *FakeService) ListEntries() ([]domain.PasswordEntry, error) {
	return s.entries, nil
}

func (s *FakeService) CreateEntry(entry domain.PasswordEntry) error {
	for _, existing := range s.entries {
		if existing.ID == entry.ID {
			return fmt.Errorf("password entry %s already exists", entry.ID)
		}
	}

	s.entries = append(s.entries, entry)
	return nil
}

func (s *FakeService) UpdateEntry(entry domain.PasswordEntry) error {
	for i, existing := range s.entries {
		if existing.ID == entry.ID {
			s.entries[i] = entry
			return nil
		}
	}

	return fmt.Errorf("password entry %s not found", entry.ID)
}

func (s *FakeService) DeleteEntry(id uuid.UUID) error {
	for i, entry := range s.entries {
		if entry.ID == id {
			s.entries = append(s.entries[:i], s.entries[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("password entry %s not found", id)
}
