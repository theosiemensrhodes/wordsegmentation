package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBigramGetKey(t *testing.T) {
	b := Bigram{
		First:  "foo",
		Second: "bar",
		Rating: 22,
	}
	assert.Equal(t, b.GetKey(), "foo#bar")
}

func TestBigrams(t *testing.T) {
	collection := NewBigrams()
	bigram := Bigram{
		First:  "foo",
		Second: "bar",
		Rating: 22,
	}
	collection.Add(bigram)

	assert.Equal(t, collection.ScoreForBigram(bigram), bigram.Rating)
	assert.Equal(t, collection.ScoreForBigram(bigram), 0.0)
}
