package processor

import (
	"github.com/ismacaulay/chip8/pkg/emu/registers"
)

type instruction4 struct {
	registers registers.ReaderWriter
}

func newInstruction4(r registers.ReaderWriter) *instruction4 {
	return &instruction4{r}
}

func (i *instruction4) execute(opcode uint16) {
	vx := extractNibbleOne(opcode)
	value := extractByteOne(opcode)
	vxValue := i.registers.GetRegisterValue(vx)

	if value != vxValue {
		i.registers.IncrementProgramCounter(uint16(2))
	} else {
		i.registers.IncrementProgramCounter(uint16(1))
	}
}
