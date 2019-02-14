package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/gdamore/tcell"
	chip8 "github.com/ismacaulay/chip8/pkg/emu"
	"github.com/ismacaulay/chip8/pkg/emu/keyboard"
	"github.com/rivo/tview"
)

func main() {
	display := NewDisplay()
	display.SetBorder(true)
	display.SetRect(0, 0, 64, 32)

	grid := tview.NewGrid().SetRows(-1, 34, -1).SetColumns(-1, 66, -1)
	grid.AddItem(display, 1, 1, 1, 1, 0, 100, true)

	app := tview.NewApplication().SetRoot(grid, true)

	// path := "/roms/programs/Keypad Test [Hap, 2006].ch8"
	// path := "/roms/games/Tic-Tac-Toe [David Winter].ch8"
	path := "/roms/games/Tetris [Fran Dachille, 1991].ch8"
	// path := "/roms/games/Space Invaders [David Winter].ch8"
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Error: unable to read file", path)
		os.Exit(1)
	}

	emulator := chip8.NewEmulator()
	emulator.LoadRom(data)

	var clearEventTimer *time.Timer
	var prevKey keyboard.Key
	display.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		keys := map[rune]keyboard.Key{
			'1': keyboard.Key1,
			'2': keyboard.Key2,
			'3': keyboard.Key3,
			'4': keyboard.KeyC,
			'q': keyboard.Key4,
			'w': keyboard.Key5,
			'e': keyboard.Key6,
			'r': keyboard.KeyD,
			'a': keyboard.Key7,
			's': keyboard.Key8,
			'd': keyboard.Key9,
			'f': keyboard.KeyE,
			'z': keyboard.KeyA,
			'x': keyboard.Key0,
			'c': keyboard.KeyB,
			'v': keyboard.KeyF,
		}

		key := keys[event.Rune()]
		if clearEventTimer != nil {
			emulator.HandleKeyEvent(keyboard.EventKeyReleased, prevKey)
			clearEventTimer.Reset(100 * time.Millisecond)
		} else {
			clearEventTimer = time.AfterFunc(100*time.Millisecond, func() {
				emulator.HandleKeyEvent(keyboard.EventKeyReleased, prevKey)
			})
		}
		emulator.HandleKeyEvent(keyboard.EventKeyPressed, key)
		prevKey = key

		return event
	})

	emulator.ListenForDisplayUpdate(func(buf []uint8) {
		app.QueueUpdateDraw(func() {
			display.SetBuffer(buf)
		})
	})
	emulator.Run()

	if err := app.Run(); err != nil {
		panic(err)
	}
}

type Display struct {
	*tview.Box
	buffer []uint8
}

func NewDisplay() *Display {
	return &Display{
		Box:    tview.NewBox(),
		buffer: make([]uint8, 64*32),
	}
}

func (d *Display) Draw(screen tcell.Screen) {
	d.Box.Draw(screen)
	x, y, _, _ := d.GetInnerRect()
	st := tcell.StyleDefault

	for row := 0; row < 32; row++ {
		for col := 0; col < 64; col++ {
			pixel := d.buffer[(row*64)+col]
			if pixel == 1 {
				screen.SetContent(x+col, y+row, ' ', nil, st.Background(tcell.ColorWhite))
			} else {
				screen.SetContent(x+col, y+row, ' ', nil, st.Background(tcell.ColorBlack))
			}
		}
	}
}

func (d *Display) SetBuffer(buf []uint8) {
	copy(d.buffer, buf)
}
