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
	f, err := os.Open(args[0])
	if err != nil {
		panic(err)
	}
	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}
	_ = f.Close()
	size := img.Bounds().Size()

	// Create a new window with the image
	glui.NewGLWin(size.X, size.Y, "Image", img, true)

	// Allow the system to render it
	glui.Loop(nil)
}
