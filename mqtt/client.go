package mqtt

import (
	"encoding/json"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"iteragit.iteratec.de/traze/goclient/util/log"
	"github.com/spf13/viper"
)

var (
	gamesTopicName string

	gridQueue         chan Grid
	gamesQueue        chan []Game

	topicsToSubscribe map[string]byte
)

func initTopics() {

	gamesTopicName = viper.GetString("topics.games")
	log.Info.Printf("all conf keys: %v\n", viper.AllKeys())

	gridQueue = make(chan Grid)
	topicsToSubscribe = map[string]byte{
		"traze/1/grid": Qos(),
		gamesTopicName:  Qos(),
	}
	gamesQueue = make(chan []Game)
}

func HandleTrazeEvents() {

	initTopics()

	client := GetClient()
	defer client.Disconnect(250)

	msgQueue := make(chan [2]string)

	log.Info.Printf("subscribe to following topics: %v\n", topicsToSubscribe)

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
		switch incoming[0] {
		case gamesTopicName:
			games := []Game{}
			unmarshallJson(incoming[1], &games)
			log.Info.Printf("unmarshalled games: %v\n", games)
			gamesQueue <- games
		case "traze/1/grid":
			grid := Grid{}
			unmarshallJson(incoming[1], &grid)
			gridQueue <- grid
		}
	}

}

func unmarshallJson(fromMqtt string, target interface{}) {
	err := json.Unmarshal([]byte(fromMqtt), &target)
	if err != nil {
		log.Error.Printf("An error occurred unmarshalling json from mqtt topic:\n\tjson=%v\n\terr=%v\n", fromMqtt, err)
	}
}

func GridQueue() chan Grid {
	return gridQueue
}
func GamesQueue() chan []Game{
	return gamesQueue
}