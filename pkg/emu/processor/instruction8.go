package processor

import (
	"github.com/ismacaulay/chip8/pkg/emu/registers"
)

type instruction8 struct {
	registers registers.ReaderWriter
}

func newInstruction8(r registers.ReaderWriter) *instruction8 {
	return &instruction8{r}
}

func (i *instruction8) execute(opcode uint16) {
	vx := extractNibbleOne(opcode)
	vxValue := i.registers.GetRegisterValue(vx)
	vy := extractNibbleTwo(opcode)
	vyValue := i.registers.GetRegisterValue(vy)
	subInstruction := extractNibbleThree(opcode)

	switch subInstruction {
	case 0x0:
		i.registers.SetRegisterValue(vx, vyValue)
	case 0x1:
		i.registers.SetRegisterValue(vx, vxValue|vyValue)
	case 0x2:
		i.registers.SetRegisterValue(vx, vxValue&vyValue)
	case 0x3:
		i.registers.SetRegisterValue(vx, vxValue^vyValue)
	case 0x4:
		vfValue := 0
		if vyValue > 0xFF-vxValue {
			vfValue = 1
		}
		i.registers.SetRegisterValue(0x0F, uint8(vfValue))
		i.registers.SetRegisterValue(vx, vxValue+vyValue)
	case 0x5:
		vfValue := 0
		if vxValue > vyValue {
			vfValue = 1
		}
		i.registers.SetRegisterValue(0x0F, uint8(vfValue))
		i.registers.SetRegisterValue(vx, vxValue-vyValue)
	case 0x6:
		vfValue := 0
		if (vxValue & 0x01) == 0x01 {
			vfValue = 1
		}
		i.registers.SetRegisterValue(0x0F, uint8(vfValue))
		i.registers.SetRegisterValue(vx, vxValue>>1)
	case 0x7:
		vfValue := 0
		if vyValue > vxValue {
			vfValue = 1
		}
		i.registers.SetRegisterValue(0x0F, uint8(vfValue))
		i.registers.SetRegisterValue(vx, vyValue-vxValue)
	case 0xE:
		vfValue := 0
		if (vxValue & 0x80) == 0x80 {
			vfValue = 1
		}
		i.registers.SetRegisterValue(0x0F, uint8(vfValue))
		i.registers.SetRegisterValue(vx, vxValue<<1)
	}

	i.registers.IncrementProgramCounter(uint16(1))
}
