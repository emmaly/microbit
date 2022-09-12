package sprite

import (
	"image/color"

	"tinygo.org/x/drivers/microbitmatrix"
)

var (
	y = color.RGBA{255, 255, 255, 255}
	n = color.RGBA{0, 0, 0, 0}
)

var (
	SPRITE_UP = [5][5]color.RGBA{
		{n, n, y, n, n},
		{n, y, y, y, n},
		{y, n, y, n, y},
		{n, n, y, n, n},
		{n, n, y, n, n},
	}
	SPRITE_O = [5][5]color.RGBA{
		{n, n, n, n, n},
		{n, y, y, y, n},
		{n, y, n, y, n},
		{n, y, y, y, n},
		{n, n, n, n, n},
	}
	SPRITE_X = [5][5]color.RGBA{
		{n, n, n, n, n},
		{n, y, n, y, n},
		{n, n, y, n, n},
		{n, y, n, y, n},
		{n, n, n, n, n},
	}
)

func Draw(display *microbitmatrix.Device, sprite [5][5]color.RGBA) {
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			display.SetPixel(int16(x), int16(y), sprite[x][y])
		}
	}
}
