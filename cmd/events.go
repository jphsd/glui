package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

/* Copyright notice from original: github.com/glfw/glfw/examples/event.c
//========================================================================
// Event linter (event spewer)
// Copyright (c) Camilla Löwy <elmindreda@glfw.org>
//
// This software is provided 'as-is', without any express or implied
// warranty. In no event will the authors be held liable for any damages
// arising from the use of this software.
//
// Permission is granted to anyone to use this software for any purpose,
// including commercial applications, and to alter it and redistribute it
// freely, subject to the following restrictions:
//
// 1. The origin of this software must not be misrepresented; you must not
//    claim that you wrote the original software. If you use this software
//    in a product, an acknowledgment in the product documentation would
//    be appreciated but is not required.
//
// 2. Altered source versions must be plainly marked as such, and must not
//    be misrepresented as being the original software.
//
// 3. This notice may not be removed or altered from any source
//    distribution.
//
//========================================================================
//
// This test hooks every available callback and outputs their arguments
//
// Log messages go to stdout, error messages to stderr
//
// Every event also gets a (sequential) number to aid discussion of logs
//
//========================================================================
*/

func init() {
	// GLFW event handling *must* run on the main OS thread
	runtime.LockOSThread()
}

func get_key_name(key glfw.Key) string {
	switch key {
	// Printable keys
	case glfw.KeyA:
		return "A"
	case glfw.KeyB:
		return "B"
	case glfw.KeyC:
		return "C"
	case glfw.KeyD:
		return "D"
	case glfw.KeyE:
		return "E"
	case glfw.KeyF:
		return "F"
	case glfw.KeyG:
		return "G"
	case glfw.KeyH:
		return "H"
	case glfw.KeyI:
		return "I"
	case glfw.KeyJ:
		return "J"
	case glfw.KeyK:
		return "K"
	case glfw.KeyL:
		return "L"
	case glfw.KeyM:
		return "M"
	case glfw.KeyN:
		return "N"
	case glfw.KeyO:
		return "O"
	case glfw.KeyP:
		return "P"
	case glfw.KeyQ:
		return "Q"
	case glfw.KeyR:
		return "R"
	case glfw.KeyS:
		return "S"
	case glfw.KeyT:
		return "T"
	case glfw.KeyU:
		return "U"
	case glfw.KeyV:
		return "V"
	case glfw.KeyW:
		return "W"
	case glfw.KeyX:
		return "X"
	case glfw.KeyY:
		return "Y"
	case glfw.KeyZ:
		return "Z"
	case glfw.Key1:
		return "1"
	case glfw.Key2:
		return "2"
	case glfw.Key3:
		return "3"
	case glfw.Key4:
		return "4"
	case glfw.Key5:
		return "5"
	case glfw.Key6:
		return "6"
	case glfw.Key7:
		return "7"
	case glfw.Key8:
		return "8"
	case glfw.Key9:
		return "9"
	case glfw.Key0:
		return "0"
	case glfw.KeySpace:
		return "SPACE"
	case glfw.KeyMinus:
		return "MINUS"
	case glfw.KeyEqual:
		return "EQUAL"
	case glfw.KeyLeftBracket:
		return "LEFT BRACKET"
	case glfw.KeyRightBracket:
		return "RIGHT BRACKET"
	case glfw.KeyBackslash:
		return "BACKSLASH"
	case glfw.KeySemicolon:
		return "SEMICOLON"
	case glfw.KeyApostrophe:
		return "APOSTROPHE"
	case glfw.KeyGraveAccent:
		return "GRAVE ACCENT"
	case glfw.KeyComma:
		return "COMMA"
	case glfw.KeyPeriod:
		return "PERIOD"
	case glfw.KeySlash:
		return "SLASH"
	case glfw.KeyWorld1:
		return "WORLD 1"
	case glfw.KeyWorld2:
		return "WORLD 2"

	// Function keys
	case glfw.KeyEscape:
		return "ESCAPE"
	case glfw.KeyF1:
		return "F1"
	case glfw.KeyF2:
		return "F2"
	case glfw.KeyF3:
		return "F3"
	case glfw.KeyF4:
		return "F4"
	case glfw.KeyF5:
		return "F5"
	case glfw.KeyF6:
		return "F6"
	case glfw.KeyF7:
		return "F7"
	case glfw.KeyF8:
		return "F8"
	case glfw.KeyF9:
		return "F9"
	case glfw.KeyF10:
		return "F10"
	case glfw.KeyF11:
		return "F11"
	case glfw.KeyF12:
		return "F12"
	case glfw.KeyF13:
		return "F13"
	case glfw.KeyF14:
		return "F14"
	case glfw.KeyF15:
		return "F15"
	case glfw.KeyF16:
		return "F16"
	case glfw.KeyF17:
		return "F17"
	case glfw.KeyF18:
		return "F18"
	case glfw.KeyF19:
		return "F19"
	case glfw.KeyF20:
		return "F20"
	case glfw.KeyF21:
		return "F21"
	case glfw.KeyF22:
		return "F22"
	case glfw.KeyF23:
		return "F23"
	case glfw.KeyF24:
		return "F24"
	case glfw.KeyF25:
		return "F25"

		// Cursor keys
	case glfw.KeyUp:
		return "UP"
	case glfw.KeyDown:
		return "DOWN"
	case glfw.KeyLeft:
		return "LEFT"
	case glfw.KeyRight:
		return "RIGHT"

		// Shift/Control etc.
	case glfw.KeyLeftShift:
		return "LEFT SHIFT"
	case glfw.KeyRightShift:
		return "RIGHT SHIFT"
	case glfw.KeyLeftControl:
		return "LEFT CONTROL"
	case glfw.KeyRightControl:
		return "RIGHT CONTROL"
	case glfw.KeyLeftAlt:
		return "LEFT ALT"
	case glfw.KeyRightAlt:
		return "RIGHT ALT"
	case glfw.KeyTab:
		return "TAB"
	case glfw.KeyEnter:
		return "ENTER"
	case glfw.KeyBackspace:
		return "BACKSPACE"
	case glfw.KeyInsert:
		return "INSERT"
	case glfw.KeyDelete:
		return "DELETE"
	case glfw.KeyPageUp:
		return "PAGE UP"
	case glfw.KeyPageDown:
		return "PAGE DOWN"
	case glfw.KeyHome:
		return "HOME"
	case glfw.KeyEnd:
		return "END"
	case glfw.KeyPrintScreen:
		return "PRINT SCREEN"
	case glfw.KeyNumLock:
		return "NUM LOCK"
	case glfw.KeyCapsLock:
		return "CAPS LOCK"
	case glfw.KeyScrollLock:
		return "SCROLL LOCK"
	case glfw.KeyPause:
		return "PAUSE"
	case glfw.KeyLeftSuper:
		return "LEFT SUPER"
	case glfw.KeyRightSuper:
		return "RIGHT SUPER"
	case glfw.KeyMenu:
		return "MENU"

		// Keypad
	case glfw.KeyKP0:
		return "KEYPAD 0"
	case glfw.KeyKP1:
		return "KEYPAD 1"
	case glfw.KeyKP2:
		return "KEYPAD 2"
	case glfw.KeyKP3:
		return "KEYPAD 3"
	case glfw.KeyKP4:
		return "KEYPAD 4"
	case glfw.KeyKP5:
		return "KEYPAD 5"
	case glfw.KeyKP6:
		return "KEYPAD 6"
	case glfw.KeyKP7:
		return "KEYPAD 7"
	case glfw.KeyKP8:
		return "KEYPAD 8"
	case glfw.KeyKP9:
		return "KEYPAD 9"
	case glfw.KeyKPDivide:
		return "KEYPAD DIVIDE"
	case glfw.KeyKPMultiply:
		return "KEYPAD MULTIPLY"
	case glfw.KeyKPSubtract:
		return "KEYPAD SUBTRACT"
	case glfw.KeyKPAdd:
		return "KEYPAD ADD"
	case glfw.KeyKPDecimal:
		return "KEYPAD DECIMAL"
	case glfw.KeyKPEqual:
		return "KEYPAD EQUAL"
	case glfw.KeyKPEnter:
		return "KEYPAD ENTER"

	default:
		return "UNKNOWN"
	}
}

