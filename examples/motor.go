package main

import (
	"os"
	"time"

	"github.com/claudiu-persoiu/godroid/motor"
	"github.com/stianeikeland/go-rpio/v4"
)

func main() {
	err := rpio.Open()
	if err != nil {
		os.Exit(1)
	}
	defer func() {
		rpio.Close()
	}()

	motor := motor.NewMotor(rpio.Pin(23), rpio.Pin(24), rpio.Pin(18))

	// motor2 := motor.NewMotor(rpio.Pin(19), rpio.Pin(26), rpio.Pin(13))

	motor.Forward(1)
	time.Sleep(time.Second * 2)
	motor.Backword(1)
	time.Sleep(time.Second * 2)
	motor.Stop()

	// motor2.Forward(1)
	// time.Sleep(time.Second * 2)
	// motor2.Backword(1)
	// time.Sleep(time.Second * 2)
	// motor2.Stop()
}
