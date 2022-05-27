package pkg

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Wordle struct {
    word string
    times int
}

func NewWordle(words []string) Wordle {
    rand.Seed(time.Now().UnixNano())
    word := words[rand.Intn(len(words))]

    return Wordle{
        word: word,
        times: 0,
    }
}

func (w *Wordle) In(character rune) int {
    for i, c := range w.word {
        if c == character {
            return i
        }
    }
    return -1
}

func (w *Wordle) Try(word string) string {
    status := []rune("wwwww")

    total := map[rune]int{}
    for _, c := range w.word {
        total[c]++
    }

    count := map[rune]int{}

    for i, c := range word {
        if w.word[i] == byte(c) {
            status[i] = 'e'
            count[c]++
        }
    }

    for i, c := range word {
        if pos := w.In(c); pos != -1 && pos != i && count[c] < total[c] && status[i] == 'w' {
            status[i] = 'o'
        }
    }

    return string(status)
}

func (w *Wordle) Run() {
    reader := bufio.NewReader(os.Stdin)

    for i := 0; i < 6; i++ {
        word := readString(reader)
        status := w.Try(word)
        w.times++
        fmt.Println(status)
        if status == "eeeee" {
            return
        }
    }
}
