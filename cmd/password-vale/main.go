package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/th3oth3rjak3/password-vale/internal/ui"
	"github.com/th3oth3rjak3/password-vale/internal/vault"
)

func main() {
	a := app.New()

	vaultService := vault.NewFakeService()
	window := a.NewWindow("Password Vale")

	window.SetContent(ui.NewVaultView(vaultService))
	window.Resize(fyne.NewSize(1100, 700))
	window.ShowAndRun()
}
