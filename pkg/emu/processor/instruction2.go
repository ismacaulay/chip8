package processor

import (
	"github.com/ismacaulay/chip8/pkg/emu/registers"
)

type instruction2 struct {
	registers registers.Registers
}

func newInstruction2(r registers.Registers) *instruction2 {
	return &instruction2{r}
}

func (i *instruction2) execute(opcode uint16) {
	address := opcode & 0x0FFF
	i.registers.PushProgramCounter()
	i.registers.SetProgramCounter(address)
}
