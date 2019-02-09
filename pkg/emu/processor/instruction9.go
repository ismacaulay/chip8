package processor

import (
	"github.com/ismacaulay/chip8/pkg/emu/registers"
)

type instruction9 struct {
	registers registers.Registers
}

func newInstruction9(r registers.Registers) *instruction9 {
	return &instruction9{r}
}

func (i *instruction9) execute(opcode uint16) {
	vx := extractNibbleOne(opcode)
	vy := extractNibbleTwo(opcode)
	vxValue := i.registers.GetRegisterValue(vx)
	vyValue := i.registers.GetRegisterValue(vy)

	if vxValue != vyValue {
		i.registers.IncrementProgramCounter(2)
	} else {
		i.registers.IncrementProgramCounter(1)
	}
}
