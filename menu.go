package additions

import (
	. "github.com/Chownie/Termbox-Additions/utils"
	"github.com/nsf/termbox-go"
)

const (
	AL_LEFT = iota
	AL_RIGHT
	AL_CENTER
)

func GetMenuSize(text string, options []string) (int, int) {
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

func DrawMenu(x, y int, text string, options []string, style, mode int) int {
	selection := 0
	width, height := GetMenuSize(text, options)
	posx := 0
	posy := 0
loop:
	for {
		if style == AL_LEFT {
			posx = 0
		} else if style == AL_CENTER {
			posx = (x/2) - (width/2)
			posy = (y/2) - (height/2)
		} else if style == AL_RIGHT {
			posx = x-width
			posy = y-height
		}
		
		DrawBox(posx, posy, width, height, mode)     // Draw the surrounding box
		DrawSelection(posx, posy, selection)   // Draw the selector arrow
		DrawOptions(posx, posy, options, text) // Draw the option text + title
		

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
	termbox.Flush()
	return selection
}
