# chip8
A chip-8 interpreter written in Go

### What is a chip-8

Chip-8 is a simple interpreted programming language that was first used in the late 1970s. It was typically used in do-it-yourself
computer systems, that included a display, a hex keyboard, and 4K ram.

The specification I followed can be found at http://devernay.free.fr/hacks/chip8/C8TECH10.HTM#1nnn

### Why

Since the chip-8 instruction set only consists of 35 instructions, it is a good
introduction to building emulators. It is a project
that has a clear definintion of done, and can be implemented in a short period of time. I
also have been enjoying using Go and wanted to continue learning how to use it in
different ways.

### Goals

There are a few goals of the project that I hope to complete by the time it is finish.

- [ ] Understand how to build an emulator
- [ ] Implement a TUI application that can run chip-8 roms
- [ ] Compile the emulator for wasm, to run chip-8 roms in the browser

Stretch goals
- [ ] Implement a chip-8 program, and successfully run it on the emulator
- [ ] Build a debugger that can control the flow of the emulator, and inspect the memory

### Design

Since the goal is to be able to target both terminal and browsers, this introduces some design constraints to the emulator.

#### Filesystem

Browsers do not have the filesystem available to them, even when compiling for wasm. This
limitation means the emulation core should not be trying to load the ROM itself, and
instead should be passed the rom binary data to process. This leaves the users of the
emulation the choice of how to load the rom data, for example the browser can pull it down
from an API, and the TUI application can load it from the filesystem

#### Graphics

The different target platforms support differnet ways of rendering to the screen. To
support this, the enumlation will implement an observer pattern to pass the display buffer
back to users of the emulation when it is updated.

#### Keyboard

Just like the display, the different target platforms handle user input differently. To
support this, the emulator will need to be able to accept an input buffer from the
applications to indicate the current keyboard state. Since the keyboard is a hex keyboard,
this can be done with a 16 bit integer.

### References

Specification: http://devernay.free.fr/hacks/chip8/C8TECH10.HTM#1nnn
Wikipedia: https://en.wikipedia.org/wiki/CHIP-8#Memory
How to write an emulator: http://www.multigesture.net/articles/how-to-write-an-emulator-chip-8-interpreter/
