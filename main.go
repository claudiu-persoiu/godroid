package main

import (
	"fmt"
	"log"

	"github.com/claudiu-persoiu/godroid/gpio"
	"github.com/claudiu-persoiu/godroid/server"
	"github.com/stianeikeland/go-rpio/v4"
)

func main() {
	address := ":3333"

	gpio.NewMotor(rpio.Pin(23), rpio.Pin(16), rpio.Pin(18))

	fmt.Println("Starting server: " + address)
	err := server.StartServer(address)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
