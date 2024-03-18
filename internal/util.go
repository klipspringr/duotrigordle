package internal

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const dictionaryTag = "duotrigordle.20240309"

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

func ReadDictionary(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	ret := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := strings.ToUpper(scanner.Text())
		if len(s) == Length {
			ret = append(ret, s)
		} else {
			log.Fatalf("Unexpected dictionary entry, len %d", len(s))
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	return ret
}

func LoadDictionaries() ([]string, []string) {
	exe, err := os.Executable()
	if err != nil {
		log.Fatalln(err)
	}
	exePath := filepath.Dir(exe)

	pathGuesses := filepath.Join(exePath, "data", dictionaryTag, "guesses.txt")
	pathSolutions := filepath.Join(exePath, "data", dictionaryTag, "solutions.txt")

	dictGuesses := ReadDictionary(pathGuesses)
	dictSolutions := ReadDictionary(pathSolutions)

	fmt.Printf("Dictionaries: %d valid guesses, %d solutions\n", len(dictGuesses), len(dictSolutions))

	return dictGuesses, dictSolutions
}
