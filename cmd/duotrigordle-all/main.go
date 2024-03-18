package main

import (
	"fmt"
	"math"

	"github.com/klipspringr/duotrigordle/internal"
)

func main() {
	dictGuesses, dictSolutions := internal.LoadDictionaries("data/guesses.duotrigordle.20240309.txt", "data/solutions.duotrigordle.20240309.txt")

	// test every guess against all possible solutions (dictGuesses)
	// build map counting number of solutions for each response
	// we are looking for guess-response pairs with one solution
	mapCountSolutions := make(map[string]map[internal.Response]int)
	for _, i := range dictGuesses {
		mapCountSolutions[i] = make(map[internal.Response]int)
		for _, j := range dictGuesses {
			response := internal.CheckGuess(i, j)
			mapCountSolutions[i][response] += 1
		}
	}

	bestGuess := ""
	bestSolvableBoards := 0
	bestTotalSolutions := 0

	for _, guess := range dictGuesses {
		solvableBoards := 0
		totalSolutions := 0
		for _, solution := range dictSolutions {
			response := internal.CheckGuess(guess, solution)
			if mapCountSolutions[guess][response] == 1 {
				solvableBoards++
			}
			totalSolutions += mapCountSolutions[guess][response]
		}

		if solvableBoards >= bestSolvableBoards {
			bestSolvableBoards = solvableBoards
			bestGuess = guess
			bestTotalSolutions = totalSolutions
			//fmt.Printf("New best guess %s (%d, %.1f)\n", bestGuess, bestSolvableBoards, float64(bestTotalSolutions)/float64(len(dictSolutions)))
		}
	}

	fmt.Printf("Best guess: %s\n", bestGuess)

	solvableFraction := float64(bestSolvableBoards) / float64(len(dictSolutions))
	fmt.Printf("%d of %d boards (%.1f%%) are solvable on next guess\n",
		bestSolvableBoards,
		len(dictSolutions),
		100*solvableFraction)

	const boards = 32
	fmt.Printf("%.1f%% chance of solvable board on grid of %d\n",
		100*(1-math.Pow(1-solvableFraction, float64(boards))),
		boards)

	fmt.Printf("%.1f average possible solutions per board\n",
		float64(bestTotalSolutions)/float64(len(dictSolutions)))
}
