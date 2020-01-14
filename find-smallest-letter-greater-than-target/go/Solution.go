package main

import (
	"fmt"
)

func nextGreatestLetter(letters []byte, target byte) byte {
	letters_len := len(letters)

	// Before the first
	if target < letters[0] {
		return letters[0]
	}
	// After the last
	if target >= letters[letters_len-1] {
		return letters[0]
	}
	// Somewhere in the middle
	next := 0
	for l, r := 0, letters_len; l < r; {
		next = l + ((r - l) / 2)
		if letters[next] > target {
			if next == 0 || letters[next - 1] <= target {
				break
			}
			r = next
		} else {
			l = next
		}
	}

  return letters[next]
}

func main() {
	fmt.Println(string(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'a')))
	fmt.Println(string(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'c')))
	fmt.Println(string(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'd')))
	fmt.Println(string(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'g')))
	fmt.Println(string(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'j')))
	fmt.Println(string(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'k')))
}