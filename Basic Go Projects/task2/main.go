package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func wordCount(input string) map[string]int {
	input = removePunctuation(strings.TrimSpace(input))
	input = strings.ToLower(input)

	inputMap := make(map[string]int)
	inputArr := strings.Split(input, " ")

	for _, value := range inputArr {
		inputMap[value]++
	}
	return inputMap
}

// function to remove the punctuations
func removePunctuation(s string) string {
	var builder strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || unicode.IsSpace(r) {
			builder.WriteRune(r)
		}
	}
	return builder.String()
}

func isPalindrome(input string) {
	input = removePunctuation(strings.TrimSpace(input))
	input = strings.ToLower(input)
	i := 0
	n := len(input) - 1

	for i < n {
		if input[i] != input[n] {
			fmt.Println("Not Palindrome!")
			return
		}
		i++
		n--
	}
	fmt.Println("Palindrome!")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a string: ")
	input, _ := reader.ReadString('\n')
	isPalindrome(input)
	//wordCount(input)

}
