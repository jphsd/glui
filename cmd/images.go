package main

import (
	"flag"
	"image"
	"os"
	"time"

	"github.com/go-gl/glfw/v3.3/glfw"

	"jphsd.org/glui"

	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"
	_ "image/jpeg"
	_ "image/png"
)

func main() {
	// Read in files indicated on command line
	flag.Parse()
	args := flag.Args()
	images := make([]image.Image, len(args))

	for i, str := range args {
		f, err := os.Open(str)
		if err != nil {
			panic(err)
		}
		img, _, err := image.Decode(f)
		if err != nil {
			panic(err)
		}
		f.Close()
		images[i] = img
	}

	win := glui.NewGLWin(800, 800, "Images", images[0], true)
	// Close the window if ESC pressed
	win.Win.SetKeyCallback(keyCallback)

	go update(win, images)

	glui.Loop(nil)
}

func keyCallback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	// Close window on ESC
	if key == glfw.KeyEscape {
		w.SetShouldClose(true)
	}
}

func update(w *glui.GLWin, images []image.Image) {
	n := len(images)
	c := 0
	d, _ := time.ParseDuration("10s")
	for true {
		time.Sleep(d)
		c++
		if c == n {
			c = 0
		}
		w.SetImage(images[c])
	}
}
