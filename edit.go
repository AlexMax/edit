package main

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

func main() {
	err := termbox.Init()
	if err != nil {
		fmt.Print(err.Error())
	}
	defer termbox.Close()

	aboutText := NewText()
	aboutText.SetText("edit 0.0\nCopyright (C) 2016  Alex Mayfield")
	aboutText.Refresh()

	aboutModal := NewModal()
	aboutModal.SetTitle("About Edit")
	aboutModal.SetBody(aboutText)
	aboutModal.Refresh()

	var x, y int
	for {
		termbox.Clear(termbox.ColorWhite, termbox.ColorBlue)
		aboutModal.Draw(x, y)
		termbox.Flush()

		event := termbox.PollEvent()

		if event.Type == termbox.EventKey {
			switch event.Key {
			case termbox.KeyArrowUp:
				y -= 1
			case termbox.KeyArrowDown:
				y += 1
			case termbox.KeyArrowLeft:
				x -= 1
			case termbox.KeyArrowRight:
				x += 1
			default:
				return
			}
		}
	}
}
