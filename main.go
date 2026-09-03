package main

import (
	"fmt"
	"strings"
)

func Add(a, b int) int {
	fmt.Println("hotfix applied")
	return a + b
}

func NormalizeName(name string) string {
	return strings.TrimSpace(name)
}
