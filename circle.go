// Make an image of a circle.

// First let's make a blank image
package main

//import "io"
import "os"
import "log"
import "image"
import "image/color"
import "image/png"
import "math"

const width, height = 512, 512

func savePng (img image.Image, fname string) {
	// Open the file
	fo,err := os.Create(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = fo.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// Write the image
	if err = png.Encode(fo, img); err != nil {
		log.Fatal(err)
	}
}

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

	savePng (img, "circle.png")
}
