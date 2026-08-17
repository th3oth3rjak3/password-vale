package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

const (
	tagHorizontalPadding = 8
	tagVerticalPadding   = 4
)

type tagWidget struct {
	widget.BaseWidget
	text string
}

func newTagWidget(text string) *tagWidget {
	tag := &tagWidget{
		text: text,
	}

	tag.ExtendBaseWidget(tag)

	return tag
}

func (t *tagWidget) CreateRenderer() fyne.WidgetRenderer {
	text := canvas.NewText(t.text, theme.Color(theme.ColorNameForeground))
	background := canvas.NewRectangle(theme.Color(theme.ColorNameInputBackground))

	return &tagWidgetRenderer{
		tag:        t,
		text:       text,
		background: background,
	}
}

type tagWidgetRenderer struct {
	tag        *tagWidget
	text       *canvas.Text
	background *canvas.Rectangle
}

func (r *tagWidgetRenderer) Layout(size fyne.Size) {
	r.background.Resize(size)

	textSize := r.text.MinSize()

	r.text.Resize(textSize)
	r.text.Move(fyne.NewPos(
		(size.Width-textSize.Width)/2,
		(size.Height-textSize.Height)/2,
	))
}

func (r *tagWidgetRenderer) MinSize() fyne.Size {
	textSize := r.text.MinSize()

	return fyne.NewSize(
		textSize.Width+tagHorizontalPadding*2,
		textSize.Height+tagVerticalPadding*2,
	)
}

func (r *tagWidgetRenderer) Refresh() {
	r.text.Text = r.tag.text
	r.text.Refresh()

	r.background.FillColor = theme.Color(theme.ColorNameInputBackground)
	r.background.Refresh()

	canvas.Refresh(r.tag)
}

func (r *tagWidgetRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{
		r.background,
		r.text,
	}
}

func (r *tagWidgetRenderer) Destroy() {
}

func (r *tagWidgetRenderer) FocusGained() {
}

func (r *tagWidgetRenderer) FocusLost() {
}

func (r *tagWidgetRenderer) Tapped(*fyne.PointEvent) {
}
