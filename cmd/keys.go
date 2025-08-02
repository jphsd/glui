package main

import (
	"fmt"
	"github.com/jphsd/glui"
	"image"
)

// Example use of KeyListener and CharacterListener

func main() {
	width, height := 600, 600

	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{width, height}})
	win := glui.NewGLWin(width, height, "Keyboard Test", img, true)
	glui.NewKeyListener(win, func(k, sc int, act glui.Action, mods glui.ModifierKey) {
		fmt.Printf("Key: %s (%d), %d, %s %s\n", glui.GetKeyName(k), k, sc, GetActionName(act), GetModsName(mods))
	})
	glui.NewCharacterListener(win, func(r rune) {
		fmt.Printf("Char: %c\n", r)
	})
	glui.Loop(nil)
}

func GetModsName(mods glui.ModifierKey) string {
	if mods == glui.ModifierKey(0) {
		return " no mods"
	}

	var name string

	if mods&glui.ModShift > 0 {
		name += " shift"
	}
	if mods&glui.ModControl > 0 {
		name += " control"
	}
	if mods&glui.ModAlt > 0 {
		name += " alt"
	}
	if mods&glui.ModSuper > 0 {
		name += " super"
	}
	if mods&glui.ModCapsLock > 0 {
		name += " capslock-on"
	}
	if mods&glui.ModNumLock > 0 {
		name += " numlock-on"
	}

	return name
}

func GetActionName(action glui.Action) string {
	switch action {
	case glui.Press:
		return "pressed"
	case glui.Release:
		return "released"
	case glui.Repeat:
		return "repeated"
	}

	return "caused unknown action"
}
