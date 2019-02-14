package timers

import "github.com/ismacaulay/chip8/pkg/emu/registers"

type Timers struct {
	registers registers.ReaderWriter
}

func NewTimers(r registers.ReaderWriter) *Timers {
	return &Timers{r}
}

func (t *Timers) Step() {
	delayTimer := t.registers.GetDelayTimer()
	if delayTimer > 0 {
		t.registers.SetDelayTimer(delayTimer - 1)
	}

	soundTimer := t.registers.GetSoundTimer()
	if soundTimer > 0 {
		t.registers.SetSoundTimer(soundTimer - 1)
	}
}
