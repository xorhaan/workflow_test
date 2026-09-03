package main

import (
	"fmt"
	"strings"
)

func Add(a, b int) int {
	return a + b
	fmt.Println("added", a, b)
}

func NormalizeName(name string) string {
	return strings.TrimSpace(name)
}
