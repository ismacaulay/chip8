package display

type Display interface {
	Clear()
	DisplaySprites(x, y uint8, data []uint8) bool
}

type Chip8Display struct {
}

func (d *Chip8Display) Clear() {
}

func (d *Chip8Display) DisplaySprites(x, y uint8, data []uint8) bool {
	return false
}
