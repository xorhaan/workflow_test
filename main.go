package main

import "strings"

func Add(a, b int) int {
	return a + b
}

func NormalizeName(name string) string {
	return strings.TrimSpace(name)
}