func get_action_name(action glfw.Action) string {
	switch action {
	case glfw.Press:
		return "pressed"
	case glfw.Release:
		return "released"
	case glfw.Repeat:
		return "repeated"
	}

	return "caused unknown action"
}

func get_button_name(button glfw.MouseButton) string {
	switch button {
	case glfw.MouseButtonLeft:
		return "left button"
	case glfw.MouseButtonRight:
		return "right button"
	case glfw.MouseButtonMiddle:
		return "middle button"
	case glfw.MouseButtonLast:
		return "last button"
	case glfw.MouseButton4:
		return "button 4"
	case glfw.MouseButton5:
		return "button 5"
	case glfw.MouseButton6:
		return "button 6"
	case glfw.MouseButton7:
		return "button 7"
	default:
		return "unknown button"
	}
}

func get_mods_name(mods glfw.ModifierKey) string {
	if mods == glfw.ModifierKey(0) {
		return " no mods"
	}

	var name string

	if mods&glfw.ModShift > 0 {
		name += " shift"
	}
	if mods&glfw.ModControl > 0 {
		name += " control"
	}
	if mods&glfw.ModAlt > 0 {
		name += " alt"
	}
	if mods&glfw.ModSuper > 0 {
		name += " super"
	}
	if mods&glfw.ModCapsLock > 0 {
		name += " capslock-on"
	}
	if mods&glfw.ModNumLock > 0 {
		name += " numlock-on"
	}

	return name
}

