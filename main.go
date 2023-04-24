package main

import (
	"fmt"
)

// The following cases can be tested
// {[}] = NO
// {}[]() = YES
const brackets = "{[}]"

func main() {
	// Convert to slice.
	characters := []rune(brackets)

	// String initial characters.
	str := []rune{}
	err := false

	for _, character := range characters {
		// Validate open characters.
		if character == '{' || character == '(' || character == '[' {
			str = append(str, character)
			continue
		}

		// Validate close characters.
		if character == '}' || character == ')' || character == ']' {
			str, err = removeCharacter(character, str)
		}

		// When we find a closed character we finish the whole operation.
		if err {
			fmt.Println("NO")
			return
		}
	}

	// When str is empty it is because all characters are balanced..
	if len(str) == 0 {
		fmt.Println("YES")
		return
	}

	// When the characters are incorrect we return
	fmt.Println("NO")
}

func removeCharacter(c rune, str []rune) ([]rune, bool) {
	// Create a map relating closing characters to opening characters.
	mc := map[rune]rune{
		']': '[',
		'}': '{',
		')': '(',
	}

	// Character length.
	cl := len(str)

	// When it is the last character we finish the operation..
	if cl == 0 {
		return str, true
	}

	// If the character is the opposite, we eliminate the last character.
	if str[cl-1] == mc[c] {
		return str[:cl-1], false
	}

	return str, false
}
