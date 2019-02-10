package processor

import (
	"github.com/ismacaulay/chip8/pkg/emu/registers"
)

type instructionA struct {
	registers registers.RegisterReaderWriter
}

func newInstructionA(r registers.RegisterReaderWriter) *instructionA {
	return &instructionA{r}
}

func (i *instructionA) execute(opcode uint16) {
	value := opcode & 0x0FFF
	i.registers.SetRegisterI(value)
	i.registers.IncrementProgramCounter(uint16(1))
}
