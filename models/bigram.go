package models

type Bigram struct {
	First  string
	Second string
	Rating float64
}

type Bigrams struct {
	data map[[2]string]float64
}

// Get a unique identifier for a Bigram.
func (b *Bigram) GetKey() [2]string {
	return [2]string{b.First, b.Second}
}

// Create a new collection of bigrams.
func NewBigrams() Bigrams {
	return Bigrams{
		data: make(map[[2]string]float64),
	}
}

// Add another Bigram to the collection.
func (b *Bigrams) Add(other Bigram) {
	b.data[other.GetKey()] = other.Rating
}

// Get the score of a Bigram. If the Bigram was
// not found, the score will be 0.
func (b *Bigrams) ScoreForBigram(other Bigram) float64 {
	score, has := b.data[other.GetKey()]
	if !has {
		return 0
	}
	return score
}