// The callbacks

// The error callback is handled internally - see glfw/error.go

// Event index
var counter uint = 0

// Slots for windows
type Slot struct {
	window    *glfw.Window
	number    int
	closeable bool
}

var slots = make(map[*glfw.Window]*Slot)

func window_pos_callback(window *glfw.Window, x, y int) {
	slot := slots[window]
	fmt.Printf("%08x to %d at %0.3f: Window position: %d %d\n",
		counter, slot.number, glfw.GetTime(), x, y)
	counter++
}

func window_size_callback(window *glfw.Window, width, height int) {
	slot := slots[window]
	fmt.Printf("%08x to %d at %0.3f: Window size: %d %d\n",
		counter, slot.number, glfw.GetTime(), width, height)
	counter++
}

func framebuffer_size_callback(window *glfw.Window, width, height int) {
	slot := slots[window]
	fmt.Printf("%08x to %d at %0.3f: Framebuffer size: %d %d\n",
		counter, slot.number, glfw.GetTime(), width, height)
	counter++
}

func window_content_scale_callback(window *glfw.Window, xscale, yscale float32) {
	slot := slots[window]
	fmt.Printf("%08x to %d at %0.3f: Window content scale: %0.3f %0.3f\n",
		counter, slot.number, glfw.GetTime(), xscale, yscale)
	counter++
}

func window_close_callback(window *glfw.Window) {
	slot := slots[window]
	fmt.Printf("%08x to %d at %0.3f: Window close\n",
		counter, slot.number, glfw.GetTime())
	counter++

	if !slot.closeable {
		fmt.Printf("( closing is disabled, press %s to re-enable )\n",
			glfw.GetKeyName(glfw.KeyC, 0))
	}

	window.SetShouldClose(slot.closeable)
}

func window_refresh_callback(window *glfw.Window) {
	slot := slots[window]
	fmt.Printf("%08x to %d at %0.3f: Window refresh\n",
		counter, slot.number, glfw.GetTime())
	counter++

	window.MakeContextCurrent()
	gl.Clear(gl.COLOR_BUFFER_BIT) // default color is black
	window.SwapBuffers()
}

func window_focus_callback(window *glfw.Window, focused bool) {
	slot := slots[window]
	str := "focused"
	if !focused {
		str = "defocused"
	}
	fmt.Printf("%08x to %d at %0.3f: Window %s\n",
		counter, slot.number, glfw.GetTime(), str)
	counter++
}

func window_iconify_callback(window *glfw.Window, iconified bool) {
	slot := slots[window]
	str := "iconified"
	if !iconified {
		str = "uniconified"
	}
	fmt.Printf("%08x to %d at %0.3f: Window was %s\n",
		counter, slot.number, glfw.GetTime(), str)
	counter++
}

func window_maximize_callback(window *glfw.Window, maximized bool) {
	slot := slots[window]
	str := "maximized"
	if !maximized {
		str = "unmaximized"
	}
	fmt.Printf("%08x to %d at %0.3f: Window was %s\n",
		counter, slot.number, glfw.GetTime(), str)
	counter++
}

