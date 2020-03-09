package main

import (
	"fmt"
	"log"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

var (
	data  = rpio.Pin(17)
	clock = rpio.Pin(22)
	latch = rpio.Pin(27)
)

func main() {
	if err := rpio.Open(); err != nil {
		log.Fatal(err)
	}

	data.Output()
	data.Low()
	clock.Output()
	clock.Low()
	latch.Output()
	latch.Low()

	sendData(0x09) // address
	sendData(0x00) // no decode
	pulseLatch()

	// set intensity
	sendData(0x0A) // address
	sendData(0x02) // 9/32
	pulseLatch()

	// set scan limit 0-7
	sendData(0x0B) // address
	sendData(0x07) // 8 digits
	pulseLatch()

	// set for normal operation
	sendData(0x0C) // address
	sendData(0x01) // On
	pulseLatch()

	fmt.Println("incepe scrisul: ")

	for n := 0; n <= 8; n++ {
		sendData(n)
		sendData(0)
		pulseLatch()
	}

	// a := 0b00010000
	// fmt.Println(a)

	data := [8]int{16, 56, 124, 56, 124, 254, 56, 0}

	for n := 0; n < len(data); n++ {
		fmt.Println(data[n])
		writeMAX7219(data[n], n+1)
	}

	time.Sleep(5 * time.Second)

	// set for normal operation
	sendData(0x0C) // address
	sendData(0x00) // On
	pulseLatch()

	fmt.Println("gata")
}

func pulseClock() {
	clock.High()
	clock.Low()
}

func pulseLatch() {
	latch.High()
	latch.Low()
}

func sendData(value int) {
	for n := 0; n < 8; n++ {
		temp := value & 0x80
		if temp == 0x80 {
			data.High()
		} else {
			data.Low()
		}
		pulseClock()
		value = value << 0x01
	}
}

func writeMAX7219(data int, location int) {
	sendData(location)
	sendData(data)
	pulseLatch()
}
