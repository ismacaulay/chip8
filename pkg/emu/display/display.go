package display

// Writer provides functions to write to the display
type Writer interface {
	Clear()
	DisplaySprites(x, y uint8, data []uint8) bool
}

// Display is the display for the emulation
type Display struct {
	width, height uint
	buffer        []uint8
	observers     []func([]uint8)
}

// NewDisplay creates a new instance of Display
func NewDisplay() *Display {
	width := uint(64)
	height := uint(32)
	return &Display{
		width:     uint(width),
		height:    uint(height),
		buffer:    make([]uint8, width*height),
		observers: make([]func([]uint8), 0),
	}
}

// Buffer returns a copy of the current display buffer
func (d *Display) Buffer() []uint8 {
	buf := make([]uint8, len(d.buffer))
	copy(buf, d.buffer)
	return buf
}

// Width returns the width of the display
func (d *Display) Width() uint {
	return d.width
}

// Height returns the height of the display
func (d *Display) Height() uint {
	return d.height
}

// Clear clears the display
func (d *Display) Clear() {
	d.buffer = make([]uint8, d.width*d.height)
}

// DisplaySprites writes data to the display. Returns true if a collision happens
func (d *Display) DisplaySprites(x, y uint8, data []uint8) bool {
	pixelCleared := false
	row := uint(y)
	col := uint(x)

	for _, b := range data {
		for i := 7; i >= 0; i-- {
			pixelIdx := (d.width * row) + col
			if (1<<uint8(i))&b != 0 {
				currentPixel := d.buffer[pixelIdx]
				newPixel := 1 ^ currentPixel

				if !pixelCleared && newPixel == 0 {
					pixelCleared = true
				}

				d.buffer[pixelIdx] = newPixel
			}

			col++
			if col >= d.width {
				col = 0
			}
		}

		row++
		col = uint(x)
	}

	for _, observer := range d.observers {
		observer(d.Buffer())
	}
	return pixelCleared
}

func (d *Display) AddObserver(cb func([]uint8)) {
	d.observers = append(d.observers, cb)
}
