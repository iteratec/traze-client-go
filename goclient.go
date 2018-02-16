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
)

func main(){
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	initConfig()

	go mock.Run()

	go mqtt.ReadFromTopic(mock.MOCK_TOPIC)

	go func() {
		signal := <-sigs
		fmt.Printf("Got os signal: %v\n", signal)
		done <- true
	}()

	<-done
	fmt.Println("Exiting.")
}

func initConfig(){
	viper.SetConfigName("goclient")
	viper.AddConfigPath("/etc/goclient")
	viper.AddConfigPath("./conf/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}