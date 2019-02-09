package processor

import (
	"github.com/ismacaulay/chip8/pkg/emu/registers"
)

type instruction3 struct {
	registers registers.Registers
}

func newInstruction3(r registers.Registers) *instruction3 {
	return &instruction3{r}
}

func (i *instruction3) execute(opcode uint16) {
	vx := extractNibbleOne(opcode)
	value := extractByteOne(opcode)
	vxValue := i.registers.GetRegisterValue(vx)

	if value == vxValue {
		i.registers.IncrementProgramCounter(2)
	} else {
		i.registers.IncrementProgramCounter(1)
	}
}
