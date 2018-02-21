package main

import (
	"os"
	"os/signal"
	"syscall"
	"fmt"
	_ "iteragit.iteratec.de/traze/goclient/mqtt"
	"github.com/spf13/viper"
	"iteragit.iteratec.de/traze/goclient/mqtt/mock"
	"iteragit.iteratec.de/traze/goclient/mqtt"
	"github.com/nsf/termbox-go"
	"strings"
)

var current string
var curev termbox.Event

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	initConfig()

	go mock.Run()

	go mqtt.ReadFromTopic(mock.MOCK_TOPIC)

	initTermbox()

	go func() {
		signal := <-sigs
		fmt.Printf("Got os signal: %v\n", signal)
		done <- true
	}()

	<-done
	fmt.Println("Exiting.")
}

func initTermbox() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputAlt)
	redraw_all()

	data := make([]byte, 0, 64)

mainloop:

	for {
		if cap(data)-len(data) < 32 {
			newdata := make([]byte, len(data), len(data)+32)
			copy(newdata, data)
			data = newdata
		}
		beg := len(data)
		d := data[beg : beg+32]
		switch ev := termbox.PollRawEvent(d); ev.Type {
		case termbox.EventRaw:
			data = data[:beg+ev.N]
			current = fmt.Sprintf("%q", data)
			if current == `"q"` {
				break mainloop
			}

			for {
				ev := termbox.ParseEvent(data)
				if ev.N == 0 {
					break
				}
				curev = ev
				copy(data, data[curev.N:])
				data = data[:len(data)-curev.N]
			}
		case termbox.EventError:
			panic(ev.Err)
		}
		redraw_all()
	}
}

func redraw_all() {
	const coldef = termbox.ColorDefault
	termbox.Clear(coldef, coldef)
	tbprint(0, 0, termbox.ColorGreen, coldef, "Press 'q' to quit")
	tbprint(0, 1, coldef, coldef, current)
	switch curev.Type {
	case termbox.EventKey:
		tbprint(0, 2, coldef, coldef,
			fmt.Sprintf("EventKey: k: %d, c: %c, mod: %s", curev.Key, curev.Ch, mod_str(curev.Mod)))
	case termbox.EventNone:
		tbprint(0, 2, coldef, coldef, "EventNone")
	}
	tbprint(0, 3, coldef, coldef, fmt.Sprintf("%d", curev.N))
	termbox.Flush()
}

func tbprint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

func mod_str(m termbox.Modifier) string {
	var out []string
	if m&termbox.ModAlt != 0 {
		out = append(out, "ModAlt")
	}
	if m&termbox.ModMotion != 0 {
		out = append(out, "ModMotion")
	}
	return strings.Join(out, " | ")
}

func initConfig() {
	viper.SetConfigName("goclient")
	viper.AddConfigPath("/etc/goclient")
	viper.AddConfigPath("./conf/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
