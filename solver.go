
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Friedchicken-42/cli"
)

type Rule struct {
	char     byte
	position int
	status   byte
	done     bool
}

func (r Rule) String() string {
    return fmt.Sprintf("[%c %d %c]", r.char, r.position, r.status)
}

func filter(words []string, cond func(string) bool) []string {
    result := make([]string, 0)

    for _, word := range words {
        if cond(word) {
            result = append(result, word)
        }
    }

    return result
}

func Filter(words []string, rules []Rule) []string {
    for _, rule := range rules {
        if rule.done {
            continue
        }

        if rule.status == 'e' {
            words = filter(words, func(word string) bool {
                return word[rule.position] == rule.char
            })
        } else if rule.status == 'o' {
            words = filter(words, func(word string) bool {
                count := 0
                for i, w := range word {
                    if w == rune(rule.char) {
                        if i == rule.position {
                            return false
                        }
                        count++;
                    }
                }
                return count > 0
            })
        } else if rule.status == 'w' {
            ok := false
             for _, r := range rules {
                 if r.status != 'w' && r.char == rule.char {
                     ok = true
                 }
             }

             if !ok {
                 words = filter(words, func(word string) bool {
                     for _, c := range word {
                         if byte(c) == rule.char {
                             return false
                         }
                     }
                     return true
                 })
             }
        }

        rule.done = true
    }

	return words
}

func GenRules(word string, status string) []Rule {
    rules := make([]Rule, 5)

    for i, w := range word {
        rules[i] = Rule{
        	char:     byte(w),
        	position: i,
        	status:   status[i],
        	done:     false,
        }
    }

    return rules
}

func CountLetters(words []string) map[rune]int {
    frequency := map[rune]int{}

    for _, w := range words {
        for _, c := range w {
            frequency[c] += 1
        }
    }

    return frequency
}

func WordsFrequency(words []string, frequency map[rune]int) map[string]int {
    m := map[string]int{}

    for _, word := range words {
        x := map[rune]bool{}
        for _, c := range word {
            if !x[c] {
                m[word] += frequency[c]
            }
            x[c] = true
        }
    }

    return m
}

func GetHigher(wordsFrequency map[string]int) string {
    var word string
    var value int

    for w, v := range wordsFrequency {
        if v > value {
            word = w
            value = v
        }
    }

    return word
}

func FindMostCommon(words []string) string {
    frequency := CountLetters(words)
    for k, v := range frequency {
        fmt.Print(string(k))
        fmt.Printf(" %5d ", v)
        fmt.Println()
    }

    wordsFrequency := WordsFrequency(words, frequency)

    higher := GetHigher(wordsFrequency)

    return higher
}

func readString(reader *bufio.Reader) string {
    data, _ := reader.ReadString('\n')
    data = strings.ReplaceAll(data, "\n", "")
    data = strings.ReplaceAll(data, "\r", "")
    return data
}

func Init() []string {
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

func UserInput(words []string) {
    reader := bufio.NewReader(os.Stdin)

    rules := []Rule{}

    for {
        common := FindMostCommon(words)
        fmt.Println(common)

        res := readString(reader)

        if res == "eeeee" {
            break
        }

        rules = append(rules, GenRules(common, res)...)
	    words = Filter(words, rules)
        fmt.Printf("words remaining: %d\n", len(words))
        if len(words) == 1 {
            fmt.Printf("only one word remains: \"%s\"\n", words[0])
            break
        } else if len(words) <= 10 {
            fmt.Println(words)
        }
    }

}

func Puzzle(words []string) {
    reader := bufio.NewReader(os.Stdin)

    rules := make([]Rule, 0)

    for {
        line := readString(reader)
        if line == "" {
            break
        }
        x := strings.Split(line, " ")
        word, status := x[0], x[1]

        rules = append(rules, GenRules(word, status)...)
	    words = Filter(words, rules)
    }

    fmt.Println(rules)
    fmt.Println(words)
}

func main() {
    app := &cli.App{
        Options: cli.Options{
            &cli.Option{
                Name: "puzzle",
                Prompt: "puzzle",
                IsFlag: true,
            },
        },
        Action: func(c *cli.Context) error {
            words := Init()
            if _, ok := c.Get("puzzle"); ok {
                Puzzle(words)
            } else {
                UserInput(words)
            }
            return nil
        },
    }

    if err := app.Run(os.Args); err != nil {
        panic(err)
    }
}
