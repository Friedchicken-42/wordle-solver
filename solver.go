package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
    Char rune
    Position int
}

type Wordle struct {
    Exact []Pair
    Ok []Pair
    Wrong []rune
    Attemps int
}

func filter(arr []string, condition func(string) bool) []string {
    result := make([]string, 0)
    for _, el := range arr {
        if condition(el) {
            result = append(result, el)
        }
    }
    return result
}

func letterFrequency(words []string) map[rune]int {
    m := map[rune]int{}

    for _, word := range words {
        for _, c := range word {
            m[c] += 1
        }
    }

    return m
}

func wordFrequency(words []string, frequency map[rune]int) map[string]int {
    m := map[string]int{}

    for _, word := range words {
        letters := map[rune]bool{}
        res := 0
        for _, c := range word {
            if !letters[c] {
                res += frequency[rune(c)]
            }
            letters[c] = true
        }
        m[word] = res
    }

    return m
}

func findMostCommon(m map[string]int) string {
    var word string
    var value int

    for w, v := range m {
        if (v > value) {
            word = w
            value = v
        }
    }

    return word
}

func (wordle *Wordle) clean() {
    for _, a := range wordle.Exact {
        for i, b := range wordle.Ok {
            if a.Char == b.Char && a.Position == b.Position {
                wordle.Ok = append(wordle.Ok[:i], wordle.Ok[i+1:]...)
            }
        }
    }
}

func readString(reader *bufio.Reader) string {
    data, _ := reader.ReadString('\n')
    data = strings.ReplaceAll(data, "\n", "")
    data = strings.ReplaceAll(data, "\r", "")
    return data
}

func main() {
    file, err := os.Open("words.txt")
    if(err != nil){
        panic(err)
    }

    defer file.Close()

    reader := bufio.NewReader(os.Stdin)

    var word_length int
    fmt.Print("word length: ")
    text := readString(reader)
    word_length, _ = strconv.Atoi(text)

    scanner := bufio.NewScanner(file)

    words := make([]string, 0)

    for scanner.Scan() {
        word := scanner.Text()
        if len(word) == word_length {
            words = append(words, word)
        }
    }
    
    wordle := Wordle{}

    for {
        frequency_letters := letterFrequency(words)

        for _, pair := range wordle.Exact {
            frequency_letters[pair.Char] = 0
        }
        for _, pair := range wordle.Ok {
            frequency_letters[pair.Char] = 0
        }
        for _, char := range wordle.Wrong {
            frequency_letters[char] = 0
        }

        frequency_words := wordFrequency(words, frequency_letters)
        common := findMostCommon(frequency_words)

        if(len(words) == 1) {
            fmt.Println("this is the only possible word: ", words[0])
            readString(reader)
            return

        }else if (len(words) <= 10) {
            fmt.Println("only these words remain: ", words)
        }
        fmt.Printf("try this word: \"%s\"\n", common)
        fmt.Println("write the result with e/o/w [exact]/[ok]/[wrong]")
        fmt.Println(common)
        result := readString(reader)

        if (result == "eeeee") {
            fmt.Println("congrats")
            return
        }

        for i, c := range result {
            char := rune(common[i])
            switch(c){
            case 'e':
                wordle.Exact = append(wordle.Exact, Pair{char, i})
            case 'o':
                wordle.Ok = append(wordle.Ok, Pair{char, i})
            case 'w':
                wordle.Wrong = append(wordle.Wrong, char)
            default:
                panic("wrong input")
            }
        }
        
        wordle.clean()

        //filter Wrong
        words = filter(words, func(word string) bool {
            for _, a := range wordle.Wrong {
                for _, w := range word {
                    if (w == a) { return false }
                }
            }
            return true
        })

        //filter exact
        words = filter(words, func(word string) bool {
            count := 0
            for _, pair := range wordle.Exact {
                if(rune(word[pair.Position]) == pair.Char){
                    count++
                }
            }

            return count == len(wordle.Exact)
        })

        //filter ok

        words = filter(words, func(word string) bool {
            count := 0
            for _, pair := range wordle.Ok {
                in := false
                for _, w := range word {
                    if (w == pair.Char){
                        in = true
                    }
                }
                if(!in) { return false }

                if(rune(word[pair.Position]) == pair.Char) { return false }

                count++
            }

            return count == len(wordle.Ok)
        })
    }
}
