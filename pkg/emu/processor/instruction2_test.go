package processor

import (
	"testing"

	"github.com/golang/mock/gomock"
	mock_registers "github.com/ismacaulay/chip8/pkg/emu/registers/mock"
)

func TestInstruction2(t *testing.T) {
	ctrl := gomock.NewController(t)
	registers := mock_registers.NewMockReaderWriter(ctrl)
	instruction := newInstruction2(registers)

	t.Run("[2nnn] Call subroutine at nnn", func(t *testing.T) {
		gomock.InOrder(
			registers.EXPECT().PushProgramCounter(),
			registers.EXPECT().SetProgramCounter(uint16(0x0123)),
		)

		opcode := uint16(0x123)

		instruction.execute(opcode)
	})

	ctrl.Finish()
}
