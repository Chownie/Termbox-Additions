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

func DrawForm(x, y int, text string) string {
	input := ""
	width, height := GetFormSize(text)

loop:
	for {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		DrawBox(x, y, width, height)
		drawSelection(x, y, len(input)-1)
		drawOptions(x, y, text)
		drawInput(x, y, input)
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
			case termbox.KeySpace:
				input += " "
			}
			switch ev.Ch {
			default:
				input += string(ev.Ch)
			}
		}
	}
	return input
}
