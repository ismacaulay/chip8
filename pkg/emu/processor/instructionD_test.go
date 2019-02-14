package processor

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"

	mock_display "github.com/ismacaulay/chip8/pkg/emu/display/mock"
	mock_memory "github.com/ismacaulay/chip8/pkg/emu/memory/mock"
	mock_registers "github.com/ismacaulay/chip8/pkg/emu/registers/mock"
)

func TestInstructionD(t *testing.T) {
	ctrl := gomock.NewController(t)
	display := mock_display.NewMockWriter(ctrl)
	memory := mock_memory.NewMockReaderWriter(ctrl)
	registers := mock_registers.NewMockReaderWriter(ctrl)
	instruction := newInstructionD(display, memory, registers)

	cases := []struct {
		collision  bool
		expectedVf uint8
	}{
		{true, uint8(1)},
		{false, uint8(0)},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("[Dxyn] Display n-byte sprite starting at memory location I at (Vx, Vy), set VF = collision, (collision: %t)", c.collision), func(t *testing.T) {
			n := uint8(12)
			vx := uint8(0x2)
			vxValue := uint8(53)
			vy := uint8(0xA)
			vyValue := uint8(27)
			iValue := uint16(1024)
			data := []uint8{23, 123, 255, 42, 0, 27}

			registers.EXPECT().GetRegisterValue(vx).Return(vxValue)
			registers.EXPECT().GetRegisterValue(vy).Return(vyValue)
			registers.EXPECT().GetRegisterI().Return(iValue)
			memory.EXPECT().ReadBytes(iValue, n).Return(data)
			display.EXPECT().DisplaySprites(vxValue, vyValue, data).Return(c.collision)
			registers.EXPECT().SetRegisterValue(uint8(0x0F), c.expectedVf)
			registers.EXPECT().IncrementProgramCounter(uint16(1))

			opcode := uint16(0x2A0) | uint16(n)

			instruction.execute(opcode)
		})
	}

	ctrl.Finish()
}
