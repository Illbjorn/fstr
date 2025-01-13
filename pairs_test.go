package fstr_test

import (
	"testing"

	"github.com/illbjorn/fstr"
	"github.com/stretchr/testify/assert"
)

func TestPairs(t *testing.T) {
	var (
		// Good Value 1
		// Just a basic interpolation test
		goodInput1  = "The {val1} brown fox {val2} over the lazy {val3}."
		goodValues1 = []string{"val1", "quick", "val2", "jumped", "val3", "dog"}
		goodOutput1 = "The quick brown fox jumped over the lazy dog."

		// Good Value 2
		// Testing interpolation token escaping
		goodInput2  = `The \{quick} brown fox \{jumped} over the lazy \{dog}.`
		goodValues2 = []string{"quick", "fast", "jumped", "leaped", "dog", "bear"}
		goodOutput2 = "The {quick} brown fox {jumped} over the lazy {dog}."

		// Good Value 3
		// Interpolation token present for which there is no `pair` value to replace
		// it with
		goodInput3  = "The {quick} brown fox {jumped} over the lazy {dog}."
		goodValues3 = []string{"quick", "fast", "jumped", "leaped"}
		goodOutput3 = "The fast brown fox leaped over the lazy dog."

		// Good Value 4
		// Tests the early return (no values to interpolate) code path
		goodInput4  = "The {quick} brown fox {jumped} over the lazy {dog}."
		goodValues4 = []string{}
		goodOutput4 = "The {quick} brown fox {jumped} over the lazy {dog}."

		// Good Value 5
		// Testing curly brace escaping
		goodInput5  = `The \{quick} brown fox \{jumped} over the lazy \{dog}.`
		goodValues5 = []string{"quick", "fast", "jumped", "leaped", "dog", "bear"}
		goodOutput5 = "The {quick} brown fox {jumped} over the lazy {dog}."

		// Good Value 6
		// Testing curly brace escaping with immediately following interpolation token
		goodInput6  = `The \{{quick}} brown fox \{{jumped}} over the lazy \{{dog}}.`
		goodValues6 = []string{"quick", "fast", "jumped", "leaped", "dog", "bear"}
		goodOutput6 = "The {fast} brown fox {leaped} over the lazy {bear}."

		// Good Value 7
		// Testing trailing escaped character
		goodInput7  = `The {quick} brown fox {jumped} over the lazy {dog}.\n`
		goodValues7 = []string{"quick", "fast", "jumped", "leaped", "dog", "bear"}
		goodOutput7 = `The fast brown fox leaped over the lazy bear.\n`

		// Bad Value 1
		// Uneven number of pairs (key with no value or vice versa)
		badInput1  = `The {quick} brown fox {jumped} over the lazy {dog}.`
		badValues1 = []string{"quick", "fast", "jumped", "leaped", "dog"}
		badOutput1 = "The fast brown fox leaped over the lazy dog."

		res string
	)

	// Good Value 1
	res = fstr.Pairs(goodInput1, goodValues1...)
	assert.Equal(t, goodOutput1, res)

	// Good Value 2
	res = fstr.Pairs(goodInput2, goodValues2...)
	assert.Equal(t, goodOutput2, res)

	// Good Value 3
	res = fstr.Pairs(goodInput3, goodValues3...)
	assert.Equal(t, goodOutput3, res)

	// Good Value 4
	res = fstr.Pairs(goodInput4, goodValues4...)
	assert.Equal(t, goodOutput4, res)

	// Good Value 5
	res = fstr.Pairs(goodInput5, goodValues5...)
	assert.Equal(t, goodOutput5, res)

	// Good Value 6
	res = fstr.Pairs(goodInput6, goodValues6...)
	assert.Equal(t, goodOutput6, res)

	// Good Value 7
	res = fstr.Pairs(goodInput7, goodValues7...)
	assert.Equal(t, goodOutput7, res)

	// Bad Value 1
	res = fstr.Pairs(badInput1, badValues1...)
	assert.Equal(t, badOutput1, res)
}

var (
	// Benchmark Pairs Template
	benchPairsInput  = "The {val1} brown {val2} jumped over the lazy {val3}."
	benchPairsValues = []string{"val1", "quick", "val2", "fox", "val3", "dog"}
)

func BenchmarkPairs(b *testing.B) {
	for range b.N {
		fstr.Pairs(benchPairsInput, benchPairsValues...)
	}
}
