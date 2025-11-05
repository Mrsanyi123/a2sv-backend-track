// Write a Go function that takes a string as input
// and returns a dictionary containing the frequency of
// each word in the string. Treat words in a case-insensitive
// manner and ignore punctuation marks.
// [Optional]: Write test for your function

package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"

)
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Text: ")
	text, _ := reader.ReadString('\n') 
	text  = strings.TrimSpace(text)
	text = strings.ToLower(text)
    

	// remove  punctuation
	cleaned := ""
	for _, c := range text {
		if (c >= 'a' && c <= 'z') || c == ' ' {
			cleaned += string(c)
		}
	}

	words := strings.Split(cleaned, " ")

	freq := make(map[string]int)
	for _, word := range words {
		if word != "" {
			freq[word]++
		}
	}

	fmt.Println("Word Frequency Count:", freq) 

}