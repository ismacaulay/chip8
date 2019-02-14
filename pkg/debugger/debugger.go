package debugger

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"text/tabwriter"
	"unsafe"

	"github.com/ismacaulay/chip8/pkg/emu"
	"github.com/ismacaulay/chip8/pkg/emu/display"
	"github.com/ismacaulay/chip8/pkg/emu/keyboard"
	"github.com/ismacaulay/chip8/pkg/emu/memory"
	"github.com/ismacaulay/chip8/pkg/emu/registers"
)

// Debugger used to inspect the emulation state
type Debugger struct {
	emulator *emu.Emulator
}

// NewDebugger returns a new Debugger
func NewDebugger(e *emu.Emulator) *Debugger {
	return &Debugger{e}
}

// PrintRegisters prints the current values of the registers
func (d *Debugger) PrintRegisters() {
	emulator := reflect.ValueOf(d.emulator).Elem()

	registersPtr := getFieldFromValueByName(emulator, "registers")
	registers := reflect.ValueOf(registersPtr.Interface().(*registers.Registers)).Elem()

	v := getFieldFromValueByName(registers, "v").Interface().([]uint8)
	i := getFieldFromValueByName(registers, "i").Uint()
	pc := getFieldFromValueByName(registers, "pc").Uint()
	sp := getFieldFromValueByName(registers, "sp").Uint()
	stack := getFieldFromValueByName(registers, "stack").Interface().([]uint16)
	delayTimer := getFieldFromValueByName(registers, "delayTimer").Uint()
	soundTimer := getFieldFromValueByName(registers, "soundTimer").Uint()

	const padding = 3
	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', 0)
	fmt.Fprintf(w, "v\t%v\n", toHexSlice(v))
	fmt.Fprintf(w, "i\t%04X\n", i)
	fmt.Fprintf(w, "pc\t%04X\n", pc)
	fmt.Fprintf(w, "sp\t%04X\n", sp)
	fmt.Fprintf(w, "stack\t%v\n", toHexSlice(stack))
	fmt.Fprintf(w, "delayTimer\t%d\n", delayTimer)
	fmt.Fprintf(w, "soundTimer\t%d\n", soundTimer)
	w.Flush()
}

// PrintDisplay prints the current display buffer
func (d *Debugger) PrintDisplay() {
	emulator := reflect.ValueOf(d.emulator).Elem()

	displayPtr := getFieldFromValueByName(emulator, "display")
	display := displayPtr.Interface().(*display.Display)

	width := display.Width()
	height := display.Height()
	buffer := display.Buffer()

	for row := uint(0); row < height; row++ {
		for col := uint(0); col < width; col++ {
			fmt.Print(buffer[(row*uint(width))+col])
		}
		fmt.Print("\n")
	}
}

func (d *Debugger) PrintNextInstructions(num int) {
	emulator := reflect.ValueOf(d.emulator).Elem()

	memoryPtr := getFieldFromValueByName(emulator, "memory")
	memory := reflect.ValueOf(memoryPtr.Interface().(*memory.Memory)).Elem()
	buffer := getFieldFromValueByName(memory, "buffer").Interface().([]uint8)

	registersPtr := getFieldFromValueByName(emulator, "registers")
	registers := reflect.ValueOf(registersPtr.Interface().(*registers.Registers)).Elem()
	pc := getFieldFromValueByName(registers, "pc").Uint()

	const padding = 3
	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', 0)

	for i := 0; i < num; i++ {
		addr := uint(pc) + uint(i*2)
		byte0 := buffer[addr]
		byte1 := buffer[addr+1]
		opcode := (uint16(byte0) << 8) | uint16(byte1)
		fmt.Fprintf(w, "%04X\t%04X\n", addr, opcode)
	}
	w.Flush()
}

func (d *Debugger) PrintKeyboard() {
	emulator := reflect.ValueOf(d.emulator).Elem()

	keyboardPtr := getFieldFromValueByName(emulator, "keyboard")
	keyboard := reflect.ValueOf(keyboardPtr.Interface().(*keyboard.Keyboard)).Elem()

	state := getFieldFromValueByName(keyboard, "state").Uint()
	waitingForKeyPress := getFieldFromValueByName(keyboard, "waitingForKeyPress").Bool()
	const padding = 3

	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', 0)
	fmt.Fprintf(w, "state\t%s\n", strconv.FormatInt(int64(state), 2))
	fmt.Fprintf(w, "waitingForKeyPress\t%v\n", waitingForKeyPress)
	w.Flush()
}

func getFieldFromValueByName(v reflect.Value, name string) reflect.Value {
	f := v.FieldByName(name)
	return reflect.NewAt(f.Type(), unsafe.Pointer(f.UnsafeAddr())).Elem()
}

func toHexSlice(list interface{}) []string {
	out := []string{}
	switch list.(type) {
	case []uint8:
		for _, v := range list.([]uint8) {
			out = append(out, fmt.Sprintf("%04X", v))
		}
	case []uint16:
		for _, v := range list.([]uint16) {
			out = append(out, fmt.Sprintf("%04X", v))
		}
	}

	return out
}
