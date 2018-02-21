package mqtt

import (
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func ReadFromTopic(topic string) {

	fmt.Printf("ReadFromTopic: %v\n", topic)

	client := GetClient()
	defer client.Disconnect(250)

	msgQueue := make(chan [2]string)

	fmt.Printf("Subscribe on topic: %v\n", topic)
	client.Subscribe(topic, byte(GetQos()), func(i MQTT.Client, msg MQTT.Message) {
		msgQueue <- [2]string{msg.Topic(), string(msg.Payload())}
	})

	//TODO: use token to secure connection
	//if token := client.Connect(); token.Wait() && token.Error() != nil {
	//	panic(token.Error())
	//}
	//
	//if token := client.Subscribe(topic, byte(GetQos()), nil); token.Wait() && token.Error() != nil {
	//	fmt.Println(token.Error())
	//	os.Exit(1)
	//}

	fmt.Printf("Waiting for incoming messages on topic '%v'.\n", topic)
	for {
		incoming := <-msgQueue
		fmt.Printf("RECEIVED TOPIC: %s MESSAGE: %s\n", incoming[0], incoming[1])
	}

}
