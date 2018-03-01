package mqtt

import (
	"encoding/json"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"iteragit.iteratec.de/traze/goclient/util/log"
)

var gridQueue chan Grid
var topicsToSubscribe map[string]byte

func init() {
	gridQueue = make(chan Grid)
	topicsToSubscribe = map[string]byte{
		"traze/mockedgame/grid": Qos(),
		"traze/games":           Qos(),
	}
}

func HandleTrazeEvents() {

	client := GetClient()
	defer client.Disconnect(250)

	msgQueue := make(chan [2]string)

	client.SubscribeMultiple(topicsToSubscribe, func(i MQTT.Client, msg MQTT.Message) {
		msgQueue <- [2]string{msg.Topic(), string(msg.Payload())}
	})

	//TODO: use token to secure connection
	//if token := client.Connect(); token.Wait() && token.Error() != nil {
	//	panic(token.Error())
	//}
	//
	//if token := client.Subscribe(topic, Qos(), nil); token.Wait() && token.Error() != nil {
	//	fmt.Println(token.Error())
	//	os.Exit(1)
	//}

	for {
		incoming := <-msgQueue
		log.Info.Printf("TOPIC '%s' -> received MESSAGE: %s\n", incoming[0], incoming[1])
		if incoming[0] == "traze/mockedgame/grid" {
			grid := Grid{}
			err := json.Unmarshal([]byte(incoming[1]), &grid)
			if err != nil {
				log.Error.Printf("The following error occurred unmarshalling grid json: %v\n", err)
			}
			gridQueue <- grid
		}
	}

}

func GridQueue() chan Grid {
	return gridQueue
}
