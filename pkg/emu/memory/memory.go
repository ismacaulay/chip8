package memory

// ReaderWriter provides functions to interact with the memory
type ReaderWriter interface {
	ReadValue(addr uint16) uint8
	WriteValue(addr uint16, value uint8)
	ReadNBytes(start uint16, n uint8) []uint8
	GetHexDigitAddress(digit uint8) uint16
}

// Memory is the memory for the emulation
type Memory struct {
	buffer []uint8
}

// NewMemory creates a new instance of Memory
func NewMemory() *Memory {
	buffer := make([]uint8, 4096)

	// digit sprites
	copy(buffer[0:], []uint8{0xF0, 0x90, 0x90, 0x90, 0xF0})  // 0
	copy(buffer[5:], []uint8{0x20, 0x60, 0x20, 0x20, 0x70})  // 1
	copy(buffer[10:], []uint8{0xF0, 0x10, 0xF0, 0x80, 0xF0}) // 2
	copy(buffer[15:], []uint8{0xF0, 0x10, 0xF0, 0x10, 0xF0}) // 3
	copy(buffer[20:], []uint8{0x90, 0x90, 0xF0, 0x10, 0x10}) // 4
	copy(buffer[25:], []uint8{0xF0, 0x80, 0xF0, 0x10, 0xF0}) // 5
	copy(buffer[30:], []uint8{0xF0, 0x80, 0xF0, 0x90, 0xF0}) // 6
	copy(buffer[35:], []uint8{0xF0, 0x10, 0x20, 0x40, 0x40}) // 7
	copy(buffer[40:], []uint8{0xF0, 0x90, 0xF0, 0x90, 0xF0}) // 8
	copy(buffer[45:], []uint8{0xF0, 0x90, 0xF0, 0x10, 0xF0}) // 9
	copy(buffer[50:], []uint8{0xF0, 0x90, 0xF0, 0x90, 0x90}) // A
	copy(buffer[55:], []uint8{0xE0, 0x90, 0xE0, 0x90, 0xE0}) // B
	copy(buffer[60:], []uint8{0xF0, 0x80, 0x80, 0x80, 0xF0}) // C
	copy(buffer[65:], []uint8{0xE0, 0x90, 0x90, 0x90, 0xE0}) // D
	copy(buffer[70:], []uint8{0xF0, 0x80, 0xF0, 0x80, 0xF0}) // E
	copy(buffer[75:], []uint8{0xF0, 0x80, 0xF0, 0x80, 0x80}) // F

	return &Memory{buffer}
}

// ReadValue returns byte at given address
func (m *Memory) ReadValue(addr uint16) uint8 {
	return m.buffer[addr]
}

// WriteValue writes value to given address
func (m *Memory) WriteValue(addr uint16, value uint8) {
	m.buffer[addr] = value
}

// ReadNBytes reads n bytes starting from start address
func (m *Memory) ReadNBytes(start uint16, n uint8) []uint8 {
	return m.buffer[start : start+uint16(n)]
}

// GetHexDigitAddress returns the address of the first byte of the
// requested digit sprite
func (m *Memory) GetHexDigitAddress(digit uint8) uint16 {
	return uint16(digit * 5)
}
