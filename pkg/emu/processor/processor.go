package processor

import (
	"github.com/ismacaulay/chip8/pkg/emu/display"
	"github.com/ismacaulay/chip8/pkg/emu/keyboard"
	"github.com/ismacaulay/chip8/pkg/emu/memory"
	"github.com/ismacaulay/chip8/pkg/emu/registers"
)

type instruction interface {
	execute(opcode uint16)
}

type Processor struct {
	instructions []instruction
}

func NewProcessor(d display.Display, k keyboard.Keyboard, m memory.Memory, r registers.Registers) *Processor {
	processor := Processor{}
	processor.instructions = []instruction{
		newInstruction0(d, r),
		newInstruction1(r),
		newInstruction2(r),
		newInstruction3(r),
		newInstruction4(r),
		newInstruction5(r),
		newInstruction6(r),
		newInstruction7(r),
		newInstruction8(r),
		newInstruction9(r),
		newInstructionA(r),
		newInstructionB(r),
		newInstructionC(r),
		newInstructionD(d, m, r),
		newInstructionE(k, r),
		newInstructionF(k, m, r),
	}
	return &processor
}

func (p *Processor) ProcessOpcode(opcode uint16) {
	instructionType := extractNibbleZero(opcode)
	p.instructions[instructionType].execute(opcode)
}
