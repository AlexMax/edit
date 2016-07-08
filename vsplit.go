package main

import "github.com/nsf/termbox-go"

type VSplit struct {
	Widget
	widgets []IWidget
}

func NewVSplit() *VSplit {
	vsplit := &VSplit{}

	return vsplit
}

func (v *VSplit) Draw(x int, y int) {
	tw, th := termbox.Size()

	CellCopyRect(
		v.buffer, v.actualWidth, v.actualHeight,
		termbox.CellBuffer(), tw, th,
		0, 0, x, y, x+v.actualWidth, y+v.actualHeight)

	// Draw each widget
	for _, widget := range v.widgets {
		widget.Draw(x, y)
		y += widget.GetActualHeight()
	}
}

func (v *VSplit) Refresh() {
	// Determine actual width and height of entire buffer.
	if v.desiredWidth == 0 {
		// Determine largest width
		maxw := 0
		for _, widget := range v.widgets {
			width := widget.GetActualWidth()
			if width > maxw {
				maxw = width
			}
		}
		v.actualWidth = maxw
	} else {
		v.actualWidth = v.desiredWidth
	}

	if v.desiredHeight == 0 {
		v.actualHeight = 0
		for _, widget := range v.widgets {
			v.actualHeight += widget.GetActualHeight()
		}
	} else {
		v.actualHeight = v.desiredHeight
	}

	// Create buffer.
	v.buffer = make([]termbox.Cell, v.actualWidth*v.actualHeight)
	v.dirty = false
}

func (v *VSplit) AddWidget(widget IWidget) {
	v.widgets = append(v.widgets, widget)
	v.dirty = true
}
