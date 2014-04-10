package utils

import (
	"github.com/nsf/termbox-go"
	"strings"
)

const (
	// Box Drawing
	TOPLEFT		= "┌"
	TOPRIGHT	= "┐"
	BOTTOMLEFT	= "└"
	BOTTOMRIGHT	= "┘"
	VERTICAL	= "│"
	HORIZONTAL	= "─"
	// Joining Marks
	JOIN_TOP	= "┬"
	JOIN_LEFT	= "├"
	JOIN_RIGHT	= "┤"
	JOIN_BOT	= "┴"
	JOIN_ALL	= "┼"
	// Modes
	CONNECT_NONE	= 0
	CONNECT_TOP	= 1
	CONNECT_BOT	= 4
	CONNECT_LEFT	= 8
	CONNECT_RIGHT	= 2
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
	TRCorner := TOPRIGHT
	TLCorner := TOPLEFT
	
	if mode == CONNECT_TOP {
		TLCorner = JOIN_LEFT
		TRCorner = JOIN_RIGHT
	} else if mode == CONNECT_TOP+CONNECT_LEFT {
		TLCorner = JOIN_ALL
	} else if mode == CONNECT_TOP+CONNECT_RIGHT {
		TRCorner = JOIN_ALL
	} else if mode == CONNECT_LEFT {
		TLCorner = JOIN_TOP
	} else if mode == CONNECT_RIGHT {
		TRCorner = JOIN_TOP
	}
	
	
	DrawText(x, y, TLCorner+strings.Repeat(HORIZONTAL, width+2)+TRCorner)
	
	for i := 1; i < height+1; i++ {
		DrawText(x, y+i, VERTICAL+strings.Repeat(" ", width+2)+VERTICAL)
	}

	BRCorner := BOTTOMRIGHT
	BLCorner := BOTTOMLEFT	
	if mode == CONNECT_BOT {
		BLCorner = JOIN_LEFT
		BRCorner = JOIN_RIGHT
	} else if mode == CONNECT_BOT+CONNECT_LEFT {
		BLCorner = JOIN_ALL
	} else if mode == CONNECT_BOT+CONNECT_RIGHT {
		BRCorner = JOIN_ALL
	} else if mode == CONNECT_LEFT {
		BLCorner = JOIN_BOT
	} else if mode == CONNECT_RIGHT {
		BRCorner = JOIN_BOT
	}
	
	DrawText(x, y+height, BLCorner+strings.Repeat(HORIZONTAL, width+2)+BRCorner)
}
