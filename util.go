package wordgraph

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"time"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

// LoadFile ...
func LoadFile(filename string) []string {
	inputFile, err := os.Open(filename)
	if err != nil {
		log.Fatal("Error opening input file:", err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	var wordlist []string
	for scanner.Scan() {
		wordlist = append(wordlist, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(scanner.Err())
	}
	return wordlist
}

// Shuffle is a simple Fisher-Yates implementation
func Shuffle(list []string) []string {
	random := rand.New(rand.NewSource(time.Now().Unix()))
	output := append([]string{}, list...)

	for i := len(output) - 1; i > 1; i-- {
		j := random.Intn(i + 1)
		output[i], output[j] = output[j], output[i]
	}
	return output
}
