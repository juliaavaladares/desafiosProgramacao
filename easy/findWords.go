package easy

import (
	"fmt"
	"strings"
)

// I would like you to write a function which takes two arguments: a list of words, each separated by comma, and a string of letters.
//The function should output the longest words from which can be made using the string of letters.

// ---------
// Inputs
// ---------
// Words: aab,cab,abba,abacus,scabs,scab,bacca,ab
// Letters: aabbscb

// ----------------------
// Expected Output
// ----------------------
// abba, scab

func EncontraPalavras(letters string, words string) {

	wordsArray := strings.Split(words, ",")
	lenLongestWord := 0
	var longestsWords []string

	for _, word := range wordsArray {
		isValid := true
		lettersCopy := letters

		for _, c := range word {
			character := string(c)
			isIn := strings.Contains(lettersCopy, character)
			if !isIn {
				isValid = false
				break
			}
			lettersCopy = strings.Replace(lettersCopy, character, "", 1)

		}

		if isValid {
			if len(word) >= lenLongestWord {
				lenLongestWord = len(word)
				longestsWords = append(longestsWords, word)
			}
		}
	}

	for _, w := range longestsWords {
		if len(w) == lenLongestWord {
			fmt.Println(w)
		}
	}
}
