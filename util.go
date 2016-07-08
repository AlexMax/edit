package main

import (
	"strings"

	"github.com/mattn/go-runewidth"
)

// StringCenter returns a string with the input string centered horizontally in
// a space specified by width, using the padding string pad.

// This function uses rune width to ensure that the string is actually
// visibly centered.  The padding string is currently assumed to be one
// character wide.
func StringCenter(str string, width int, pad string) string {
	strWidth := runewidth.StringWidth(str)
	leftPad := width/2 - strWidth/2
	rightPad := width - leftPad - strWidth

	// Ensure uneven padding is shorter on the left side.
	if leftPad > rightPad {
		leftPad, rightPad = rightPad, leftPad
	}

	result := strings.Repeat(pad, leftPad) + str + strings.Repeat(pad, rightPad)

	return result
}
