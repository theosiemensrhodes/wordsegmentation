package corpus

import (
	"regexp"
	"strings"

	"github.com/kennygrant/sanitize"
	"github.com/theosiemensrhodes/wordsegmentation/data/english"
	m "github.com/theosiemensrhodes/wordsegmentation/models"
	"github.com/theosiemensrhodes/wordsegmentation/parsers"
)

type EnglishCorpus struct {
	unigrams      m.Unigrams
	bigrams       m.Bigrams
	maxWordLength int
}

// Create a new EnglishCorpus by loading unigrams
// and bigrams from TSV files.
func NewEnglishCorpus() *EnglishCorpus {
	var bigrams m.Bigrams
	var unigrams m.Unigrams
	var maxWordLength int

	// Load unigrams and bigrams from data files.
	done := make(chan int)
	go func() {
		bigrams = parsers.Bigrams(english.Bigrams_data)
		done <- 1
	}()
	go func() {
		unigrams, maxWordLength = parsers.Unigrams(english.Unigrams_data)
		done <- 1
	}()
	<-done
	<-done

	return &EnglishCorpus{unigrams, bigrams, maxWordLength}
}

// Get bigrams from the corpus.
func (corpus *EnglishCorpus) Bigrams() *m.Bigrams {
	return &corpus.bigrams
}

// Get unigrams from the corpus.
func (corpus *EnglishCorpus) Unigrams() *m.Unigrams {
	return &corpus.unigrams
}

// Get the total number of words in the corpus.
func (corpus *EnglishCorpus) Total() float64 {
	return 1024908267229.0
}

// Get the longest word in the corpus.
func (corpus *EnglishCorpus) MaxWordLength() int {
	return corpus.maxWordLength
}

// Clean a string from special characters.
func (corpus *EnglishCorpus) Clean(s string) string {
	s = sanitize.Accents(strings.ToLower(s))
	return regexp.MustCompile("[^a-zA-Z0-9]+").ReplaceAllString(s, "")
}
