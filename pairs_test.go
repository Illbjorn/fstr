package fstr_test

import (
	"testing"

	"github.com/illbjorn/fstr"
)

var (
	// Good Value 1
	// Just a basic interpolation test.
	pairsInput1  = "The {val1} brown fox {val2} over the lazy {val3}."
	pairsValues1 = []string{"val1", "quick", "val2", "jumped", "val3", "dog"}
	pairsOutput1 = "The quick brown fox jumped over the lazy dog."

	// Good Value 2
	// Testing interpolation token escaping.
	pairsInput2  = `The \{quick} brown fox \{jumped} over the lazy \{dog}.`
	pairsValues2 = []string{"quick", "fast", "jumped", "leaped", "dog", "bear"}
	pairsOutput2 = "The {quick} brown fox {jumped} over the lazy {dog}."

	// Good Value 3
	// Interpolation token present for which there is no `pair` value to replace
	// it with.
	pairsInput3  = "The {quick} brown fox {jumped} over the lazy {dog}."
	pairsValues3 = []string{"quick", "fast", "jumped", "leaped"}
	pairsOutput3 = "The fast brown fox leaped over the lazy dog."

	// Good Value 4
	// Tests the early return (no values to interpolate) code path.
	pairsInput4  = "The {quick} brown fox {jumped} over the lazy {dog}."
	pairsValues4 = []string{}
	pairsOutput4 = "The {quick} brown fox {jumped} over the lazy {dog}."

	// Bad Value 1
	// Uneven number of pairs (key with no value or vice versa).
	pairsInput5  = `The {quick} brown fox {jumped} over the lazy {dog}.`
	pairsValues5 = []string{"quick", "fast", "jumped", "leaped", "dog"}
	pairsOutput5 = "The fast brown fox leaped over the lazy dog."
)

func TestPairs(t *testing.T) {
	// Good Value 1
	if res := fstr.Pairs(pairsInput1, pairsValues1...); res != pairsOutput1 {
		t.Error("Pairs input 1 did not produce expected results.")
		t.Error("Expected:", pairsOutput1)
		t.Error("Received:", res)
		t.Fail()
	}

	// Good Value 2
	if res := fstr.Pairs(pairsInput2, pairsValues2...); res != pairsOutput2 {
		t.Error("Pairs input 2 did not produce expected results.")
		t.Error("Expected:", pairsOutput2)
		t.Error("Received:", res)
		t.Fail()
	}

	// Good Value 3
	if res := fstr.Pairs(pairsInput3, pairsValues3...); res != pairsOutput3 {
		t.Error("Pairs input 3 did not produce expected results.")
		t.Error("Expected:", pairsOutput3)
		t.Error("Received:", res)
		t.Fail()
	}

	// Good Value 4
	if res := fstr.Pairs(pairsInput4, pairsValues4...); res != pairsOutput4 {
		t.Error("Pairs input 4 did not produce expected results.")
		t.Error("Expected:", pairsOutput4)
		t.Error("Received:", res)
		t.Fail()
	}

	// Bad Value 1
	if res := fstr.Pairs(pairsInput5, pairsValues5...); res != pairsOutput5 {
		t.Error("Pairs input 5 did not produce expected results.")
		t.Error("Expected:", pairsOutput5)
		t.Error("Received:", res)
		t.Fail()
	}
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
