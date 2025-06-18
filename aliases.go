package glui

import "github.com/go-gl/glfw/v3.3/glfw"

type Action = glfw.Action

const (
	Press   Action = glfw.Press
	Repeat  Action = glfw.Repeat
	Release Action = glfw.Release
)

type MouseButton = glfw.MouseButton

const (
	MouseButton1      MouseButton = glfw.MouseButton1
	MouseButton2      MouseButton = glfw.MouseButton2
	MouseButton3      MouseButton = glfw.MouseButton3
	MouseButton4      MouseButton = glfw.MouseButton4
	MouseButton5      MouseButton = glfw.MouseButton5
	MouseButton6      MouseButton = glfw.MouseButton6
	MouseButton7      MouseButton = glfw.MouseButton7
	MouseButton8      MouseButton = glfw.MouseButton8
	MouseButtonLast   MouseButton = glfw.MouseButtonLast
	MouseButtonLeft   MouseButton = glfw.MouseButtonLeft
	MouseButtonRight  MouseButton = glfw.MouseButtonRight
	MouseButtonMiddle MouseButton = glfw.MouseButtonMiddle
)
