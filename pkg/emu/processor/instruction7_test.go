package processor

import (
	"testing"

	"github.com/golang/mock/gomock"
	mock_registers "github.com/ismacaulay/chip8/pkg/emu/registers/mock"
)

func TestInstruction7(t *testing.T) {
	ctrl := gomock.NewController(t)
	registers := mock_registers.NewMockReaderWriter(ctrl)
	instruction := newInstruction7(registers)

	t.Run("[7xkk] Set Vx = Vx + kk", func(t *testing.T) {
		vxValue := uint8(42)
		value := uint8(103)
		registers.EXPECT().GetRegisterValue(uint8(0x01)).Return(vxValue)
		registers.EXPECT().SetRegisterValue(uint8(0x01), vxValue+value)
		registers.EXPECT().IncrementProgramCounter(uint16(1))

		opcode := uint16(0x100) | uint16(value)

		instruction.execute(opcode)
	})

	ctrl.Finish()
}
