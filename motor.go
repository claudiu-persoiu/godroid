package main

import (
	"fmt"
        "os"
        "time"
        "github.com/stianeikeland/go-rpio/v4"
)

func main() {
        err := rpio.Open()
        if err != nil {
                os.Exit(1)
        }
        defer rpio.Close()
	 
	out := rpio.Pin(23)
	out.Output()
	out.High()

        pin := rpio.Pin(18)
        pin.Mode(rpio.Pwm)
        pin.Freq(200)
        pin.DutyCycle(0, 32)
        // the LED will be blinking at 2000Hz
        // (source frequency divided by cycle length => 64000/32 = 2000)

        // five times smoothly fade in and out
        for i := 0; i < 5; i++ {
                for i := uint32(16); i < 32; i++ { // increasing brightness
			fmt.Println(i)
			pin.DutyCycle(i, 32)
                        time.Sleep(time.Second/16)
                }
                for i := uint32(32); i > 16; i-- { // decreasing brightness
			fmt.Println(i)
                        pin.DutyCycle(i, 32)
                        time.Sleep(time.Second/16)
                }
        }

	out.Low()
}
