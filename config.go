package main

import "github.com/nsf/termbox-go"

var Config = struct {
	DialogFG       termbox.Attribute
	DialogBG       termbox.Attribute
	DialogTitleFG  termbox.Attribute
	DialogTitleBG  termbox.Attribute
	DialogButtonFG termbox.Attribute
	DialogButtonBG termbox.Attribute
}{
	termbox.ColorBlack,
	termbox.ColorWhite,
	termbox.ColorBlack,
	termbox.ColorCyan,
	termbox.ColorBlack,
	termbox.ColorCyan,
}
