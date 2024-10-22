package fstr_test

import (
	"fmt"
	"testing"

	"github.com/illbjorn/fstr"
)

var (
	// Good Value 1
	// Just test basic string interpolation.
	input1  = "The {val1} brown fox {val2} over the lazy {val3}."
	map1    = map[string]string{"val1": "quick", "val2": "jumped", "val3": "dog"}
	output1 = "The quick brown fox jumped over the lazy dog."
	// Good Value 2
	// Testing interpolation token escaping.
	input2  = `The \{quick} brown fox \{jumped} over the lazy \{dog}.`
	map2    = map[string]string{"quick": "fast", "jumped": "leaped", "dog": "bear"}
	output2 = "The {quick} brown fox {jumped} over the lazy {dog}."
)

func TestMap(t *testing.T) {
	// Good Value 1
	if res := fstr.Map(input1, map1); res != output1 {
		println("Input1 did not produce expected results.")
		println(res)
		t.Fail()
	}

	// Good Value 2
	if res := fstr.Map(input2, map2); res != output2 {
		println("Input2 did not produce expected results.")
		println(res)
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
