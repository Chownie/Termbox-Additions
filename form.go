package additions

import (
	. "github.com/Chownie/Termbox-Additions/utils"
	"github.com/nsf/termbox-go"
)

func GetFormSize(text string) (int, int) {
	width := len(text)
	height := 4
	return width, height
}

func drawOptions(x, y int, title string) {
	DrawText(x+2, y, title)
}

func drawSelection(x, y int, length int) {
	DrawText(x+length+2, y+1, "_")
}

func drawInput(x, y int, input string) {
	DrawText(x+1, y+1, input)
}

func DrawForm(x, y int, text string, style int) string {
	input := ""
	width, height := GetFormSize(text)
	posx := 0
	posy := 0
	
	if style == AL_LEFT {
		posx = 0 
	} else if style == AL_CENTER {
		posx = (x/2) - (width/2)
		posy = (y/2) - (height/2)
	} else if style == AL_RIGHT {
		posx = x-width
		posy = y-height
	}
	
loop:
	for {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		DrawBox(posx, posy, width, height)
		drawSelection(posx, posy, len(input)-1)
		drawOptions(posx, posy, text)
		drawInput(posx, posy, input)
		termbox.Flush()
		
		
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEnter:
				break loop
			case termbox.KeyBackspace:
				input = input[:len(input)-1]
			case termbox.KeyBackspace2:
				input = input[:len(input)-1]
			}
			switch ev.Ch {
			case '\b':
			default:
				input += string(ev.Ch)
			}
		}
	}
	termbox.Flush()
	return input
}
