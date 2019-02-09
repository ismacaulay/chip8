package processor

func extractNibbleZero(value uint16) uint8 {
	return uint8((value & 0xF000) >> 12)
}

func extractNibbleOne(value uint16) uint8 {
	return uint8((value & 0x0F00) >> 8)
}

func extractNibbleTwo(value uint16) uint8 {
	return uint8((value & 0x00F0) >> 4)
}

func extractNibbleThree(value uint16) uint8 {
	return uint8(value & 0x000F)
}

func extractByteOne(value uint16) uint8 {
	return uint8(value & 0x00FF)
}
