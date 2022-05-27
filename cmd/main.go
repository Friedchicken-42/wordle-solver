package main

import (
	"os"
	"strconv"

	"github.com/Friedchicken-42/cli"
	"github.com/Friedchicken-42/wordle-solver/pkg"
)

func main() {
    app := &cli.App{
        Options: cli.Options{
            &cli.Option{
                Name: "puzzle",
                Prompt: "puzzle",
                IsFlag: true,
            },
            &cli.Option{
                Name: "test",
                Prompt: "test",
            },
            &cli.Option{
                Name: "play",
                Prompt: "play",
                IsFlag: true,
            },
        },
        Action: func(c *cli.Context) error {
            words := pkg.GetWords()
            if _, ok := c.Get("puzzle"); ok {
                pkg.Puzzle(words)
            } else if times, ok := c.Get("test"); ok {
                t, _ := strconv.Atoi(times)
                pkg.TestSolver(words, t)
            } else if _, ok := c.Get("play"); ok {
                wordle := pkg.NewWordle(words)
                wordle.Run()
            } else {
                pkg.UserInput(words)
            }
            return nil
        },
    }

    if err := app.Run(os.Args); err != nil {
        panic(err)
    }
}