func mouse_button_callback(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
	slot := slots[window]
	fmt.Printf("%08x to %d at %0.3f: Mouse button %d (%s) (with%s) was %s\n",
		counter, slot.number, glfw.GetTime(), button,
		get_button_name(button),
		get_mods_name(mods),
		get_action_name(action))
	counter++
}

var track = true

func cursor_position_callback(window *glfw.Window, x, y float64) {
	if !track {
		return
	}
	slot := slots[window]
	fmt.Printf("%08x to %d at %0.3f: Cursor position: %f %f\n",
		counter, slot.number, glfw.GetTime(), x, y)
	counter++
}

func cursor_enter_callback(window *glfw.Window, entered bool) {
	slot := slots[window]
	str := "entered"
	if !entered {
		str = "left"
	}
	fmt.Printf("%08x to %d at %0.3f: Cursor %s window\n",
		counter, slot.number, glfw.GetTime(), str)
	counter++
}

func scroll_callback(window *glfw.Window, x, y float64) {
	slot := slots[window]
	fmt.Printf("%08x to %d at %0.3f: Scroll: %0.3f %0.3f\n",
		counter, slot.number, glfw.GetTime(), x, y)
	counter++
}

func key_callback(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	slot := slots[window]

	if key != glfw.KeyUnknown {
		fmt.Printf("%08x to %d at %0.3f: Key 0x%04x Scancode 0x%04x (%s) (%s) (with%s) was %s\n",
			counter, slot.number, glfw.GetTime(), key, scancode,
			get_key_name(key),
			glfw.GetKeyName(key, scancode),
			get_mods_name(mods),
			get_action_name(action))
	} else {
		fmt.Printf("%08x to %d at %0.3f: Key 0x%04x Scancode 0x%04x (%s) (with%s) was %s\n",
			counter, slot.number, glfw.GetTime(), key, scancode,
			get_key_name(key),
			get_mods_name(mods),
			get_action_name(action))
	}
	counter++

	if action != glfw.Press {
		return
	}

	switch key {
	case glfw.KeyC:
		slot.closeable = !slot.closeable
		cstr := "enabled"
		if !slot.closeable {
			cstr = "disabled"
		}
		fmt.Printf("(( closing %s ))\n", cstr)
	case glfw.KeyL:
		state := window.GetInputMode(glfw.LockKeyMods)
		sstr := "disabled"
		if state == 0 {
			sstr = "enabled"
			state = 1
		} else {
			state = 0
		}
		window.SetInputMode(glfw.LockKeyMods, state)
		fmt.Printf("(( lock key mods %s ))\n", sstr)
	case glfw.KeyEscape:
		window.SetShouldClose(true)
	case glfw.KeyT:
		track = !track
		if track {
			fmt.Printf("(( cursor tracking ON ))\n")
		} else {
			fmt.Printf("(( cursor tracking OFF ))\n")
		}
	}
}

func char_callback(window *glfw.Window, codepoint rune) {
	slot := slots[window]

	fmt.Printf("%08x to %d at %0.3f: Character 0x%08x (%c) input\n",
		counter, slot.number, glfw.GetTime(), codepoint, codepoint)
	counter++
}

func drop_callback(window *glfw.Window, paths []string) {
	slot := slots[window]

	fmt.Printf("%08x to %d at %0.3f: Drop input\n",
		counter, slot.number, glfw.GetTime())
	counter++

	for i, s := range paths {
		fmt.Printf("  %d: \"%s\"\n", i, s)
	}
}

func monitor_callback(monitor *glfw.Monitor, event glfw.PeripheralEvent) {
	if event == glfw.Connected {
		mode := monitor.GetVideoMode()
		x, y := monitor.GetPos()
		widthMM, heightMM := monitor.GetPhysicalSize()

		fmt.Printf("%08x at %0.3f: Monitor %s (%dx%d at %dx%d, %dx%d mm) was connected\n",
			counter,
			glfw.GetTime(),
			monitor.GetName(),
			mode.Width, mode.Height,
			x, y,
			widthMM, heightMM)
	} else if event == glfw.Disconnected {
		fmt.Printf("%08x at %0.3f: Monitor %s was disconnected\n",
			counter,
			glfw.GetTime(),
			monitor.GetName())
	}
	counter++
}

