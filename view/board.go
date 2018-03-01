package view

import (
	"github.com/nsf/termbox-go"
	"iteragit.iteratec.de/traze/goclient/mqtt"
	"iteragit.iteratec.de/traze/goclient/mqtt/mock"
	"time"
	"sync"
	"os"
)

const arrowLeft = "←"
const arrowRight = "→"
const arrowUp = "↑"
const arrowDown = "↓"

func InitBoard(wg sync.WaitGroup) {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	defer wg.Done()

	event_queue := make(chan termbox.Event)
	go func() {
		for {
			event_queue <- termbox.PollEvent()
		}
	}()

	RedrawBoard(mock.GetGrid(), ".")

	for {
		select {
		case ev := <-event_queue:
			if ev.Type == termbox.EventKey && ev.Key == termbox.KeyEsc {
				termbox.Close()
				os.Exit(0)
			}
		case grid := <-mqtt.GridQueue():
			RedrawBoard(grid, " ")
		default:
			time.Sleep(10 * time.Millisecond)
		}
	}
}

func redraw_header() {
	const coldef = termbox.ColorDefault
	tbprint(0, 0, termbox.ColorRed, coldef, "Press ESC to quit")
	tbprint(0, 2, termbox.ColorGreen, coldef, "Press "+arrowUp+" to turn NORTH")
	tbprint(0, 3, termbox.ColorGreen, coldef, "Press "+arrowLeft+" to turn WEST")
	tbprint(0, 4, termbox.ColorGreen, coldef, "Press "+arrowDown+" to turn SOUTH")
	tbprint(0, 5, termbox.ColorGreen, coldef, "Press "+arrowRight+" to turn EAST")
}

func RedrawBoard(grid mqtt.Grid, char string) {
	const coldef = termbox.ColorDefault
	termbox.Clear(coldef, coldef)
	redraw_header()
	tiles := grid.Tiles
	for rowIndex, row := range tiles {
		for columnIndex, value := range row {
			if value != 0 {
				tbprint(columnIndex*2, (rowIndex*2)+7, coldef, termbox.ColorRed, char)
			} else {
				tbprint(columnIndex*2, (rowIndex*2)+7, coldef, termbox.ColorWhite, char)
			}
		}
	}
	termbox.Flush()
}

func tbprint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

func CloseTermbox() {
	termbox.Close()
}
