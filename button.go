package main

import (
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

type Button struct {
	Widget
	text string
}

func NewButton() *Button {
	button := &Button{}

	return button
}

func (b *Button) Draw(x int, y int) {
	tw, th := termbox.Size()

	CellCopyRect(
		b.buffer, b.actualWidth, b.actualHeight,
		termbox.CellBuffer(), tw, th,
		0, 0, x, y, x+b.actualWidth, y+b.actualHeight)
}

func (b *Button) Refresh() {
	// Determine actual width and height of entire buffer.
	if b.desiredWidth == 0 {
		b.actualWidth = runewidth.StringWidth(b.text) + 4
	} else {
		b.actualWidth = b.desiredWidth
	}
	b.actualHeight = 1

	// Create buffer.
	b.buffer = make([]termbox.Cell, b.actualWidth*b.actualHeight)

	// Render text.
	x, y := 2, 0
	for _, r := range b.text {
		b.SetCell(x, y, r, Config.DialogButtonFG, Config.DialogButtonBG)
		x += runewidth.RuneWidth(r)
	}
}

func (b *Button) SetText(text string) {
	b.text = text
	b.dirty = true
}
