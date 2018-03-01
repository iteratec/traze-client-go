package main

import (
	"fmt"
	"github.com/spf13/viper"
	"iteragit.iteratec.de/traze/goclient/mqtt"
	"iteragit.iteratec.de/traze/goclient/mqtt/mock"
	"iteragit.iteratec.de/traze/goclient/view"
	"sync"
)

var wg sync.WaitGroup

func main() {

	defer view.CloseTermbox()

	initConfig()
	mqtt.InitClient()

	go mqtt.HandleTrazeEvents()

	wg.Add(1)
	go view.InitBoard(wg)

	go mock.Run()

	wg.Wait()
}

func initConfig() {
	viper.AddConfigPath("/etc/goclient")
	viper.AddConfigPath("./conf/")
	viper.SetConfigName("goclient")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}