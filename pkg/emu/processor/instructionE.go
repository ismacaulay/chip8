package processor

import (
	"github.com/ismacaulay/chip8/pkg/emu/keyboard"
	"github.com/ismacaulay/chip8/pkg/emu/registers"
)

type instructionE struct {
	keyboard  keyboard.Keyboard
	registers registers.ReaderWriter
}

func newInstructionE(k keyboard.Keyboard, r registers.ReaderWriter) *instructionE {
	return &instructionE{k, r}
}

func (i *instructionE) execute(opcode uint16) {
	subInstruction := extractByteOne(opcode)

	vx := extractNibbleOne(opcode)
	vxValue := i.registers.GetRegisterValue(vx)

	switch subInstruction {
	case 0x9E:
		if i.keyboard.IsPressed(vxValue) {
			i.registers.IncrementProgramCounter(uint16(2))
		} else {
			i.registers.IncrementProgramCounter(uint16(1))
		}
	case 0xA1:
		if !i.keyboard.IsPressed(vxValue) {
			i.registers.IncrementProgramCounter(uint16(2))
		} else {
			i.registers.IncrementProgramCounter(uint16(1))
		}
	}

}
