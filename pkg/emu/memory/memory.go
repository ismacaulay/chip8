package memory

// ReaderWriter provides functions to interact with the memory
type ReaderWriter interface {
	ReadValue(addr uint16) uint8
	WriteValue(addr uint16, value uint8)
	ReadNBytes(start uint16, n uint8) []uint8
	GetHexDigitAddress(digit uint8) uint16
}
