package glui

import "github.com/go-gl/glfw/v3.3/glfw"

type Action = glfw.Action

const (
	Press   Action = glfw.Press
	Repeat  Action = glfw.Repeat
	Release Action = glfw.Release
)
