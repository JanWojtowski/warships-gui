//go:build windows

package gui

import (
	tl "github.com/grupawp/termloop"
	termbox "github.com/nsf/termbox-go"
)

// Color represents an RGB color.
type Color tl.Attr

// NewColor returns a new color. Parameters are red, green and blue values.
func NewColor(r, g, b uint8) Color {
	panic("not supported on Windows")
}

func (c Color) toAttr() tl.Attr {
	return tl.Attr(c)
}

var (
	White   = Color(termbox.ColorWhite)
	Black   = Color(termbox.ColorBlack)
	Blue    = Color(termbox.ColorBlue)
	Red     = Color(termbox.ColorRed)
	Grey    = Color(termbox.ColorDarkGray)
	Green   = Color(termbox.ColorGreen)
	Magenta = Color(termbox.ColorMagenta)
)
