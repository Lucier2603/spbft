package main

import (
	"csic/pbft/pbft"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
)


// curl -X POST 'http://127.0.0.1:8081/newTrx' -H 'Content-Type: application/json' -d '{"SequenceId": 99,"From": 999,"To": 1,"ViewId": 1,"MsgType": 19,"Digest": "xxxx"}'

func main() {

	// read config
	cfg_f, err := os.Open("cfg.yaml")
	if err != nil {
		panic(err)
	}

	bytes, err := ioutil.ReadAll(cfg_f)
	if err != nil {
		panic(err)
	}

	cfg := pbft.Config{}
	if err = yaml.Unmarshal(bytes, &cfg); err != nil {
		panic(err)
	}

	// init log
	pbft.InitializeLogger("default", "log4go.xml")

	// build node
	node := pbft.NewPbftNode(&cfg)

	if err = node.Start(); err != nil {
		panic(err)
	}


	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGKILL)
	s := <-c
	fmt.Println("Killed", s)
}

