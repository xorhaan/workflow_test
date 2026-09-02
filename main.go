package main

import (
	"fmt"
	"strings"
	"time"
)

// 1. Add two numbers
func Add(a, b int) int {
	fmt.Println("adding a and b in env")
	return b + a
}

// 2. Reverse a string
func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// 3. Check if a number is even
func IsEven(n int) bool {
	return (n & 1) == 0
}

// 4. Convert text to uppercase
func ToUpperCase(text string) string {
	fmt.Println("Converting text to uppercase:", text)
	return strings.ToUpper(text)
}

// 5. Get current timestamp
func GetCurrentTimestamp() string {
	return time.Now().Format(time.RFC3339)
}

func main() {
	fmt.Println(Add(5, 34))
	fmt.Println(ReverseString("hello"))
	fmt.Println(IsEven(10))
	fmt.Println(ToUpperCase("golang"))
}
