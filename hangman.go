package hangman

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type HangManData struct {
	Word        string
	ToFind      string
	Attempts    int
	UsedLetters []string
}

func Setword(words []string, hangmanData *HangManData) {
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(words))
	word := words[randomIndex]
	guessWord := make([]byte, len(word))
	for i := range guessWord {
		guessWord[i] = '_'
	}
	hangmanData.Word = string(guessWord)
	hangmanData.ToFind = word
}

func Meca(guess string, hangmanData *HangManData) bool {
	hangmanData.UsedLetters = append(hangmanData.UsedLetters, guess)

	if len(guess) > 1 {
		if !Advancedword(hangmanData.ToFind, guess, hangmanData) {
			return true
		}
	} else if strings.Contains(hangmanData.ToFind, guess) {
		hangmanData.Word = updateGuessWord(hangmanData.ToFind, guess, []byte(hangmanData.Word))
	} else {
		hangmanData.Attempts--
	}
	return false
}

func updateGuessWord(word, guess string, guessWord []byte) string {
	for i := 0; i < len(word); i++ {
		if word[i] == guess[0] {
			guessWord[i] = guess[0]
		}
	}
	return string(guessWord)
}

func Advancedword(word, guess string, hangmanData *HangManData) bool {
	for i, letter := range word {
		if letter != rune(guess[i]) {
			hangmanData.Attempts = hangmanData.Attempts - 2
			return true
		}
	}
	return false
}

func LetterAlreadyUsed(guess string, usedLetters []string) bool {
	for _, letter := range usedLetters {
		if letter == guess {
			return true
		}
	}
	return false
}

func EndGame(hangmanData *HangManData) bool {
	if hangmanData.Attempts <= 0 {
		return true
	}
	if hangmanData.Word == hangmanData.ToFind {
		return true
	}
	return false
}

func SetData(hangmanData *HangManData) {
	hangmanData.Attempts = 10
	hangmanData.UsedLetters = make([]string, 0)
}

func ReadDico(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Erreur lors de l'ouverture du fichier : %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Erreur lors de la lecture du fichier : %v\n", err)
		os.Exit(1)
	}
	return words
}