func joystick_callback(jid glfw.Joystick, event glfw.PeripheralEvent) {
	if event == glfw.Connected {
		axes := jid.GetAxes()
		actions := jid.GetButtons()
		hatStates := jid.GetHats()

		fmt.Printf("%08x at %0.3f: Joystick %d (%s) was connected with %d axes, %d buttons, and %d hats\n",
			counter, glfw.GetTime(),
			jid,
			jid.GetName(),
			len(axes),
			len(actions),
			len(hatStates))

		if jid.IsGamepad() {
			fmt.Printf("  Joystick %d (%s) has a gamepad mapping (%s)\n",
				jid,
				jid.GetGUID(),
				jid.GetGamepadName())
		} else {
			fmt.Printf("  Joystick %d (%s) has no gamepad mapping\n",
				jid,
				jid.GetGUID())
		}
	} else {
		fmt.Printf("%08x at %0.3f: Joystick %d was disconnected\n",
			counter, glfw.GetTime(), jid)
	}
	counter++
}

func main() {
	ffp := flag.Int("f", -1, "full screen")
	nfp := flag.Int("n", 1, "number of windows")
	flag.Parse()

	var monitor *glfw.Monitor
	var width, height int
	count := *nfp

	// Inititalize graphics
	if err := glfw.Init(); err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	fmt.Printf("C - toggle window closing enabled\n")
	fmt.Printf("L - toggle modifier lock\n")
	fmt.Printf("ESC - to exit\n")
	fmt.Printf("Library initialized\n")

	glfw.SetMonitorCallback(monitor_callback)
	glfw.SetJoystickCallback(joystick_callback)

	monitors := glfw.GetMonitors()
	if *ffp != -1 {
		m := *ffp
		if m > len(monitors) {
			m = 0
		}
		monitor = monitors[m]
	}

	if monitor != nil {
		mode := monitor.GetVideoMode()

		glfw.WindowHint(glfw.RefreshRate, mode.RefreshRate)
		glfw.WindowHint(glfw.RedBits, mode.RedBits)
		glfw.WindowHint(glfw.GreenBits, mode.GreenBits)
		glfw.WindowHint(glfw.BlueBits, mode.BlueBits)

		width = mode.Width
		height = mode.Height
	} else {
		width = 640
		height = 480
	}

	for i := 0; i < count; i++ {
		title := fmt.Sprintf("Event Linter (Window %d)", i+1)

		if monitor != nil {
			fmt.Printf("Creating full screen window %d (%dx%d on %s)\n",
				i+1,
				width, height,
				monitor.GetName())
		} else {
			fmt.Printf("Creating windowed mode window %d (%dx%d)\n",
				i+1,
				width, height)
		}

		window, err := glfw.CreateWindow(width, height, title, monitor, nil)
		if err != nil {
			panic(err)
		}

		slot := Slot{window, i + 1, true}
		slots[window] = &slot

		window.SetPosCallback(window_pos_callback)
		window.SetSizeCallback(window_size_callback)
		window.SetFramebufferSizeCallback(framebuffer_size_callback)
		window.SetContentScaleCallback(window_content_scale_callback)
		window.SetCloseCallback(window_close_callback)
		window.SetRefreshCallback(window_refresh_callback)
		window.SetFocusCallback(window_focus_callback)
		window.SetIconifyCallback(window_iconify_callback)
		window.SetMaximizeCallback(window_maximize_callback)
		window.SetMouseButtonCallback(mouse_button_callback)
		window.SetCursorPosCallback(cursor_position_callback)
		window.SetCursorEnterCallback(cursor_enter_callback)
		window.SetScrollCallback(scroll_callback)
		window.SetKeyCallback(key_callback)
		window.SetCharCallback(char_callback)
		window.SetDropCallback(drop_callback)

		window.MakeContextCurrent()
	}
	// Initialize GL bindings (uses the last active context)
	if err := gl.Init(); err != nil {
		panic(err)
	}
	glfw.SwapInterval(1)

	fmt.Printf("Main loop starting\n")

	for {
		// Closing all windows quits the program
		i := 0
		for w := range slots {
			if w.ShouldClose() {
				w.Destroy()
				delete(slots, w)
				// Pass focus to another visible window, if any
				for wf := range slots {
					if wf.GetAttrib(glfw.Visible) == 1 {
						wf.Focus()
						break
					}
				}
			} else {
				i++
			}
		}

		if i < 1 {
			break
		}

		glfw.WaitEvents()
	}
}
