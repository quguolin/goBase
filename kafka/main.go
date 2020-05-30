package main

import (
	"encoding/json"
	"fmt"

	"github.com/Shopify/sarama"
)


func main()  {
	addrs := []string{
		"kafka1:9092",
		"kafka2:9093",
		"kafka3:9094",
	}
	config := sarama.NewConfig()
	config.Version = sarama.V2_5_0_0
	admin,err := sarama.NewClusterAdmin(addrs,config)
	if err != nil{
		panic(err)
	}
	topics,err := admin.ListTopics()
	if err != nil{
		panic(err)
	}
	bs,err := json.Marshal(topics)
	if err != nil{
		panic(err)
	}
	fmt.Println(string(bs))
}

