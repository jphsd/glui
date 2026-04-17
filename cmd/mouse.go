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
	win := glui.NewGLWin(width, height, "Mouse Click/Move/Scroll Test", img, true)
	cb := frodo{}
	glui.NewMouseMoveListener(win, cb)
	glui.NewMouseClickListener(win, cb)
	glui.NewMouseScrollListener(win, cb)
	glui.Loop(nil)
}

type frodo struct{}

func (m frodo) OnMove(pt []float64) {
	fmt.Printf("Mouse moved to %.2f,%.2f\n", pt[0], pt[1])
}

func (c frodo) OnClick(but glui.MouseButton, pt []float64, mods glui.ModifierKey) {
	fmt.Printf("Mouse button %d  clicked at %.2f,%.2f with mods 0x%x\n", but, pt[0], pt[1], mods)
}

func (s frodo) OnScroll(x, y float64) {
	fmt.Printf("Mouse scroll %.2f,%.2f\n", x, y)
}
