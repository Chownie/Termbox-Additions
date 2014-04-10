package utils

import (
	"github.com/nsf/termbox-go"
	"strings"
)

const (
	TOPLEFT       = "┌"
	TOPRIGHT      = "┐"
	BOTTOMLEFT    = "└"
	BOTTOMRIGHT   = "┘"
	VERTICAL      = "│"
	HORIZONTAL    = "─"
	JOIN_LEFT  = "├"
	JOIN_RIGHT = "┤"
	JOIN_TOP   = "┬"
	JOIN_BOT   = "┴"
	//Modes
	CONNECT_TOP   = iota
	CONNECT_BOT
	CONNECT_LEFT
	CONNECT_RIGHT
	CONNECT_NONE
)

func DrawRichTextMulti(x, y int, text string, fgColor termbox.Attribute, bgColor termbox.Attribute) {
	lines := strings.SplitAfterN(text, "\n", -1)
	for i := 0; i < len(lines); i++ {
		DrawRichText(x, y+i, lines[i], fgColor, bgColor)
	}
}

func DrawRichText(x, y int, text string, fgColor termbox.Attribute, bgColor termbox.Attribute) {
	j := 0
	for _, r := range text {
		termbox.SetCell(x+j, y, r, fgColor, bgColor)
		j += 1
	}
}

func DrawText(x, y int, text string) {
	j := 0
	for _, r := range text {
		termbox.SetCell(x+j, y, r, termbox.ColorDefault, termbox.ColorDefault)
		j += 1
	}
}

func DrawTextMulti(x, y int, text string) {
	lines := strings.SplitAfterN(text, "\n", -1)
	for i := 0; i < len(lines); i++ {
		DrawText(x, y+i, lines[i])
	}
}

func DrawBox(x, y, width, height, mode int) {
	if mode == CONNECT_TOP {
		DrawText(x, y, JOIN_LEFT+strings.Repeat(HORIZONTAL, width+2)+JOIN_RIGHT)
	} else {
		DrawText(x, y, TOPLEFT+strings.Repeat(HORIZONTAL, width+2)+TOPRIGHT)
	}
	
	for i := 1; i < height+1; i++ {
		DrawText(x, y+i, VERTICAL+strings.Repeat(" ", width+2)+VERTICAL)
	}
	
	if mode == CONNECT_BOT {
		DrawText(x, y+height, JOIN_LEFT+strings.Repeat(HORIZONTAL, width+2)+JOIN_RIGHT)
	} else {
		DrawText(x, y+height, BOTTOMLEFT+strings.Repeat(HORIZONTAL, width+2)+BOTTOMRIGHT)
	}
}
