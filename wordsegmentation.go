package wordsegmentation

import (
	"math"

	m "github.com/theosiemensrhodes/wordsegmentation/models"
)

// Corpus  lets access bigrams,
// unigrams, the total number of words from the corpus
// and a function to clean a string.
//
// This is the interface you will need to implement if
// you want to use a custom corpus.
type Corpus interface {
	Bigrams() *m.Bigrams
	Unigrams() *m.Unigrams
	Total() float64
	Clean(string) string
}

// Segmentor holds word segmentation functionality for specific corpus
type Segmentor struct {
	corpus        Corpus
	candidates    m.Candidates
	maxWordLength int
}

/*
NewSegmentor creates new segmentor instance that holds the language corpus,  candidate table and max word length
  - 99.997% of english words are less than or equal to 24 letters
  - 99% of english words are less than or equal to 12 letters
*/
func NewSegmentor(languageCorpus Corpus, maxWordLength int) *Segmentor {
	if maxWordLength == 0 {
		maxWordLength = 24
	}

	return &Segmentor{
		corpus:        languageCorpus,
		maxWordLength: maxWordLength,
	}
}

// Segment Return a list of words that is the best segmentation of a given text.
func (s *Segmentor) Segment(text string) []string {
	return s.search(s.corpus.Clean(text), "<s>", s.maxWordLength).Words
}

// Score a word in the context of the previous word.
func (s *Segmentor) score(current, previous string) float64 {
	if len(previous) == 0 {
		unigramScore := s.corpus.Unigrams().ScoreForWord(current)
		if unigramScore > 0 {
			// Probability of the current word
			return unigramScore / s.corpus.Total()
		}
		// Penalize words not found in the unigrams according to their length
		return 10.0 / (s.corpus.Total() * math.Pow(10, float64(len(current))))
	}
	// We've got a bigram
	unigramScore := s.corpus.Unigrams().ScoreForWord(previous)
	if unigramScore > 0 {
		bigramScore := s.corpus.Bigrams().ScoreForBigram(m.Bigram{
			First:  previous,
			Second: current,
			Rating: 0,
		})
		if bigramScore > 0 {
			// Conditional probability of the word given the previous
			// word. The technical name is 'stupid backoff' and it's
			// not a probability distribution
			return bigramScore / s.corpus.Total() / s.score(previous, "<s>")
		}
	}

	return s.score(current, "")
}

// Find candidates for a given text and an optional previous chunk of letters.
func (s *Segmentor) findCandidates(text, prev string, maxWordLength int) <-chan m.Arrangement {
	ch := make(chan m.Arrangement)

	go func() {
		for p := range divide(text, maxWordLength) {
			prefixScore := math.Log10(s.score(p.Prefix, prev))
			arrangement := s.candidates.ForPossibility(p)
			if len(arrangement.Words) == 0 {
				arrangement = s.search(p.Suffix, p.Prefix, maxWordLength)
				s.candidates.Add(m.Candidate{
					P: p,
					A: arrangement,
				})
			}

			var slice []string
			slice = append(slice, p.Prefix)
			slice = append(slice, arrangement.Words...)
			ch <- m.Arrangement{
				Words:  slice,
				Rating: prefixScore + arrangement.Rating,
			}
		}
		close(ch)
	}()

	return ch
}

// Search for the best arrangement for a text in the context of a previous phrase.
func (s *Segmentor) search(text, prev string, maxWordLength int) (ar m.Arrangement) {
	if len(text) == 0 {
		return m.Arrangement{}
	}

	max := -10000000.0

	// Find the best candidate by finding the best arrangement rating
	for a := range s.findCandidates(text, prev, maxWordLength) {
		if a.Rating > max {
			max = a.Rating
			ar = a
		}
	}

	return
}

// Create multiple (prefix, suffix) pairs from a text.
// The length of the prefix should not exceed the 'limit'.
func divide(text string, limit int) <-chan m.Possibility {
	ch := make(chan m.Possibility)
	bound := min(len(text), limit)

	go func() {
		for i := 1; i <= bound; i++ {
			ch <- m.Possibility{
				Prefix: text[:i],
				Suffix: text[i:],
			}
		}
		close(ch)
	}()

	return ch
}
