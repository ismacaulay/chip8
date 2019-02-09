package keyboard

type Keyboard interface {
	IsPressed(key uint8) bool
	GetKeyPress() uint8
}
