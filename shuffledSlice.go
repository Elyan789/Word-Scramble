package main

import (
	"bufio"
	"os"
)

func GetShuffled(fileName string) ([]string, error) {
	var shuffledWords []string
	file, err := os.Open(fileName)
	if err != nil {
		return shuffledWords, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		shuffled := scanner.Text()
		shuffledWords = append(shuffledWords, shuffled)
	}
	err = file.Close()

	if err != nil {
		return shuffledWords, err
	}

	if scanner.Err() != nil {
		return shuffledWords, scanner.Err()
	}
	return shuffledWords, nil
}
