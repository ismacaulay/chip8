package processor

import (
	"github.com/ismacaulay/chip8/pkg/emu/registers"
)

type instructionB struct {
	registers registers.ReaderWriter
}

func newInstructionB(r registers.ReaderWriter) *instructionB {
	return &instructionB{r}
}

func (i *instructionB) execute(opcode uint16) {
	value := opcode & 0x0FFF
	v0 := i.registers.GetRegisterValue(0x0)
	i.registers.SetProgramCounter(value + uint16(v0))
}
