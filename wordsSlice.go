package main

import (
	"bufio"
	"os"
)

func GetWords(fileName string) ([]string, error) {
	var words []string
	file, err := os.Open(fileName)
	if err != nil {
		return words, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		aWord := scanner.Text()
		words = append(words, aWord)
	}
	err = file.Close()

	if err != nil {
		return words, err
	}

	if scanner.Err() != nil {
		return words, scanner.Err()
	}
	return words, nil
}
