package main

import (
	"fmt"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/jphsd/glui"
	"image"
)

// Example use of MouseDragListener

func main() {
	width, height := 600, 600

	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{width, height}})
	win := glui.NewGLWin(width, height, "Mouse Drag Test", img, true)
	glui.NewMouseDragListener(win, glfw.MouseButtonLeft, func(pt []float64, dx, dy float64, act glfw.Action) {
		switch act {
		case glfw.Press:
			fmt.Printf("Drag started at %.2f,%.2f\n", pt[0], pt[1])
		case glfw.Release:
			fmt.Printf("Drag ended at %.2f,%.2f (%.2f,%.2f)\n", pt[0], pt[1], dx, dy)
		case glfw.Repeat:
			fmt.Printf("Dragging %.2f,%.2f (%.2f,%.2f)\n", pt[0], pt[1], dx, dy)
		}
	})
	glui.Loop(nil)
}
