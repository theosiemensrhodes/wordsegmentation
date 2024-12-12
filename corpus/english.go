package corpus

import (
	"github.com/theosiemensrhodes/wordsegmentation/data/english"
	help "github.com/theosiemensrhodes/wordsegmentation/helpers"
	m "github.com/theosiemensrhodes/wordsegmentation/models"
	"github.com/theosiemensrhodes/wordsegmentation/parsers"
)

type EnglishCorpus struct {
	unigrams m.Unigrams
	bigrams  m.Bigrams
}

// Create a new EnglishCorpus by loading unigrams
// and bigrams from TSV files.
func NewEnglishCorpus() EnglishCorpus {
	var bigrams m.Bigrams
	var unigrams m.Unigrams

	// Load unigrams and bigrams from data files.
	done := make(chan int)
	go func() {
		bigrams = parsers.Bigrams(english.Bigrams_data)
		done <- 1
	}()
	go func() {
		unigrams = parsers.Unigrams(english.Unigrams_data)
		done <- 1
	}()
	<-done
	<-done

	return EnglishCorpus{unigrams, bigrams}
}

// Get bigrams from the corpus.
func (corpus EnglishCorpus) Bigrams() *m.Bigrams {
	return &corpus.bigrams
}

// Get unigrams from the corpus.
func (corpus EnglishCorpus) Unigrams() *m.Unigrams {
	return &corpus.unigrams
}

// Get the total number of words in the corpus.
func (corpus EnglishCorpus) Total() float64 {
	return 1024908267229.0
}

// Clean a string from special characters.
func (corpus EnglishCorpus) Clean(s string) string {
	return help.CleanString(s)
}
