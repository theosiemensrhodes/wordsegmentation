package parsers

import (
	"bufio"
	"strconv"
	"strings"

	m "github.com/theosiemensrhodes/wordsegmentation/models"
)

// Parse unigrams from a given TSV file.
func Unigrams(path string) (m.Unigrams, int) {
	scanner := bufio.NewScanner(strings.NewReader(path))

	unigrams := m.NewUnigrams()
	var maxWordLength int
	var fields []string
	for scanner.Scan() {
		fields = strings.Split(scanner.Text(), "\t")
		rating, _ := strconv.ParseFloat(fields[1], 64)

		u := m.Unigram{
			Word:   fields[0],
			Rating: rating,
		}
		unigrams.Add(u)

		if len(u.Word) > maxWordLength {
			maxWordLength = len(u.Word)
		}
	}

	return unigrams, maxWordLength
}

// Parse bigrams from a given TSV file.
func Bigrams(path string) m.Bigrams {
	scanner := bufio.NewScanner(strings.NewReader(path))

	bigrams := m.NewBigrams()
	var fields []string
	for scanner.Scan() {
		fields = strings.Split(scanner.Text(), "\t")
		rating, _ := strconv.ParseFloat(fields[2], 64)

		bigrams.Add(m.Bigram{
			First:  fields[0],
			Second: fields[1],
			Rating: rating,
		})
	}

	return bigrams
}
