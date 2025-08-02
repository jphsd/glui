package main

import (
	"fmt"
	"github.com/jphsd/glui"
	"image"
)

// Example use of MouseDragListener

func main() {
	width, height := 600, 600

	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{width, height}})
	win := glui.NewGLWin(width, height, "Mouse Drag Test", img, true)
	glui.NewMouseDragListener(win, glui.MouseButtonLeft, func(pt []float64, dx, dy float64, act glui.Action) {
		switch act {
		case glui.Press:
			fmt.Printf("Drag started at %.2f,%.2f\n", pt[0], pt[1])
		case glui.Release:
			fmt.Printf("Drag ended at %.2f,%.2f (%.2f,%.2f)\n", pt[0], pt[1], dx, dy)
		case glui.Repeat:
			fmt.Printf("Dragging %.2f,%.2f (%.2f,%.2f)\n", pt[0], pt[1], dx, dy)
		}
	})
	glui.Loop(nil)
}
