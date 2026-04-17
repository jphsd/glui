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
	glui.NewMouseDragListener(win, drag{})
	glui.Loop(nil)
}

type drag struct{}

func (d drag) OnDrag(mdl *glui.MouseDragListener, but glui.MouseButton, pt []float64, dx, dy float64, act glui.Action, mods glui.ModifierKey) {
	switch act {
	case glui.Press:
		fmt.Printf("Drag started at %.2f,%.2f\n", pt[0], pt[1])
	case glui.Release:
		fmt.Printf("Drag ended at %.2f,%.2f (%.2f,%.2f)\n", pt[0], pt[1], dx, dy)
	case glui.Repeat:
		fmt.Printf("Dragging %.2f,%.2f (%.2f,%.2f)\n", pt[0], pt[1], dx, dy)
	}
}
