package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/ismacaulay/chip8/pkg/debugger"
	chip8 "github.com/ismacaulay/chip8/pkg/emu"
	"github.com/ismacaulay/chip8/pkg/emu/keyboard"
)

var roms = map[string]string{
	"keypad_test": "/roms/programs/Keypad Test [Hap, 2006].ch8",
}

func main() {
	emulator := chip8.NewEmulator()
	debugger := debugger.NewDebugger(emulator)
	reader := bufio.NewReader(os.Stdin)

loop:
	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		input := strings.Split(strings.TrimSpace(text), " ")

		cmd := strings.ToLower(input[0])

		switch cmd {
		case "q", "quit":
			break loop
		case "help":
			printHelp()
		case "l", "load":
			name := "keypad_test"
			data := loadRom(name)
			if data != nil {
				emulator.LoadRom(data)
			}
		case "s", "step":
			steps := 1
			if len(input) == 2 {
				value, _ := strconv.ParseUint(input[1], 10, 32)
				steps = int(value)
			}
			for i := 0; i < steps; i++ {
				emulator.Step()
			}
		case "r", "registers":
			debugger.PrintRegisters()
		case "d", "display":
			debugger.PrintDisplay()
		case "n":
			debugger.PrintNextInstructions(10)
		case "k":
			debugger.PrintKeyboard()
		case "kp", "keypress":
			if len(input) == 1 {
				fmt.Println("Error: missing key hex value. cmd: kp [0-9A-F]")
				break
			}

			keyStr := input[1]
			key, err := strconv.ParseUint(keyStr, 16, 8)
			if err != nil {
				fmt.Println("Error parsing key:", keyStr)
				break
			}

			if key > 0xF {
				fmt.Println("Key must be a hex digit [0-9A-F]")
				break
			}
			emulator.HandleKeyEvent(keyboard.EventKeyPressed, keyboard.Key(key))
		default:
			fmt.Println("Unknown command:", cmd, ". Use help for list of commands")
		}
	}
}

func printHelp() {
	fmt.Println("Help..... todo implement me")
}

func loadRom(name string) []uint8 {
	path, ok := roms[name]
	if !ok {
		fmt.Println("Error: unknown rom", name)
		return nil
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Error: unable to read file", path)
		return nil
	}
	return data
}
