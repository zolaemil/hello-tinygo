# BBC Micro:bit

Experiments with running go on BBC micro:bit

## Blink

This is an extended version of the [official example](https://github.com/tinygo-org/tinygo/tree/master/src/examples/microbit-blink) of Tinygo for the Micro:bit. 

It blinks the entire array of built-in LEDs on the board.

### Usage

Build and deploy:

`tinygo build -o=/Volumes/MICROBIT/blink.hex -target=microbit examples/microbit-blink`

(Replace `Volumes/MICROBIT` with the path to the mounted Micro:bit on your system)
