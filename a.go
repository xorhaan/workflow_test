package main

import (
	"sort"
	"strings"
)

func Double(value int) int {
	return value * 2
}

func IsEven(value int) bool {
	return value%2 == 0
}

func Reverse(text string) string {
	runes := []rune(text)
	for left, right := 0, len(runes)-1; left < right; left, right = left+1, right-1 {
		runes[left], runes[right] = runes[right], runes[left]
	}
	return string(runes)
}

func WordFrequency(text string) map[string]int {
	frequency := make(map[string]int)
	for _, word := range strings.Fields(strings.ToLower(text)) {
		frequency[word]++
	}
	return frequency
}

func MergeIntervals(intervals [][2]int) [][2]int {
	if len(intervals) == 0 {
		return nil
	}

	sorted := append([][2]int(nil), intervals...)
	sort.Slice(sorted, func(left, right int) bool {
		return sorted[left][0] < sorted[right][0]
	})

	merged := [][2]int{sorted[0]}
	for _, current := range sorted[1:] {
		last := &merged[len(merged)-1]
		if current[0] <= last[1] {
			if current[1] > last[1] {
				last[1] = current[1]
			}
			continue
		}
		merged = append(merged, current)
	}

	return merged
}
