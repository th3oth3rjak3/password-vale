package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/th3oth3rjak3/password-vale/internal/domain"
	"github.com/th3oth3rjak3/password-vale/internal/vault"
)

// NewVaultView creates the main vault view.
func NewVaultView(service vault.Service) fyne.CanvasObject {
	entries, err := service.ListEntries()
	if err != nil {
		return widget.NewLabel(err.Error())
	}

	// Main Content
	mainContent := container.NewBorder(
		newVaultHeader(),
		nil,
		nil,
		nil,
		newVaultContent(entries),
	)

	split := container.NewHSplit(
		newVaultSidebar(entries),
		mainContent,
	)

	split.SetOffset(0.30)

	return container.NewBorder(
		nil,
		nil,
		nil,
		nil,
		split,
	)
}

func newVaultSidebar(entries []domain.PasswordEntry) fyne.CanvasObject {
	tags := make(map[string]struct{})

	for _, entry := range entries {
		for _, tag := range entry.Tags {
			tags[tag] = struct{}{}
		}
	}

	tagList := container.NewVBox()

	for tag := range tags {
		tagList.Add(widget.NewButton(tag, nil))
	}

	tagSection := container.NewBorder(
		widget.NewLabel("Tags"),
		nil,
		nil,
		nil,
		container.NewVScroll(tagList),
	)

	return container.NewBorder(
		container.NewVBox(
			widget.NewLabel("Vault"),
			widget.NewButton("All Passwords", nil),
			widget.NewButton("Favorites", nil),
		),
		nil,
		nil,
		nil,
		tagSection,
	)
}

func newVaultHeader() fyne.CanvasObject {
	headerTitle := widget.NewLabel("Passwords")

	addButton := widget.NewButtonWithIcon(
		"",
		theme.ContentAddIcon(),
		nil,
	)

	lockButton := widget.NewButtonWithIcon(
		"",
		theme.LogoutIcon(),
		nil,
	)

	settingsButton := widget.NewButtonWithIcon(
		"",
		theme.SettingsIcon(),
		nil,
	)

	actions := container.NewHBox(
		addButton,
		lockButton,
		settingsButton,
	)

	return container.NewBorder(
		nil,
		nil,
		headerTitle,
		actions,
		nil,
	)
}

func newVaultContent(entries []domain.PasswordEntry) fyne.CanvasObject {
	entryList := container.NewVBox()

	for _, entry := range entries {
		entryList.Add(newPasswordEntryItem(entry))
	}

	return container.NewVScroll(entryList)
}

func newPasswordEntryItem(entry domain.PasswordEntry) fyne.CanvasObject {
	name := widget.NewLabel(entry.Name)
	username := widget.NewLabel(entry.Username)

	return container.NewBorder(
		nil,
		nil,
		name,
		nil,
		username,
	)
}
