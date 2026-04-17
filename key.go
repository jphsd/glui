package glui

import "github.com/go-gl/glfw/v3.3/glfw"

type Keyable interface {
	OnKey(key Key, scancode int, act Action, mods ModifierKey)
}

// KeyListener adds a wrapper around glfw's key callback.
type KeyListener struct {
	Window    *GLWin
	Observers []Keyable
}

// NewKeyListener adds a key listener to the supplied window
// and the function to be called when a key action is detected.
func NewKeyListener(win *GLWin, onKey ...Keyable) *KeyListener {
	res := &KeyListener{win, onKey}
	win.Win.SetKeyCallback(func(w *glfw.Window, k Key, sc int, act Action, mods ModifierKey) {
		for _, ocf := range res.Observers {
			if ocf == nil {
				continue
			}
			go ocf.OnKey(k, sc, act, mods)
		}
	})
	return res
}

func GetKeyName(key Key) string {
	switch key {
	// Printable keys
	case KeyA:
		return "A"
	case KeyB:
		return "B"
	case KeyC:
		return "C"
	case KeyD:
		return "D"
	case KeyE:
		return "E"
	case KeyF:
		return "F"
	case KeyG:
		return "G"
	case KeyH:
		return "H"
	case KeyI:
		return "I"
	case KeyJ:
		return "J"
	case KeyK:
		return "K"
	case KeyL:
		return "L"
	case KeyM:
		return "M"
	case KeyN:
		return "N"
	case KeyO:
		return "O"
	case KeyP:
		return "P"
	case KeyQ:
		return "Q"
	case KeyR:
		return "R"
	case KeyS:
		return "S"
	case KeyT:
		return "T"
	case KeyU:
		return "U"
	case KeyV:
		return "V"
	case KeyW:
		return "W"
	case KeyX:
		return "X"
	case KeyY:
		return "Y"
	case KeyZ:
		return "Z"
	case Key1:
		return "1"
	case Key2:
		return "2"
	case Key3:
		return "3"
	case Key4:
		return "4"
	case Key5:
		return "5"
	case Key6:
		return "6"
	case Key7:
		return "7"
	case Key8:
		return "8"
	case Key9:
		return "9"
	case Key0:
		return "0"
	case KeySpace:
		return "SPACE"
	case KeyMinus:
		return "MINUS"
	case KeyEqual:
		return "EQUAL"
	case KeyLeftBracket:
		return "LEFT BRACKET"
	case KeyRightBracket:
		return "RIGHT BRACKET"
	case KeyBackslash:
		return "BACKSLASH"
	case KeySemicolon:
		return "SEMICOLON"
	case KeyApostrophe:
		return "APOSTROPHE"
	case KeyGraveAccent:
		return "GRAVE ACCENT"
	case KeyComma:
		return "COMMA"
	case KeyPeriod:
		return "PERIOD"
	case KeySlash:
		return "SLASH"
	case KeyWorld1:
		return "WORLD 1"
	case KeyWorld2:
		return "WORLD 2"

	// Function keys
	case KeyEscape:
		return "ESCAPE"
	case KeyF1:
		return "F1"
	case KeyF2:
		return "F2"
	case KeyF3:
		return "F3"
	case KeyF4:
		return "F4"
	case KeyF5:
		return "F5"
	case KeyF6:
		return "F6"
	case KeyF7:
		return "F7"
	case KeyF8:
		return "F8"
	case KeyF9:
		return "F9"
	case KeyF10:
		return "F10"
	case KeyF11:
		return "F11"
	case KeyF12:
		return "F12"
	case KeyF13:
		return "F13"
	case KeyF14:
		return "F14"
	case KeyF15:
		return "F15"
	case KeyF16:
		return "F16"
	case KeyF17:
		return "F17"
	case KeyF18:
		return "F18"
	case KeyF19:
		return "F19"
	case KeyF20:
		return "F20"
	case KeyF21:
		return "F21"
	case KeyF22:
		return "F22"
	case KeyF23:
		return "F23"
	case KeyF24:
		return "F24"
	case KeyF25:
		return "F25"

		// Cursor keys
	case KeyUp:
		return "UP"
	case KeyDown:
		return "DOWN"
	case KeyLeft:
		return "LEFT"
	case KeyRight:
		return "RIGHT"

		// Shift/Control etc.
	case KeyLeftShift:
		return "LEFT SHIFT"
	case KeyRightShift:
		return "RIGHT SHIFT"
	case KeyLeftControl:
		return "LEFT CONTROL"
	case KeyRightControl:
		return "RIGHT CONTROL"
	case KeyLeftAlt:
		return "LEFT ALT"
	case KeyRightAlt:
		return "RIGHT ALT"
	case KeyTab:
		return "TAB"
	case KeyEnter:
		return "ENTER"
	case KeyBackspace:
		return "BACKSPACE"
	case KeyInsert:
		return "INSERT"
	case KeyDelete:
		return "DELETE"
	case KeyPageUp:
		return "PAGE UP"
	case KeyPageDown:
		return "PAGE DOWN"
	case KeyHome:
		return "HOME"
	case KeyEnd:
		return "END"
	case KeyPrintScreen:
		return "PRINT SCREEN"
	case KeyNumLock:
		return "NUM LOCK"
	case KeyCapsLock:
		return "CAPS LOCK"
	case KeyScrollLock:
		return "SCROLL LOCK"
	case KeyPause:
		return "PAUSE"
	case KeyLeftSuper:
		return "LEFT SUPER"
	case KeyRightSuper:
		return "RIGHT SUPER"
	case KeyMenu:
		return "MENU"

		// Keypad
	case KeyKP0:
		return "KEYPAD 0"
	case KeyKP1:
		return "KEYPAD 1"
	case KeyKP2:
		return "KEYPAD 2"
	case KeyKP3:
		return "KEYPAD 3"
	case KeyKP4:
		return "KEYPAD 4"
	case KeyKP5:
		return "KEYPAD 5"
	case KeyKP6:
		return "KEYPAD 6"
	case KeyKP7:
		return "KEYPAD 7"
	case KeyKP8:
		return "KEYPAD 8"
	case KeyKP9:
		return "KEYPAD 9"
	case KeyKPDivide:
		return "KEYPAD DIVIDE"
	case KeyKPMultiply:
		return "KEYPAD MULTIPLY"
	case KeyKPSubtract:
		return "KEYPAD SUBTRACT"
	case KeyKPAdd:
		return "KEYPAD ADD"
	case KeyKPDecimal:
		return "KEYPAD DECIMAL"
	case KeyKPEqual:
		return "KEYPAD EQUAL"
	case KeyKPEnter:
		return "KEYPAD ENTER"

	default:
		return "UNKNOWN"
	}
}

//window.SetCharCallback(char_callback)
//func char_callback(window *glfw.Window, codepoint rune) {

type Characterable interface {
	OnCharacter(c rune)
}

// CharaterListener adds a wrapper around glfw's character callback.
type CharacterListener struct {
	Window    *GLWin
	Observers []Characterable
}

// NewCharacterListener adds a key listener to the supplied window
// and the function to be called when a character is detected.
func NewCharacterListener(win *GLWin, onChar ...Characterable) *CharacterListener {
	res := &CharacterListener{win, onChar}
	win.Win.SetCharCallback(func(w *glfw.Window, r rune) {
		for _, ocf := range res.Observers {
			if ocf == nil {
				continue
			}
			go ocf.OnCharacter(r)
		}
	})
	return res
}
