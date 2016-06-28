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

func (t *Text) SetText(text string) {
	t.text = text
	t.dirty = true
}

func (t *Text) Draw() {
	if t.dirty == false {
		return
	}

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

	t.buffer = make([]termbox.Cell, t.actualWidth*t.actualHeight)

	var x, y int
	for _, r := range t.text {
		if r == '\n' {
			x = 0
			y += 1
		} else {
			t.SetCell(x, y, r, termbox.ColorDefault, termbox.ColorDefault)
			x += runewidth.RuneWidth(r)
		}
	}

	t.dirty = false
}
