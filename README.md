# chip8

A [CHIP-8](https://en.m.wikipedia.org/wiki/Chip-8) emulator written in Go.

## Usage

```
Usage of chip8:
  -scaleFactor int
        Display scale factor (default 10)
  rom
        Path to ROM (mandatory)
```

The original CHIP-8 had  a 64x32 pixel display, which is too small for modern
screens. The optional `scaleFactor` argument can be used to upscale it to a
more sensible size. With the default value of 10 the emulated display will have
a size of 640x320 pixels and each CHIP-8 pixel will be rendered as 10x10px
rectangle.

### Keyboard mapping

CHIP-8 systems used a hexadecimal keyboard with the layout shown on the left.
This is mapped to the physical keyboard as shown on the right.

```
|---|---|---|---|               |---|---|---|---|
| 1 | 2 | 3 | C |               | 1 | 2 | 3 | 4 |
|---|---|---|---|               |---|---|---|---|
| 4 | 5 | 6 | D |               | Q | W | E | R |
|---|---|---|---|               |---|---|---|---|
| 7 | 8 | 9 | E |               | A | S | D | F |
|---|---|---|---|               |---|---|---|---|
| A | 0 | B | F |               | Z | X | C | V |
|---|---|---|---|               |---|---|---|---|
```

So to get "deadbeef" inside the emulator you'd have to type "rfzrcffv".

## License

MIT License Copyright (c) 2022 Michael Kohl

For the full license text see [LICENSE](./LICENSE).
