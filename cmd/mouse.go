package main

import (
	"fmt"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/jphsd/glui"
	"image"
)

// Example use of MouseMoveListener and MouseClickListener

func main() {
	width, height := 600, 600

	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{width, height}})
	win := glui.NewGLWin(width, height, "Mouse Click Test", img, true)
	glui.NewMouseMoveListener(win, func(pt []float64) {
		fmt.Printf("Mouse moved to %.2f,%.2f\n", pt[0], pt[1])
	})
	glui.NewMouseClickListener(win, glfw.MouseButtonLeft, func(pt []float64) {
		fmt.Printf("Mouse clicked at %.2f,%.2f\n", pt[0], pt[1])
	})
	glui.Loop(nil)
}
