package pkg

import (
	"bufio"
	"os"
	"strings"
)

func readString(reader *bufio.Reader) string {
    data, _ := reader.ReadString('\n')
    data = strings.ReplaceAll(data, "\n", "")
    data = strings.ReplaceAll(data, "\r", "")
    return data
}

func GetWords() []string {
	file, _ := os.Open("words.txt")
	defer file.Close()

	words := make([]string, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		word := scanner.Text()[:5]
		words = append(words, word)
	}

    return words
}

