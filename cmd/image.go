package main

import (
	"flag"
	"image"
	"os"

	"github.com/jphsd/glui"

	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"
	_ "image/jpeg"
	_ "image/png"
)

// Read in an image from the file name on the command line and display it in a window.
func main() {
	// Read in image file indicated in command line
	flag.Parse()
	args := flag.Args()
	img := readImage(args[0])

	// Create a new window with the image
	size := img.Bounds().Size()
	glui.NewGLWin(size.X, size.Y, args[0], img, true)

	// Allow the system to render it
	glui.Loop(nil)
}

func readImage(str string) image.Image {
	f, err := os.Open(str)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return img
}
