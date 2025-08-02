package main

import (
	"fmt"
	"github.com/jphsd/glui"
	"image"
)

// Example use of MouseMoveListener, MouseClickListener and MouseScrollListener.

func main() {
	width, height := 600, 600

	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{width, height}})
	win := glui.NewGLWin(width, height, "Mouse Click Test", img, true)
	glui.NewMouseMoveListener(win, func(pt []float64) {
		fmt.Printf("Mouse moved to %.2f,%.2f\n", pt[0], pt[1])
	})
	glui.NewMouseClickListener(win, glui.MouseButtonLeft, func(pt []float64) {
		fmt.Printf("Mouse clicked at %.2f,%.2f\n", pt[0], pt[1])
	})
	glui.NewMouseScrollListener(win, func(x, y float64) {
		fmt.Printf("Mouse scroll %.2f,%.2f\n", x, y)
	})
	glui.Loop(nil)
}
