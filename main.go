package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"strconv"

	"github.com/claudiu-persoiu/godroid/max7219"
	"github.com/claudiu-persoiu/godroid/motor"
	"github.com/claudiu-persoiu/godroid/server"
	"github.com/stianeikeland/go-rpio/v4"
)

var motorLeft *motor.Motor
var motorRight *motor.Motor
var matrix *max7219.Display

func main() {
	address := flag.String("address", ":3333", "server address")
	noGpio := flag.Bool("noGpio", false, "disable gpio init")
	flag.Parse()

	if !*noGpio {
		err := rpio.Open()
		if err != nil {
			log.Fatal("could not initiate GPIO: ", err)
		}
		defer func() {
			rpio.Close()
		}()
		motorLeft = motor.NewMotor(rpio.Pin(24), rpio.Pin(23), rpio.Pin(18))
		motorRight = motor.NewMotor(rpio.Pin(26), rpio.Pin(19), rpio.Pin(13))
		matrix = max7219.NewDisplay(rpio.Pin(17), rpio.Pin(22), rpio.Pin(27), 1)
	}

	fmt.Println("Starting server: " + *address)
	if err := server.StartServer(*address, callback); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func callback(msg server.Message) {
	switch msg.Action {
	case "left":
		dataToAction(msg.Data, motorLeft)
	case "right":
		dataToAction(msg.Data, motorRight)
	case "text":
		matrix.ShowText(msg.Data)
	case "symbol":
		matrix.ShowSymbol(msg.Data)
	}
}

func dataToAction(data string, motor *motor.Motor) {
	if motor == nil {
		return
	}

	if data == "stop" {
		motor.Stop()
		return
	}

	var arr []string

	if err := json.Unmarshal([]byte(data), &arr); err != nil {
		log.Println("Error parcing message:")
		log.Fatal(err)
	}

	speed, _ := strconv.Atoi(arr[1])

	switch arr[0] {
	case "up":
		motor.Forward(speed)
	case "down":
		motor.Backword(speed)
	}
}
