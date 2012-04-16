package additions

import (
	. "github.com/chownplusx/Termbox-Additions/utils"
	"github.com/nsf/termbox-go"
)

func getDimensions(text string, options []string) (int, int) {
	width := len(text)
	for i := 0; i < len(options); i++ {
		if len(options[i]) > width {
			width = len(options[i])
		}
	}
	height := len(options) + 1
	return width, height
}

func DrawOptions(x, y int, options []string, title string) {
	DrawText(x+2, y, title)
	y += 1
	for i := 0; i < len(options); i++ {
		DrawText(x+2, y+i, options[i])
	}
}

func DrawSelection(x, y int, sel int) {
	DrawText(x+1, y+1+sel, ">")
}

func DrawMenu(x, y int, text string, options []string) int {
	selection := 0
	width, height := getDimensions(text, options)

loop:
	for {

		DrawBox(x, y, width, height)     // Draw the surrounding box
		DrawSelection(x, y, selection)   // Draw the selector arrow
		DrawOptions(x, y, options, text) // Draw the option text + title
		termbox.Flush()

		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEnter:
				break loop
			case termbox.KeyArrowDown:
				if selection < len(options)-1 {
					selection += 1
				}
			case termbox.KeyArrowUp:
				if selection > 0 {
					selection -= 1
				}
			}
		}
	}
	return selection
}
