package main

import (
	"strings"
	"testing"
)

func JoinWordsBuilder(words []string) string {
	var builder strings.Builder
	for i, word := range words {
		if i > 0 {
			builder.WriteString(",")
		}
		builder.WriteString(word)
	}

	return builder.String()
}

var benchmarkWords = []string{
	"go", "bench", "practice", "profile", "improve",
	"measure", "repeat", "learn", "ship", "refine",
}

func BenchmarkJoinWords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = JoinWords(benchmarkWords)
	}
}

func BenchmarkJoinWordsBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = JoinWordsBuilder(benchmarkWords)
	}
}
