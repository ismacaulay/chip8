package processor

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"

	mock_registers "github.com/ismacaulay/chip8/pkg/emu/registers/mock"
)

func TestInstruction8(t *testing.T) {
	ctrl := gomock.NewController(t)
	registers := mock_registers.NewMockRegisters(ctrl)
	instruction := newInstruction8(registers)

	t.Run("[8xy0] Set Vx = Vy", func(t *testing.T) {
		vxValue := uint8(42)
		vyValue := uint8(103)
		registers.EXPECT().GetRegisterValue(uint8(0x01)).Return(vxValue)
		registers.EXPECT().GetRegisterValue(uint8(0x0C)).Return(vyValue)
		registers.EXPECT().SetRegisterValue(uint8(0x01), vyValue)
		registers.EXPECT().IncrementProgramCounter(1)

		opcode := uint16(0x1C0)

		instruction.execute(opcode)
	})

	t.Run("[8xy1] Set Vx = Vx OR Vy", func(t *testing.T) {
		vxValue := uint8(42)
		vyValue := uint8(103)
		registers.EXPECT().GetRegisterValue(uint8(0x01)).Return(vxValue)
		registers.EXPECT().GetRegisterValue(uint8(0x0C)).Return(vyValue)
		registers.EXPECT().SetRegisterValue(uint8(0x01), vxValue|vyValue)
		registers.EXPECT().IncrementProgramCounter(1)

		opcode := uint16(0x1C1)

		instruction.execute(opcode)
	})

	t.Run("[8xy2] Set Vx = Vx AND Vy", func(t *testing.T) {
		vxValue := uint8(42)
		vyValue := uint8(103)
		registers.EXPECT().GetRegisterValue(uint8(0x01)).Return(vxValue)
		registers.EXPECT().GetRegisterValue(uint8(0x0C)).Return(vyValue)
		registers.EXPECT().SetRegisterValue(uint8(0x01), vxValue&vyValue)
		registers.EXPECT().IncrementProgramCounter(1)

		opcode := uint16(0x1C2)

		instruction.execute(opcode)
	})

	t.Run("[8xy3] Set Vx = Vx XOR Vy", func(t *testing.T) {
		vxValue := uint8(42)
		vyValue := uint8(103)
		registers.EXPECT().GetRegisterValue(uint8(0x01)).Return(vxValue)
		registers.EXPECT().GetRegisterValue(uint8(0x0C)).Return(vyValue)
		registers.EXPECT().SetRegisterValue(uint8(0x01), vxValue^vyValue)
		registers.EXPECT().IncrementProgramCounter(1)

		opcode := uint16(0x1C3)

		instruction.execute(opcode)
	})

	cases := []struct {
		vxValue, vyValue uint8
		expectedVf       uint8
	}{
		{uint8(1), uint8(254), uint8(0)},
		{uint8(1), uint8(255), uint8(1)},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("[8xy4] Set Vx = Vx + Vy, set VF = carry. (vx: %d, vy: %d)", c.vxValue, c.vyValue), func(t *testing.T) {
			registers.EXPECT().GetRegisterValue(uint8(0x01)).Return(c.vxValue)
			registers.EXPECT().GetRegisterValue(uint8(0x0C)).Return(c.vyValue)
			registers.EXPECT().SetRegisterValue(uint8(0x01), uint8(c.vxValue+c.vyValue))
			registers.EXPECT().SetRegisterValue(uint8(0x0F), c.expectedVf)
			registers.EXPECT().IncrementProgramCounter(1)

			opcode := uint16(0x1C4)

			instruction.execute(opcode)
		})
	}

	cases = []struct {
		vxValue, vyValue uint8
		expectedVf       uint8
	}{
		{uint8(103), uint8(255), uint8(0)},
		{uint8(255), uint8(42), uint8(1)},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("[8xy5] Set Vx = Vx - Vy, set VF = NOT borrow. (vx: %d, vy: %d)", c.vxValue, c.vyValue), func(t *testing.T) {
			registers.EXPECT().GetRegisterValue(uint8(0x01)).Return(c.vxValue)
			registers.EXPECT().GetRegisterValue(uint8(0x0C)).Return(c.vyValue)
			registers.EXPECT().SetRegisterValue(uint8(0x01), uint8(c.vxValue-c.vyValue))
			registers.EXPECT().SetRegisterValue(uint8(0x0F), c.expectedVf)
			registers.EXPECT().IncrementProgramCounter(1)

			opcode := uint16(0x1C5)

			instruction.execute(opcode)
		})
	}

	cases = []struct {
		vxValue, vyValue uint8
		expectedVf       uint8
	}{
		{uint8(0x11), uint8(255), uint8(1)},
		{uint8(0x10), uint8(255), uint8(0)},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("[8xy6] Set Vx = Vx SHR 1, (vx: %d)", c.vxValue), func(t *testing.T) {
			registers.EXPECT().GetRegisterValue(uint8(0x01)).Return(c.vxValue)
			registers.EXPECT().GetRegisterValue(uint8(0x0C)).Return(c.vyValue)
			registers.EXPECT().SetRegisterValue(uint8(0x01), uint8(c.vxValue>>1))
			registers.EXPECT().SetRegisterValue(uint8(0x0F), c.expectedVf)
			registers.EXPECT().IncrementProgramCounter(1)

			opcode := uint16(0x1C6)

			instruction.execute(opcode)
		})
	}

	cases = []struct {
		vxValue, vyValue uint8
		expectedVf       uint8
	}{
		{uint8(255), uint8(103), uint8(0)},
		{uint8(42), uint8(255), uint8(1)},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("[8xy7] Set Vx = Vy - Vx, set VF = NOT borrow, (vx: %d, vy: %d)", c.vxValue, c.vyValue), func(t *testing.T) {
			registers.EXPECT().GetRegisterValue(uint8(0x01)).Return(c.vxValue)
			registers.EXPECT().GetRegisterValue(uint8(0x0C)).Return(c.vyValue)
			registers.EXPECT().SetRegisterValue(uint8(0x01), uint8(c.vyValue-c.vxValue))
			registers.EXPECT().SetRegisterValue(uint8(0x0F), c.expectedVf)
			registers.EXPECT().IncrementProgramCounter(1)

			opcode := uint16(0x1C7)

			instruction.execute(opcode)
		})
	}

	cases = []struct {
		vxValue, vyValue uint8
		expectedVf       uint8
	}{
		{uint8(0x83), uint8(255), uint8(1)},
		{uint8(0x18), uint8(255), uint8(0)},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("[8xyE] Set Vx = Vx SHL 1, (vx: %d)", c.vxValue), func(t *testing.T) {
			registers.EXPECT().GetRegisterValue(uint8(0x01)).Return(c.vxValue)
			registers.EXPECT().GetRegisterValue(uint8(0x0C)).Return(c.vyValue)
			registers.EXPECT().SetRegisterValue(uint8(0x01), uint8(c.vxValue<<1))
			registers.EXPECT().SetRegisterValue(uint8(0x0F), c.expectedVf)
			registers.EXPECT().IncrementProgramCounter(1)

			opcode := uint16(0x1CE)

			instruction.execute(opcode)
		})
	}

	ctrl.Finish()
}
