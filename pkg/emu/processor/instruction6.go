package processor

import (
	"github.com/ismacaulay/chip8/pkg/emu/registers"
)

type instruction6 struct {
	registers registers.ReaderWriter
}

func newInstruction6(r registers.ReaderWriter) *instruction6 {
	return &instruction6{r}
}

func (i *instruction6) execute(opcode uint16) {
	vx := extractNibbleOne(opcode)
	value := extractByteOne(opcode)

	i.registers.SetRegisterValue(vx, value)
	i.registers.IncrementProgramCounter(uint16(1))
}
