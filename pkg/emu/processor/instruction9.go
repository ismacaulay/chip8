package processor

import (
	"github.com/ismacaulay/chip8/pkg/emu/registers"
)

type instruction9 struct {
	registers registers.ReaderWriter
}

func newInstruction9(r registers.ReaderWriter) *instruction9 {
	return &instruction9{r}
}

func (i *instruction9) execute(opcode uint16) {
	vx := extractNibbleOne(opcode)
	vy := extractNibbleTwo(opcode)
	vxValue := i.registers.GetRegisterValue(vx)
	vyValue := i.registers.GetRegisterValue(vy)

	if vxValue != vyValue {
		i.registers.IncrementProgramCounter(uint16(2))
	} else {
		i.registers.IncrementProgramCounter(uint16(1))
	}
}
