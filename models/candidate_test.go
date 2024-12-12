package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForPossibility(t *testing.T) {
	candidates := Candidates{}
	arrangement := Arrangement{
		Words:  []string{"a", "b"},
		Rating: 42,
	}
	possibility := Possibility{
		Prefix: "prefix",
		Suffix: "suffix",
	}
	candidate := Candidate{
		P: possibility,
		A: arrangement,
	}
	candidates.Add(candidate)

	assert.Equal(t, candidates.ForPossibility(possibility), arrangement)
	possibilityNotFound := Possibility{
		Prefix: "not",
		Suffix: "found",
	}
	assert.Equal(t, len(candidates.ForPossibility(possibilityNotFound).Words), 0)
}
