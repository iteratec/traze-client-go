package view

import (
	"github.com/nsf/termbox-go"
	"iteragit.iteratec.de/traze/goclient/mqtt"
	"iteragit.iteratec.de/traze/goclient/mqtt/mock"
	"time"
	"sync"
	"os"
	"iteragit.iteratec.de/traze/goclient/util/log"
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

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	drawStaticHeader()

	redrawBoard(mock.GetGrid())

	for {
		select {
		case ev := <-event_queue:
			if ev.Type == termbox.EventKey {
				switch ev.Key {
				case termbox.KeyEsc:
					termbox.Close()
					os.Exit(0)
				case termbox.KeyArrowLeft:
					//TODO: handle move left
				case termbox.KeyArrowRight:
					//TODO: handle move right
				case termbox.KeyArrowUp:
					//TODO: handle move up
				case termbox.KeyArrowDown:
					//TODO: handle move down
				}
			}
		case grid := <-mqtt.GridQueue():
			log.Info.Printf("got grid from channel")
			redrawBoard(grid)
		case games := <- mqtt.GamesQueue():
			log.Info.Printf("got games from channel")
			redrawGames(games)
		default:
			time.Sleep(10 * time.Millisecond)
		}
	}
}

func drawStaticHeader() {
	const coldef = termbox.ColorDefault
	tbprint(0, 0, termbox.ColorRed, coldef, "Press ESC to quit")
	tbprint(0, 2, termbox.ColorGreen, coldef, "Press "+arrowUp+" "+arrowLeft+" "+arrowDown+" "+arrowRight+" to move your bike")
}

func redrawGames(games []mqtt.Game){
	log.Info.Printf("Got games to draw on board: %v\n", games)
	maxGames := 5
	for index, game := range games {
		if index < maxGames-1 {
			tbprint(200, 2 + index, termbox.ColorGreen, termbox.ColorDefault, game.Name + " (" + string(game.ActivePlayers) + " active players)")
		}
	}
}

func redrawBoard(grid mqtt.Grid) {
	tiles := grid.Tiles

	var rowCounter int
	for columnIndex, column := range tiles {
		rowCounter = 0
		for rowIndex := len(column); rowIndex > 0; rowIndex-- {
			value := tiles[columnIndex][rowIndex-1]
			log.Debug.Printf("Printing value %v at position %v|%v\n", value, columnIndex, rowIndex)
			if value != 0 {
				tbprint(columnIndex*2, rowCounter+4, termbox.ColorRed, termbox.ColorDefault, "●")
			} else {
				tbprint(columnIndex*2, rowCounter+4, termbox.ColorWhite, termbox.ColorDefault, "·")
			}
			rowCounter++
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
