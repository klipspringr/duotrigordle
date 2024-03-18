package internal

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func FormatResponse(word string, response Response) string {
	var b strings.Builder
	for i, c := range word {
		if response&(1<<(i+Length)) != 0 {
			fmt.Fprintf(&b, "\x1b[42;30m") // Correct
		} else if response&(1<<i) != 0 {
			fmt.Fprintf(&b, "\x1b[43;30m") // Present
		} else {
			fmt.Fprintf(&b, "\x1b[40;37m") // Absent
		}
		fmt.Fprintf(&b, "%s\x1b[0m", string(c))
	}
	return b.String()
}

func FormatGuesses(guesses []string, solution string) string {
	r := []string{}
	for _, guess := range guesses {
		response := CheckGuess(guess, solution)
		r = append(r, FormatResponse(guess, response))
	}
	return strings.Join(r, " ")
}

func ReadDictionary(path string, requiredLength int) []string {
	exe, err := os.Executable()
	if err != nil {
		log.Fatalln(err)
	}
	absolutePath := filepath.Join(filepath.Dir(exe), path)
	f, err := os.Open(absolutePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	ret := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := strings.ToUpper(scanner.Text())
		if len(s) == requiredLength {
			ret = append(ret, s)
		} else {
			log.Fatalf("Unexpected line in dictionary, len %d", len(s))
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	return ret
}

func LoadDictionaries(pathGuesses string, pathSolutions string) ([]string, []string) {
	dictGuesses := ReadDictionary(pathGuesses, Length)
	dictSolutions := ReadDictionary(pathSolutions, Length)

	fmt.Printf("Dictionaries: %d valid guesses, %d solutions\n", len(dictGuesses), len(dictSolutions))

	return dictGuesses, dictSolutions
}
