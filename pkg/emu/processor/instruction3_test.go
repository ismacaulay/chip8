package processor

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	mock_registers "github.com/ismacaulay/chip8/pkg/emu/registers/mock"
)

func TestInstruction3(t *testing.T) {
	ctrl := gomock.NewController(t)
	registers := mock_registers.NewMockReaderWriter(ctrl)
	instruction := newInstruction3(registers)

	cases := []struct {
		vxValue, kkValue  uint8
		expectedIncrement uint16
	}{
		{uint8(42), uint8(42), 2},
		{uint8(42), uint8(43), 1},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("[3xkk] Skip next instruction if Vx = kk. (vx: %d, kk: %d)", c.vxValue, c.kkValue), func(t *testing.T) {
			registers.EXPECT().GetRegisterValue(uint8(0x01)).Return(c.vxValue)
			registers.EXPECT().IncrementProgramCounter(c.expectedIncrement)

			opcode := uint16(0x100) | uint16(c.kkValue)

			instruction.execute(opcode)
		})
	}

	ctrl.Finish()
}
