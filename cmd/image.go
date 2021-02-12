package main

import (
	"flag"
	"image"
	"os"

	"jphsd.org/glui"

	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"
	_ "image/jpeg"
	_ "image/png"
)

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
	f.Close()
	size := img.Bounds().Size()

	glui.NewGLWin(size.X, size.Y, "Image", img, true)
	glui.Loop(nil)
}
