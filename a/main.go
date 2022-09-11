package main

import (
	"image/color"
	"math/rand"
	"time"

	"tinygo.org/x/drivers/microbitmatrix"
)

var display microbitmatrix.Device

func main() {
	display := microbitmatrix.New()
	display.Configure(microbitmatrix.Config{})
	display.ClearDisplay()

	xs := []int16{0, 1, 2, 3, 4}
	ys := []int16{0, 1, 2, 3, 4}

	off := color.RGBA{0, 0, 0, 0}
	on := color.RGBA{255, 255, 255, 255}

	last := time.Now()

	for {
		if time.Since(last) > 100*time.Millisecond {
			last = time.Now()
			x := xs[rand.Intn(len(xs))]
			y := ys[rand.Intn(len(ys))]
			if pixelIsOn := display.GetPixel(x, y); pixelIsOn == true {
				display.SetPixel(x, y, off)
			} else {
				display.SetPixel(x, y, on)
			}
		}
		display.Display()
	}
}
