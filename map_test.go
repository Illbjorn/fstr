package fstr_test

import (
	"fmt"
	"testing"

	"github.com/illbjorn/fstr"
)

var (
	// Good Value 1
	// Just test basic string interpolation.
	mapInput1  = "The {val1} brown fox {val2} over the lazy {val3}."
	mapValues1 = map[string]string{"val1": "quick", "val2": "jumped", "val3": "dog"}
	mapOutput1 = "The quick brown fox jumped over the lazy dog."

	// Good Value 2
	// Testing interpolation token escaping.
	mapInput2  = `The \{quick} brown fox \{jumped} over the lazy \{dog}.`
	mapValues2 = map[string]string{"quick": "fast", "jumped": "leaped", "dog": "bear"}
	mapOutput2 = "The {quick} brown fox {jumped} over the lazy {dog}."

	// Good Value 3
	// Interpolation token present for which there is no `pair` value to replace
	// it with.
	mapInput3  = "The {quick} brown fox {jumped} over the lazy {dog}."
	mapValues3 = map[string]string{"quick": "fast", "jumped": "leaped"}
	mapOutput3 = "The fast brown fox leaped over the lazy dog."

	// Good Value 4
	// Tests the early return (no values to interpolate) code path.
	mapInput4  = "The {quick} brown fox {jumped} over the lazy {dog}."
	mapValues4 = map[string]string{}
	mapOutput4 = "The {quick} brown fox {jumped} over the lazy {dog}."
)

func TestMap(t *testing.T) {
	// Good Value 1
	if res := fstr.Map(mapInput1, mapValues1); res != mapOutput1 {
		t.Error("Map input 1 did not produce expected results.")
		t.Error("Expected:", mapOutput1)
		t.Error("Received:", res)
		t.Fail()
	}

	// Good Value 2
	if res := fstr.Map(mapInput2, mapValues2); res != mapOutput2 {
		t.Error("Map input 2 did not produce expected results.")
		t.Error("Expected:", mapOutput2)
		t.Error("Received:", res)
		t.Fail()
	}

	// Good Value 3
	if res := fstr.Map(mapInput3, mapValues3); res != mapOutput3 {
		t.Error("Map input 3 did not produce expected results.")
		t.Error("Expected:", mapOutput3)
		t.Error("Received:", res)
		t.Fail()
	}

	// Good Value 4
	if res := fstr.Map(mapInput4, mapValues4); res != mapOutput4 {
		t.Error("Map input 4 did not produce expected results.")
		t.Error("Expected:", mapOutput4)
		t.Error("Received:", res)
		t.Fail()
	}
}

var (
	// Benchmark Map Template
	benchMapInput  = "The {val1} brown {val2} jumped over the lazy {val3}."
	benchMapValues = map[string]string{"val1": "quick", "val2": "fox", "val3": "dog"}
	// Benchmark Sprintf Template
	benchSprintfInput  = "The %s brown %s jumped over the lazy %s."
	benchSprintfValues = []string{"quick", "fox", "dog"}
)

func BenchmarkMap(b *testing.B) {
	for range b.N {
		fstr.Map(benchMapInput, benchMapValues)
	}
}

func BenchmarkSprintf(b *testing.B) {
	for range b.N {
		fmt.Sprintf(benchSprintfInput, benchSprintfValues)
	}
}
