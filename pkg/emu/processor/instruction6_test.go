package processor

import (
	"testing"

	"github.com/golang/mock/gomock"
	mock_registers "github.com/ismacaulay/chip8/pkg/emu/registers/mock"
)

func TestInstruction6(t *testing.T) {
	ctrl := gomock.NewController(t)
	registers := mock_registers.NewMockReaderWriter(ctrl)
	instruction := newInstruction6(registers)

	t.Run("[6xkk] Set Vx = kk", func(t *testing.T) {
		value := uint8(42)
		registers.EXPECT().SetRegisterValue(uint8(0x01), value)
		registers.EXPECT().IncrementProgramCounter(uint16(1))

		opcode := uint16(0x100) | uint16(value)

		instruction.execute(opcode)
	})

	ctrl.Finish()
}
