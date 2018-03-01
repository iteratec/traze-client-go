package mock

import (
	"iteragit.iteratec.de/traze/goclient/mqtt"
	"time"
	"encoding/json"
	"iteragit.iteratec.de/traze/goclient/util/log"
)

func Run() {

	client := mqtt.GetClient()

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	for {

		client.Publish("traze/mockedgame/grid", mqtt.Qos(), false, getGrid())
		client.Publish("traze/games", mqtt.Qos(), false, getGames())

		time.Sleep(1000 * time.Millisecond)

	}

	client.Disconnect(250)
	log.Info.Print("Sample Publisher Disconnected")

}

func GetGrid() mqtt.Grid {
	return mqtt.Grid{
		Height: 20,
		Width:  30,
		Tiles: [][]int{
			{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 4, 4, 4, 0},
			{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 0},
			{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 0},
			{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 0},
			{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 0},
			{0, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 3, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 0},
			{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 3, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 0},
			{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 3, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 0},
			{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 3, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 0},
			{0, 0, 1, 1, 1, 1, 1, 0, 0, 0, 0, 3, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 0},
			{0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 0, 0, 4, 0},
			{0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 4, 0},
			{0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 4, 0},
			{0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 2, 2, 2, 2, 2, 2, 2, 2, 2, 0, 0, 4, 0},
			{0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 0},
			{0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 4, 4, 4, 0},
			{0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		Bikes: []mqtt.Bike{
			mqtt.Bike{
				PlayerId:        1,
				CurrentLocation: [2]int{1, 0},
				Direction:       "E",
				Trail:           [][2]int{{9, 0}, {8, 0}},
			},
			mqtt.Bike{
				PlayerId:        2,
				CurrentLocation: [2]int{1, 0},
				Direction:       "E",
				Trail:           [][2]int{{9, 0}, {8, 0}},
			},
			mqtt.Bike{
				PlayerId:        3,
				CurrentLocation: [2]int{1, 0},
				Direction:       "E",
				Trail:           [][2]int{{9, 0}, {8, 0}},
			},
			mqtt.Bike{
				PlayerId:        4,
				CurrentLocation: [2]int{1, 0},
				Direction:       "E",
				Trail:           [][2]int{{9, 0}, {8, 0}},
			},
		},
	}
}

func getGrid() string {
	grid := GetGrid()
	var jsonGrid, _ = json.Marshal(grid)
	return string(jsonGrid)
}

func getGames() string {
	games := []mqtt.Game{
		mqtt.Game{
			Name:          "1",
			ActivePlayers: 3,
		},
		mqtt.Game{
			Name:          "2",
			ActivePlayers: 2,
		},
		mqtt.Game{
			Name:          "3",
			ActivePlayers: 5,
		},
	}
	var jsonGames, _ = json.Marshal(games)
	return string(jsonGames)
}
