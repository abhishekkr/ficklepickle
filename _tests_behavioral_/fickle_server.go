package main

import (
	"fmt"
	"time"

	"github.com/abhishekkr/ficklepickle"
	"github.com/abhishekkr/ficklepickle/config"
)

func init() {
	config.PickleDir = "./_tests_behavioral_/server"
}

func main() {
	go ficklepickle.StartTcpServer()
	time.Sleep(time.Second * 30)
	ficklepickle.CloseTcpServer()
	fmt.Println("[+] done.")
}
