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

	aboutBox := NewText()
	aboutBox.SetText("edit 0.0\nCopyright (C) 2016  Alex Mayfield")
	aboutBox.Draw()

	w, _ := termbox.Size()

	copy(termbox.CellBuffer(), aboutBox.buffer[:aboutBox.actualWidth])
	copy(termbox.CellBuffer()[w:], aboutBox.buffer[aboutBox.actualWidth:])
	termbox.Flush()

	select {}
}
