package mock

import (
	"fmt"
	"iteragit.iteratec.de/traze/goclient/mqtt"
	"time"
	"encoding/json"
)

var MOCK_TOPIC string

func init() {
	MOCK_TOPIC = "traze/mockedgame/grid"
}

func Run() {

	client := mqtt.GetClient()

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	fmt.Println("Mock Publisher Started.")
	for {
		grid := mqtt.Grid{
			Height: 3,
			Width:  3,
			Tiles: [][]int{
				{0, 1, 0},
				{0, 2, 0},
				{0, 2, 0},
			},
		}
		var jsonGrid, _ = json.Marshal(grid)
		fmt.Printf("Publishing on topic '%v' now: %v\n", MOCK_TOPIC, grid)
		token := client.Publish(MOCK_TOPIC, byte(mqtt.GetQos()), false, string(jsonGrid))
		token.Wait()

		time.Sleep(500 * time.Millisecond)
	}

	client.Disconnect(250)
	fmt.Println("Sample Publisher Disconnected")

}
