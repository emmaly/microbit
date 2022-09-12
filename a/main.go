package main

import (
	"image/color"
	"machine"
	"math/rand"
	"time"

	"github.com/emmaly/microbit/a/sprite"
	"tinygo.org/x/drivers/mag3110"
	"tinygo.org/x/drivers/microbitmatrix"
)

var framerate = 800 * time.Millisecond

func main() {
	machine.I2C0.Configure(machine.I2CConfig{})

	// display config
	var display = microbitmatrix.New()
	display.Configure(microbitmatrix.Config{Rotation: 1})
	sprite.Draw(&display, sprite.SPRITE_O)

	// mag3110 config
	var mag = mag3110.New(machine.I2C0)
	mag.Configure()

	last := time.Now()
	for {
		if time.Since(last) >= framerate {
			last = time.Now()
			displayHeight, displayWidth := display.Size()
			x := rand.Intn(int(displayHeight))
			y := rand.Intn(int(displayWidth))
			togglePixel(&display, x, y)
		}

		display.Display()
	}
}

func togglePixel(display *microbitmatrix.Device, x int, y int) {
	if !display.GetPixel(int16(x), int16(y)) {
		display.SetPixel(int16(x), int16(y), color.RGBA{255, 255, 255, 255})
	} else {
		display.SetPixel(int16(x), int16(y), color.RGBA{0, 0, 0, 0})
	}
}
