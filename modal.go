package main

import "github.com/nsf/termbox-go"
import "github.com/mattn/go-runewidth"

type Modal struct {
	Widget
	title string
	body  IWidget
}

func NewModal() *Modal {
	modal := &Modal{}

	return modal
}

func (m *Modal) Draw(x int, y int) {
	tw, th := termbox.Size()

	CellCopyRect(
		m.buffer, m.actualWidth, m.actualHeight,
		termbox.CellBuffer(), tw, th,
		0, 0, x, y, x+m.actualWidth, y+m.actualHeight)

	if m.body != nil {
		m.body.Draw(x, y+1)
	}
}

func (m *Modal) Refresh() {
	// Determine actual width and height of modal.
	if m.desiredWidth == 0 {
		m.actualWidth = m.body.GetActualWidth()
	} else {
		m.actualWidth = m.desiredWidth
	}

	if m.desiredHeight == 0 {
		m.actualHeight = m.body.GetActualHeight()
		m.actualHeight += 1
	} else {
		m.actualHeight = m.desiredHeight
	}

	// Create buffer.
	m.buffer = make([]termbox.Cell, m.actualWidth*m.actualHeight)

	// Render title.
	titleWidth := runewidth.StringWidth(m.title)
	x := (m.actualWidth / 2) - (titleWidth / 2)
	for i := 0; i < x; i++ {
		m.SetCell(i, 0, ' ', Config.DialogTitleFG, Config.DialogTitleBG)
	}
	for _, r := range m.title {
		m.SetCell(x, 0, r, Config.DialogTitleFG, Config.DialogTitleBG)
		x += runewidth.RuneWidth(r)
	}
	for ; x < m.actualWidth; x++ {
		m.SetCell(x, 0, ' ', Config.DialogTitleFG, Config.DialogTitleBG)
	}

	m.dirty = false
}

func (m *Modal) SetTitle(title string) {
	m.title = title
	m.dirty = true
}

func (m *Modal) SetBody(body IWidget) {
	m.body = body
	m.dirty = true
}
