package processor

import (
	"github.com/ismacaulay/chip8/pkg/emu/registers"
	"math/rand"
)

type instructionC struct {
	registers registers.ReaderWriter
}

func newInstructionC(r registers.ReaderWriter) *instructionC {
	return &instructionC{r}
}

func (i *instructionC) execute(opcode uint16) {
	vx := extractNibbleOne(opcode)
	value := extractByteOne(opcode)
	randValue := uint8(rand.Intn(256))

	i.registers.SetRegisterValue(vx, value&randValue)
	i.registers.IncrementProgramCounter(uint16(1))
}
