package main

import (
	"bufio"
	"os"
	"testing"
)

func TestSplitStrByNewLines(t *testing.T) {
	input := "Hello\nWorld"
	expected := []string{"Hello", "\n", "World"}

	result := splitStrByNewLines(input)

	if len(result) != len(expected) {
		t.Fatalf("expected %d elements, got %d", len(expected), len(result))
	}

	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("at index %d: expected %q, got %q", i, expected[i], result[i])
		}
	}
}

func TestGetBannerMapping(t *testing.T) {
	file, err := os.Open("standard.txt")
	if err != nil {
		t.Fatalf("failed to open standard.txt: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	bannerMap := getBannerMapping(scanner)

	if len(bannerMap['A']) != 8 {
		t.Errorf("expected 8 lines for 'A', got %d", len(bannerMap['A']))
	}

	if _, exists := bannerMap[' ']; !exists {
		t.Errorf("space character not found in banner map")
	}
}
