// blink program for the BBC micro:bit based on the microbit-blink example
// of the tinygo repo https://github.com/tinygo-org/tinygo/tree/master/src/examples/microbit-blink
package main

import (
	"machine"
	"time"
)

// The LED matrix in the micro:bit is a multiplexed display: https://en.wikipedia.org/wiki/Multiplexed_display
// Driver for easier control: https://github.com/tinygo-org/drivers/tree/master/microbitmatrix

var ledcols [9]machine.Pin = [9]machine.Pin{
	machine.LED_COL_1,
	machine.LED_COL_2,
	machine.LED_COL_3,
	machine.LED_COL_4,
	machine.LED_COL_5,
	machine.LED_COL_6,
	machine.LED_COL_7,
	machine.LED_COL_8,
	machine.LED_COL_9,
}

var ledrows [3]machine.Pin = [3]machine.Pin{
	machine.LED_ROW_1,
	machine.LED_ROW_2,
	machine.LED_ROW_3,
}

func confCols() {
	for _, ledcol := range ledcols {
		ledcol.Configure(machine.PinConfig{Mode: machine.PinOutput})
		ledcol.Low()
	}
}

func confRows() {
	for _, ledrow := range ledrows {
		ledrow.Configure(machine.PinConfig{Mode: machine.PinOutput})
	}
}

func on() {
	for _, ledrow := range ledrows {
		ledrow.High()
	}
}

func off() {
	for _, ledrow := range ledrows {
		ledrow.Low()
	}
}

func main() {
	confCols()
	confRows()

	for {
		on()
		time.Sleep(time.Millisecond * 500)

		off()
		time.Sleep(time.Millisecond * 200)
	}
}
