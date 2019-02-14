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
	memory       memory.ReaderWriter
	registers    registers.ReaderWriter
	instructions []instruction
}

func NewProcessor(d display.Writer, k keyboard.Reader, m memory.ReaderWriter, r registers.ReaderWriter) *Processor {
	processor := Processor{
		memory:    m,
		registers: r,
		instructions: []instruction{
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
		},
	}
	return &processor
}

// Cycle executes the next instruction that is stored in memory
func (p *Processor) Cycle() {
	address := p.registers.GetProgramCounter()
	byte0 := p.memory.ReadValue(address)
	byte1 := p.memory.ReadValue(address + 1)
	opcode := (uint16(byte0) << 8) | uint16(byte1)
	instructionType := extractNibbleZero(opcode)
	p.instructions[instructionType].execute(opcode)
}
