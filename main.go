package main

import (
	"fmt"
	"strings"
	"time"
)

// Add returns the sum of two integers.
func Add(a, b int) int {
	fmt.Println("adding a and b dev-experiment")
	return b + a
}

// ReverseString reverses a string.
func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// IsEven checks whether a number is even.
func IsEven(n int) bool {
	return n%2 == 0
}

// DEV-only function.
func IsOdd(n int) bool {
	return n%2 != 0
}

// ToUpperCase converts text to uppercase.
func ToUpperCase(text string) string {
	fmt.Println("dev logging:", text)
	return strings.ToUpper(text)
}

// DEV-only helper.
func ToLowerCase(text string) string {
	return strings.ToLower(text)
}

// GetCurrentTimestamp returns UTC timestamp.
func GetCurrentTimestamp() string {
	return time.Now().UTC().Format(time.RFC3339)
}

func main() {
	fmt.Println(Add(5, 34))
	fmt.Println(ReverseString("hello"))
	fmt.Println(IsEven(10))
	fmt.Println(IsOdd(11))
	fmt.Println(ToUpperCase("golang"))
	fmt.Println(ToLowerCase("GOLANG"))
	fmt.Println(GetCurrentTimestamp())
}
