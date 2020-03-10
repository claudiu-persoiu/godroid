package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/claudiu-persoiu/godroid/motor"
	"github.com/claudiu-persoiu/godroid/server"
	"github.com/stianeikeland/go-rpio/v4"
)

func main() {
	address := flag.String("address", ":3333", "server address")
	noGpio := flag.Bool("noGpio", false, "disable gpio init")
	flag.Parse()

	var motorLeft motor.Motor
	var motorRight motor.Motor

	if !*noGpio {
		err := rpio.Open()
		if err != nil {
			log.Fatal("could not initiate GPIO: ", err)
		}
		defer func() {
			rpio.Close()
		}()
		motorLeft = motor.NewRealMotor(rpio.Pin(24), rpio.Pin(23), rpio.Pin(18))
		motorRight = motor.NewRealMotor(rpio.Pin(26), rpio.Pin(19), rpio.Pin(13))
	} else {
		motorLeft = motor.NewFakeMotor("left")
		motorRight = motor.NewFakeMotor("right")
	}

	fmt.Println("Starting server: " + *address)
	if err := server.StartServer(*address, motorLeft, motorRight); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
