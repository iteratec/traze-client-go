package mqtt

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"fmt"
)

var broker string
var client MQTT.Client
var store string
var qos int

const(
	id = "GO_TRAZE_CLIENT"
)

func GetBroker() string{
	return broker
}
func GetQos() int{
	return qos
}

func init(){

	//broker = viper.GetString("broker")
	broker = "tcp://localhost:1883"
	fmt.Printf("broker=%v\n", broker)
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
	fmt.Printf("Initialize mqtt client with: broker=%v\n", broker)
	return opts
}
