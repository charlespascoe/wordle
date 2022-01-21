package main

import (
	"bufio"
	"strings"
	"os"
)

func main() {
	term := NewTerminal(os.Stdout)
	game := NewGame(term)
	game.Init()

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := strings.ToLower(scanner.Text())

		if len(line) == 0 {
			game.Render()
			continue
		}

		if line == "q" {
			break
		}

		if won := game.Input(line); won {
			break
		}
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}
}
