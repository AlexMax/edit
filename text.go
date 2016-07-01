package main

import (
	"strings"

	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

type Text struct {
	Widget
	text string
}

func NewText() *Text {
	text := &Text{}

	return text
}

func (t *Text) Draw(x int, y int) {
	tw, th := termbox.Size()

	CellCopyRect(
		t.buffer, t.actualWidth, t.actualHeight,
		termbox.CellBuffer(), tw, th,
		0, 0, x, y, x+t.actualWidth, y+t.actualHeight)
}

func (t *Text) Refresh() {
	// No subwidgets, so we can return early if not dirty.
	if t.dirty == false {
		return
	}

	// Determine actual width and height of text.
	if t.desiredWidth == 0 {
		var maxlen int = 0

		rows := strings.Split(t.text, "\n")
		for _, row := range rows {
			rowlen := runewidth.StringWidth(row)
			if rowlen > maxlen {
				maxlen = rowlen
			}
		}

		t.actualWidth = maxlen
	} else {
		t.actualWidth = t.desiredWidth
	}

	if t.desiredHeight == 0 {
		t.actualHeight = strings.Count(t.text, "\n") + 1
	} else {
		t.actualHeight = t.desiredHeight
	}

	// Create buffer.
	t.buffer = make([]termbox.Cell, t.actualWidth*t.actualHeight)

	// Color empty buffer.
	for by := 0; by < t.actualHeight; by++ {
		for bx := 0; bx < t.actualWidth; bx++ {
			t.SetCell(bx, by, 0, Config.DialogFG, Config.DialogBG)
		}
	}

	// Render text.
	var x, y int
	for _, r := range t.text {
		if r == '\n' {
			x = 0
			y += 1
		} else {
			t.SetCell(x, y, r, Config.DialogFG, Config.DialogBG)
			x += runewidth.RuneWidth(r)
		}
	}

	t.dirty = false
}

func (t *Text) SetText(text string) {
	t.text = text
	t.dirty = true
}
