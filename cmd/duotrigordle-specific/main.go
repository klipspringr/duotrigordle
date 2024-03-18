package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"

	"github.com/klipspringr/duotrigordle/internal"
)

func main() {
	dictGuesses, dictSolutions := internal.LoadDictionaries("data/guesses.duotrigordle.20240309.txt", "data/solutions.duotrigordle.20240309.txt")

	var guesses []string
	for _, arg := range os.Args[1:] {
		if len(arg) == internal.Length {
			guesses = append(guesses, strings.ToUpper(arg))
		}
	}
	if len(guesses) == 0 {
		log.Fatalln("No valid guesses provided")
	} else {
		fmt.Printf("Guesses: %s\n", guesses)
	}

	solvableBoards := 0
	totalPossibleSolutions := 0
	for _, solution := range dictSolutions {
		possibleSolutions := dictGuesses
		for _, guess := range guesses {
			response := internal.CheckGuess(guess, solution)
			remaining := []string{}
			for _, possibleSolution := range possibleSolutions {
				if internal.CheckGuess(guess, possibleSolution) == response {
					remaining = append(remaining, possibleSolution)
				}
			}
			possibleSolutions = remaining
		}

		totalPossibleSolutions += len(possibleSolutions)
		if len(possibleSolutions) == 1 {
			fmt.Println(internal.FormatGuesses(append(guesses, solution), solution))
			solvableBoards++
		}
	}

	solvableFraction := float64(solvableBoards) / float64(len(dictSolutions))
	fmt.Printf("%d of %d boards (%.1f%%) are solvable on next guess\n",
		solvableBoards,
		len(dictSolutions),
		100*solvableFraction)

	const boards = 32
	fmt.Printf("%.1f%% chance of solvable board on grid of %d\n",
		100*(1-math.Pow(1-solvableFraction, float64(boards))),
		boards)

	fmt.Printf("%.1f average possible solutions per board\n",
		float64(totalPossibleSolutions)/float64(len(dictSolutions)))
}
