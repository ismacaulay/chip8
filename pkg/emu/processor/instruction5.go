package processor

import (
	"github.com/ismacaulay/chip8/pkg/emu/registers"
)

type instruction5 struct {
	registers registers.Registers
}

func newInstruction5(r registers.Registers) *instruction5 {
	return &instruction5{r}
}

func (i *instruction5) execute(opcode uint16) {
	vx := extractNibbleOne(opcode)
	vy := extractNibbleTwo(opcode)
	vxValue := i.registers.GetRegisterValue(vx)
	vyValue := i.registers.GetRegisterValue(vy)

	if vxValue == vyValue {
		i.registers.IncrementProgramCounter(2)
	} else {
		i.registers.IncrementProgramCounter(1)
	}
}
