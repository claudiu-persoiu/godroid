package max7219

import (
	"fmt"
	"strings"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

type Display struct {
	data  rpio.Pin
	clock rpio.Pin
	latch rpio.Pin
	dir   int
}

func NewDisplay(data, clock, latch rpio.Pin, dir int) *Display {
	data.Output()
	data.Low()
	clock.Output()
	clock.Low()
	latch.Output()
	latch.Low()

	d := &Display{data, clock, latch, dir}

	// https://datasheets.maximintegrated.com/en/ds/MAX7219-MAX7221.pdf
	// set decode mode
	d.sendDataInt(0x09, 0)

	// set led intensity
	d.sendDataInt(0x0A, 2)

	// set scan limit 0-7
	d.sendDataInt(0x0B, 0x07)

	return d
}

func (d *Display) ShowText(text string) error {

	text = strings.TrimSpace(strings.ToLower(text))
	var data = [8]string{}

	for _, r := range []rune(text) {
		if alphabet[r] != nil {
			for n := 0; n < len(alphabet[r]); n++ {
				data[n] = data[n] + alphabet[r][n]
			}
		}
	}

	for n := 0; n < len(data); n++ {
		data[n] = "00000000" + data[n] + "00000000"
	}

	d.displayData(data, time.Millisecond*300)
	return nil
}

func (d *Display) ShowSymbol(symbol string) error {
	if symbol, ok := shapesMap[symbol]; ok {
		d.displayData(symbol, time.Second*2)
	}
	return nil
}

func (d *Display) displayData(data [8]string, timeout time.Duration) {
	// turn on
	d.sendDataInt(0x0C, 1)

	var matrixToDisplay = [8]string{}
	for o := 0; o <= len(data[0])-8; o++ {
		for n := 0; n < len(data); n++ {
			matrixToDisplay[n] = data[n][o:(o + 8)]
		}
		d.displayMatrix(matrixToDisplay)
		time.Sleep(timeout)
	}

	// turn off
	d.sendDataInt(0x0C, 0)
}

func (d *Display) displayMatrix(matrix [8]string) {
	for i := 0; i < d.dir; i++ {
		matrix = transposeMatrix(matrix)
	}

	for i := 0; i < len(matrix); i++ {
		d.processData(fmt.Sprintf("%08b%s", i+1, matrix[i]))
	}
}

func (d *Display) sendDataInt(address, value int) {
	d.processData(fmt.Sprintf("%08b%08b", address, value))
}

func transposeMatrix(matrix [8]string) [8]string {
	newMatrix := [8]string{}

	for i := 0; i < len(matrix); i++ {
		newMatrix[i] = "00000000"
	}

	for i := 0; i < len(matrix); i++ {
		for j, r := range matrix[i] {
			if r == '1' {
				newMatrix[7-j] = newMatrix[7-j][:i] + "1" + newMatrix[7-j][i+1:]
			}
		}
	}

	return newMatrix
}

func (d *Display) processData(value string) {
	fmt.Println(value)
	for _, r := range []rune(value) {
		if r == '1' {
			d.data.High()
		} else {
			d.data.Low()
		}
		d.clock.High()
		d.clock.Low()
	}
	d.latch.High()
	d.latch.Low()
}

func intToBinStr(value int) string {
	return fmt.Sprintf("%08b", value)
}

var shapesMap = map[string][8]string{
	"tree": {
		"00010000",
		"00111000",
		"01111100",
		"00111000",
		"01111100",
		"11111110",
		"00111000",
		"00000000",
	},
	"smile": {
		"00111100",
		"01000010",
		"10100101",
		"10000001",
		"10100101",
		"10011001",
		"01000010",
		"00111100",
	},
	"sad": {
		"00111100",
		"01000010",
		"10100101",
		"10000001",
		"10011001",
		"10100101",
		"01000010",
		"00111100",
	},
	"neutral": {
		"00111100",
		"01000010",
		"10100101",
		"10000001",
		"10000001",
		"10111101",
		"01000010",
		"00111100",
	},
	"heart": {
		"01100110",
		"11111111",
		"11111111",
		"11111111",
		"01111110",
		"00111100",
		"00011000",
		"00000000",
	},
}

var alphabet = map[rune][]string{
	'a': {
		"00000",
		"01100",
		"10010",
		"10010",
		"11110",
		"10010",
		"10010",
		"00000",
	},
	'b': {
		"00000",
		"11100",
		"10010",
		"11100",
		"10010",
		"10010",
		"11100",
		"00000",
	},
	'c': {
		"00000",
		"01100",
		"10010",
		"10000",
		"10000",
		"10010",
		"01100",
		"00000",
	},
	'd': {
		"00000",
		"11100",
		"10010",
		"10010",
		"10010",
		"10010",
		"11100",
		"00000",
	},
	'e': {
		"0000",
		"1110",
		"1000",
		"1110",
		"1000",
		"1000",
		"1110",
		"0000",
	},
	'f': {
		"0000",
		"1110",
		"1000",
		"1110",
		"1000",
		"1000",
		"1000",
		"0000",
	},
	'g': {
		"00000",
		"01100",
		"10010",
		"10000",
		"10110",
		"10010",
		"01100",
		"00000",
	},
	'h': {
		"00000",
		"10010",
		"10010",
		"11110",
		"10010",
		"10010",
		"10010",
		"00000",
	},
	'i': {
		"00",
		"10",
		"00",
		"10",
		"10",
		"10",
		"10",
		"00",
	},
	'j': {
		"00000",
		"00010",
		"00010",
		"00010",
		"00010",
		"10010",
		"01100",
		"00000",
	},
	'k': {
		"00000",
		"10010",
		"10100",
		"11000",
		"10100",
		"10010",
		"10010",
		"00000",
	},
	'l': {
		"0000",
		"1000",
		"1000",
		"1000",
		"1000",
		"1000",
		"1110",
		"0000",
	},
	'm': {
		"000000",
		"100010",
		"110110",
		"101010",
		"100010",
		"100010",
		"100010",
		"000000",
	},
	'n': {
		"00000",
		"10010",
		"11010",
		"10110",
		"10010",
		"10010",
		"10010",
		"00000",
	},
	'o': {
		"00000",
		"01100",
		"10010",
		"10010",
		"10010",
		"10010",
		"01100",
		"00000",
	},
	'p': {
		"00000",
		"11100",
		"10010",
		"10010",
		"11100",
		"10000",
		"10000",
		"00000",
	},
	'q': {
		"00000",
		"01100",
		"10010",
		"10010",
		"10010",
		"10010",
		"01110",
		"00000",
	},
	'r': {
		"00000",
		"11100",
		"10010",
		"10010",
		"11100",
		"10100",
		"10010",
		"00000",
	},
	's': {
		"00000",
		"01110",
		"10000",
		"01100",
		"00010",
		"10010",
		"01100",
		"00000",
	},
	't': {
		"000000",
		"111110",
		"001000",
		"001000",
		"001000",
		"001000",
		"001000",
		"000000",
	},
	'u': {
		"00000",
		"10010",
		"10010",
		"10010",
		"10010",
		"10010",
		"01100",
		"00000",
	},
	'v': {
		"000000",
		"100010",
		"100010",
		"100010",
		"100010",
		"010100",
		"001000",
		"000000",
	},
	'w': {
		"000000",
		"100010",
		"100010",
		"100010",
		"101010",
		"110110",
		"100010",
		"000000",
	},
	'x': {
		"00000",
		"10010",
		"10010",
		"01100",
		"01100",
		"10010",
		"10010",
		"00000",
	},
	'y': {
		"000000",
		"100010",
		"100010",
		"011100",
		"001000",
		"001000",
		"001000",
		"000000",
	},
	'z': {
		"00000",
		"11110",
		"00100",
		"01000",
		"10000",
		"10000",
		"11110",
		"00000",
	},
	' ': {
		"00",
		"00",
		"00",
		"00",
		"00",
		"00",
		"00",
		"00",
	},
	'.': {
		"00",
		"00",
		"00",
		"00",
		"00",
		"00",
		"10",
		"00",
	},
	'!': {
		"00",
		"10",
		"10",
		"10",
		"10",
		"00",
		"10",
		"00",
	},
	'?': {
		"00000",
		"01100",
		"10010",
		"00100",
		"01000",
		"00000",
		"01000",
		"00000",
	},
}
