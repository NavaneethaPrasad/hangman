package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"unicode"
)

type Hangman struct {
	secretWord     string
	guesses        []byte
	chancesLeft    uint
	correctGuesses []byte
}

func NewHangman(secretWord string) Hangman {
	return Hangman{
		secretWord:     secretWord,
		guesses:        []byte{},
		chancesLeft:    7,
		correctGuesses: []byte{},
	}
}

func isAllLetters(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
func getSecretWord(wordFileName string) string {
	allowedWords := []string{}
	file, err := os.Open(wordFileName)
	if err != nil {
		errMessage := fmt.Sprintf("Can't open file %s : %v\n", wordFileName, err)
		panic(errMessage)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		if word == strings.ToLower(word) && len(word) >= 6 && isAllLetters(word) {
			allowedWords = append(allowedWords, word)
		}
	}
	randomNum := rand.Intn(len(allowedWords))
	return allowedWords[randomNum]
}
func checkGuess(currentState Hangman, user_Input byte) Hangman {
	return currentState
}

func main() {
	fmt.Println(getSecretWord("/usr/share/dict/words"))
}
