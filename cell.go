package main

import "github.com/nsf/termbox-go"

// CellCopyRect copies a rectangle-shaped area between two different
// cell arrays.
func CellCopyRect(
	src []termbox.Cell, srcw int, srch int,
	dst []termbox.Cell, dstw int, dsth int,
	srcx int, srcy int, dstx int, dsty int, dstmaxx int, dstmaxy int) {

	// Determine width and height of copy based on destination rectangle.
	var width, height int
	width = dstmaxx - dstx
	height = dstmaxy - dsty

	// Modify copy based on destination buffer dimensions
	if dstx < 0 {
		srcx -= dstx
		width += dstx
		dstx = 0
	}
	if dstmaxx > dstw {
		width -= dstmaxx - dstw
	}
	if dsty < 0 {
		srcy -= dsty
		height += dsty
		dsty = 0
	}
	if dstmaxy > dsth {
		height -= dstmaxy - dsth
	}

	// If we're not copying anything, abort
	if width <= 0 || height <= 0 {
		return
	}

	for i := 0; i < height; i++ {
		srcStart := srcw*(i+srcy) + srcx
		srcEnd := srcStart + width
		dstStart := dstw*(i+dsty) + dstx
		dstEnd := dstStart + width

		//fmt.Printf("%d, %d\n", width, height)
		//fmt.Printf("%d, %d, %d, %d\n", srcStart, srcEnd, dstStart, dstEnd)

		copy(dst[dstStart:dstEnd], src[srcStart:srcEnd])
	}
}
