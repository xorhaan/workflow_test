package main

import "strings"

func Add(a, b int) int {
	return b + a
}

func NormalizeName(name string) string {
	return strings.TrimSpace(name)
}
