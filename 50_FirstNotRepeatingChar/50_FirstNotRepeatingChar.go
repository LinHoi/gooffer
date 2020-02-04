package main

import (
	"errors"
	"fmt"
)

func FirstNotRepeatingChar(str string) (rune, error) {
	if len(str) == 0 {
		return rune(' '), errors.New("string is nil")
	}
	hashTable := make(map[uint8]int, 256)
	for index := range str {
		hashTable[str[index]] ++
	}
	for index := range str {
		if hashTable[str[index]] == 1 {
			return rune(str[index]),nil
		}
	}
	return rune(' '),  errors.New("every char has repeat")
}

func main(){
	strings := []string{
		"wwqerrrf24455",
		"qwrtuifsfhbhhfds",
		"",
		"wtrwtreytuuioijdsgkasgfdgdkhsskcvdsjkhjsrl34593690JFGSJGKSHAghnh",
		"qwertyuioppoiuytrewq",
	}
	for _, string := range strings {
		res ,err := FirstNotRepeatingChar(string)
		if err != nil {
			fmt.Println(err," in ", string)
		} else {
			fmt.Printf("the first not repeating char of %s is %c \n",string,res)
		}
	}
}

/*
the first not repeating char of wwqerrrf24455 is q
the first not repeating char of qwrtuifsfhbhhfds is q
string is nil  in
the first not repeating char of wtrwtreytuuioijdsgkasgfdgdkhsskcvdsjkhjsrl34593690JFGSJGKSHAghnh is e
every char has repeat  in  qwertyuioppoiuytrewq
 */