package display

type Writer interface {
	Clear()
	DisplaySprites(x, y uint8, data []uint8) bool
}

type Display struct {
}

func (d *Display) Clear() {
}

func (d *Display) DisplaySprites(x, y uint8, data []uint8) bool {
	return false
}
