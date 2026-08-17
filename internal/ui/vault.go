package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/th3oth3rjak3/password-vale/internal/vault"
)

// NewVaultView creates the main vault view.
func NewVaultView(service vault.Service) fyne.CanvasObject {
	return container.NewCenter(
		widget.NewLabel("Password Vale"),
	)
}
