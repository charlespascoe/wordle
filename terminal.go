package main

import (
	"fmt"
	"os"
)

const (
	Up = "A"
	Down = "B"
	Right = "C"
	Left = "D"
)

type Terminal struct {
	file *os.File
}

func NewTerminal(out *os.File) *Terminal {
	return &Terminal{
		file: out,
	}
}

// Write TODO: Description.
func (t *Terminal) Write(args ...interface{}) *Terminal {
	var err error

	for _, arg := range args {
		switch val := arg.(type) {
		case string:
			_, err = t.file.WriteString(val)
		case []byte:
			_, err = t.file.Write(val)
		case rune:
			_, err = t.file.WriteString(string(val))
		default:
			_, err = t.file.WriteString(fmt.Sprintf("%s", val))
		}
	}

	if err != nil {
		panic(err)
	}

	return t
}

// SetTextColour TODO: Description.
func (t *Terminal) SetTextColour(code int) *Terminal {
	return t.Write(fmt.Sprintf("\033[38;5;%dm", code))
}

// SetBackgroundColour TODO: Description.
func (t *Terminal) SetBackgroundColour(code int) *Terminal {
	return t.Write(fmt.Sprintf("\033[48;5;%dm", code))
}

// ResetTextStyle TODO: Description.
func (t *Terminal) ResetTextStyle() *Terminal {
	return t.Write([]byte("\033[0m"))
}

// ClearScreen TODO: Description.
func (t *Terminal) ClearScreen() *Terminal {
	return t.Write([]byte("\033[2J"))
}

// ResetCursorPosition TODO: Description.
func (t *Terminal) ResetCursorPosition() *Terminal {
	return t.Write([]byte("\033[H"))
}

// MoveCursor TODO: Description.
func (t *Terminal) MoveCursor(dir string, n int) *Terminal {
	return t.Write(fmt.Sprintf("\033[%d%s", n, dir))
}
