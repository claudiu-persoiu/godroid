package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/claudiu-persoiu/godroid/max7219"
	"github.com/stianeikeland/go-rpio/v4"
)

func main() {
	// fmt.Printf("%08b", 0)
	// println()
	// a := fmt.Sprintf("%08b", 10)

	// for i, r := range []rune(a) {
	// 	fmt.Printf("i%d r %c\n", i, r)
	// }

	if err := rpio.Open(); err != nil {
		log.Fatal(err)
	}

	matrix := max7219.NewDisplay(rpio.Pin(17), rpio.Pin(22), rpio.Pin(27), 1)

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter your text: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text == "" {
			fmt.Println("Bu Bye now")
			break
		}

		matrix.Show(text)
	}
}
