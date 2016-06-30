package main

import "github.com/nsf/termbox-go"

type Widget struct {
	actualWidth   int
	actualHeight  int
	buffer        []termbox.Cell
	desiredWidth  int
	desiredHeight int
	dirty         bool
}

type IWidget interface {
	Draw(x int, y int)
	GetActualHeight() int
	GetActualWidth() int
	Refresh()
	SetCell(x int, y int, ch rune, fg termbox.Attribute, bg termbox.Attribute)
	SetHeight(height int)
	SetWidth(width int)
}

func (w *Widget) GetActualHeight() int {
	return w.actualHeight
}

func (w *Widget) GetActualWidth() int {
	return w.actualWidth
}

func (w *Widget) SetCell(x int, y int, ch rune, fg termbox.Attribute, bg termbox.Attribute) {
	if x < 0 || x >= w.actualWidth {
		return
	}

	if y < 0 || y >= w.actualHeight {
		return
	}

	w.buffer[y*w.actualWidth+x] = termbox.Cell{ch, fg, bg}
}

func (w *Widget) SetHeight(height int) {
	if w.desiredHeight == height {
		return
	}

	w.desiredHeight = height
	w.dirty = true
}

func (w *Widget) SetWidth(width int) {
	if w.desiredWidth == width {
		return
	}

	w.desiredWidth = width
	w.dirty = true
}
