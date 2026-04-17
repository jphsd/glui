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

type Key = glfw.Key

const (
	// Printable keys
	KeyA            Key = glfw.KeyA
	KeyB            Key = glfw.KeyB
	KeyC            Key = glfw.KeyC
	KeyD            Key = glfw.KeyD
	KeyE            Key = glfw.KeyE
	KeyF            Key = glfw.KeyF
	KeyG            Key = glfw.KeyG
	KeyH            Key = glfw.KeyH
	KeyI            Key = glfw.KeyI
	KeyJ            Key = glfw.KeyJ
	KeyK            Key = glfw.KeyK
	KeyL            Key = glfw.KeyL
	KeyM            Key = glfw.KeyM
	KeyN            Key = glfw.KeyN
	KeyO            Key = glfw.KeyO
	KeyP            Key = glfw.KeyP
	KeyQ            Key = glfw.KeyQ
	KeyR            Key = glfw.KeyR
	KeyS            Key = glfw.KeyS
	KeyT            Key = glfw.KeyT
	KeyU            Key = glfw.KeyU
	KeyV            Key = glfw.KeyV
	KeyW            Key = glfw.KeyW
	KeyX            Key = glfw.KeyX
	KeyY            Key = glfw.KeyY
	KeyZ            Key = glfw.KeyZ
	Key1            Key = glfw.Key1
	Key2            Key = glfw.Key2
	Key3            Key = glfw.Key3
	Key4            Key = glfw.Key4
	Key5            Key = glfw.Key5
	Key6            Key = glfw.Key6
	Key7            Key = glfw.Key7
	Key8            Key = glfw.Key8
	Key9            Key = glfw.Key9
	Key0            Key = glfw.Key0
	KeySpace        Key = glfw.KeySpace
	KeyMinus        Key = glfw.KeyMinus
	KeyEqual        Key = glfw.KeyEqual
	KeyLeftBracket  Key = glfw.KeyLeftBracket
	KeyRightBracket Key = glfw.KeyRightBracket
	KeyBackslash    Key = glfw.KeyBackslash
	KeySemicolon    Key = glfw.KeySemicolon
	KeyApostrophe   Key = glfw.KeyApostrophe
	KeyGraveAccent  Key = glfw.KeyGraveAccent
	KeyComma        Key = glfw.KeyComma
	KeyPeriod       Key = glfw.KeyPeriod
	KeySlash        Key = glfw.KeySlash
	KeyWorld1       Key = glfw.KeyWorld1
	KeyWorld2       Key = glfw.KeyWorld2

	// Function keys
	KeyEscape Key = glfw.KeyEscape
	KeyF1     Key = glfw.KeyF1
	KeyF2     Key = glfw.KeyF2
	KeyF3     Key = glfw.KeyF3
	KeyF4     Key = glfw.KeyF4
	KeyF5     Key = glfw.KeyF5
	KeyF6     Key = glfw.KeyF6
	KeyF7     Key = glfw.KeyF7
	KeyF8     Key = glfw.KeyF8
	KeyF9     Key = glfw.KeyF9
	KeyF10    Key = glfw.KeyF10
	KeyF11    Key = glfw.KeyF11
	KeyF12    Key = glfw.KeyF12
	KeyF13    Key = glfw.KeyF13
	KeyF14    Key = glfw.KeyF14
	KeyF15    Key = glfw.KeyF15
	KeyF16    Key = glfw.KeyF16
	KeyF17    Key = glfw.KeyF17
	KeyF18    Key = glfw.KeyF18
	KeyF19    Key = glfw.KeyF19
	KeyF20    Key = glfw.KeyF20
	KeyF21    Key = glfw.KeyF21
	KeyF22    Key = glfw.KeyF22
	KeyF23    Key = glfw.KeyF23
	KeyF24    Key = glfw.KeyF24
	KeyF25    Key = glfw.KeyF25

	// Cursor keys
	KeyUp    Key = glfw.KeyUp
	KeyDown  Key = glfw.KeyDown
	KeyLeft  Key = glfw.KeyLeft
	KeyRight Key = glfw.KeyRight

	// Shift/Control etc.
	KeyLeftShift    Key = glfw.KeyLeftShift
	KeyRightShift   Key = glfw.KeyRightShift
	KeyLeftControl  Key = glfw.KeyLeftControl
	KeyRightControl Key = glfw.KeyRightControl
	KeyLeftAlt      Key = glfw.KeyLeftAlt
	KeyRightAlt     Key = glfw.KeyRightAlt
	KeyTab          Key = glfw.KeyTab
	KeyEnter        Key = glfw.KeyEnter
	KeyBackspace    Key = glfw.KeyBackspace
	KeyInsert       Key = glfw.KeyInsert
	KeyDelete       Key = glfw.KeyDelete
	KeyPageUp       Key = glfw.KeyPageUp
	KeyPageDown     Key = glfw.KeyPageDown
	KeyHome         Key = glfw.KeyHome
	KeyEnd          Key = glfw.KeyEnd
	KeyPrintScreen  Key = glfw.KeyPrintScreen
	KeyNumLock      Key = glfw.KeyNumLock
	KeyCapsLock     Key = glfw.KeyCapsLock
	KeyScrollLock   Key = glfw.KeyScrollLock
	KeyPause        Key = glfw.KeyPause
	KeyLeftSuper    Key = glfw.KeyLeftSuper
	KeyRightSuper   Key = glfw.KeyRightSuper
	KeyMenu         Key = glfw.KeyMenu

	// Keypad
	KeyKP0        Key = glfw.KeyKP0
	KeyKP1        Key = glfw.KeyKP1
	KeyKP2        Key = glfw.KeyKP2
	KeyKP3        Key = glfw.KeyKP3
	KeyKP4        Key = glfw.KeyKP4
	KeyKP5        Key = glfw.KeyKP5
	KeyKP6        Key = glfw.KeyKP6
	KeyKP7        Key = glfw.KeyKP7
	KeyKP8        Key = glfw.KeyKP8
	KeyKP9        Key = glfw.KeyKP9
	KeyKPDivide   Key = glfw.KeyKPDivide
	KeyKPMultiply Key = glfw.KeyKPMultiply
	KeyKPSubtract Key = glfw.KeyKPSubtract
	KeyKPAdd      Key = glfw.KeyKPAdd
	KeyKPDecimal  Key = glfw.KeyKPDecimal
	KeyKPEqual    Key = glfw.KeyKPEqual
	KeyKPEnter    Key = glfw.KeyKPEnter
)

type ModifierKey = glfw.ModifierKey

const (
	ModShift    ModifierKey = glfw.ModShift
	ModControl  ModifierKey = glfw.ModControl
	ModAlt      ModifierKey = glfw.ModAlt
	ModSuper    ModifierKey = glfw.ModSuper
	ModCapsLock ModifierKey = glfw.ModCapsLock
	ModNumLock  ModifierKey = glfw.ModNumLock
)
