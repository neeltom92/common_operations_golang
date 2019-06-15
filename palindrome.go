package main

import (
	"bufio"
	"fmt"
	"os"
)

func pal_drom(pal_word string) {

	runes := []rune(pal_word)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]

	}

	orig_pal := pal_word
	rev_pal := string(runes)

	if orig_pal == rev_pal {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Not Palindome")
	}

}

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	textv1 := text[:len(text)-1]
	pal_drom(textv1)

}
