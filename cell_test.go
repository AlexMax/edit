package main

import (
	"testing"

	"github.com/nsf/termbox-go"
)

func TestCellCopyRect(t *testing.T) {
	// 012345
	// 6789ab
	// cdefgh
	// ijklmn
	// opqrst
	// uvwxyz
	var testrunes string = "0123456789abcdefghijklmnopqrstuvwxyz"

	var src [36]termbox.Cell
	var srcw int = 6
	var srch int = 6

	var dst [36]termbox.Cell
	var dstw int = 6
	var dsth int = 6

	for i, r := range testrunes {
		src[i].Ch = r
	}

	// Test basic copy.
	CellCopyRect(src[:], srcw, srch, dst[:], dstw, dsth,
		0, 0, 0, 0, 6, 6)
	if dst[0].Ch != '0' {
		t.FailNow()
	}
	if dst[35].Ch != 'z' {
		t.FailNow()
	}

	// Test copy to destination offset.
	CellCopyRect(src[:], srcw, srch, dst[:], dstw, dsth,
		0, 0, 1, 1, 6, 6)
	if dst[7].Ch != '0' {
		t.Fatalf("Incorrect rune at offset 7 (expected %c, actual %c)", '0', dst[7].Ch)
	}
	if dst[35].Ch != 's' {
		t.Fatalf("Incorrect rune at offset 35 (expected %c, actual %c)", '0', dst[35].Ch)
	}

	// Test copy to negative destination offset.
	CellCopyRect(src[:], srcw, srch, dst[:], dstw, dsth,
		0, 0, -1, -1, 6, 6)
	if dst[7].Ch != '7' {
		t.Fatalf("Incorrect rune at offset 0 (expected %c, actual %c)", '7', dst[0].Ch)
	}
	if dst[28].Ch != 'z' {
		t.Fatalf("Incorrect rune at offset 28 (expected %c, actual %c)", 'z', dst[28].Ch)
	}
}
