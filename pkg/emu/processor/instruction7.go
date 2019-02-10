package processor

import (
	"github.com/ismacaulay/chip8/pkg/emu/registers"
)

type instruction7 struct {
	registers registers.RegisterReaderWriter
}

func newInstruction7(r registers.RegisterReaderWriter) *instruction7 {
	return &instruction7{r}
}

func (i *instruction7) execute(opcode uint16) {
	vx := extractNibbleOne(opcode)
	value := extractByteOne(opcode)
	vxValue := i.registers.GetRegisterValue(vx)

	i.registers.SetRegisterValue(vx, value+vxValue)
	i.registers.IncrementProgramCounter(uint16(1))
}
