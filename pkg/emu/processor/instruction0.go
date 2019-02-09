package processor

import (
	"github.com/ismacaulay/chip8/pkg/emu/display"
	"github.com/ismacaulay/chip8/pkg/emu/registers"
)

type instruction0 struct {
	display   display.Display
	registers registers.Registers
}

func newInstruction0(d display.Display, r registers.Registers) *instruction0 {
	return &instruction0{d, r}
}

func (i *instruction0) execute(opcode uint16) {
	switch opcode {
	case 0x00E0:
		i.display.Clear()
		i.registers.IncrementProgramCounter(1)
	case 0x00EE:
		address := i.registers.PopProgramCounter()
		i.registers.SetProgramCounter(address)
	default:
		i.registers.IncrementProgramCounter(1)
	}
}
