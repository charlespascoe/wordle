package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

var inputRegexp = regexp.MustCompile(`^[a-z]{5}$`)

const ActualMatch = 3
const CorrectLetter = 2
const NoMatch = 1

const ActualMatchColour = 34
const CorrectLetterColour = 220
const NoMatchColour = 239
const ErrorTextColour = 160

// Game TODO: Description.
type Game struct {
	attempts    [][]rune
	errMsg      string
	rnd         *rand.Rand
	term        *Terminal
	word        []rune
	wordLetters map[rune]struct{}
	words       map[string]struct{}
	won         bool
}

func NewGame(term *Terminal) *Game {
	game := &Game{
		words: make(map[string]struct{}),
		term:  term,
		rnd:   rand.New(rand.NewSource(time.Now().UnixNano())),
	}

	for _, word := range Words {
		game.words[word] = struct{}{}
	}

	return game
}

// Init TODO: Description.
func (game *Game) Init() {
	game.word = []rune(game.RandWord())
	game.wordLetters = make(map[rune]struct{})

	for _, char := range game.word {
		game.wordLetters[char] = struct{}{}
	}

	game.attempts = nil

	game.Render()
}

// RandWord TODO: Description.
func (game *Game) RandWord() string {
	return Words[game.rnd.Intn(len(Words))]
}

// Input TODO: Description.
func (game *Game) Input(input string) (won bool) {
	if !inputRegexp.MatchString(input) {
		game.errMsg = "that isn't a five letter word"
	} else if !game.IsWord(input) {
		game.errMsg = "that isn't a word in the word list"
	} else {
		game.errMsg = ""
		game.attempts = append(game.attempts, []rune(input))

		if input == string(game.word) {
			game.won = true
		}
	}

	game.Render()

	return game.won
}

// IsWord TODO: Description.
func (game *Game) IsWord(input string) bool {
	_, exists := game.words[input]

	return exists
}

// Render TODO: Description.
func (game *Game) Render() {
	game.term.
		ClearScreen().
		ResetCursorPosition().
		ResetTextStyle().
		Write("guess the five letter word", "\n\n")

	matches := make(map[rune]int)

	for _, attempt := range game.attempts {
		for i, char := range attempt {
			if game.word[i] == char {
				game.term.SetTextColour(ActualMatchColour)

				if matches[char] < ActualMatch {
					matches[char] = ActualMatch
				}
			} else if _, exists := game.wordLetters[char]; exists {
				game.term.SetTextColour(CorrectLetterColour)

				if matches[char] < CorrectLetter {
					matches[char] = CorrectLetter
				}
			} else {
				game.term.SetTextColour(NoMatchColour)

				if matches[char] < NoMatch {
					matches[char] = NoMatch
				}
			}

			game.term.Write(char)
		}

		game.term.ResetTextStyle().Write("\n")
	}

	if len(game.errMsg) > 0 {
		game.term.
			Write("\n\n").
			SetTextColour(ErrorTextColour).
			Write(game.errMsg).
			ResetTextStyle()
	}

	game.term.Write(
		"\n\n",
		fmt.Sprintf("attempts: %d", len(game.attempts)),
		"\n\n",
	)

	if game.won {
		game.term.
			SetTextColour(ActualMatchColour).
			Write("well done!").
			ResetTextStyle().
			Write("\n")

		return
	}

	for i, char := range "abcdefghijklmnopqrstuvwxyz" {
		if i == 13 {
			game.term.Write("\n")
		}

		switch matches[char] {
		case ActualMatch:
			game.term.SetTextColour(ActualMatchColour)
		case CorrectLetter:
			game.term.SetTextColour(CorrectLetterColour)
		case NoMatch:
			game.term.SetTextColour(NoMatchColour)
		default:
			game.term.ResetTextStyle()
		}

		game.term.Write(char, " ")
	}

	game.term.
		ResetCursorPosition().
		MoveCursor(Down, 2+len(game.attempts))
}
