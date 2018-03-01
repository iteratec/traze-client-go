package mqtt

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/spf13/viper"
	"iteragit.iteratec.de/traze/goclient/util/log"
)

var broker string
var client MQTT.Client
var store string
var qos byte

const (
	id = "GO_TRAZE_CLIENT"
)

func Qos() byte {
	return qos
}

func InitClient() {
	broker = viper.GetString("broker")
	//password := flag.String("password", "", "The password (optional)")
	//user := flag.String("user", "", "The User (optional)")
	//id := flag.String("id", "testgoid", "The ClientID (optional)")
	qos = 0
	store = ":memory:"

	client = MQTT.NewClient(getMqttOptions())
}

func GetClient() MQTT.Client {
	return client
}

func getMqttOptions() *MQTT.ClientOptions {
	opts := MQTT.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(id)
	opts.SetKeepAlive(5)
	//opts.SetUsername(*user)
	//opts.SetPassword(*password)
	opts.SetCleanSession(false)
	if store != ":memory:" {
		opts.SetStore(MQTT.NewFileStore(store))
	}
	log.Info.Printf("Initialize mqtt client with: broker=%v\n", broker)
	return opts
}
