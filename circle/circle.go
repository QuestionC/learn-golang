// Make an image of a circle.

// First let's make a blank image
package main

// import "io"
// import "log"
import "image"
import "image/color"
import "math"
import "learn-golang/util"

const width, height = 512, 512

func main() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			delta_x := x - width/2
			delta_y := y - height/2
			distance := math.Sqrt(float64(delta_x*delta_x + delta_y*delta_y))

			var c color.RGBA
			if distance > 49 && distance < 51 {
				c = color.RGBA {0, 0, 0, 255}
			} else {
				c = color.RGBA {255, 255, 255, 255}
			}
			img.Set(x, y, c)
		}
	}

	util.SavePng (img, "circle.png")
}
