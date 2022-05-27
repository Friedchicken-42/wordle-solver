package pkg

import (
	"fmt"
	"strings"
	"time"
)

type Result struct {
	Attemps []string
	Results []string
	Word    string
	Tries   int
	Done    bool
}

func (r *Result) Add(word string, status string) {
	r.Attemps = append(r.Attemps, word)
	r.Results = append(r.Results, status)
	r.Tries++
}

func (r Result) String() string {
    data := make([]string, 0)

    data = append(data, fmt.Sprintf("-- %s --", r.Word))
    for i := range r.Attemps {
        data = append(data, fmt.Sprintf("%s %s", r.Attemps[i], r.Results[i]))
    }
    if r.Done {
        data = append(data, fmt.Sprintf(":: Tries: %d", r.Tries))
    } else {
        data = append(data, "++ Failed  ++")
    }
    data = append(data, "")

    return strings.Join(data, "\n")
}

func testSolver(words []string) Result { // return stuff if fails
	wordle := NewWordle(words)
	result := Result{
		Attemps: make([]string, 0),
		Results: make([]string, 0),
		Word:    wordle.word,
	}

	rules := make([]Rule, 0)

	for i := 0; i < 6; i++ {
		word := FindMostCommon(words)
		status := wordle.Try(word)
		result.Add(word, status)

		if status == "eeeee" {
			result.Done = true
			return result
		}

		rules = append(rules, GenRules(word, status)...)
		words = Filter(words, rules)
	}

	return result
}

func TestSolver(words []string, times int) {
    failed := make(map[string]bool, 0)
    count := 0
    start := time.Now()

	for i := 0; i < times; i++ {
        res := testSolver(words)

        if !res.Done {
            failed[res.Word] = true
            count++
        }
        // fmt.Println(res)

        percentage := float32(count) / float32(i + 1) * 100
        fmt.Printf("Failed: %d/%d (%f%%)\r", count, i + 1, percentage)
	}
    fmt.Println("")

    elapsed := time.Since(start)
    ops := float32(times) / float32(elapsed) * float32(time.Second)
    fmt.Printf("Took: %s, ~%f ops\n", elapsed, ops)

    for word := range failed {
        fmt.Printf("%s ", word)
    }
    fmt.Println("")
}
