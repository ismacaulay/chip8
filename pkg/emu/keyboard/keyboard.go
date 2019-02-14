package keyboard

type Reader interface {
	IsPressed(key uint8) bool
	GetKeyPress(cb func(uint8))
}

type Event uint8
type Key uint8

const (
	EventKeyPressed Event = iota
	EventKeyReleased
)

const (
	Key0 Key = iota
	Key1
	Key2
	Key3
	Key4
	Key5
	Key6
	Key7
	Key8
	Key9
	KeyA
	KeyB
	KeyC
	KeyD
	KeyE
	KeyF
)

// Keyboard is a hexidecimal keyboard
type Keyboard struct {
	state              uint16
	waitingForKeyPress bool
	inputChan          chan Key
}

// NewKeyboard returns a new instance of a Keyboard
func NewKeyboard() *Keyboard {
	return &Keyboard{uint16(0), false, make(chan Key)}
}

// HandleKeyEvent updates the keyboard state
func (k *Keyboard) HandleKeyEvent(event Event, key Key) {
	switch event {
	case EventKeyPressed:
		k.state = k.state | (uint16(1) << key)
		if k.waitingForKeyPress {
			k.inputChan <- key
		}
	case EventKeyReleased:
		k.state = k.state &^ (uint16(1) << key)
	}
}

// WaitingForKeyPress returns whether the keyboard is blocking for a keypress
func (k *Keyboard) WaitingForKeyPress() bool {
	return k.waitingForKeyPress
}

// IsPressed returns whether a given key is pressed
func (k *Keyboard) IsPressed(key uint8) bool {
	return k.state&(uint16(1)<<key) > 0
}

// GetKeyPress blocks until a key is pressed
func (k *Keyboard) GetKeyPress(cb func(uint8)) {
	k.waitingForKeyPress = true
	go func() {
		key := <-k.inputChan
		cb(uint8(key))
		k.waitingForKeyPress = false
	}()
}
