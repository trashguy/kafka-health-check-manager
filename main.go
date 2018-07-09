package main

import (
	"github.com/kataras/iris"
	"flag"
	"strings"
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type Partition struct {
	Status string
	OSR []int
}

type Topic struct {
	Topic string
	Status string
	Partitions map[string]Partition
}

type BrokerStatus struct {
	Broker string
	Status string
	Topics []Topic
}

var brokersFlag string
var refreshFlag int

// TODO Config, auto parse add http if does not exist. Recognize pattern for brokers like kafka1-30 auto add ports
// default is 8000

// TODO Add notifer class for slack.

// TODO Add logging metrics for polling. Store in small DB

func main() {
	// TODO Switch form cli to config. Flag to point to config file
	flag.StringVar(&brokersFlag,"brokers", "http://localhost:8000,http://localhost:8000","comma seprated list of brokers")
	flag.IntVar(&refreshFlag, "refresh", 5, "Time to poll brokers")
	flag.Parse()

	app := iris.New()
	app.RegisterView(iris.Handlebars("./templates", ".html"))

	app.Get("/", func(ctx iris.Context) {
		ctx.ViewData("brokers", getClusterHealth())
		ctx.View("index.html")
	})

	app.Get("/broker/{id:int}", func(ctx iris.Context) {
		ctx.ViewData("broker", getBrokerInfo(ctx.Params().Get("id")))
		ctx.View("broker.html")
	})

	app.Run(iris.Addr(":8080"))

}

func getBrokerPath(brokerId int) string {

}

func getBrokerInfo(brokerId int) BrokerStatus {
	var status BrokerStatus
	response, err := http.Get(broker + "/cluster")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		json.Unmarshal([]byte(data), &status)
		status.Broker = broker
	}
	return status
}

func getClusterHealth() []BrokerStatus {
	brokers := strings.Split(brokersFlag, ",")
	var clusterStatus []BrokerStatus
	for _, broker := range brokers {
		var status BrokerStatus
		response, err := http.Get(broker)
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			json.Unmarshal([]byte(data), &status)
			status.Broker = broker
			clusterStatus = append(clusterStatus, status)
		}
	}
	return clusterStatus
}
