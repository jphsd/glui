package main

import (
	"flag"
	"image"
	"os"
	"time"

	"github.com/go-gl/glfw/v3.3/glfw"

	"github.com/jphsd/glui"

	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"
	_ "image/jpeg"
	_ "image/png"
)

// A more sophisticated version of image.go that takes a collection of images and cycles
// through them.
func main() {
	// Read in files indicated on command line
	flag.Parse()
	args := flag.Args()
	images := make([]image.Image, len(args))

	for i, str := range args {
		images[i] = readImage(str)
	}

	// Create a new window to display the images and set a callback
	win := glui.NewGLWin(800, 800, "Images", images[0], true)
	win.Win.SetKeyCallback(keyCallback)

	// Launch the image updater
	go update(win, images)

	// Allow the system to render it
	glui.Loop(nil)
}

func keyCallback(w *glfw.Window, key glfw.Key, _ int, _ glfw.Action, _ glfw.ModifierKey) {
	// Close window on ESC key pressed
	if key == glfw.KeyEscape {
		w.SetShouldClose(true)
	}
}

func update(w *glui.GLWin, images []image.Image) {
	n := len(images)
	c := 0
	d, _ := time.ParseDuration("5s")
	for true {
		time.Sleep(d)
		c++
		if c == n {
			c = 0
		}
		w.SetImage(images[c])
	}
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
