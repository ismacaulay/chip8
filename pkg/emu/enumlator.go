package emu

import (
	"time"

	"github.com/ismacaulay/chip8/pkg/emu/display"
	"github.com/ismacaulay/chip8/pkg/emu/keyboard"
	"github.com/ismacaulay/chip8/pkg/emu/memory"
	"github.com/ismacaulay/chip8/pkg/emu/processor"
	"github.com/ismacaulay/chip8/pkg/emu/registers"
	"github.com/ismacaulay/chip8/pkg/emu/timers"
)

// Emulator is the core of the chip8 emulation
type Emulator struct {
	display   *display.Display
	keyboard  *keyboard.Keyboard
	memory    *memory.Memory
	registers *registers.Registers
	timers    *timers.Timers
	processor *processor.Processor
}

// NewEmulator creates a new instance of the emulator
func NewEmulator() *Emulator {
	display := display.NewDisplay()
	keyboard := keyboard.NewKeyboard()
	memory := memory.NewMemory()
	registers := registers.NewRegisters()

	timers := timers.NewTimers(registers)
	processor := processor.NewProcessor(display, keyboard, memory, registers)

	return &Emulator{display, keyboard, memory, registers, timers, processor}
}

// LoadRom loads the rom data into memory and initialized the emulator
func (e *Emulator) LoadRom(data []uint8) {
	e.memory.WriteBytes(0x200, data)
	e.registers.SetProgramCounter(0x200)
}

// Run stops the emulator at 60Hz
func (e *Emulator) Run() {
	go func() {
		clock := time.NewTicker(time.Second / 60)

		for {
			<-clock.C
			e.Step()
		}
	}()
}

// Step the emulation by 1 cycle
func (e *Emulator) Step() {
	if !e.keyboard.WaitingForKeyPress() {
		for i := 0; i < 10; i++ {
			e.processor.Cycle()
		}
		e.timers.Step()
	}
}

func (e *Emulator) HandleKeyEvent(event keyboard.Event, key keyboard.Key) {
	e.keyboard.HandleKeyEvent(event, key)
}

func (e *Emulator) ListenForDisplayUpdate(cb func([]uint8)) {
	e.display.AddObserver(cb)
}
