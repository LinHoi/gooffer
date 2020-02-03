package main

import "fmt"

func LongestSubstringWithoutDuplication(str string) (string, int) {
	prePositionFrom1 := make(map[rune]int, 26)
	curLength, maxLength := 0, 0
	endIndexOfSubstring := 0
	for strIndex := 0; strIndex < len(str); strIndex++ {
		preIndexFrom1 := prePositionFrom1[rune(str[strIndex])]
		if preIndexFrom1 == 0 || (strIndex+1-preIndexFrom1) > curLength {
			curLength ++
		} else {
			curLength = strIndex + 1 - preIndexFrom1
		}
		prePositionFrom1[rune(str[strIndex])] = strIndex + 1
		if curLength > maxLength {
			maxLength = curLength
			endIndexOfSubstring = strIndex
		}
	}

	return str[endIndexOfSubstring-maxLength+1 : endIndexOfSubstring+1], maxLength
}

func main() {
	strs := []string{
		"qwrafffsqwttre",
		"ffertrtrewt",
		"q",
		"aaa",
		"asa",
		"qwertyuiopasdfghjklzxcvbnmmnbvcx",
	}
	for _, str := range strs {
		resStr, length := LongestSubstringWithoutDuplication(str)
		fmt.Printf("Longest substring of %s without duplication is %s with length %d\n", str, resStr, length)
	}
}

/* output
Longest substring of qwrafffsqwttre without duplication is qwraf with length 5
Longest substring of ffertrtrewt without duplication is fert with length 4
Longest substring of q without duplication is q with length 1
Longest substring of aaa without duplication is a with length 1
Longest substring of asa without duplication is as with length 2
Longest substring of qwertyuiopasdfghjklzxcvbnmmnbvcx without duplication is qwertyuiopasdfghjklzxcvbnm with length 26
 */
