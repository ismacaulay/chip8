package processor

import (
	"github.com/ismacaulay/chip8/pkg/emu/keyboard"
	"github.com/ismacaulay/chip8/pkg/emu/memory"
	"github.com/ismacaulay/chip8/pkg/emu/registers"
)

type instructionF struct {
	keyboard  keyboard.Reader
	memory    memory.ReaderWriter
	registers registers.ReaderWriter
}

func newInstructionF(k keyboard.Reader, m memory.ReaderWriter, r registers.ReaderWriter) *instructionF {
	return &instructionF{k, m, r}
}

func (i *instructionF) execute(opcode uint16) {
	subInstruction := extractByteOne(opcode)
	vx := extractNibbleOne(opcode)

	switch subInstruction {
	case 0x07:
		dt := i.registers.GetDelayTimer()
		i.registers.SetRegisterValue(vx, dt)
	case 0x0A:
		i.keyboard.GetKeyPress(func(key uint8) {
			i.registers.SetRegisterValue(vx, key)
		})
	case 0x15:
		vxValue := i.registers.GetRegisterValue(vx)
		i.registers.SetDelayTimer(vxValue)
	case 0x18:
		vxValue := i.registers.GetRegisterValue(vx)
		i.registers.SetSoundTimer(vxValue)
	case 0x1E:
		vxValue := i.registers.GetRegisterValue(vx)
		iValue := i.registers.GetRegisterI()
		i.registers.SetRegisterI(iValue + uint16(vxValue))
	case 0x29:
		vxValue := i.registers.GetRegisterValue(vx)
		addr := i.memory.GetHexDigitAddress(vxValue)
		i.registers.SetRegisterI(addr)
	case 0x33:
		vxValue := i.registers.GetRegisterValue(vx)
		iValue := i.registers.GetRegisterI()
		i.memory.WriteValue(uint16(iValue), uint8((vxValue/100)%10))
		i.memory.WriteValue(uint16(iValue+1), uint8((vxValue/10)%10))
		i.memory.WriteValue(uint16(iValue+2), uint8(vxValue%10))
	case 0x55:
		iValue := i.registers.GetRegisterI()
		for register := uint8(0); register <= vx; register++ {
			value := i.registers.GetRegisterValue(register)
			i.memory.WriteValue(iValue+uint16(register), value)
		}
	case 0x65:
		iValue := i.registers.GetRegisterI()
		for register := uint8(0); register <= vx; register++ {
			value := i.memory.ReadValue(iValue + uint16(register))
			i.registers.SetRegisterValue(register, value)
		}
	}

	i.registers.IncrementProgramCounter(uint16(1))
}
