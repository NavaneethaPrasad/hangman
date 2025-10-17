package main

import (
	"os"
	"strings"
	"testing"
)

func createDictFiles(words []string) (string, error) {
	f, err := os.CreateTemp("/tmp", "hangman-dict")
	data := strings.Join(words, "\n")
	_, err = f.Write([]byte(data))
	if err != nil {
		return "", err
	}
	return f.Name(), nil
}

func TestSecretWordNoCapitals(t *testing.T) {
	wordList, err := createDictFiles([]string{"Lion", "Elephant", "monkey"})
	defer os.Remove(wordList)
	if err != nil {
		t.Errorf("Couldn't create word list.Can't proceed with test:%v", err)
	}
	secretWord := getSecretWord(wordList)
	if secretWord != "monkey" {
		t.Errorf("Should get 'monkey'but got %s", secretWord)
	}

}

func TestSecretWordLength(t *testing.T) {
	wordList, err := createDictFiles([]string{"Lion", "it", "monkey"})
	defer os.Remove(wordList)
	if err != nil {
		t.Errorf("Couldn't create word list.Can't proceed with test:%v", err)
	}
	secretWord := getSecretWord(wordList)
	if len(secretWord) < 6 {
		t.Errorf("Expected word length 6 or greater than 6, but got %q (length %d)", secretWord, len(secretWord))
	}
}

func TestSecretWordNopunctuation(t *testing.T) {
	wordList, err := createDictFiles([]string{"Lion's", "Elephant's", "monkey"})
	defer os.Remove(wordList)
	if err != nil {
		t.Errorf("Couldn't create word list.Can't proceed with test:%v", err)
	}
	secretWord := getSecretWord(wordList)
	if secretWord != "monkey" {
		t.Errorf("Should not get words with punctuations. Got %s", secretWord)
	}
}

func TestCorrectGuess(t *testing.T) {
	// currentState := Hangman{
	// 	secretWord:     "elephant",
	// 	guesses:        []byte{},
	// 	chancesLeft:    7,
	// 	correctGuesses: []byte{},
	// }
	secretWord := "elephant"
	currentState := NewHangman(secretWord)
	user_Input := byte('t')
	newState := checkGuess(currentState, byte(user_Input))
	expected := Hangman{
		secretWord:     secretWord,
		guesses:        append(currentState.guesses, byte(user_Input)),
		chancesLeft:    7,
		correctGuesses: append(currentState.correctGuesses, byte(user_Input)),
	}
	if newState.secretWord != expected.secretWord {
		t.Errorf("Secreat word is modified\n")
	}
	if string(newState.guesses) != string(expected.guesses) {
		t.Errorf("Guess should be [t] but got %v\n", newState.guesses)
	}
	if string(newState.correctGuesses) != string(expected.correctGuesses) {
		t.Errorf("Guess should be [t] but got %v", newState.correctGuesses)
	}
	if newState.chancesLeft != expected.chancesLeft {
		t.Errorf("Chances left modified!\n")
	}
}
