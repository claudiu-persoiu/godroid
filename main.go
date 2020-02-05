package main

import (
	"fmt"
	"log"

	"github.com/claudiu-persoiu/godroid/gpio"
	"github.com/claudiu-persoiu/godroid/server"
)

func main() {
	address := ":3333"

	motorLeft := gpio.NewMotor(24, 23, 18)
	motorRight := gpio.NewMotor(26, 19, 13)

	fmt.Println("Starting server: " + address)
	err := server.StartServer(address)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
