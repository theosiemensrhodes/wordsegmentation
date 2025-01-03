package wordsegmentation

import (
	"math"
	"sync"

	m "github.com/theosiemensrhodes/wordsegmentation/models"
)

// Corpus interface as before
type Corpus interface {
	Bigrams() *m.Bigrams
	Unigrams() *m.Unigrams
	Total() float64
	Clean(string) string
	MaxWordLength() int
}

type Segmentor struct {
	corpus        Corpus
	maxWordLength int
}

func NewSegmentor(c Corpus) *Segmentor {
	return &Segmentor{
		corpus:        c,
		maxWordLength: c.MaxWordLength(),
	}
}

// Override the length of the maximum word to search for
func (segmentor *Segmentor) OverrideMaxWordLength(newValue int) {
	segmentor.maxWordLength = newValue
}

// Get the length of the maximum word word to search for
func (segmentor *Segmentor) MaxWordLength() int {
	return segmentor.maxWordLength
}

// Segment text
func (s *Segmentor) Segment(text string) []string {
	text = s.corpus.Clean(text)
	textLength := len(text)
	if textLength == 0 {
		return nil
	}

	return s.segmentInternal(text, s.MaxWordLength())
}

// Score computes log-prob of current given previous.
func (s *Segmentor) score(current, previous string) float64 {
	if previous == "" {
		unigramCount := s.corpus.Unigrams().ScoreForWord(current)
		total := s.corpus.Total()
		if unigramCount > 0 {
			return math.Log10(unigramCount / total)
		}

		// Penalize unknown
		penalty := 10.0 / (total * math.Pow(10, float64(len(current))))
		return math.Log10(penalty)
	}

	// Otherwise, bigram
	unigramCount := s.corpus.Unigrams().ScoreForWord(previous)
	bigramCount := s.corpus.Bigrams().ScoreForBigram(m.Bigram{
		First:  previous,
		Second: current,
	})

	if unigramCount > 0 && bigramCount > 0 {
		// log[P(current | previous)] => log[Count(current && previous)/Count(previous)]
		// Add a term to incentivize choosing pairs from the bigram dictionary
		return math.Log10(bigramCount/unigramCount) + 0.4
	}

	// Fallback to unigram
	return s.score(current, "")
}

// Uses a DP approach to optimize call and space complexity
func (s *Segmentor) segmentInternal(text string, maxWordLength int) []string {
	textLength := len(text)

	// dp[i][j]: best log-prob for text[:i], with last word = text[j:i].
	dp := make([][]float64, textLength+1)
	// prev[i][j]: the index k in [0..j] that yields dp[i][j] = dp[j][k] + Score(text[j:i])
	back := make([][]int, textLength+1)

	// Initialize arrays
	for i := 0; i <= textLength; i++ {
		dp[i] = make([]float64, i+1)
		back[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			dp[i][j] = math.Inf(-1)
			back[i][j] = -1
		}
	}
	dp[0][0] = 0.0 // No text has log prob of 0

	// Fill table row by row
	for i := 1; i <= textLength; i++ {
		// Parallelize over the j loop
		var wg sync.WaitGroup

		// j is where the last word starts => text[j:i]
		startJ := max(i-maxWordLength, 0)

		// Compute dp[i][j] for j in [startJ..i-1]
		for j := startJ; j < i; j++ {
			wg.Add(1)
			go func(j, i int) {
				defer wg.Done()

				word := text[j:i]

				// Search k in [j - maxWordLength.. j]
				startK := max(j-maxWordLength, 0)
				bestScore := math.Inf(-1)
				bestK := -1

				for k := startK; k <= j; k++ {
					if math.IsInf(dp[j][k], -1) {
						continue
					}

					// If j == 0, start of the sentence
					prevWord := "<s>"
					if j > 0 {
						prevWord = text[k:j]
					}
					newScore := dp[j][k] + s.score(word, prevWord)

					if newScore > bestScore {
						bestScore = newScore
						bestK = k
					}
				}

				// Write back the result for dp[i][j]
				dp[i][j] = bestScore
				back[i][j] = bestK
			}(j, i)
		}
		wg.Wait()
	}

	// Best solution is in dp[n][j] for j in [0..n]
	bestEnd := -1
	bestVal := math.Inf(-1)
	for j := 0; j < textLength+1; j++ {
		if dp[textLength][j] > bestVal {
			bestVal = dp[textLength][j]
			bestEnd = j
		}
	}
	if bestEnd == -1 || math.IsInf(bestVal, -1) {
		// no valid segmentation
		return nil
	}

	// Backtrack
	var words []string
	i := textLength
	j := bestEnd

	for i > 0 && j >= 0 {
		if j == 0 {
			// Means the final word is text[0:i]
			words = append(words, text[0:i])
			break
		}
		k := back[i][j]
		words = append(words, text[j:i])
		i = j
		j = k
	}

	// Reverse the slice since we appended from end->front
	for left, right := 0, len(words)-1; left < right; left, right = left+1, right-1 {
		words[left], words[right] = words[right], words[left]
	}

	return words
}
