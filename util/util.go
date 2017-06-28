package util

import "image"
import "image/png"
import "os"
import "log"

func SavePng (img image.Image, fname string) {
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
