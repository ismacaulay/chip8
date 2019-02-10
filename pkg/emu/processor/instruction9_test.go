package processor

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"

	mock_registers "github.com/ismacaulay/chip8/pkg/emu/registers/mock"
)

func TestInstruction9(t *testing.T) {
	ctrl := gomock.NewController(t)
	registers := mock_registers.NewMockReaderWriter(ctrl)
	instruction := newInstruction9(registers)

	cases := []struct {
		vxValue, vyValue  uint8
		expectedIncrement uint16
	}{
		{uint8(42), uint8(43), 2},
		{uint8(42), uint8(42), 1},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("[9xy0] Skip next instruction if Vx != Vy. (vx: %d, vy: %d)", c.vxValue, c.vyValue), func(t *testing.T) {
			registers.EXPECT().GetRegisterValue(uint8(0x01)).Return(c.vxValue)
			registers.EXPECT().GetRegisterValue(uint8(0x0C)).Return(c.vyValue)
			registers.EXPECT().IncrementProgramCounter(c.expectedIncrement)

			opcode := uint16(0x1C0)

			instruction.execute(opcode)
		})
	}

	ctrl.Finish()
}
