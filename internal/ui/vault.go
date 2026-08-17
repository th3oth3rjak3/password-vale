package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/th3oth3rjak3/password-vale/internal/vault"
)

// NewVaultView creates the main vault view.
func NewVaultView(service vault.Service) fyne.CanvasObject {
	_ = service

	// Sidebar
	sidebar := container.NewVBox(
		widget.NewLabel("Vault"),
		widget.NewButton("All Passwords", nil),
		widget.NewButton("Favorites", nil),
	)

	// Action Bar Components
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
		lockButton,
		settingsButton,
	)

	header := container.NewBorder(
		nil,
		nil,
		nil,
		actions,
		nil,
	)

	// Main Content
	mainContent := container.NewBorder(
		header,
		nil,
		nil,
		nil,
		widget.NewLabel("Vault Main Content"),
	)

	split := container.NewHSplit(
		sidebar,
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
