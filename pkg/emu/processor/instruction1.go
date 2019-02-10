package processor

import (
	"github.com/ismacaulay/chip8/pkg/emu/registers"
)

type instruction1 struct {
	registers registers.RegisterReaderWriter
}

func newInstruction1(r registers.RegisterReaderWriter) *instruction1 {
	return &instruction1{r}
}

func (i *instruction1) execute(opcode uint16) {
	address := opcode & 0x0FFF
	i.registers.SetProgramCounter(address)
}
