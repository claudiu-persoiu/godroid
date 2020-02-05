package main

import (
	"os"
	"time"

	"github.com/claudiu-persoiu/godroid/gpio"
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

	motor := gpio.NewMotor(rpio.Pin(23), rpio.Pin(24), rpio.Pin(18))

	motor2 := gpio.NewMotor(rpio.Pin(19), rpio.Pin(26), rpio.Pin(13))

	
	// out := rpio.Pin(23)
	// out.Output()
	// out.High()

	// pin := rpio.Pin(18)
	// pin.Mode(rpio.Pwm)
	// pin.Freq(200)
	// pin.DutyCycle(0, 32)
	// // the LED will be blinking at 2000Hz
	// // (source frequency divided by cycle length => 64000/32 = 2000)

	// // five times smoothly fade in and out
	// for i := 0; i < 5; i++ {
	// 	for i := uint32(16); i < 32; i++ { // increasing brightness
	// 		fmt.Println(i)
	// 		pin.DutyCycle(i, 32)
	// 		time.Sleep(time.Second / 16)
	// 	}
	// 	for i := uint32(32); i > 16; i-- { // decreasing brightness
	// 		fmt.Println(i)
	// 		pin.DutyCycle(i, 32)
	// 		time.Sleep(time.Second / 16)
	// 	}
	// }

	// out.Low()

	motor.Forward(1)
	time.Sleep(time.Second * 2)
	motor.Backword(1)
	time.Sleep(time.Second * 2)
	motor.Stop()

	motor2.Forward(1)
	time.Sleep(time.Second * 2)
	motor2.Backword(1)
	time.Sleep(time.Second * 2)
	motor2.Stop()
}
