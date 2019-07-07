package main

import (
	"fmt"
	"log"
	"net/http"

	consulapi "github.com/hashicorp/consul/api"
)

const (
	_checkPort    = 8080
	_regPort      = 8081
	_regIp        = "127.0.0.1"
	_checkTimeout = "3s"
	_checkInter   = "5s"
	_checkAfter   = "30s"
	_regID        = "service_id_1"
	_regNodeName  = "service_node_1"
	_regTag       = "servicde"
)

func consulCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "consulCheck")
}

func registerServer() error {
	config := consulapi.DefaultConfig()
	client, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatal("consul client error : ", err)
		return err
	}
	checkPort := _checkPort

	registration := new(consulapi.AgentServiceRegistration)
	registration.ID = _regID
	registration.Name = _regNodeName
	registration.Port = _regPort
	registration.Tags = []string{_regTag}
	registration.Address = _regIp
	registration.Check = &consulapi.AgentServiceCheck{
		HTTP:                           fmt.Sprintf("http://%s:%d%s", registration.Address, checkPort, "/check"),
		Timeout:                        _checkTimeout,
		Interval:                       _checkInter,
		DeregisterCriticalServiceAfter: _checkAfter, //check失败后30秒删除本服务
	}

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatal("register server error : ", err)
		return err
	}
	http.HandleFunc("/check", consulCheck)
	if err = http.ListenAndServe(fmt.Sprintf(":%d", checkPort), nil); err != nil {
		return err
	}
	return nil
}

func main() {
	if err := registerServer(); err != nil {
		panic(err)
	}
}
