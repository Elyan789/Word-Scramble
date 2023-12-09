package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

var streak int
var maxStreak int

func guessWordPrompt(scrambledWord string) {
	fmt.Println()
	fmt.Printf("Guess the correct word from this arrangment of letters in 5 attempts: %s\n", scrambledWord)
	fmt.Println("If you don't know, enter Reveal.")
}

func difficultyPrompt() {
	fmt.Println("The following difficulties are: Easy, Normal, Hard, Expert, Nightmare): ")
	fmt.Println("Easy: 5 attempts, Normal: 4 attempts, Hard: 3 Attempts, Expert: 2 Attempts, Nightmare: 1 attempt,")
	fmt.Print("Enter a difficulty: ")
}

func errorOccuring(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func exitGame() {
	err := errors.New("thanks for playing")
	errorOccuring(err)
}

func playAgain() {
	fmt.Println("Would you like to play again?")
	reader4 := bufio.NewReader(os.Stdin)
	response, err := reader4.ReadString('\n')
	errorOccuring(err)
	response = strings.TrimSpace(response)
	if response == "Yes" {
		main()
	} else if response == "No" {
		exitGame()
	} else {
		fmt.Println("Invalid response.")
	}
}

func highestStreak(streak int, maxStreak int) int {
	if streak > maxStreak {
		return streak
	}
	return maxStreak
}

func winCheck(guessedWord string, correctWord string, revealed bool, streak *int) {

	if guessedWord == correctWord {
		(*streak)++
		fmt.Println()
		fmt.Printf("Current win streak is: %d\n", *streak)
	} else if !revealed || guessedWord == "Reveal" && !revealed {
		(*streak) = 0
		fmt.Println()
		fmt.Printf("Current win streak is: %d\n", *streak)
	}

	maxStreak = highestStreak(*streak, maxStreak)
	fmt.Printf("Highest win streak is: %d\n", maxStreak)
	fmt.Println()
	playAgain()
}

func toLoop(guessedWord string, correctWord string, attempt int, streak *int) {
	var i int
	revealed := false
	for i = 1; i <= attempt; i++ {
		fmt.Println("You have", attempt-i+1, "attempts remaining.")
		fmt.Println()
		fmt.Print("Enter your word: ")
		fmt.Scanln(&guessedWord)
		if guessedWord == correctWord {
			fmt.Println("Correct word!")
			winCheck(guessedWord, correctWord, revealed, streak)
			break
		} else if guessedWord == "Reveal" && !revealed {
			fmt.Printf("The word was: %s\n", correctWord)
			winCheck(guessedWord, correctWord, revealed, streak)
			revealed = true
			break
		} else {
			fmt.Println("Incorrect word.")
		}
	}
	if !revealed {
		fmt.Printf("Sorry, you didn't guess the word correctly, it was: %s\n", correctWord)
		winCheck(guessedWord, correctWord, revealed, streak)
	}
	fmt.Println()
	playAgain()
}

func sevenLetterWord(target int, correctWord string, guessedWord string, scrambledWord string, attempt int, streak *int) {
	difficultyPrompt()
	reader6 := bufio.NewReader(os.Stdin)
	i6, err := reader6.ReadString('\n')
	errorOccuring(err)
	i6 = strings.TrimSpace(i6)

	if i6 == "Easy" {
		words7E, err := GetWords("7Word.txt")
		errorOccuring(err)
		shuffledWords7E, err := GetShuffled("7Scrambled.txt")
		errorOccuring(err)
		correctWord = words7E[target]
		scrambledWord = shuffledWords7E[target]
		attempt = 5
		guessWordPrompt(scrambledWord)
		toLoop(guessedWord, correctWord, attempt, streak)
	} else if i6 == "Normal" {
		words7N, err := GetWords("7Word.txt")
		errorOccuring(err)
		shuffledWords7N, err := GetShuffled("7Scrambled.txt")
		errorOccuring(err)
		correctWord = words7N[target]
		scrambledWord = shuffledWords7N[target]
		attempt = 4
		guessWordPrompt(scrambledWord)
		toLoop(guessedWord, correctWord, attempt, streak)
	} else if i6 == "Hard" {
		words7H, err := GetWords("7Word.txt")
		errorOccuring(err)
		shuffledWords7H, err := GetShuffled("7Scrambled.txt")
		errorOccuring(err)
		correctWord = words7H[target]
		scrambledWord = shuffledWords7H[target]
		attempt = 3
		guessWordPrompt(scrambledWord)
		toLoop(guessedWord, correctWord, attempt, streak)
	} else if i6 == "Expert" {
		words7Ext, err := GetWords("7Word.txt")
		errorOccuring(err)
		shuffledWords7Ext, err := GetShuffled("7Scrambled.txt")
		errorOccuring(err)
		correctWord = words7Ext[target]
		scrambledWord = shuffledWords7Ext[target]
		attempt = 2
		guessWordPrompt(scrambledWord)
		toLoop(guessedWord, correctWord, attempt, streak)
	} else if i6 == "Nightmare" {
		words7Night, err := GetWords("7Word.txt")
		errorOccuring(err)
		shuffledWords7Night, err := GetShuffled("7Scrambled.txt")
		errorOccuring(err)
		correctWord = words7Night[target]
		scrambledWord = shuffledWords7Night[target]
		attempt = 1
		guessWordPrompt(scrambledWord)
		toLoop(guessedWord, correctWord, attempt, streak)
	} else {
		guessedWord = ""
		correctWord = ""
		scrambledWord = ""
		fmt.Println("Invalid input.")
	}
}

func sixLetterWord(target int, correctWord string, guessedWord string, scrambledWord string, attempt int, streak *int) {
	difficultyPrompt()
	reader5 := bufio.NewReader(os.Stdin)
	i5, err := reader5.ReadString('\n')
	errorOccuring(err)
	i5 = strings.TrimSpace(i5)

	if i5 == "Easy" {
		words6E, err := GetWords("6Word.txt")
		errorOccuring(err)
		shuffledWords6E, err := GetShuffled("6Scrambled.txt")
		errorOccuring(err)
		correctWord = words6E[target]
		scrambledWord = shuffledWords6E[target]
		attempt = 5
		guessWordPrompt(scrambledWord)
		toLoop(guessedWord, correctWord, attempt, streak)
	} else if i5 == "Normal" {
		words6N, err := GetWords("6Word.txt")
		errorOccuring(err)
		shuffledWords6N, err := GetShuffled("6Scrambled.txt")
		errorOccuring(err)
		correctWord = words6N[target]
		scrambledWord = shuffledWords6N[target]
		attempt = 4
		guessWordPrompt(scrambledWord)
		toLoop(guessedWord, correctWord, attempt, streak)
	} else if i5 == "Hard" {
		words6H, err := GetWords("6Word.txt")
		errorOccuring(err)
		shuffledWords6H, err := GetShuffled("6Scrambled.txt")
		errorOccuring(err)
		correctWord = words6H[target]
		scrambledWord = shuffledWords6H[target]
		attempt = 3
		guessWordPrompt(scrambledWord)
		toLoop(guessedWord, correctWord, attempt, streak)
	} else if i5 == "Expert" {
		words6Ext, err := GetWords("6Word.txt")
		errorOccuring(err)
		shuffledWords6Ext, err := GetShuffled("6Scrambled.txt")
		errorOccuring(err)
		correctWord = words6Ext[target]
		scrambledWord = shuffledWords6Ext[target]
		attempt = 2
		guessWordPrompt(scrambledWord)
		toLoop(guessedWord, correctWord, attempt, streak)
	} else if i5 == "Nightmare" {
		words6Night, err := GetWords("6Word.txt")
		errorOccuring(err)
		shuffledWords6Night, err := GetShuffled("6Scrambled.txt")
		errorOccuring(err)
		correctWord = words6Night[target]
		scrambledWord = shuffledWords6Night[target]
		attempt = 1
		guessWordPrompt(scrambledWord)
		toLoop(guessedWord, correctWord, attempt, streak)
	} else {
		guessedWord = ""
		correctWord = ""
		scrambledWord = ""
		fmt.Println("Invalid input.")
	}
}

func fiveLetterWord(target int, correctWord string, guessedWord string, scrambledWord string, attempt int, streak *int) {
	difficultyPrompt()
	reader4 := bufio.NewReader(os.Stdin)
	i4, err := reader4.ReadString('\n')
	errorOccuring(err)
	i4 = strings.TrimSpace(i4)

	if i4 == "Easy" {
		words5E, err := GetWords("5Word.txt")
		errorOccuring(err)
		shuffledWords5E, err := GetShuffled("5Scrambled.txt")
		errorOccuring(err)
		correctWord = words5E[target]
		scrambledWord = shuffledWords5E[target]
		attempt = 5
		guessWordPrompt(scrambledWord)
		toLoop(guessedWord, correctWord, attempt, streak)
	} else if i4 == "Normal" {
		words5N, err := GetWords("5Word.txt")
		errorOccuring(err)
		shuffledWords5N, err := GetShuffled("5Scrambled.txt")
		errorOccuring(err)
		correctWord = words5N[target]
		scrambledWord = shuffledWords5N[target]
		attempt = 4
		guessWordPrompt(scrambledWord)
		toLoop(guessedWord, correctWord, attempt, streak)
	} else if i4 == "Hard" {
		words5H, err := GetWords("5Word.txt")
		errorOccuring(err)
		shuffledWords5H, err := GetShuffled("5Scrambled.txt")
		errorOccuring(err)
		correctWord = words5H[target]
		scrambledWord = shuffledWords5H[target]
		attempt = 3
		guessWordPrompt(scrambledWord)
		toLoop(guessedWord, correctWord, attempt, streak)
	} else if i4 == "Expert" {
		words5Ext, err := GetWords("5Word.txt")
		errorOccuring(err)
		shuffledWords5Ext, err := GetShuffled("5Scrambled.txt")
		errorOccuring(err)
		correctWord = words5Ext[target]
		scrambledWord = shuffledWords5Ext[target]
		attempt = 2
		guessWordPrompt(scrambledWord)
		toLoop(guessedWord, correctWord, attempt, streak)
	} else if i4 == "Nightmare" {
		words5Night, err := GetWords("5Word.txt")
		errorOccuring(err)
		shuffledWords5Night, err := GetShuffled("5Scrambled.txt")
		errorOccuring(err)
		correctWord = words5Night[target]
		scrambledWord = shuffledWords5Night[target]
		attempt = 1
		guessWordPrompt(scrambledWord)
		toLoop(guessedWord, correctWord, attempt, streak)
	} else {
		guessedWord = ""
		correctWord = ""
		scrambledWord = ""
		fmt.Println("Invalid input.")
	}
}

func fourLetterWord(target int, correctWord string, guessedWord string, scrambledWord string, attempt int, streak *int) {
	difficultyPrompt()
	reader3 := bufio.NewReader(os.Stdin)
	i3, err := reader3.ReadString('\n')
	errorOccuring(err)
	i3 = strings.TrimSpace(i3)

	if i3 == "Easy" {
		words4E, err := GetWords("4Word.txt")
		errorOccuring(err)
		shuffledWords4E, err := GetShuffled("4Scrambled.txt")
		errorOccuring(err)
		correctWord = words4E[target]
		scrambledWord = shuffledWords4E[target]
		attempt = 5
		guessWordPrompt(scrambledWord)
		toLoop(guessedWord, correctWord, attempt, streak)
	} else if i3 == "Normal" {
		words4N, err := GetWords("4Word.txt")
		errorOccuring(err)
		shuffledWords4N, err := GetShuffled("4Scrambled.txt")
		errorOccuring(err)
		correctWord = words4N[target]
		scrambledWord = shuffledWords4N[target]
		attempt = 4
		guessWordPrompt(scrambledWord)
		toLoop(guessedWord, correctWord, attempt, streak)
	} else if i3 == "Hard" {
		words4H, err := GetWords("4Word.txt")
		errorOccuring(err)
		shuffledWords4H, err := GetShuffled("4Scrambled.txt")
		errorOccuring(err)
		correctWord = words4H[target]
		scrambledWord = shuffledWords4H[target]
		attempt = 3
		guessWordPrompt(scrambledWord)
		toLoop(guessedWord, correctWord, attempt, streak)
	} else if i3 == "Expert" {
		words4Ext, err := GetWords("4Word.txt")
		errorOccuring(err)
		shuffledWords4Ext, err := GetShuffled("4Scrambled.txt")
		errorOccuring(err)
		correctWord = words4Ext[target]
		scrambledWord = shuffledWords4Ext[target]
		attempt = 2
		guessWordPrompt(scrambledWord)
		toLoop(guessedWord, correctWord, attempt, streak)
	} else if i3 == "Nightmare" {
		words4Night, err := GetWords("4Word.txt")
		errorOccuring(err)
		shuffledWords4Night, err := GetShuffled("4Scrambled.txt")
		errorOccuring(err)
		correctWord = words4Night[target]
		scrambledWord = shuffledWords4Night[target]
		attempt = 1
		guessWordPrompt(scrambledWord)
		toLoop(guessedWord, correctWord, attempt, streak)
	} else {
		guessedWord = ""
		correctWord = ""
		scrambledWord = ""
		fmt.Println("Invalid input.")
	}
}

func threeLetterWord(target int, correctWord string, guessedWord string, scrambledWord string, attempt int, streak *int) {
	difficultyPrompt()
	reader2 := bufio.NewReader(os.Stdin)
	i2, err := reader2.ReadString('\n')
	errorOccuring(err)
	i2 = strings.TrimSpace(i2)
	if i2 == "Easy" {
		words3E, err := GetWords("3Word.txt")
		errorOccuring(err)
		shuffledWords3E, err := GetShuffled("3Scrambled.txt")
		errorOccuring(err)
		correctWord = words3E[target]
		scrambledWord = shuffledWords3E[target]
		attempt = 5
		guessWordPrompt(scrambledWord)
		toLoop(guessedWord, correctWord, attempt, streak)
	} else if i2 == "Normal" {
		words3N, err := GetWords("3Word.txt")
		errorOccuring(err)
		shuffledWords3N, err := GetShuffled("3Scrambled.txt")
		errorOccuring(err)
		correctWord = words3N[target]
		scrambledWord = shuffledWords3N[target]
		attempt = 4
		guessWordPrompt(scrambledWord)
		toLoop(guessedWord, correctWord, attempt, streak)
	} else if i2 == "Hard" {
		words3H, err := GetWords("3Word.txt")
		errorOccuring(err)
		shuffledWords3H, err := GetShuffled("3Scrambled.txt")
		errorOccuring(err)
		correctWord = words3H[target]
		scrambledWord = shuffledWords3H[target]
		attempt = 3
		guessWordPrompt(scrambledWord)
		toLoop(guessedWord, correctWord, attempt, streak)
	} else if i2 == "Expert" {
		words3Ext, err := GetWords("3Word.txt")
		errorOccuring(err)
		shuffledWords3Ext, err := GetShuffled("3Scrambled.txt")
		errorOccuring(err)
		correctWord = words3Ext[target]
		scrambledWord = shuffledWords3Ext[target]
		attempt = 2
		guessWordPrompt(scrambledWord)
		toLoop(guessedWord, correctWord, attempt, streak)
	} else if i2 == "Nightmare" {
		words3Night, err := GetWords("3Word.txt")
		errorOccuring(err)
		shuffledWords3Night, err := GetShuffled("3Scrambled.txt")
		errorOccuring(err)
		correctWord = words3Night[target]
		scrambledWord = shuffledWords3Night[target]
		attempt = 1
		guessWordPrompt(scrambledWord)
		toLoop(guessedWord, correctWord, attempt, streak)
	} else {
		guessedWord = ""
		correctWord = ""
		scrambledWord = ""
		fmt.Println("Invalid input.")
	}
}

func main() {
	seconds := time.Now().UnixNano()
	rand.Seed(seconds)
	target := rand.Intn(19)
	fmt.Println()
	fmt.Println("Choose a 3, 4, 5, 6, or 7 letter word.")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	errorOccuring(err)
	input = strings.TrimSpace(input)

	var correctWord string
	var guessedWord string
	var scrambledWord string
	var attempt int

	if input == "3" {
		threeLetterWord(target, correctWord, guessedWord, scrambledWord, attempt, &streak)
	} else if input == "4" {
		fourLetterWord(target, correctWord, guessedWord, scrambledWord, attempt, &streak)
	} else if input == "5" {
		fiveLetterWord(target, correctWord, guessedWord, scrambledWord, attempt, &streak)
	} else if input == "6" {
		sixLetterWord(target, correctWord, guessedWord, scrambledWord, attempt, &streak)
	} else if input == "7" {
		sevenLetterWord(target, correctWord, guessedWord, scrambledWord, attempt, &streak)
	} else {
		fmt.Println("Invalid input.")
	}
}
