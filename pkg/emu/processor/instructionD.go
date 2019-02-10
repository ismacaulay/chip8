package processor

import (
	"github.com/ismacaulay/chip8/pkg/emu/display"
	"github.com/ismacaulay/chip8/pkg/emu/memory"
	"github.com/ismacaulay/chip8/pkg/emu/registers"
)

type instructionD struct {
	display   display.Writer
	memory    memory.ReaderWriter
	registers registers.ReaderWriter
}

func newInstructionD(d display.Writer, m memory.ReaderWriter, r registers.ReaderWriter) *instructionD {
	return &instructionD{d, m, r}
}

func (i *instructionD) execute(opcode uint16) {
	vx := extractNibbleOne(opcode)
	vxValue := i.registers.GetRegisterValue(vx)
	vy := extractNibbleTwo(opcode)
	vyValue := i.registers.GetRegisterValue(vy)
	n := extractNibbleThree(opcode)
	start := i.registers.GetRegisterI()

	data := i.memory.ReadNBytes(start, n)
	collision := i.display.DisplaySprites(vxValue, vyValue, data)

	if collision {
		i.registers.SetRegisterValue(0x0F, uint8(1))
	} else {
		i.registers.SetRegisterValue(0x0F, uint8(0))
	}

	i.registers.IncrementProgramCounter(uint16(1))
}
