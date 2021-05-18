package examples

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
	if err := rpio.Open(); err != nil {
		log.Fatal(err)
	}

	matrix := max7219.NewDisplay(rpio.Pin(17), rpio.Pin(22), rpio.Pin(27), 1)

	matrix.ShowSymbol("tree")
	matrix.ShowSymbol("smile")
	matrix.ShowSymbol("sad")
	matrix.ShowSymbol("heart")

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter your text: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text == "" {
			fmt.Println("Bu Bye now")
			break
		}

		matrix.ShowText(text)
	}
}
