package processor

import (
	"math/rand"
	"testing"

	"github.com/golang/mock/gomock"

	mock_registers "github.com/ismacaulay/chip8/pkg/emu/registers/mock"
)

func TestInstructionC(t *testing.T) {
	ctrl := gomock.NewController(t)
	registers := mock_registers.NewMockReaderWriter(ctrl)
	instruction := newInstructionC(registers)

	t.Run("[Cxkk] Set Vx = random byte AND kk", func(t *testing.T) {
		rand.Seed(42)
		randValue := uint8(rand.Intn(256))
		rand.Seed(42)
		value := uint8(116)

		registers.EXPECT().SetRegisterValue(uint8(0x3), uint8(randValue&value))
		registers.EXPECT().IncrementProgramCounter(uint16(1))

		opcode := uint16(0x300) | uint16(value)

		instruction.execute(opcode)
	})

	ctrl.Finish()
}
