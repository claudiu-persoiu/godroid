package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/claudiu-persoiu/godroid/gpio"
	"github.com/claudiu-persoiu/godroid/server"
	"github.com/stianeikeland/go-rpio/v4"
)

func main() {
	address := flag.String("address", ":3333", "server address")
	noGpio := flag.Bool("noGpio", false, "disable gpio init")
	flag.Parse()

	var motorLeft gpio.Motor
	var motorRight gpio.Motor

	if !*noGpio {
		err := rpio.Open()
		if err != nil {
			log.Fatal("could not initiate GPIO: ", err)
		}
		defer func() {
			rpio.Close()
		}()
		motorLeft = gpio.NewRealMotor(rpio.Pin(24), rpio.Pin(23), rpio.Pin(18))
		motorRight = gpio.NewRealMotor(rpio.Pin(26), rpio.Pin(19), rpio.Pin(13))
	} else {
		motorLeft = gpio.NewFakeMotor("left")
		motorRight = gpio.NewFakeMotor("right")
	}

	fmt.Println("Starting server: " + *address)
	if err := server.StartServer(*address, motorLeft, motorRight); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
